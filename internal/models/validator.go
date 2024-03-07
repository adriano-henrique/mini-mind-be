package models

import (
	"database/sql"
	"errors"
)

func ValidateNuggetInsertion(nuggetKey string, nuggetFolderID string, db *sql.DB) error {
	nugget, err := FetchNuggetByKeyAndFolder(nuggetKey, nuggetFolderID, db)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("unable to fetch nugget by key")
	}
	if nugget.Key == nuggetKey {
		return errors.New("nugget key already exists in this folder")
	}
	return nil
}
