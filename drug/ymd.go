package drug

import (
	"bytes"
	"fmt"
	"time"
)

type YearMonthDay struct {
	time.Time
}

func (ct *YearMonthDay) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, "\"")
	ct.Time, err = time.Parse("20060102", string(b))
	if err == nil {
		return nil
	}
	ct.Time, err = time.Parse("200601", string(b))
	if err == nil {
		return nil
	}
	ct.Time, err = time.Parse("2006", string(b))
	if err == nil {
		return nil
	}
	return err
}

func (ct *YearMonthDay) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%04d%02d%02d\"", ct.Year(), ct.Month(), ct.Day())), nil
}
