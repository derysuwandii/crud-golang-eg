package jurusanmodel

import (
	"crud-golang/config"
	"crud-golang/entities"
	"log"
)

func GetAll() []entities.Jurusan {
	rows, err := config.DB.Query("Select * from tbl_jurusan order by nama asc")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var jurusans []entities.Jurusan

	for rows.Next() {
		var jurusan entities.Jurusan
		if err := rows.Scan(&jurusan.Id, &jurusan.Nama); err != nil {
			panic(err)
		}

		jurusans = append(jurusans, jurusan)
	}

	return jurusans
}

func Create(jurusan entities.Jurusan) bool {
	result, err := config.DB.Exec(`
		insert into tbl_jurusan (nama)
		VALUE (?)`, jurusan.Nama)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Jurusan {
	row := config.DB.QueryRow("Select * from tbl_jurusan where id=?", id)

	var jurusan entities.Jurusan
	if err := row.Scan(&jurusan.Id, &jurusan.Nama); err != nil {
		panic(err)
	}

	return jurusan
}

func Update(id int, jurusan entities.Jurusan) bool {
	query, err := config.DB.Exec(`
		update tbl_jurusan set nama=? where id=?`,
		jurusan.Nama, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`delete from tbl_jurusan where id=?`, id)

	return err
}

func DetailByNama(nama string) int {
	var count int
	err := config.DB.QueryRow("Select count(*) from tbl_jurusan where nama=?", nama).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
