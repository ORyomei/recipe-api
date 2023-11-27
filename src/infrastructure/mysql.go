package infrastructure

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"recipe_api/src/util"
)

// MySql is db
type MySql struct {
	db *gorm.DB
}

// Begin begin new transaction
func (mySql *MySql) Begin() (*gorm.DB, func() error, func(), error) {
	tx := mySql.db.Begin()
	if tx.Error != nil {
		return nil, nil, nil, tx.Error
	}
	commit := func() error {
		return tx.Commit().Error
	}
	rollback := func() {
		tx.Rollback()
	}
	return tx, commit, rollback, tx.Error
}

// GetQuerier get queyer with no transaction
func (mySql *MySql) GetQuerier() *gorm.DB {
	return mySql.db
}

// WithContext add context to db
func (mySql *MySql) WithContext(ctx context.Context) *MySql {
	return &MySql{
		db: mySql.db.WithContext(ctx),
	}
}

// Ping check connection to db
func (mySql *MySql) Ping() error {
	var result int64
	if err := mySql.db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return err
	}
	return nil
}

// OpenDB open db
func OpenMySql(config *util.MySqlConfig) (*MySql, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/recipedb?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	)
	fmt.Println(dsn)
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			SkipDefaultTransaction: false,
			Logger:                 logger.Default.LogMode(logger.Silent),
			PrepareStmt:            true,
			ConnPool:               nil,
		},
	)
	if err != nil {
		log.Panicf("Opendb failed: %s", err)
		return nil, err
	}
	/* 	if err := gormDB.Callback().Query().Before("gorm:query").Register("auto_loader", gateways.AutoLoaderCallBack); err != nil {
		log.Panicf("failed to callback: %s", err)
	} */
	return &MySql{
		db: gormDB,
	}, nil
}
