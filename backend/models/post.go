package models

type Post struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	ImageURL  string `json:"image_url"`
	Likes     int    `json:"likes"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
