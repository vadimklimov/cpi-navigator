package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Styles struct {
	Log log.Styles

	ContentPackagesPane struct {
		Inactive lipgloss.Style
		Active   lipgloss.Style
		Title    lipgloss.Style
		Dataset  struct {
			Area    lipgloss.Style
			NoItems lipgloss.Style
			Item    struct {
				Normal   lipgloss.Style
				Selected lipgloss.Style
			}
		}
	}

	IntegrationArtifactsPane struct {
		Inactive lipgloss.Style
		Active   lipgloss.Style
		Tabs     struct {
			Area lipgloss.Style
			Tab  struct {
				Inactive  lipgloss.Style
				Active    lipgloss.Style
				Separator lipgloss.Style
			}
		}
		Dataset struct {
			Area    lipgloss.Style
			NoItems lipgloss.Style
			Item    struct {
				Normal   lipgloss.Style
				Selected lipgloss.Style
			}
		}
	}

	AttributesPane struct {
		Area      lipgloss.Style
		Attribute struct {
			Key   lipgloss.Style
			Value lipgloss.Style
		}
	}

	TitleBar struct {
		Area  lipgloss.Style
		Title lipgloss.Style
	}

	StatusBar struct {
		Area    lipgloss.Style
		Tenant  lipgloss.Style
		Message lipgloss.Style
	}

	Error struct {
		Area    lipgloss.Style
		Title   lipgloss.Style
		Details lipgloss.Style
	}
}

