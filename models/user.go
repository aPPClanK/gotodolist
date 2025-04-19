package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"unique;not null" json:"name"`
	Password string `gorm:"not null" json:"-"`
	Tasks    []Task `json:"tasks,omitempty"`
}
