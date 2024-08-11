package common

type SQLModel struct {
	ID        int    `json:"id" gorm:"primary_key"`
	CreatedAt string `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updated_at"`
}
