package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func fomartAsTime(t string) (string, error) {
	layout, layoutDefined := os.LookupEnv("DATETIME_LAYOUT")
	if !layoutDefined {
		return "", errors.New("could not find the 'DATETIME_LAYOUT' environment variable")
	}
	var ti, err = time.Parse(layout, t)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%02d:%02d", ti.Hour(), ti.Minute()), nil
}
