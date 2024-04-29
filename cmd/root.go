package cmd

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadimklimov/cpi-navigator/internal"
	"github.com/vadimklimov/cpi-navigator/internal/ui"
	"github.com/vadimklimov/cpi-navigator/internal/util"
)

const (
	UserCfgDir              = ".config"
	AppCfgDir               = "cpi-navigator"
	DefaultCfgFileName      = "config"
	DefaultCfgFileExtension = "yaml"
	DefaultLogLevel         = "info"
)

// Set using command flags at runtime.
var (
	cfgFile  string
	logLevel string
)

var (
	mandatoryConfigParams = []string{
		"tenant.base_url",
		"tenant.token_url",
		"tenant.client_id",
		"tenant.client_secret",
	}

	optionalConfigParams = []string{
		"tenant.name",
		"tenant.webui_url",
	}
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
		Use:           "cpi-navigator",
		Version:       internal.AppVersion,
		Short:         internal.AppShortName,
		Long:          internal.AppLongName,
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(_ *cobra.Command, _ []string) {
			if err := ui.Start(); err != nil {
				log.Fatal("Program failed to start", "err", err)
			}
		},
	}

	cmd.SetVersionTemplate(version())

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		fmt.Sprintf("configuration file [default: ./%[2]s, ~/%[1]s/%[2]s]",
			filepath.Join(UserCfgDir, AppCfgDir),
			strings.Join([]string{DefaultCfgFileName, DefaultCfgFileExtension}, "."),
		),
	)

	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "",
		fmt.Sprintf("log level (supported: %s) [default: %s]",
			strings.Join(logLevels, ", "), DefaultLogLevel),
	)

	cobra.OnInitialize(
		initLogger,
		initConfig,
		checkMandatoryConfig,
		setDefaultConfig,
	)

	return cmd
}

func Execute() {
	cmd := NewCmd()
	if err := cmd.Execute(); err != nil {
		logger.Fatal("Initialization failed", "err", err)
	}
}

func version() string {
	terminalColourProfile := func() string {
		switch termenv.EnvColorProfile() {
		case termenv.TrueColor:
			return "True Color"
		case termenv.ANSI256:
			return "ANSI256"
		case termenv.ANSI:
			return "ANSI"
		case termenv.Ascii:
			return "ASCII (Uncolored)"
		default:
			return ""
		}
	}

	label := func(text string) string {
		var textWidth, textIndent uint = 20, 2

		return indent.String(padding.String(text+":", textWidth), textIndent)
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s:\n", internal.AppShortName))
	builder.WriteString(fmt.Sprintf("%s %s\n", label("Version"), cmd.Version))
	builder.WriteString("\n")
	builder.WriteString("Runtime environment:\n")
	builder.WriteString(fmt.Sprintf("%s %s/%s\n", label("Platform"), runtime.GOOS, runtime.GOARCH))
	builder.WriteString(fmt.Sprintf("%s %s\n", label("Color profile"), terminalColourProfile()))

	return builder.String()
}

func initLogger() {
	logLvl, _ := log.ParseLevel(logLevel)
	logger.SetLevel(logLvl)

	if logger.GetLevel() == log.DebugLevel {
		logger.SetReportCaller(true)
	}

	log.SetDefault(logger)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		workDir, err := os.Getwd()
		if err != nil {
			log.Fatal("Unable to determine current (working) directory", "err", err)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Unable to determine user's home directory", "err", err)
		}

		viper.AddConfigPath(workDir)
		viper.AddConfigPath(filepath.Join(homeDir, UserCfgDir, AppCfgDir))
		viper.SetConfigName(DefaultCfgFileName)
		viper.SetConfigType(DefaultCfgFileExtension)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read configuration", "err", err)
	}
}

func checkMandatoryConfig() {
	missingConfigParams := make([]string, 0)

	for _, param := range mandatoryConfigParams {
		if !viper.IsSet(param) {
			missingConfigParams = append(missingConfigParams, param)
		}
	}

	if len(missingConfigParams) > 0 {
		log.Fatal(
			"Mandatory configuration parameters were not provided",
			"required", strings.Join(mandatoryConfigParams, ", "),
			"optional", strings.Join(optionalConfigParams, ", "),
			"provided", strings.Join(viper.AllKeys(), ", "),
			"not provided", strings.Join(missingConfigParams, ", "),
		)
	}
}

func setDefaultConfig() {
	// Set tenant name.
	var tenantName string

	if viper.IsSet("tenant.webui_url") {
		tenantURL, err := url.Parse(viper.GetString("tenant.webui_url"))
		if err != nil {
			log.Fatal("The value provided for the configuration parameter tenant.webui_url is incorrect", "err", err)
		}

		tenantName = strings.Split(tenantURL.Hostname(), ".")[0]
	} else {
		tenantName = "SAP Cloud Integration"
	}

	viper.SetDefault("tenant.name", tenantName)
}
