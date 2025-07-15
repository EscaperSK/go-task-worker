package database

import (
	"fmt"
	"log"
)

func GetTaskIds(number int) ([]int, error) {
	checkConnection()

	rows, err := db.Query(fmt.Sprintf(`
		SELECT id FROM tasks
		WHERE status = '%s'
		LIMIT %d
	`, New, number))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			continue
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func UpdateTask(id int, status status) {
	checkConnection()

	_, err := db.Exec(fmt.Sprintf(`
		UPDATE tasks
		SET status = '%s'
		WHERE id = %d
	`, status, id))
	if err != nil {
		log.Printf("Failed to update task #%d status. Error: %v\n", id, err)
	}
}

func checkConnection() {
	if db == nil {
		log.Fatal("Not connected to database.")
	}
}
