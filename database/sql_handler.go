package database

import (
	"log"

	"github.com/dionomusuko/go-api-hands-on/model"
)

func List() (*model.Tasks, error) {
	var tasks model.Tasks
	rows, err := DB.Query("SELECT Id, Title, Description FROM Tasks")
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

func FindById(id int) (*model.Task, error) {
	row := DB.QueryRow("SELECT Id, Title, Description FROM Tasks WHERE Id = ?", id)
	var task model.Task
	if err := row.Scan(&task.Id, &task.Title, &task.Description); err != nil {
		log.Printf("sql error: %v", err)
		return nil, err
	}
	return &task, nil
}

func Store(task model.Task) error {
	_, err := DB.Exec("INSERT INTO Tasks (Title, Description) Values(?,?)", task.Title, task.Description)
	if err != nil {
		log.Printf("insert error: %v", err)
		return err
	}
	return nil
}

func Update(task model.Task) error {
	_, err := DB.Exec("UPDATE Tasks SET Title = ?, Description = ? WHERE id = ?", task.Title, task.Description, task.Id)
	if err != nil {
		log.Printf("update error: %v", err)
		return err
	}
	return nil
}

func Delete(id int) error {
	_, err := DB.Exec("DELETE FROM Tasks WHERE id = ?", id)
	if err != nil {
		log.Printf("delete error: %v", err)
		return err
	}
	return nil
}
