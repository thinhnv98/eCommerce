package repository

import (
	"context"
	"database/sql"

	"eCommerce/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserRepo interface {
	Create(ctx context.Context, u models.User) (int, error)
	Login(ctx context.Context, userName string, passWordHashed string) (bool, int, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (_self UserRepo) Create(ctx context.Context, u models.User) (int, error) {
	err := u.Insert(ctx, _self.Db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (_self UserRepo) Login(ctx context.Context, userName string, passWordHashed string) (bool, int, error) {
	u, err := models.Users(
		models.UserWhere.UserName.EQ(userName),
		models.UserWhere.PasswordHash.EQ(passWordHashed),
	).One(ctx, _self.Db)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, 0, err
		}
		return false, 0, nil
	}

	return true, u.ID, nil
}
