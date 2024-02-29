package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type EncantoApk struct {
	gorm.Model
	FileName  string `json:"fileName"`
	ApkBase64 string `gorm:"-" json:"apkBase63"`
	Version   string `json:"version"`
}

func ApkHandler(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {

		apkFile := EncantoApk{}
		dbResult := db.First(&apkFile)
		if dbResult.Error != nil {
			// handle error
		}
		fmt.Fprint(w, apkFile.Version)

	}).Methods("GET")

	r.HandleFunc("/api/download", func(w http.ResponseWriter, r *http.Request) {
		apkInfo := EncantoApk{}

		dbResult := db.First(&apkInfo, "id = ?", 1)
		if dbResult.Error != nil {
			w.WriteHeader(http.StatusNoContent)
			fmt.Println(dbResult.Error)
			return
		}
		fmt.Println("load file...")
		tempFile, err := os.ReadFile("/files/latest.apk")
		if err != nil {
			fmt.Println(err)
		}
		base64File := base64.StdEncoding.EncodeToString(tempFile)
		apkInfo.ApkBase64 = base64File

		json.NewEncoder(w).Encode(&apkInfo)
	}).Methods("GET")

	r.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		newApk := EncantoApk{}

		err := json.NewDecoder(r.Body).Decode(&newApk)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error devode Json: %v", err)
			return
		}

		data, errBase := base64.StdEncoding.DecodeString(newApk.ApkBase64)
		if errBase != nil {
			log.Fatal("error", errBase)
		}

		os.WriteFile("/files/latest.apk", data, os.ModePerm)
		db.Save(&newApk)
		json.NewEncoder(w).Encode(newApk.Version)

	}).Methods("POST")
}
