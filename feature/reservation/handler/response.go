package handler

type rsvResponse struct {
	ID              uint   `json:"id"`
	Ticket          string `json:"ticket"`
	Name            string `json:"homestay_name"`
	Image           string `json:"image"`
	TotalPrice      int    `json:"total_price"`
	ReservationDate string `json:"reservation_date"`
}

type rsvDetailResponse struct {
	ID              uint   `json:"id"`
	Ticket          string `json:"ticket"`
	Name            string `json:"homestay_name"`
	Image           string `json:"image"`
	Address         string `json:"address"`
	Checkin         string `json:"check_in"`
	Checkout        string `json:"check_out"`
	Guest           int    `json:"guest"`
	TotalPrice      int    `json:"total_price"`
	ReservationDate string `json:"reservation_date"`
	Bank            string `json:"bank"`
	VAnumber        string `json:"va_number"`
}
