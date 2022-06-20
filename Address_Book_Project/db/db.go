package db

import (
	"log"

	"Address_Book_Project/proto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	//dbURL := "postgres://root:password@localhost:5432/crud"
	//dbURL := "postgres://root:password@database/crud?sslmode=disable"
	//dbURL := "postgres://root:password@database:5432/crud"
	dbURL := "host=database user=root password=password dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&proto.User{})

	return db
}
