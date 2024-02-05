package main

import (
	"fmt"
	"time"
)

func fomartAsTime(t string) (string, error) {
	// TODO: export this layout in a constants file
	layout := "2006-01-02T15:04:05.000Z"
	var ti, err = time.Parse(layout, t)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%02d:%02d", ti.Hour(), ti.Minute()), nil
}
