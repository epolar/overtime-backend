package repository

import "orderStatistics/data"

func migration(db *DBCli) {
	db.AutoMigrate(
		&data.User{},
		&data.Overtime{},
		&data.OvertimeRecord{},
	)
}
