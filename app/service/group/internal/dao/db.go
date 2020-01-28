package dao

import (
	pb "chat/app/service/group/api"
	"chat/app/service/group/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	_getAllGroups         = "select id,name,description from `group`"
	_getGroupByIdSql      = "select id,description,name  from `group` where id =? "
	_getMembersOfGroupSQL = `select uid from membership  where gid =?`
	_getGroupsByUserId    = "select id,name,description from `group` where id in (select gid from membership where uid =?)"
	_getGroupsByNameSql   = "select id,name,description from `group` where name like ?"
	_createGroupSql       = "INSERT INTO `group` (name, creator_id, description) VALUES (?,?,?)"
	_insertMemberSql      = "insert into membership values (?,?)"

	_mysqlErrCodeForeignConstraint = 1216
	_mysqlErrCodeDuplicateKeyEntry = 1062
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() { db.Close() }
	return
}

func (d *dao) RawGroup(ctx context.Context, gid int64) (g *model.Group, err error) {
	var row *sql.Row
	row = d.db.QueryRow(ctx, _getGroupByIdSql, gid)
	if row == nil {
		return
	}

	g = new(model.Group)
	if err = row.Scan(&g.Id, &g.Description, &g.Name); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
	}

	var rows *sql.Rows
	if rows, err = d.db.Query(ctx, _getMembersOfGroupSQL, gid); err != nil {
		log.Error("Match:d.db.Query error(%v)", err)
		return
	}

	for rows.Next() {
		var uid int64
		if err = rows.Scan(&uid); err != nil {
			log.Error("Match:row.Scan() error(%v)", err)
			return
		}
		g.Members = append(g.Members, uid)
	}
	return
}

func (d *dao) GetAllGroupsByUserId(ctx context.Context, uid int64) (groups []*model.Group, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(ctx, _getGroupsByUserId, uid); err != nil {
		log.Error("Match:d.db.Query error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		g := new(model.Group)
		if err = rows.Scan(&g.Id, &g.Name, &g.Description); err != nil {
			log.Error("Match:row.Scan() error(%v)", err)
			return
		}
		groups = append(groups, g)
	}
	return
}

func (d *dao) GetAllGroupsByName(ctx context.Context, name string) (groups []*model.Group, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(ctx, _getGroupsByNameSql, "%"+name+"%"); err != nil {
		log.Error("Match:d.db.Query error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		g := new(model.Group)
		if err = rows.Scan(&g.Id, &g.Name, &g.Description); err != nil {
			log.Error("Match:row.Scan() error(%v)", err)
			return
		}
		groups = append(groups, g)
	}
	return
}

func (d *dao) GetAllGroups(ctx context.Context) (groups []*model.Group, err error) {
	var rows *sql.Rows
	if rows, err = d.db.Query(ctx, _getAllGroups); err != nil {
		log.Error("Match:d.db.Query error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		g := new(model.Group)
		if err = rows.Scan(&g.Id, &g.Name, &g.Description); err != nil {
			log.Error("Match:row.Scan() error(%v)", err)
			return
		}

		groups = append(groups, g)
	}
	return
}

func (d *dao) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (info *pb.GroupInfo, err error) {
	if req.Description == "" {
		req.Description = "快来聊天吧~"
	}
	res, err := d.db.Exec(ctx, _createGroupSql, req.Name, req.Uid, req.Description)
	if err != nil {
		return
	}
	info = new(pb.GroupInfo)
	info.Name = req.Name
	info.Description = req.Description
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return nil, ecode.Error(10000, "fatal error when get auto increment id")
	}
	info.Gid = lastInsertedId

	// delete the possible empty cache
	_ = d.DeleteGroupCache(ctx, info.Gid)
	var members []int64
	members = append(members, req.Uid)
	info.Members = members
	info.MemberNumber = 1

	_, err = d.db.Exec(ctx, _insertMemberSql, req.Uid, lastInsertedId)
	if err != nil {
		return
	}
	return
}

// example for error handling
func (d *dao) AddMember(ctx context.Context, uid int64, gid int64) error {
	_, err := d.db.Exec(ctx, _insertMemberSql, uid, gid)
	if err != nil {
		switch nErr := errors.Cause(err).(type) {
		case *mysql.MySQLError:
			if nErr.Number == _mysqlErrCodeForeignConstraint {
				return ecode.Errorf(pb.GroupToAddNotExist, "group with id %d doesn't exist", gid)
			}
			if nErr.Number == _mysqlErrCodeDuplicateKeyEntry {
				return ecode.Errorf(pb.UserAlreadyInGroup, "user with id %d already exists in group %d", uid, gid)
			}
		}
		return err
	}
	return nil
}
