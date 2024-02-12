package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Nugget struct {
	ID     int
	Key    string
	Value  string
	UserID int
}

func (nugget *Nugget) Save(db *sql.DB) error {
	sqlStatement := `
		INSERT INTO nuggets (id, key, value, user_id)
		VALUES ($1, $2, $3, $5)
	`
	_, err := db.Exec(sqlStatement, &nugget.ID, &nugget.Key, &nugget.Value, &nugget.UserID)

	if err != nil {
		fmt.Fprint(os.Stderr, "failed to insert value into table, \n", err)
		return err
	}
	return nil
}

func FetchAllNuggets(db *sql.DB) ([]Nugget, error) {
	sqlStatement := "SELECT * FROM nuggets"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %v, \n", err)
		return make([]Nugget, 0), err
	}
	defer rows.Close()

	var nuggets []Nugget
	for rows.Next() {
		var nugget Nugget

		if err := rows.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.UserID); err != nil {
			fmt.Println("failed scannig row: ", err)
			return nuggets, err
		}

		nuggets = append(nuggets, nugget)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row iteration:", err)
		return nuggets, err
	}

	return nuggets, nil
}
