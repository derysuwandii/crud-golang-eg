package mahasiswacontroller

import (
	"crud-golang/entities"
	"crud-golang/models/jurusanmodel"
	"crud-golang/models/mahasiswamodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	mahasiswas := mahasiswamodel.GetAll()

	data := map[string]any{
		"mahasiswas": mahasiswas,
	}

	temp, err := template.ParseFiles("views/mahasiswa/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/mahasiswa/create.html")
		if err != nil {
			panic(err)
		}
		jurusans := jurusanmodel.GetAll()
		data := map[string]any{
			"jurusans": jurusans,
		}
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var mahasiswa entities.Mahasiswa

		mahasiswa.NamaMahasiswa = r.FormValue("NamaMahasiswa")
		mahasiswa.Alamat = r.FormValue("Alamat")
		mahasiswa.JenisKelamin = r.FormValue("JenisKelamin")
		mahasiswa.TglMasuk = r.FormValue("TglMasuk")

		idJurusanString := r.FormValue("Jurusan")
		idJurusan, err := strconv.Atoi(idJurusanString)
		if err != nil {
			panic(err)
		}
		mahasiswa.Jurusan.Id = uint(idJurusan)

		if ok := mahasiswamodel.Create(mahasiswa); !ok {
			temp, _ := template.ParseFiles("views/mahasiswa/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/mahasiswa/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		jurusans := jurusanmodel.GetAll()
		mahasiswa := mahasiswamodel.Detail(id)
		data := map[string]any{
			"mahasiswa": mahasiswa,
			"jurusans":  jurusans,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var mahasiswa entities.Mahasiswa
		idString := r.FormValue("Id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		mahasiswa.NamaMahasiswa = r.FormValue("NamaMahasiswa")
		mahasiswa.Alamat = r.FormValue("Alamat")
		mahasiswa.JenisKelamin = r.FormValue("JenisKelamin")
		mahasiswa.TglMasuk = r.FormValue("TglMasuk")

		idJurusanString := r.FormValue("Jurusan")
		idJurusan, err := strconv.Atoi(idJurusanString)
		if err != nil {
			panic(err)
		}
		mahasiswa.Jurusan.Id = uint(idJurusan)

		if ok := mahasiswamodel.Update(id, mahasiswa); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := mahasiswamodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
}
