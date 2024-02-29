package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Test go")

	// Update the connection string to match your MySQL database credentials
	dsn := "root:@tcp(127.0.0.1:3306)/encanto?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get generic database object:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Migrate the schema
	db.AutoMigrate(&EncantoApk{})

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Execute Hello world")
		fmt.Fprint(w, "Hello World!")
	}).Methods("GET")

	ApkHandler(r, db)

	log.Fatal(http.ListenAndServe(":8080", r))

}
