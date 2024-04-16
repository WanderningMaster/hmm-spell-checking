package main

import (
	"fmt"

	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func main() {
	pairs := []string{
		"aeriplane	aeroplane",
		"aerolane	aeroplane",
	}

	for _, s := range pairs {
		records, _ := utils.MapWordPair(s)
		fmt.Println(records)
	}
}
