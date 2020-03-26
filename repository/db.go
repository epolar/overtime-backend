package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

var db *DBCli
var dbOnce sync.Once

func DB() *DBCli {
	dbOnce.Do(func() {
		dbParams := viper.GetStringMapString("db")
		db = NewDBClient(dbParams)
		migration(db)
	})
	return db
}

type DBCli struct {
	*gorm.DB
}

func NewDBClient(params map[string]string) *DBCli {
	user := params["user"]
	password := params["password"]
	host := params["host"]
	port := params["port"]
	dbname := params["db"]

	u := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		user,
		password,
		host,
		port,
		dbname,
	)

	cli, err := gorm.Open("mysql", u)
	if err != nil {
		panic(err)
	}

	cli.LogMode(viper.GetBool("enable-db-log"))

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "td_" + defaultTableName
	}

	return &DBCli{cli}
}
