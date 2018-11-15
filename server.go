package main

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/web/module/handler"
	"github.com/web/module/page"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Echo instance
	e := echo.New()

	db := initDB("storage.db")
	migrate(db)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", page.Index)
	e.GET("/tasks", handler.GetTasks(db))
	e.POST("/tasks", handler.PutTask(db))
	e.PUT("/tasks", handler.EditTask(db))
	e.DELETE("/tasks/:id", handler.DeleteTask(db))
	//log.Printf("", var)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

//func hello(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
//}
