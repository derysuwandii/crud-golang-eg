package entities

type Mahasiswa struct {
	Id            uint
	NamaMahasiswa string
	Alamat        string
	JenisKelamin  string
	TglMasuk      string
	Jurusan       Jurusan
}
