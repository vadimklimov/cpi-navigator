package util

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/styles"
)

func NewLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
	})

	logger.SetStyles(&styles.DefaultStyles().Log)

	return logger
}
