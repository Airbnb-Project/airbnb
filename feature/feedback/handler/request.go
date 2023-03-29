package handler

type AddRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Rating     uint   `json:"rating" form:"rating"`
	Note       string `json:"note" form:"note"`
}

type ListRequest struct {
	HomestayID uint `json:"homestay_id" form:"homestay_id"`
}
