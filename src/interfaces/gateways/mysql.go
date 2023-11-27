package gateways

import (
	"context"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// MySql is db
type MySql interface {
	Begin() (*gorm.DB, func() error, func(), error)
	GetQuerier() *gorm.DB
}

// ContextlessDB is db with context
type ContextlessDB interface {
	WithContext(context.Context) MySql
}

type autoLoader interface {
	AutoLoad() *autoLoadSetting
}

type autoLoadSetting struct {
	Preloads    []string
	ArgPreloads []autoLoadSettingArgQuery
	Joins       []string
	ArgJoins    []autoLoadSettingArgQuery
}

type autoLoadSettingArgQuery struct {
	Query string
	Args  []interface{}
}

func AutoLoaderCallBack(db *gorm.DB) {
	autoLoader, ok := db.Statement.Dest.(autoLoader)
	if !ok {
		return
	}
	if db.Statement.Preloads == nil {
		db.Statement.Preloads = map[string][]interface{}{}
	}
	s := autoLoader.AutoLoad()
	for _, p := range s.Preloads {
		db.Statement.Preloads[p] = []interface{}{}
	}
	for _, p := range s.ArgPreloads {
		db.Statement.Preloads[p.Query] = p.Args
	}
	for _, j := range s.Joins {
		db.Statement.Joins = db.Joins(j).Statement.Joins
	}
	for _, j := range s.ArgJoins {
		db.Statement.Joins = db.Joins(j.Query, j.Args...).Statement.Joins
	}
}

func getTableName(db *gorm.DB, table interface{}) (string, error) {
	s, err := schema.Parse(table, &sync.Map{}, db.NamingStrategy)
	if err != nil {
		return "", err
	}
	return s.Table, nil
}
