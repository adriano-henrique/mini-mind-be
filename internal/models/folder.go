package models

import (
	"database/sql"
	"fmt"
	"os"
)

type Folder struct {
	ID         string `json:"id"`
	FolderName string `json:"folderName"`
	UserID     string `json:"userID"`
}

func (folder *Folder) Save(db *sql.DB) (*Folder, error) {
	sqlStatement := `
		INSERT INTO folder (folder_name, user_id)
		VALUES (?, ?)
	`

	result, err := db.Exec(sqlStatement, &folder.FolderName, &folder.UserID)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to insert value into table, \n", err)
		return &Folder{}, err
	}

	folderID, err := result.LastInsertId()
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to get recently inserted data, \n", err)
		return &Folder{}, err
	}
	return FetchFolder(fmt.Sprint(folderID), db)
}

func (folder *Folder) UpdateFolder(db *sql.DB) (*Folder, error) {
	sqlStatement := `
		UPDATE folder
		SET folder_name = ?
		WHERE id = ?
	`

	_, err := db.Exec(sqlStatement, &folder.FolderName, &folder.ID)
	if err != nil {
		fmt.Println("failed to update table folder, \n", err)
		return &Folder{}, err
	}

	return FetchFolder(folder.ID, db)
}

func FetchFolder(folderID string, db *sql.DB) (*Folder, error) {
	sqlStatement := `
		SELECT * FROM folder
		WHERE id = ?
	`

	var folder Folder
	row := db.QueryRow(sqlStatement, folderID)
	err := row.Scan(&folder.ID, &folder.FolderName, &folder.UserID)
	if err != nil {
		fmt.Println("Failed to get query with id: ", folderID)
		return &Folder{}, err
	}

	return &folder, nil
}

func FetchAllFolders(db *sql.DB) ([]Folder, error) {
	sqlStatement := "SELECT * FROM folder"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %v, \n", err)
		return make([]Folder, 0), err
	}
	defer rows.Close()

	var folders []Folder
	for rows.Next() {
		var folder Folder

		if err := rows.Scan(&folder.ID, &folder.FolderName, &folder.UserID); err != nil {
			fmt.Println("failed scanning row: ", err)
			return folders, err
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row iteration: ", err)
		return folders, err
	}

	return folders, nil
}

func DeleteFolder(folderID string, db *sql.DB) error {
	sqlStatement := "DELETE FROM folder WHERE id = ?"
	_, err := db.Exec(sqlStatement, folderID)
	if err != nil {
		fmt.Println("failed to delete nugget with id: ", folderID)
		return err
	}
	return nil
}
