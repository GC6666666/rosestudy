package repo

import (
	"context"
	"rose/internal/conf"
	"rose/internal/repo/db"
)

type Repo struct {
	db *db.DB
}

func NewRepo(conf *conf.Conf) *Repo {
	return &Repo{
		db: db.NewDB(conf.DB),
	}
}

func (r *Repo) GetUserByID(ctx context.Context, userID int64) (*db.UserModel, error) {
	return r.db.GetUserByID(ctx, userID)
}

func (r *Repo) CreatUser(ctx context.Context, userName string) (int64, error) {
	return r.db.InsertUser(ctx, userName)
}

func (r *Repo) GetUserByIDs(ctx context.Context, userIDs []string) ([]*db.UserModel, error) {
	return r.db.GetUserByIDs(ctx, userIDs)

}
