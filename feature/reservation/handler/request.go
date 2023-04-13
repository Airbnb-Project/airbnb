package handler

type rsvRequest struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Checkin    string `json:"checkin" form:"checkin"`
	Checkout   string `json:"checkout" form:"checkout"`
	Guest      int    `json:"guest" form:"guest"`
	TotalPrice int    `json:"total_price" form:"total_price"`
	Bank       string `json:"bank" form:"bank"`
}

type Callback struct {
	OrderID           string `json:"order_id" form:"order_id"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
}
