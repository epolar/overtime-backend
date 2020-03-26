package data

type Overtime struct {
	Model
	Title string `gorm:"varchar:10;index"`
}

type OvertimeRecord struct {
	Model
	Overtime uint64 `gorm:"index"`
	User     uint64 `gorm:"index"`
}
