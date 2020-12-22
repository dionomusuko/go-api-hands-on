package database

import (
	"log"

	"github.com/dionomusuko/go-api-hands-on/model"
)

// TODO DBからの取得するSQLを書いていく
func List() (*model.Tasks, error) {
	var tasks model.Tasks
	rows, err := Conn.Query("SELECT Id, Title, Description FROM Tasks")
	if err != nil {
		log.Printf("sql error: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			continue
		}
		param := model.Task{
			Id:          task.Id,
			Title:       task.Title,
			Description: task.Description,
		}
		tasks = append(tasks, param)
	}
	return &tasks, nil
}
