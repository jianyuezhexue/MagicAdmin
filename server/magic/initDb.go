package magic

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库
func InitDb() {
	config := Config.Mysql
	// db, err := sql.Open("mysql", config.UserName+":"+config.Password+"@/mysql")
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Path + ":3306)/" + "@/mysql"
	Print(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	check := db.Ping()
	if check != nil {
		panic(check)
	}

	_, err = db.Exec("CREATE DATABASE dbMagicAdmin;")
	if err != nil {
		panic(err)
	}
}
