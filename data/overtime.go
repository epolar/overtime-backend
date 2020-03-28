package data

type Overtime struct {
	Model
	Title string `gorm:"varchar:10;index"`
}

type OvertimeRecord struct {
	Model
	OvertimeID uint64 `gorm:"index;unique_index:ou"`
	UserID     uint64 `gorm:"index;unique_index:ou"`
}
