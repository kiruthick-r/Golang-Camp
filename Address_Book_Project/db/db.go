package db

import (
	"fmt"
	"log"

	"Address_Book_Project/proto"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	//dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", viper.Get("database.user"), viper.Get("database.password"), viper.Get("server.port"), viper.Get("database.dbname"))
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", viper.Get("server.host"), viper.Get("database.user"), viper.Get("database.password"), viper.Get("database.dbname"), viper.Get("server.port"))

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&proto.User{})

	return db
}
