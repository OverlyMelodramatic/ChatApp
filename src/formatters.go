package main

import (
	"fmt"
	"time"
)

const (
	DatetimeLayout = "2006-01-02T15:04:05.000Z"
)

func fomartAsTime(t string) (string, error) {
	// TODO: export this layout in a constants file
	var ti, err = time.Parse(DatetimeLayout, t)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%02d:%02d", ti.Hour(), ti.Minute()), nil
}
