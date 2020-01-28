package dao

import (
	pb "chat/app/service/account/api"
	"context"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	uuid "github.com/satori/go.uuid"
	"time"

	"chat/app/service/account/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	qb "github.com/didi/gendry/builder"
)

// 有点内味儿了
const (
	accountTable       = "account"
	_getAccountByIdSql = "select email,password,sign,profile_pic_url,nickname from account where id=?"
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

func (d *dao) RawAccount(ctx context.Context, uid int64) (acc *model.Account, err error) {
	row := d.db.QueryRow(ctx, _getAccountByIdSql, uid)
	acc = &model.Account{}
	acc.UID = uid
	if err = row.Scan(&acc.Email, &acc.Password, &acc.Sign, &acc.ProfilePicUrl, &acc.NickName); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			acc = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}

func (d *dao) AddAccount(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	tx, err := d.db.Begin(ctx)
	defer func() {
		if nil == err {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	acc, err := d.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return
	}
	if acc != nil {
		return nil, ecode.Errorf(pb.AccountEmailRepeated, "email %s already registered", req.Email)
	}

	var data []map[string]interface{}
	now := xtime.Time(time.Now().Unix())
	u := uuid.NewV4().String()
	data = append(data, map[string]interface{}{
		"email":           req.Email,
		"password":        req.Password,
		"ctime":           now,
		"mtime":           now,
		"nickname":        "user-" + u,
		"nickname_mtime":  now,
		"profile_pic_url": "default",
	})
	cond, vals, err := qb.BuildInsert(accountTable, data)
	log.Infoc(ctx, cond, vals)
	res, err := d.db.Exec(ctx, cond, vals...)
	if err != nil {
		log.Error("account.service.dao.addAccount.Exec(%v) Err(%+v)", cond, err)
		return
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		log.Error("account.service.dao error getting auto increment id Email(%s) Err(%v)", req.Email, err)
		return nil, err
	}

	resp = new(pb.RegisterResp)
	resp.Uid = lastInsertedId
	// delete the empty cache item
	_ = d.DeleteCacheAccount(ctx, resp.Uid)
	return
}

func (d *dao) GetAccountByEmail(ctx context.Context, email string) (acc *model.Account, err error) {
	where := map[string]interface{}{
		"email": email,
	}
	cond, vals, err := qb.BuildSelect(accountTable, where, []string{"id", "password", "sign", "profile_pic_url", "nickname"})
	if err != nil {
		panic(err)
	}

	row := d.db.QueryRow(ctx, cond, vals[0])
	acc = &model.Account{}
	acc.Email = email
	if err = row.Scan(&acc.UID, &acc.Password, &acc.Sign, &acc.ProfilePicUrl, &acc.NickName); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			acc = nil
		} else {
			log.Error("row.Scan error(%v)", err)
		}
	}
	return
}
