package models

type ProductOrderCount struct {
	ProductID   uint   `json:"product_id"`
	NamaProduct string `json:"nama_product"`
	TotalQty    int    `json:"total_qty"`
}

type CustomerSpending struct {
	CustomerID uint    `json:"customer_id"`
	Nama       string  `json:"nama"`
	TotalSpent float64 `json:"total_spent"`
}

type ProductRevenue struct {
	ProductID    uint    `json:"product_id"`
	NamaProduct  string  `json:"nama_product"`
	TotalRevenue float64 `json:"total_revenue"`
}
