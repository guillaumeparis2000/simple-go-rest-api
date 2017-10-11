package models

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB

var currentId int

func InitDb() {
	var err error
	db, err = sql.Open("mysql", "root:qwe123@/todos?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func GetAllTodos() Todos {
	var err error
	rows, err := db.Query("SELECT id, name, completed, due FROM todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos Todos
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)

		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}


	return todos
}

func RepoFindTodo(id int64) Todo {
	var err error
	var todo Todo

	// Execute the query
	err = db.QueryRow("SELECT id, name, completed, due FROM todo where id = ?", id).Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)

	if err != nil {
		if err == sql.ErrNoRows {
			// return empty Todo if not found
			return Todo{}
		} else {
			log.Fatal(err)
		}
	}

	if todo.Id != 0 {
		return todo
	}

	return Todo{}

}

func RepoCreateTodo(t Todo) Todo {
	var err error
	insert, err := db.Exec("INSERT INTO todo(name, completed, due) VALUES(?, ?, ?)", t.Name, t.Completed, t.Due)

	if err != nil {
		log.Fatal(err.Error())
	}

	id, _ := insert.LastInsertId()
	t.Id = id

	return t
}

func RepoDestroyTodo(id int64) {
	var err error
	_, err = db.Exec("DELETE FROM todo WHERE id = ?", id)

	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func RepoUpdateTodo(todo Todo) Todo {
	var err error
	_, err = db.Exec("UPDATE todo SET name = ?, completed = ?, due = ? where id = ?", todo.Name, todo.Completed, todo.Due, todo.Id)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	return todo
}