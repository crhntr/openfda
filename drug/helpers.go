package drug

import (
	"fmt"
	"time"
)

type YearMonthDay struct {
	time.Time
}

func (ct *YearMonthDay) UnmarshalJSON(b []byte) (err error) {
	ct.Time, err = time.Parse("20060102", string(b))
	return err
}

func (ct *YearMonthDay) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%04d%02d%02d", ct.Year(), ct.Month(), ct.Day())), nil
}

