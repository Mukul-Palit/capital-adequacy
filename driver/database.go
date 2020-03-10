package driver

import (
	"fmt"
	"log"
	"os"
	"runtime"

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

//WriteLogFile : error logging
func WriteLogFile(err error) {
	f := OpenLogFile()
	pc, fn, line, _ := runtime.Caller(1)
	fmt.Println(err)
	fmt.Println(pc, fn, line)
	log.SetOutput(f)
	log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	defer f.Close()
}

// OpenLogFile : this function will open log file and return the file writer
func OpenLogFile() (f *os.File) {
	f, err := os.OpenFile("logs/output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return
	}
	return f
}
