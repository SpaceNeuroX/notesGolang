package main

import (
	"awesomeProject/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func initDB() {
	var err error
	dsn := "postgres://default:ZaGKO8fjDFH7@ep-fragrant-frost-a4bbwauo-pooler.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require"
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func main() {
	initDB()
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	e.GET("/notes", handlers.GetNotes(db))
	e.POST("/notes", handlers.CreateNote(db))

	e.Logger.Fatal(e.Start(":8080"))
}
