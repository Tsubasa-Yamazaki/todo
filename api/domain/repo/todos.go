package repo

import (
	"github.com/Tsubasa-Yamazaki/todo/api/domain/model"
	"github.com/labstack/echo"
)

// データ処理のインターフェイスの定義

type TodoFile interface {
	Todos() ([]model.TodoFile, error)
	CreateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}
