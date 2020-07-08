package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *sql.DB

type todo struct {
	ID         int    `json:"id"`
	Importance string `json:"importance"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:0111@/todo283")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Open the 283todo.")

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/todos", createTodo)
	e.GET("/todos", todos)
	e.DELETE("/todos", deleteTodo)

	e.Logger.Fatal(e.Start(":8080"))

}

// 表示
func todos(c echo.Context) error {
	todos := []todo{}
	var err error
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		var todotable todo
		err = rows.Scan(&todotable.ID, &todotable.Importance, &todotable.Task, &todotable.Deadline)
		todos = append(todos, todotable)
	}
	return c.JSON(http.StatusOK, todos)
}

func createTodo(c echo.Context) error {
	var body todo
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tdl, err := db.Prepare("INSERT todolist SET importance=?,task=?,deadline=?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, err = tdl.Exec(body.Importance, body.Task, body.Deadline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "OK")
}

func deleteTodo(c echo.Context) error {
	deleteId := c.QueryParam("id")
	tdl, err := db.Prepare("DELETE FROM todolist WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, err = tdl.Exec(deleteId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
