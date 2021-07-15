package connection

import (
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	ConnType   = "postgres"
	ConnString = []string{"postgres://hammurabi:hammurabi@localhost:5432/test?sslmode=disable"}
)

func Connect() (*xorm.EngineGroup, error) {
	engine, err := xorm.NewEngineGroup(ConnType, ConnString)
	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})
	return engine, err
}
