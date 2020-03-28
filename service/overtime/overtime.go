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

	if err := repository.Transaction(func(tx *repository.DBCli) (err error) {
		overtimeRepository := repository.NewOvertimeRepository(tx)

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

		var joined bool
		if joined, err = overtimeRepository.IsJoined(overtime.ID, userInfo.ID); err != nil {
			log.Log.Errorf("assert user is joined overtime failure: %s", err)
			return
		} else if joined {
			return
		}

		overtimeRecord := &data.OvertimeRecord{
			OvertimeID: overtime.ID,
			UserID:     userInfo.ID,
		}
		if err = overtimeRepository.SaveRecord(overtimeRecord); err != nil {
			log.Log.Errorf("save overtime record failure: %s", err)
			return
		}

		return
	}); err != nil {
		return err
	}

	return
}

func (o *Overtime) GetTodayJoinedUsers() (resp []*data.User, err error) {
	overtimeRepository := repository.DefaultOvertimeRepository()
	var records []*data.OvertimeRecord
	records, err = overtimeRepository.FindRecordsByTitle(o.getTodayTitle())
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = nil
			return
		}
		log.Log.Errorf("find overtime(%s) records failure: %s", o.getTodayTitle(), err)
		return
	}

	userIDs := make([]uint64, 0, len(records))
	for _, record := range records {
		userIDs = append(userIDs, record.UserID)
	}
	userRepository := repository.DefaultUserRepository()
	if resp, err = userRepository.FindByIdIn(userIDs); err != nil {
		log.Log.Errorf("find user info by id in failure: %s", err)
		return
	}

	return
}
