package main

import (
	"crud-golang/config"
	"crud-golang/controllers/homecontroller"
	"crud-golang/controllers/jurusancontroller"
	"crud-golang/controllers/mahasiswacontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	config.ConnectDB2()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/mahasiswa", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/add", mahasiswacontroller.Add)
	http.HandleFunc("/mahasiswa/edit", mahasiswacontroller.Edit)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	http.HandleFunc("/jurusan", jurusancontroller.Index)
	http.HandleFunc("/jurusan/add", jurusancontroller.Add)
	http.HandleFunc("/jurusan/edit", jurusancontroller.Edit)
	http.HandleFunc("/jurusan/delete", jurusancontroller.Delete)

	log.Println("Server running on port 8080.....")
	http.ListenAndServe(":8080", nil)
}
