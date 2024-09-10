package appinfo

import (
	"cmp"
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
	"github.com/muesli/termenv"
)

// Set using ldflags during build.
var version string

type Application struct {
	ID       string
	Name     string
	FullName string
	Version  string
}

type Platform struct {
	OS   string
	Arch string
}

type Terminal struct {
	ColorProfile string
}

type AppInfo struct {
	Application *Application
	Platform    *Platform
	Terminal    *Terminal
}

var (
	instance *AppInfo
	once     sync.Once
)

func GetInstance() *AppInfo {
	once.Do(func() {
		application := &Application{
			ID:       "cpi-navigator",
			Name:     "CPI Navigator",
			FullName: "Cloud Integration Navigator",
			Version:  cmp.Or(version, "undefined"),
		}

		platform := &Platform{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
		}

		terminal := &Terminal{
			ColorProfile: terminalColorProfile(),
		}

		instance = &AppInfo{
			Application: application,
			Platform:    platform,
			Terminal:    terminal,
		}
	})

	return instance
}

func ID() string {
	return GetInstance().Application.ID
}

func Name() string {
	return GetInstance().Application.Name
}

func FullName() string {
	return GetInstance().Application.FullName
}

func Version() string {
	return GetInstance().Application.Version
}

func (appInfo *AppInfo) String() string {
	label := func(text string) string {
		var textWidth, textIndent uint = 20, 2

		return indent.String(padding.String(text+":", textWidth), textIndent)
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s:\n", appInfo.Application.Name))
	builder.WriteString(fmt.Sprintf("%s %s\n", label("Version"), appInfo.Application.Version))
	builder.WriteString("\n")
	builder.WriteString("Runtime environment:\n")
	builder.WriteString(fmt.Sprintf("%s %s/%s\n", label("Platform"), appInfo.Platform.OS, appInfo.Platform.Arch))
	builder.WriteString(fmt.Sprintf("%s %s\n", label("Color profile"), appInfo.Terminal.ColorProfile))

	return builder.String()
}

func terminalColorProfile() string {
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
		return "unknown"
	}
}
