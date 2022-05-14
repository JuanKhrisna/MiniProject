package domain

type Stock struct {
	Harga  string
	Qty    string
	Obat   []Obat
	User   []User
	Apotik []Apotik
}
type Obat struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
type User struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
type Apotik struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
