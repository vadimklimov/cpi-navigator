package util

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/golang-module/carbon/v2"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/styles"
)

func NewLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      carbon.DateTimeLayout,
	})

	logger.SetStyles(&styles.DefaultStyles().Log)

	return logger
}
