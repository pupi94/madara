package xtype

import (
	"database/sql/driver"
	"errors"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05.999999"
const TimeFormatISO8601 = "2006-01-02T15:04:05Z"

type UtcTime time.Time

func (t UtcTime) Value() (v driver.Value, err error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t).UTC().Format(TimeFormat), nil
}

func (t *UtcTime) Scan(value interface{}) (err error) {
	if value == nil {
		*t = UtcTime(time.Time{}.UTC())
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		*t = UtcTime(time.Time{}.UTC())
		return errors.New("invalid scan source")
	}
	tt, err := time.ParseInLocation(TimeFormat, string(s), time.UTC)
	if err != nil {
		return errors.New("invalid scan source")
	}
	*t = UtcTime(tt)
	return nil
}

func (t UtcTime) MarshalJSON() (value []byte, err error) {
	xt := time.Time(t)
	if y := xt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(TimeFormatISO8601)+2)
	b = append(b, '"')
	b = xt.UTC().AppendFormat(b, TimeFormatISO8601)
	b = append(b, '"')
	return b, nil

}

func (t *UtcTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	xt, err := time.ParseInLocation(`"`+TimeFormatISO8601+`"`, string(data), time.UTC)
	*t = UtcTime(xt)
	return err
}

func (t *UtcTime) Unix() int64 {
	return time.Time(*t).Unix()
}

type LocalTimex int64

func (t LocalTimex) Value() (v driver.Value, err error) {
	if t == 0 {
		return nil, nil
	}
	return time.Unix(int64(t), 0).UTC().Format(TimeFormat), nil
}

func (t *LocalTimex) Scan(value interface{}) (err error) {
	if value == nil {
		*t = LocalTimex(0)
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		*t = LocalTimex(0)
		return errors.New("invalid scan source")
	}
	tt, err := time.ParseInLocation(TimeFormat, string(s), time.UTC)
	if err != nil {
		return errors.New("invalid scan source")
	}
	*t = LocalTimex(tt.Unix())
	return nil
}
