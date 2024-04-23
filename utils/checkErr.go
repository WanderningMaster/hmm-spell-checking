package utils

import (
	"os"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
)

func Require(err error) {
	logger := logger.GetLogger()
	if err != nil {
		logger.Error("Error: ", err)
		os.Exit(1)
	}
}
