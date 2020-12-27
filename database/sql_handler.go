package database

import (
	"context"
	"log"

	"github.com/dionomusuko/go-api-hands-on/model"
)

func List(ctx context.Context) (*model.Tasks, error) {
	var tasks model.Tasks
	rows, err := DB.QueryContext(ctx, "SELECT Id, Title, Description FROM Tasks")
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

func FindById(ctx context.Context, id int) (*model.Task, error) {
	row := DB.QueryRowContext(ctx, "SELECT Id, Title, Description FROM Tasks WHERE Id = ?", id)
	var task model.Task
	if err := row.Scan(&task.Id, &task.Title, &task.Description); err != nil {
		log.Printf("sql error: %v", err)
		return nil, err
	}
	return &task, nil
}

func Store(ctx context.Context, task model.Task) error {
	_, err := DB.ExecContext(ctx, "INSERT INTO Tasks (Title, Description) Values(?,?)", task.Title, task.Description)
	if err != nil {
		log.Printf("insert error: %v", err)
		return err
	}
	return nil
}

func Update(ctx context.Context, task model.Task) error {
	_, err := DB.ExecContext(ctx, "UPDATE Tasks SET Title = ?, Description = ? WHERE id = ?", task.Title, task.Description, task.Id)
	if err != nil {
		log.Printf("update error: %v", err)
		return err
	}
	return nil
}

func Delete(ctx context.Context, id int) error {
	_, err := DB.ExecContext(ctx, "DELETE FROM Tasks WHERE id = ?", id)
	if err != nil {
		log.Printf("delete error: %v", err)
		return err
	}
	return nil
}
