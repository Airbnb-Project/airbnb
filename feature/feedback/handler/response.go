package handler

type FeedbackResponse struct {
	ID              uint   `json:"id"`
	Rating          uint   `json:"rating"`
	Note            string `json:"note"`
	UserID          uint   `json:"user_id"`
	UserName        string `json:"user_name"`
	HomestayID      uint   `json:"homestay_id"`
	HomestayName    string `json:"homestay_name"`
	HomestayAddress string `json:"homestay_address"`
}
