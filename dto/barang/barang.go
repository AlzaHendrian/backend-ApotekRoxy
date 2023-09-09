package barangdto

type CreateBarangRequest struct {
	ID    int    `json:"id" form:"id" validate:"required"`
	Nama  string `json:"nama" form:"nama" validate:"required"`
	Harga int    `json:"harga" form:"harga" validate:"required"`
	Qty   int    `json:"qty" form:"qty" validate:"required"`
}
