package config

import (
	"comments_api/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("comments.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&models.Comment{})
    if err != nil {
        return nil, err
    }

		var count int64
		err = db.Model(&models.Comment{}).Count(&count).Error
		if err != nil {
			log.Fatalf("Failed to count records: %v", err)
		}

		if count == 0 {
			if err := loadInitialData(db); err != nil {
					return nil, err
			}
		} else {
				log.Println("Database is not empty, skipping initial data insertion")
		}


    return db, nil
}

func loadInitialData(db *gorm.DB) error {
	// Read JSON file
	comments, err := readJSONFile("data/comments.json")
	if err != nil {
			return fmt.Errorf("failed to read initial data: %v", err)
	}

	// Insert data if there are records
	if len(comments) > 0 {
			result := db.Create(&comments)
			if result.Error != nil {
					return fmt.Errorf("failed to insert initial data: %v", result.Error)
			}
			log.Printf("Successfully inserted %d records to database", len(comments))
	}

	return nil
}

func readJSONFile(filePath string) ([]models.Comment, error) {
	file, err := os.Open(filePath)
	if err != nil {
			return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var comments []models.Comment
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&comments); err != nil {
			return nil, fmt.Errorf("could not decode JSON: %v", err)
	}

	return comments, nil
}
