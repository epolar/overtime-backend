package overtime

import (
	"errors"
	"github.com/jinzhu/gorm"
	"orderStatistics/data"
	"orderStatistics/repository"
	"orderStatistics/runtime/log"
	"sync"
	"time"
)

var overtimeService Overtime
var overtimeServiceOnce sync.Once

func Service() *Overtime {
	overtimeServiceOnce.Do(func() {
		overtimeService = Overtime{}
	})
	return &overtimeService
}

type Overtime struct{}

func (o *Overtime) getTodayTitle() string {
	now := time.Now().In(data.ChainZone)
	return now.Format("2006-01-02")
}

func (o *Overtime) JoinToday(userID uint64) (err error) {
	userRepository := repository.DefaultUserRepository()

	var userInfo *data.User
	if userInfo, err = userRepository.FindByID(userID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("user not found")
		} else {
			log.Log.Errorf("get user info failure: %s", err)
			return err
		}
	}

	overtimeRepository := repository.DefaultOvertimeRepository()

	overtimeTitle := o.getTodayTitle()
	var overtime *data.Overtime
	if overtime, err = overtimeRepository.FindByTitle(overtimeTitle); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			overtime = &data.Overtime{
				Title: overtimeTitle,
			}
			if err = overtimeRepository.SaveOvertime(overtime); err != nil {
				log.Log.Errorf("save overtime info failure: %s", err)
				return
			}
		} else {
			log.Log.Errorf("find overtime by title(%s) failure: %s", overtimeTitle, err)
			return
		}
	}

	overtimeRecord := &data.OvertimeRecord{
		Overtime: overtime.ID,
		User:     userInfo.ID,
	}
	if err = overtimeRepository.SaveRecord(overtimeRecord); err != nil {
		log.Log.Errorf("save overtime record failure: %s", err)
		return
	}

	return
}

func (o *Overtime) GetTodayRecords() (resp []*data.User, err error) {
	overtimeRepository := repository.DefaultOvertimeRepository()
	resp, err = overtimeRepository.FindRecordsByTitle(o.getTodayTitle())
	if err != nil {
		log.Log.Errorf("find overtime(%s) records failure: %s", o.getTodayTitle(), err)
		return
	}
	return
}
