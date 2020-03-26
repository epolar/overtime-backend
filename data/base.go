package data

type Model struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt Timestamp `json:"created_at"`
	UpdatedAt Timestamp `json:"updated_at"`
	DeletedAt Timestamp `json:"-" gorm:"index"`
}
