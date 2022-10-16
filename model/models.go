package model

import (
	"fmt"

	"weeee9/wire-example/config"

	_ "github.com/lib/pq"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/contexts"
)

var tables = []interface{}{
	new(User),
}

func NewEngine(cfg config.Config, hook contexts.Hook) (*xorm.Engine, error) {
	connStr := cfg.Database.ConnStr()

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "db", "db", "localhost", "9999", "db")

	x, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := x.Ping(); err != nil {
		return nil, err
	}

	x.AddHook(hook)
	x.SetMapper(core.GonicMapper{})
	x.Sync2(tables...)

	return x, nil
}
