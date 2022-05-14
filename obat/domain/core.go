package domain

type Obat struct {
	Nama        string
	Jenis       string
	Harga       string
	Description string
	Stock       []Stock
}
type Stock struct {
	Harga string
	Qty   string
}
