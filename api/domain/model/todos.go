package model

// Todoのデータ構造

type TodoFile struct {
	ID         int    `json:"id"`
	Importance string `json:"importance"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}
