package data

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const timestampLayout = "2006-01-02T15:04:05.000"
const timestampLayoutWithLocation = "2006-01-02T15:04:05.000Z07:00"

var ChainZone = time.FixedZone("utc+8", 8*60*60)

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	u := t.Unix()
	if u < 0 {
		return []byte{'0'}, nil
	}
	return []byte(strconv.Itoa(int(u))), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	str := string(data)

	// unix_timestamp
	secs, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		str = strings.ReplaceAll(str, "\"", "")
		if tl, e := time.Parse(timestampLayoutWithLocation, str); e == nil {
			*t = Timestamp{
				Time: tl,
			}
			return nil
		}
		if tl, e := time.Parse(timestampLayout, str); e != nil {
			return e
		} else {
			*t = Timestamp{
				Time: tl,
			}
		}
	} else {
		*t = Timestamp{
			Time: time.Unix(secs, 0),
		}
	}
	return nil
}

func (t Timestamp) Value() (driver.Value, error) {
	var zero time.Time
	if t.Unix() == zero.Unix() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Timestamp) Scan(v interface{}) error {
	if v == nil {
		t.Time = time.Unix(0, 0)
		return nil
	}
	value, ok := v.(time.Time)
	if ok {
		*t = Timestamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func NewTimestampOfTime(t time.Time) *Timestamp {
	return &Timestamp{Time: t}
}

func NewTimestamp() Timestamp {
	return Timestamp{Time: time.Now()}
}
