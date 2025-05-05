package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"type:varchar(191);uniqueIndex;not null" json:"email"`
	Name      string `gorm:"type:varchar(191)" json:"name"`
	AvatarURL string `gorm:"type:varchar(255)" json:"avatar_url"`
	// Add more fields as needed (e.g., Microsoft ID)
	Posts []Post `json:"posts"`
}
