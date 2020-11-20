package usersdbgorm

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysql_username = "mysql_users_username"
	mysql_password = "mysql_users_password"
	mysql_host     = "mysql_users_host"
	mysql_schema   = "mysql_users_schema"
)

var (
	GormDB   *gorm.DB
	username = os.Getenv(mysql_username)
	password = os.Getenv(mysql_password)
	host     = os.Getenv(mysql_host)
	schema   = os.Getenv(mysql_schema)
)

func init() {
	// Open the database
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	GormDB, err = gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//GormDB.AutoMigrate(&usersdomain.User)
}
