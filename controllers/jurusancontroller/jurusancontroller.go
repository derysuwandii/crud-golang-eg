package jurusancontroller

import (
	"crud-golang/entities"
	"crud-golang/models/jurusanmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	jurusans := jurusanmodel.GetAll()

	data := map[string]any{
		"jurusans": jurusans,
	}

	temp, err := template.ParseFiles("views/jurusan/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/jurusan/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var data = make(map[string]interface{})
		var jurusan entities.Jurusan

		jurusan.Nama = r.FormValue("Nama")

		data["jurusan"] = jurusan
		if ok := jurusanmodel.DetailByNama(jurusan.Nama); ok > 0 {
			temp, _ := template.ParseFiles("views/jurusan/create.html")

			data["pesan"] = "Data sudah ada"

			temp.Execute(w, data)
		} else {

			if ok := jurusanmodel.Create(jurusan); !ok {
				temp, _ := template.ParseFiles("views/jurusan/create.html")
				temp.Execute(w, nil)
			}
			http.Redirect(w, r, "/jurusan", http.StatusSeeOther)
		}
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/jurusan/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		jurusan := jurusanmodel.Detail(id)
		data := map[string]any{
			"jurusan": jurusan,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var jurusan entities.Jurusan
		idString := r.FormValue("Id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		jurusan.Nama = r.FormValue("Nama")

		if ok := jurusanmodel.Update(id, jurusan); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/jurusan", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := jurusanmodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/jurusan", http.StatusSeeOther)
}
