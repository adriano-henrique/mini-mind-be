package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Nugget struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	UserID string `json:"userID"`
}

func (nugget *Nugget) Save(db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		INSERT INTO nuggets (key, value, user_id)
		VALUES (?, ?, ?)
	`
	result, err := db.Exec(sqlStatement, &nugget.Key, &nugget.Value, &nugget.UserID)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to insert value into table, \n", err)
		return &Nugget{}, err
	}

	nuggetID, err := result.LastInsertId()
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to get recently inserted data, \n", err)
		return &Nugget{}, err
	}
	return FetchNugget(fmt.Sprint(nuggetID), db)
}

func FetchNugget(nuggetID string, db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		SELECT * FROM nuggets
		WHERE id = ?
	`

	var nugget Nugget
	row := db.QueryRow(sqlStatement, nuggetID)
	err := row.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.UserID)
	if err != nil {
		return &Nugget{}, err
	}

	return &nugget, nil
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
