package model

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

type User struct {
	ID       int64  `xorm:"pk autoincr"`
	Username string `xorm:"UNIQUE"`
	Name     string
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
}

func NewUserRepository(engine *xorm.Engine) UserRepository {
	return userRepositoryDB{
		e: engine,
	}
}

type userRepositoryDB struct {
	e *xorm.Engine
}

func (db userRepositoryDB) GetAllUsers(ctx context.Context) ([]*User, error) {
	sess := db.e.Context(ctx)

	users := make([]*User, 0)

	return users, sess.Find(&users)
}

func (db userRepositoryDB) GetUserByID(ctx context.Context, id int64) (*User, error) {
	sess := db.e.Context(ctx)

	user := &User{
		ID: id,
	}

	has, err := sess.Get(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return nil, err
	}

	if !has {
		log.Error().Msgf("user not found: [%d]", id)
		return nil, fmt.Errorf("user not found: [%d]", id)
	}

	return user, nil
}