func DefaultStyles() *Styles {
	const (
		LogLevelWidth                 = 5
		ContentPackagesPaneWidth      = 60
		IntegrationArtifactsPaneWidth = 90
		AttributesPaneWidth           = 152
		TitleBarWidth                 = 154
		StatusBarWidth                = 154
		ErrorMessageWidth             = 100
	)

	colours := DefaultColours()
	styles := new(Styles)

	baseCommonStyle := lipgloss.NewStyle().
		Background(colours.Base).
		Foreground(colours.Text)

	baseBorderStyle := lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Border(lipgloss.NormalBorder(), true).
		BorderBackground(colours.Base).
		BorderForeground(colours.Overlay0)

	styles.Log = *log.DefaultStyles()

	styles.Log.Levels[log.DebugLevel] = styles.Log.Levels[log.DebugLevel].
		Width(LogLevelWidth).
		MaxWidth(LogLevelWidth).
		Foreground(colours.Blue)

	styles.Log.Levels[log.InfoLevel] = styles.Log.Levels[log.InfoLevel].
		Width(LogLevelWidth).
		MaxWidth(LogLevelWidth).
		Foreground(colours.Green)

	styles.Log.Levels[log.WarnLevel] = styles.Log.Levels[log.WarnLevel].
		Width(LogLevelWidth).
		MaxWidth(LogLevelWidth).
		Foreground(colours.Yellow)

	styles.Log.Levels[log.ErrorLevel] = styles.Log.Levels[log.ErrorLevel].
		Width(LogLevelWidth).
		MaxWidth(LogLevelWidth).
		Foreground(colours.Red)

	styles.Log.Levels[log.FatalLevel] = styles.Log.Levels[log.FatalLevel].
		Width(LogLevelWidth).
		MaxWidth(LogLevelWidth).
		Foreground(colours.Maroon)

	styles.ContentPackagesPane.Inactive = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(ContentPackagesPaneWidth).
		Height(22)

	styles.ContentPackagesPane.Active = lipgloss.NewStyle().
		Inherit(styles.ContentPackagesPane.Inactive).
		BorderForeground(colours.Lavender)

	styles.ContentPackagesPane.Title = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(ContentPackagesPaneWidth).
		Foreground(colours.Teal).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		AlignHorizontal(lipgloss.Center)

	styles.ContentPackagesPane.Dataset.Area = lipgloss.NewStyle().
		Width(ContentPackagesPaneWidth).
		Height(20)

	styles.ContentPackagesPane.Dataset.NoItems = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(ContentPackagesPaneWidth).
		Foreground(colours.Overlay0).
		AlignHorizontal(lipgloss.Center)

	styles.ContentPackagesPane.Dataset.Item.Normal = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(ContentPackagesPaneWidth).
		MaxWidth(ContentPackagesPaneWidth)

	styles.ContentPackagesPane.Dataset.Item.Selected = lipgloss.NewStyle().
		Inherit(styles.ContentPackagesPane.Dataset.Item.Normal).
		Background(colours.Green).
		Foreground(colours.Crust)

	styles.IntegrationArtifactsPane.Inactive = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(IntegrationArtifactsPaneWidth).
		Height(22)

	styles.IntegrationArtifactsPane.Active = lipgloss.NewStyle().
		Inherit(styles.IntegrationArtifactsPane.Inactive).
		BorderForeground(colours.Lavender)

	styles.IntegrationArtifactsPane.Tabs.Area = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(IntegrationArtifactsPaneWidth).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(colours.Overlay0).
		AlignHorizontal(lipgloss.Center)

	styles.IntegrationArtifactsPane.Tabs.Tab.Inactive = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(18).
		Foreground(colours.Overlay0).
		AlignHorizontal(lipgloss.Center)

	styles.IntegrationArtifactsPane.Tabs.Tab.Active = lipgloss.NewStyle().
		Inherit(styles.IntegrationArtifactsPane.Tabs.Tab.Inactive).
		Foreground(colours.Sky)

	styles.IntegrationArtifactsPane.Tabs.Tab.Separator = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Padding(0, 1).
		Foreground(colours.Overlay0).
		SetString("|")

	styles.IntegrationArtifactsPane.Dataset.Area = lipgloss.NewStyle().
		Width(IntegrationArtifactsPaneWidth).
		Height(20)

	styles.IntegrationArtifactsPane.Dataset.NoItems = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(IntegrationArtifactsPaneWidth).
		Foreground(colours.Overlay0).
		AlignHorizontal(lipgloss.Center)

	styles.IntegrationArtifactsPane.Dataset.Item.Normal = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(IntegrationArtifactsPaneWidth).
		MaxWidth(IntegrationArtifactsPaneWidth)

	styles.IntegrationArtifactsPane.Dataset.Item.Selected = lipgloss.NewStyle().
		Inherit(styles.IntegrationArtifactsPane.Dataset.Item.Normal).
		Background(colours.Peach).
		Foreground(colours.Crust)

	styles.AttributesPane.Area = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(AttributesPaneWidth).
		Height(12).
		MaxHeight(22)

	styles.AttributesPane.Attribute.Key = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(15).
		MaxWidth(15).
		Padding(0, 1).
		Foreground(colours.Blue)

	styles.AttributesPane.Attribute.Value = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(135).
		MaxWidth(135).
		Padding(0, 1)

	styles.TitleBar.Area = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(TitleBarWidth).
		Background(colours.Sapphire).
		AlignHorizontal(lipgloss.Center)

	styles.TitleBar.Title = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Padding(0, 1).
		Background(colours.Sapphire).
		Foreground(colours.Crust)

	styles.StatusBar.Area = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(StatusBarWidth).
		Height(1).
		Background(colours.Surface0)

	styles.StatusBar.Tenant = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Padding(0, 1).
		Background(colours.Lavender).
		Foreground(colours.Crust).
		AlignHorizontal(lipgloss.Center)

	styles.StatusBar.Message = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Padding(0, 1).
		Background(colours.Surface0).
		AlignHorizontal(lipgloss.Left)

	styles.Error.Area = lipgloss.NewStyle().
		Inherit(baseBorderStyle).
		Width(ErrorMessageWidth).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(colours.Red)

	styles.Error.Title = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(ErrorMessageWidth).
		MarginBottom(1).
		Background(colours.Red).
		Foreground(colours.Crust).
		AlignHorizontal(lipgloss.Center)

	styles.Error.Details = lipgloss.NewStyle().
		Inherit(baseCommonStyle).
		Width(ErrorMessageWidth).
		Height(5).
		Padding(0, 1).
		Foreground(colours.Red).
		AlignHorizontal(lipgloss.Left).
		AlignVertical(lipgloss.Center)

	return styles
}
