package dto

import "time"

type ProductResponse struct {
	Id        int
	Nama      string
	Harga     int
	Quantity  int
	CreatedAt time.Time
}
