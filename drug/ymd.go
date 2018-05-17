package drug

import (
	"bytes"
	"time"
)

type YearMonthDay string

func (ct *YearMonthDay) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, "\"")

	(*ct) = YearMonthDay(string(b))

	_, err = time.Parse("20060102", string(b))
	if err == nil {
		return nil
	}
	_, err = time.Parse("200601", string(b))
	if err == nil {
		return nil
	}
	_, err = time.Parse("2006", string(b))
	return err
}

func (ct YearMonthDay) MarshalJSON() ([]byte, error) {
	return []byte(ct), nil
}
