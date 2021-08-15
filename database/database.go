package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func InitDatabase() {

	// Load environment variables
	if fileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	//Load Database
	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("DB_PORT")
	DBUSER := os.Getenv("DB_USER")
	DBNAME := os.Getenv("DB_NAME")
	DBPASS := os.Getenv("DB_PASS")

	connection_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DBHOST, DBPORT, DBUSER, DBNAME, DBPASS)

	retries := 5
	database, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	// fmt.Println(err)
	for err != nil {
		// log.Fatal("error: ", err)
		count := 0
		count++
		if retries > 1 {
			fmt.Println("Retrying connect to database again ...", count)
			retries--
			time.Sleep(10 * time.Second)
			database, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
			continue
		}
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxIdleTime(time.Hour)

	RunMigration(db)

}

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(
		User{},
	)
}

func GetDatabase() *gorm.DB {
	return db
}
