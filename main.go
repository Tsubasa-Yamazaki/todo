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
	Id         int    `json:"id"`
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

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/post", post)
	http.HandleFunc("/todoList", todoList)
	http.HandleFunc("/todoDelete", todoDelete)
	http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 登録
func post(w http.ResponseWriter, r *http.Request) {
	var body todo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tdl, err := db.Prepare("INSERT todolist SET importance=?,task=?,deadline=?")
	if err != nil {
		return
	}

	_, err = tdl.Exec(body.Importance, body.Task, body.Deadline)
	if err != nil {
		return
	}
}

// 表示
func todoList(w http.ResponseWriter, r *http.Request) {
	var todos []todo
	var err error
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		return
	}

	for rows.Next() {
		var todotable todo
		err = rows.Scan(&todotable.Id, &todotable.Importance, &todotable.Task, &todotable.Deadline)
		todos = append(todos, todotable)
	}

	todoJson, err := json.Marshal(todos)
	if err != nil {
		return
	}

	w.Write(todoJson)
}

// 削除
func todoDelete(w http.ResponseWriter, r *http.Request) {
	var body todo
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteId := body.Id
	tdl, err := db.Prepare("DELETE FROM todolist WHERE id = ?")
	if err != nil {
		return
	}

	_, err = tdl.Exec(deleteId)
	if err != nil {
		return
	}
}
