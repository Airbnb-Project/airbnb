package handler

type HomeResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
	Price    float64 `json:"price"`
	Facility string  `json:"facility"`
	Image    string  `json:"image"`
}

type homeImage struct {
	ImageURL string `json:"image"`
}

type HomeDetailResponse struct {
	HomeResponse
	Images []homeImage `json:"image"`
}
