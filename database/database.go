package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Prepare(name string, user string, pass string, tasks int) {
	connect(name, user, pass)

	createTable(name)

	populateTasks(tasks)
}

func connect(name string, user string, pass string) {
	conf := mysql.NewConfig()
	conf.DBName = name
	conf.User = user
	conf.Passwd = pass

	database, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	db = database
}

func createTable(dbName string) {
	rows, err := db.Query(fmt.Sprintf(`
		SELECT * FROM information_schema.TABLES
		WHERE table_schema = '%s' AND table_name = 'tasks'
	`, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return
	}

	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE tasks (
			id BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
			status ENUM('%s', '%s', '%s') NOT NULL DEFAULT '%s'
		)
	`, New, Processing, Processed, New))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tasks table created.")
}

func populateTasks(tasks int) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return
	}

	statuses := fmt.Sprintf("('%s')", New)
	for range tasks - 1 {
		statuses += fmt.Sprintf(", ('%s')", New)
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO tasks (status) VALUES %s", statuses))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tasks table populated.")
}
