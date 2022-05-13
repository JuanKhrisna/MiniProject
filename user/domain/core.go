package domain

type User struct {
	Username string
	Password string
	Email    string
	Alamat   string
	Obat     []Obat
}

type Obat struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
