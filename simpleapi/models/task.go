package models

import (
	"database/sql"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price 	float64    `json:"price"`
	Description string `json:"description"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// Exit jika terjadi error
	if err != nil {
		panic(err)
	}
	// clean rows
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name, &task.Price, &task.Description)
		// Exit jika error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string, price float64, description string) (int64, error) {
	sql := "INSERT INTO tasks(name, price, description) VALUES(?,?,?);"

	// membuat prepared SQL statement
	stmt, err := db.Prepare(sql)

	// Exit jika error
	if err != nil {
		panic(err)
	}

	// memastikan statement ditutup setelah selesai
	defer stmt.Close()

	result, err2 := stmt.Exec(name, price, description)

	// Exit jika error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditTask(db *sql.DB, taskId int, name string, price float64, description string) (int64, error) {
	sql := "UPDATE tasks set name = ?, price = ?, description = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(name, price, description, taskId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// buat prepare statement
	stmt, err := db.Prepare(sql)
	// Exit jika error
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	// Exit jika error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
