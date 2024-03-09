package models

import (
	"database/sql"
	"fmt"
	"os"
)

// Mind is a struct that represents ALL nuggets for a user
type FolderNuggets struct {
	FolderID string `json:"folderID"`
	Nuggets  []Nugget
}

type Mind struct {
	FolderNuggets []FolderNuggets
	UserID        string `json:"userID"`
}

func FetchMind(userID string, db *sql.DB) (*Mind, error) {
	sqlStatement := `
		SELECT * FROM folder
		JOIN nuggets ON folder.id = nuggets.folder_id
		WHERE user_id = ?
		`
	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %v, \n", err)
		return &Mind{}, err
	}
	defer rows.Close()

	var mind Mind
	var folderNuggetMap = make(map[string][]Nugget)
	for rows.Next() {
		var folder Folder
		var nugget Nugget

		mind.FolderNuggets = make([]FolderNuggets, 0)

		if err := rows.Scan(&folder.ID, &folder.FolderName, &folder.UserID, &nugget.ID, &nugget.Key, &nugget.Value, &nugget.FolderID); err != nil {
			fmt.Println("failed scanning row: ", err)
			return &Mind{}, err
		}

		if _, ok := folderNuggetMap[folder.ID]; !ok {
			folderNuggetMap[folder.ID] = make([]Nugget, 0)
		}
		folderNuggetMap[folder.ID] = append(folderNuggetMap[folder.ID], nugget)
	}

	for folderID, folderNugget := range folderNuggetMap {
		mind.FolderNuggets = append(mind.FolderNuggets, FolderNuggets{
			FolderID: folderID,
			Nuggets:  folderNugget,
		})
	}

	mind.UserID = userID

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row iteration:", err)
		return &mind, err
	}

	return &mind, nil
}
