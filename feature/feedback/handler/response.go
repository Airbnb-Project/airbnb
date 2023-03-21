package handler

type FeedbackResponse struct {
	ID         uint   `json:"id"`
	Rating     uint   `json:"rating"`
	UserID     uint   `json:"user_id"`
	HomestayID uint   `json:"homestay_id"`
	Name       string `json:"name"`
}
