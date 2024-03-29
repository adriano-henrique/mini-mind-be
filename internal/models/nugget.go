package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Nugget struct {
	ID       string `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	FolderID string `json:"folderID"`
}

func (nugget *Nugget) Save(db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		INSERT INTO nuggets (key, value, folder_id)
		VALUES (?, ?, ?)
	`
	validationError := ValidateNuggetInsertion(nugget.Key, nugget.FolderID, db)
	if validationError != nil {
		return &Nugget{}, validationError
	}
	result, err := db.Exec(sqlStatement, &nugget.Key, &nugget.Value, &nugget.FolderID)
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

func (nugget *Nugget) UpdateNugget(db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		UPDATE nuggets
		SET key = ?, value = ?
		WHERE id = ?
	`

	_, err := db.Exec(sqlStatement, &nugget.Key, &nugget.Value, &nugget.ID)
	if err != nil {
		fmt.Println("failed to update table nuggets, \n", err)
		return &Nugget{}, err
	}

	return FetchNugget(nugget.ID, db)
}

func FetchNugget(nuggetID string, db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		SELECT * FROM nuggets
		WHERE id = ?
	`

	var nugget Nugget
	row := db.QueryRow(sqlStatement, nuggetID)
	err := row.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID)
	if err != nil {
		fmt.Println("Failed to get query with id: ", nuggetID)
		return &Nugget{}, err
	}

	return &nugget, nil
}

func FetchAllNuggetsByKey(nuggetKey string, db *sql.DB) ([]Nugget, error) {
	sqlStatement := `
		SELECT * FROM nuggets
		WHERE key = ?
	`
	rows, err := db.Query(sqlStatement, nuggetKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %v, \n", err)
		return make([]Nugget, 0), err
	}
	defer rows.Close()

	var nuggets []Nugget
	for rows.Next() {
		var nugget Nugget

		if err := rows.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID); err != nil {
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

func FetchNuggetByKeyAndFolder(nuggetKey string, nuggetFolderID string, db *sql.DB) (*Nugget, error) {
	sqlStatement := `
		SELECT * FROM nuggets
		WHERE key = ?
		AND folder_id = ?
	`
	var nugget Nugget
	row := db.QueryRow(sqlStatement, nuggetKey, nuggetFolderID)
	err := row.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID)
	if err != nil {
		fmt.Println("Failed to get nuggest with key ? and folder ?", nuggetKey, nuggetFolderID)
		return &Nugget{}, err
	}

	return &nugget, nil
}

func FetchNuggetsByFolderID(folderID string, db *sql.DB) ([]Nugget, error) {
	sqlStatement := "SELECT * FROM nuggets WHERE folder_id = ?"
	rows, err := db.Query(sqlStatement, folderID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %v, \n", err)
		return make([]Nugget, 0), err
	}
	defer rows.Close()

	var nuggets []Nugget
	for rows.Next() {
		var nugget Nugget

		if err := rows.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID); err != nil {
			fmt.Println("failed scanning row: ", err)
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

		if err := rows.Scan(&nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID); err != nil {
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

func DeleteNugget(nuggetID string, db *sql.DB) error {
	sqlStatement := "DELETE FROM nuggets WHERE id = ?"
	_, err := db.Exec(sqlStatement, nuggetID)
	if err != nil {
		fmt.Println("failed to delete nugget with id: ", nuggetID)
		return err
	}
	return nil
}
