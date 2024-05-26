package mahasiswamodel

import (
	"crud-golang/config"
	"crud-golang/entities"
)

func GetAll() []entities.Mahasiswa {
	rows, err := config.DB.Query(`
	Select mhs.id, mhs.nama_mahasiswa, mhs.alamat, 
	mhs.jenis_kelamin, mhs.tgl_masuk, jur.nama as nama_jurusan
	from tbl_mahasiswa mhs JOIN tbl_jurusan jur ON mhs.jurusan_id=jur.id
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var mahasiswas []entities.Mahasiswa

	for rows.Next() {
		var mahasiswa entities.Mahasiswa
		if err := rows.Scan(&mahasiswa.Id, &mahasiswa.NamaMahasiswa, &mahasiswa.Alamat,
			&mahasiswa.JenisKelamin, &mahasiswa.TglMasuk, &mahasiswa.Jurusan.Nama); err != nil {
			panic(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas
}

func Create(mahasiswa entities.Mahasiswa) bool {
	result, err := config.DB.Exec(`
		insert into tbl_mahasiswa (nama_mahasiswa, alamat, jenis_kelamin, tgl_masuk, jurusan_id)
		VALUE (?, ?, ?, ?, ?)`,
		mahasiswa.NamaMahasiswa, mahasiswa.Alamat, mahasiswa.JenisKelamin, mahasiswa.TglMasuk,
		mahasiswa.Jurusan.Id)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Mahasiswa {
	row := config.DB.QueryRow(`
	Select * from tbl_mahasiswa where id=?`, id)

	var mahasiswa entities.Mahasiswa
	if err := row.Scan(&mahasiswa.Id, &mahasiswa.NamaMahasiswa, &mahasiswa.Alamat,
		&mahasiswa.JenisKelamin, &mahasiswa.TglMasuk, &mahasiswa.Jurusan.Id); err != nil {
		panic(err)
	}

	return mahasiswa
}

func Update(id int, mahasiswa entities.Mahasiswa) bool {
	query, err := config.DB.Exec(`
		update tbl_mahasiswa set nama_mahasiswa=?, alamat=?, jenis_kelamin=?, tgl_masuk=?, 
		jurusan_id=? where id=?`,
		mahasiswa.NamaMahasiswa, mahasiswa.Alamat, mahasiswa.JenisKelamin, mahasiswa.TglMasuk,
		mahasiswa.Jurusan.Id, id)

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
	_, err := config.DB.Exec(`delete from tbl_mahasiswa where id=?`, id)

	return err
}
