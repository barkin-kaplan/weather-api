package helper

import (
	"fmt"
	"time"
)

type JsonDate struct {
	time.Time
}

const customDateLayout = "02/01/2006"

func (cd *JsonDate) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	// Remove quotes
	strInput = strInput[1 : len(strInput)-1]
	t, err := time.Parse(customDateLayout, strInput)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", cd.Time.Format(customDateLayout))), nil
}