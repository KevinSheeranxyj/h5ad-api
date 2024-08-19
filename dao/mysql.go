package dao

import (
	"database/sql"
	"fmt"
	"time"
	"uy0/h5ad/config"

	_ "github.com/go-sql-driver/mysql"
)

// Link 连接
var db *sql.DB
var err error

// Db 数据库
func Db() *sql.DB {
	if db == nil {
		jdbcS := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
		jdbc := fmt.Sprintf(jdbcS, config.Config.Mysql.Username, config.Config.Mysql.Password,
			config.Config.Mysql.Hostname, config.Config.Mysql.HostPort,
			config.Config.Mysql.Database)
		fmt.Println("mysql:", config.Config.Mysql.Hostname)
		fmt.Println("-------------------------------------")
		db, err = sql.Open("mysql", jdbc)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3) // 连接最长存活期，超过这个时间连接将不再被复用
		db.SetMaxOpenConns(10)                 // 数据库最大连接数
		db.SetMaxIdleConns(10)                 // 最大空闲连接数
		return db
	}
	return db
}
