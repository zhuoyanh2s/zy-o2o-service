package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func init() {
	var _ error
	engine, _ = xorm.NewEngine("postgres", "postgres://ceshi:ceshi@39.96.30.230:5432/zy-o2o-server?sslmode=disable")
	err := engine.Sync2(new(User))
	fmt.Println(err)
}
