package domain

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	Alamat   string
	Obat     []Obat
}

// Error implements error
func (User) Error() string {
	panic("unimplemented")
}

type Obat struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
