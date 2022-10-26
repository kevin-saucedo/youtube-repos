package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/user0608/ytgorm/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GetConn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=maira002 dbname=gorm_curso_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, err
}

func main() {
	conn, err := GetConn()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("OK!")
	e := echo.New()
	e.HideBanner = true

	handlers.Start(e, conn)

	if err := e.Start("localhost:90"); err != nil {
		log.Fatalln(err)
	}
}
