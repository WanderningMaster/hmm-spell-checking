package utils

import "log"

func Require(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
