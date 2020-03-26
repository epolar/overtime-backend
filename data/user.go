package data

type User struct {
	Model
	Name  string `gorm:"type:varchar(8);not null;"`
	Label string `gorm:"type:varchar(16);"`
}
