package repository

import (
	"orderStatistics/data"
)

type OvertimeRepository struct {
	*Repository
}

func NewOvertimeRepository(db *DBCli) *OvertimeRepository {
	return &OvertimeRepository{Repository: NewRepository(db)}
}

func DefaultOvertimeRepository() *OvertimeRepository {
	return NewOvertimeRepository(DB())
}

////////////////// overtime ////////////////////

func (r *OvertimeRepository) FindByTitle(title string) (*data.Overtime, error) {
	var overtime data.Overtime
	if err := r.db.First(&overtime, "title = ?", title).Error; err != nil {
		return nil, err
	} else {
		return &overtime, nil
	}
}

func (r *OvertimeRepository) SaveOvertime(v *data.Overtime) error {
	return r.save(v)
}

/////////////////// record ///////////////////////

func (r *OvertimeRepository) SaveRecord(v *data.OvertimeRecord) error {
	return r.save(v)
}

func (r *OvertimeRepository) FindRecordsByTitle(title string) (resp []*data.OvertimeRecord, err error) {
	var overtime data.Overtime
	if err = r.db.First(&overtime, "title = ?", title).
		Related(&resp, "overtime_id").
		Error; err != nil {
		return
	}
	return
}

func (r *OvertimeRepository) IsJoined(overtimeID uint64, userID uint64) (bool, error) {
	var count uint32

	tableName := r.db.NewScope(&data.OvertimeRecord{}).TableName()
	err := r.db.
		Table(tableName).
		Where("overtime_id = ? and user_id = ?", overtimeID, userID).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
