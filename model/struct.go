package model

import "time"

type CaloriesConsumed struct {
	Calories    string    `json:"calories"`
	CreatedAt   time.Time `json:"date"`
	Food        string    `json:"food"`
	MoreDetails string    `json:"more_details"`
	UserID      string    `json:"user_id"`
	ID          int64
}
