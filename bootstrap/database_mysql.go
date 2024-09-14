package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDatabase(env *Env) gorm.DB {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	dbHost := env.MySQLDBHost
	dbPort := env.MySQLDBPort
	dbUser := env.MySQLDBUser
	dbPass := env.MySQLDBPass
	dbName := env.MySQLDBName

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return *db
}

func CloseMySQLDBConnection(client gorm.DB) {
	// 不需要关闭
	log.Println("call CloseMySQLDBConnection")
}
