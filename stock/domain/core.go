package domain

type Stock struct {
	Harga  string
	Qty    string
	Obat   []Obat
	Apotik []Apotik
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
