package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTodo(w, r)
		case http.MethodGet:
			todos(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 登録
func createTodo(w http.ResponseWriter, r *http.Request) {
	var body todo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tdl, err := db.Prepare("INSERT todolist SET importance=?,task=?,deadline=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tdl.Exec(body.Importance, body.Task, body.Deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 表示
func todos(w http.ResponseWriter, r *http.Request) {
	todos := []todo{}
	var err error
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var todotable todo
		err = rows.Scan(&todotable.ID, &todotable.Importance, &todotable.Task, &todotable.Deadline)
		todos = append(todos, todotable)
	}

	todoJson, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(todoJson)
}

// 削除
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	deleteId := r.URL.Query().Get("id")
	tdl, err := db.Prepare("DELETE FROM todolist WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tdl.Exec(deleteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
