package driver

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //blank import
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//DbConn : function for database connection
func DbConn() (db *gorm.DB) {
	godotenv.Load()
	dbDriver := os.Getenv("DATABASE_DRIVER")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dbIP := os.Getenv("DATABASE_IP")
	dbPort := os.Getenv("DATABASE_PORT")
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
