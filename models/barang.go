package models

type Barang struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Nama  string `json:"nama" gorm:"type:varchar(100);"`
	Harga int    `json:"harga"`
	Qty   int    `json:"qty"`
}
