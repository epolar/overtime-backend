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

func (r *OvertimeRepository) FindRecordsByTitle(title string) (resp []*data.User, err error) {
	records := new([]*data.OvertimeRecord)
	var overtime data.Overtime
	if err = r.db.First(&overtime, "title = ?", title).
		Related(&records, "id").
		Related(&resp, "id").
		Error; err != nil {
		return nil, err
	}
	return
}
