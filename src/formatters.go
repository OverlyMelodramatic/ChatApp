package main

import (
	"fmt"
	"time"
)

func fomartAsTime(t string) (string, error) {
	layout := "2006-01-02T15:04:05.000Z"
	var ti, err = time.Parse(layout, t)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%02d:%02d", ti.Local().Hour(), ti.Minute()), nil
}
