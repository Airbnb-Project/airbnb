package data

import (
	"airbnb/feature/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating     uint
	UserID     uint
	HomestayID uint
}

func DataToCore(data Feedback) feedback.Core {
	return feedback.Core{
		ID:         data.ID,
		Rating:     data.Rating,
		UserID:     data.UserID,
		HomestayID: data.HomestayID,
	}
}

func CoreToData(data feedback.Core) Feedback {
	return Feedback{
		Model:      gorm.Model{ID: data.ID},
		Rating:     data.Rating,
		UserID:     data.UserID,
		HomestayID: data.HomestayID,
	}
}
