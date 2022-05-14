package domain

type Resep struct {
	Qty    string
	User   []User
	Obat   []Obat
	Apotik []Apotik
}
type User struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
type Obat struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
}
type Apotik struct {
	Nama   string
	Alamat string
}
