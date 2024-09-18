package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/vadimklimov/cpi-navigator/internal/appinfo"
	"github.com/vadimklimov/cpi-navigator/internal/config"
	"github.com/vadimklimov/cpi-navigator/internal/ui"
	"github.com/vadimklimov/cpi-navigator/internal/util"
)

const DefaultLogLevel = "info"

// Set using command flags at runtime.
var (
	configFile string
	logLevel   string
)

var logLevels = []string{
	strings.ToLower(log.DebugLevel.String()),
	strings.ToLower(log.InfoLevel.String()),
	strings.ToLower(log.WarnLevel.String()),
	strings.ToLower(log.ErrorLevel.String()),
	strings.ToLower(log.FatalLevel.String()),
}

var (
	cmd    *cobra.Command
	logger = util.NewLogger()
)

func NewCmd() *cobra.Command {
	cmd = &cobra.Command{
		Use:           appinfo.ID(),
		Version:       appinfo.Version(),
		Short:         appinfo.Name(),
		Long:          appinfo.FullName(),
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(_ *cobra.Command, _ []string) {
			if err := ui.Start(); err != nil {
				log.Fatal("Program failed to start", "err", err)
			}
		},
	}

	cmd.SetVersionTemplate(appinfo.GetInstance().String())

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		fmt.Sprintf("configuration file [default: ./%[2]s, ~/%[1]s/%[2]s]",
			filepath.Join(config.DefaultUserConfigDir, config.DefaultAppConfigDir),
			strings.Join([]string{config.DefaultConfigFileName, config.DefaultConfigFileExt}, "."),
		),
	)

	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "",
		fmt.Sprintf("log level (supported: %s) [default: %s]",
			strings.Join(logLevels, ", "), DefaultLogLevel),
	)

	cobra.OnInitialize(
		initLogger,
		func() { config.Init(configFile) },
	)

	return cmd
}

func Execute() {
	cmd := NewCmd()
	if err := cmd.Execute(); err != nil {
		logger.Fatal("Initialization failed", "err", err)
	}
}

func initLogger() {
	logLvl, _ := log.ParseLevel(logLevel)
	logger.SetLevel(logLvl)

	if logger.GetLevel() == log.DebugLevel {
		logger.SetReportCaller(true)
	}

	log.SetDefault(logger)
}
