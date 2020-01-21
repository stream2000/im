package dao

import (
	pb "chat/app/service/account/api"
	"context"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"

	"chat/app/service/account/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	qb "github.com/didi/gendry/builder"
)

// 有点内味儿了
const (
	accountTable = "account"
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

func (d *dao) RawAccount(ctx context.Context, email string) (acc *model.Account, err error) {
	mp := map[string]interface{}{
		"email": email,
	}
	cond, vals, err := qb.BuildSelect(accountTable, mp, []string{"id", "password", "sign", "profile_pic_url", "nickname"})
	log.Infoc(ctx, "conds %s vals %s", cond, vals[0])
	row := d.db.QueryRow(ctx, cond, vals[0])
	acc = &model.Account{}
	acc.Email = email
	if err = row.Scan(&acc.UID, &acc.Password, &acc.Sign, &acc.ProfilePicUrl, &acc.NickName); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			acc = nil
		} else {
			err = errors.WithStack(err)
		}
		return
	}
	return
}

func (d *dao) AddAccount(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	// if the key is cached, the account will get a empty response
	resp = new(pb.RegisterResp)
	acc, err := d.Account(ctx, req.Email)
	if err != nil {
		return
	}
	if acc != nil {
		return nil, pb.AccountEmailRepeated
	} else {
		// delete the empty cache item
		_ = d.DeleteCacheAccount(ctx, req.Email)
	}

	var data []map[string]interface{}
	now := xtime.Time(time.Now().Unix())
	uid := uuid.NewV4().String()
	data = append(data, map[string]interface{}{
		"id":              uid,
		"email":           req.Email,
		"password":        req.Password,
		"ctime":           now,
		"mtime":           now,
		"nickname":        "user-" + uid,
		"nickname_mtime":  now,
		"profile_pic_url": "default",
	})

	cond, vals, err := qb.BuildInsert(accountTable, data)
	log.Infoc(ctx, cond, vals)
	_, err = d.db.Exec(ctx, cond, vals...)
	if err != nil {
		return
	}

	resp.Uid = uid
	return
}
