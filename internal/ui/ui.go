package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/vadimklimov/cpi-navigator/internal/appinfo"
	"github.com/vadimklimov/cpi-navigator/internal/config"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/err"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/artifactspane/integrationartifact"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/artifactspane/tab"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/attributespane/attribute"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/packagespane/contentpackage"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/statusbar"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/titlebar"
	"github.com/vadimklimov/cpi-navigator/internal/ui/tools/browser"
)

func Start() error {
	program := tea.NewProgram(NewModel(), tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		log.Fatal("Program failed to start", "err", err)
	}

	return nil
}

type Model struct {
	common        common.Common
	packages      *contentpackage.Model
	artifacts     *integrationartifact.Model
	attributes    *attribute.Model
	tabs          *tab.Model
	titlebar      *titlebar.Model
	statusbar     *statusbar.Model
	layout        int
	activePane    int
	showArtifacts bool
	err           error
}

type LayoutMsg int

const (
	LayoutNormal = iota
	LayoutCompact
)

const (
	PackagesPane = iota
	ArtifactsPane
	AttributesPane
	NoPane
)

func NewModel() *Model {
	var layout int

	switch config.UILayout() {
	case config.LayoutNormal:
		layout = LayoutNormal
	case config.LayoutCompact:
		layout = LayoutCompact
	}

	return &Model{
		common:        common.New(),
		packages:      contentpackage.New(),
		artifacts:     integrationartifact.New(),
		attributes:    attribute.New(),
		tabs:          tab.New(),
		titlebar:      titlebar.New(),
		statusbar:     statusbar.New(),
		layout:        layout,
		activePane:    NoPane,
		showArtifacts: false,
		err:           nil,
	}
}

func (model Model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle(appinfo.Name()),
		model.packages.Init(),
		model.artifacts.Init(),
		model.attributes.Init(),
		model.tabs.Init(),
		model.titlebar.Init(),
		model.statusbar.Init(),
	)
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.common.KeyMap.Quit):
			return model, tea.Quit

		case key.Matches(msg, model.common.KeyMap.Up), key.Matches(msg, model.common.KeyMap.Down):
			switch model.activePane {
			case PackagesPane:
				model.showArtifacts = false
				model.packages.Update(msg)
				cmds = append(cmds,
					model.artifacts.Init(),
					model.tabs.Init(),
					model.attributes.AttributesCmd(model.packages.SelectedPackageAttributes()),
				)

			case ArtifactsPane:
				model.showArtifacts = true
				model.artifacts.Update(msg)
				cmds = append(cmds,
					model.attributes.AttributesCmd(model.artifacts.SelectedArtifactAttributes()),
				)
			}

		case key.Matches(msg, model.common.KeyMap.Left), key.Matches(msg, model.common.KeyMap.Right):
			if model.activePane == ArtifactsPane {
				model.showArtifacts = true
				t, cmd := model.tabs.Update(msg)
				model.tabs = t.(*tab.Model)

				if cmd != nil {
					cmds = append(cmds, cmd)
				}
			}

		case key.Matches(msg, model.common.KeyMap.Enter):
			if model.activePane == PackagesPane {
				model.showArtifacts = true
				cmds = append(cmds,
					model.artifacts.Init(),
					model.tabs.Init(),
				)

				if model.packages.SelectedPackageID() != nil {
					cmds = append(cmds,
						model.artifacts.IntegrationArtifactsByPackageCmd(*model.packages.SelectedPackageID()),
					)
				}
			}

		case key.Matches(msg, model.common.KeyMap.Tab):
			switch model.activePane {
			case PackagesPane:
				model.activePane = ArtifactsPane
				model.showArtifacts = true
				cmds = append(cmds,
					model.attributes.AttributesCmd(model.artifacts.SelectedArtifactAttributes()),
				)

			case ArtifactsPane:
				model.activePane = PackagesPane
				model.showArtifacts = false
				cmds = append(cmds,
					model.attributes.AttributesCmd(model.packages.SelectedPackageAttributes()),
				)
			}

		case key.Matches(msg, model.common.KeyMap.Refresh):
			switch model.activePane {
			case PackagesPane:
				cmds = append(cmds,
					model.artifacts.Init(),
					model.tabs.Init(),
					model.packages.ContentPackagesCmd,
				)

			case ArtifactsPane:
				if model.packages.SelectedPackageID() != nil {
					cmds = append(cmds,
						model.artifacts.IntegrationArtifactsByPackageCmd(*model.packages.SelectedPackageID()),
					)
				}
			}

		case key.Matches(msg, model.common.KeyMap.Open):
			switch model.activePane {
			case PackagesPane:
				if model.packages.SelectedPackageID() != nil {
					cmds = append(cmds,
						browser.OpenURLCmd(model.packages.SelectedPackageWebUIURL()),
					)
				}

			case ArtifactsPane:
				if model.artifacts.SelectedArtifactID() != nil {
					cmds = append(cmds,
						browser.OpenURLCmd(model.artifacts.SelectedArtifactWebUIURL()),
					)
				}
			}

		case key.Matches(msg, model.common.KeyMap.Layout):
			cmds = append(cmds, model.ToggleLayoutCmd())
		}

	case LayoutMsg:
		model.layout = int(msg)

	case contentpackage.ContentPackagesMsg:
		model.activePane = PackagesPane
		model.showArtifacts = false
		p, cmd := model.packages.Update(msg)
		model.packages = p.(*contentpackage.Model)

		if cmd != nil {
			cmds = append(cmds, cmd)
		}

		cmds = append(cmds,
			model.attributes.AttributesCmd(model.packages.SelectedPackageAttributes()),
		)

	case tab.ActiveTabMsg:
		if model.activePane == ArtifactsPane {
			model.showArtifacts = true
			model.tabs.Update(msg)
			model.artifacts.Update(msg)
			cmds = append(cmds,
				model.attributes.AttributesCmd(model.artifacts.SelectedArtifactAttributes()),
			)
		}

	case integrationartifact.IntegrationFlowsMsg,
		integrationartifact.ValueMappingsMsg,
		integrationartifact.MessageMappingsMsg,
		integrationartifact.ScriptCollectionsMsg:
		a, cmd := model.artifacts.Update(msg)
		model.artifacts = a.(*integrationartifact.Model)

		if cmd != nil {
			cmds = append(cmds, cmd)
		}

		if model.activePane == ArtifactsPane {
			cmds = append(cmds,
				model.attributes.AttributesCmd(model.artifacts.SelectedArtifactAttributes()),
			)
		}

	case attribute.AttributesMsg:
		a, cmd := model.attributes.Update(msg)
		model.attributes = a.(*attribute.Model)

		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case statusbar.StatusMsg:
		s, cmd := model.statusbar.Update(msg)
		model.statusbar = s.(*statusbar.Model)

		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case err.ErrorMsg:
		model.err = msg
	}

	return model, tea.Batch(cmds...)
}

func (model Model) View() string {
	if model.err != nil {
		return model.common.Styles.Error.Area.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				model.common.Styles.Error.Title.Render("Error"),
				model.common.Styles.Error.Details.Render(model.err.Error()),
			),
		)
	}

	var (
		packagesPaneStyle, artifactsPaneStyle             lipgloss.Style
		packagesPane, artifactsPane, artifactsPaneContent string
	)

	switch model.activePane {
	case PackagesPane:
		packagesPaneStyle = model.common.Styles.ContentPackagesPane.Active
		artifactsPaneStyle = model.common.Styles.IntegrationArtifactsPane.Inactive
	case ArtifactsPane:
		packagesPaneStyle = model.common.Styles.ContentPackagesPane.Inactive
		artifactsPaneStyle = model.common.Styles.IntegrationArtifactsPane.Active
	default:
		packagesPaneStyle = model.common.Styles.ContentPackagesPane.Inactive
		artifactsPaneStyle = model.common.Styles.IntegrationArtifactsPane.Inactive
	}

	packagesPane = packagesPaneStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			model.common.Styles.ContentPackagesPane.Title.Render("Packages"),
			model.packages.View(),
		),
	)

	if model.showArtifacts {
		artifactsPaneContent = lipgloss.JoinVertical(
			lipgloss.Center,
			model.common.Styles.IntegrationArtifactsPane.Tabs.Area.Render(model.tabs.View()),
			model.artifacts.View(),
		)
	}

	artifactsPane = artifactsPaneStyle.Render(artifactsPaneContent)

	if model.layout == LayoutCompact {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, packagesPane, artifactsPane),
			model.common.Styles.AttributesPane.Area.Render(model.attributes.View()),
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		model.common.Styles.TitleBar.Area.Render(model.titlebar.View()),
		lipgloss.JoinHorizontal(lipgloss.Top, packagesPane, artifactsPane),
		model.common.Styles.AttributesPane.Area.Render(model.attributes.View()),
		model.common.Styles.StatusBar.Area.Render(model.statusbar.View()),
	)
}

func (model *Model) ToggleLayoutCmd() tea.Cmd {
	return func() tea.Msg {
		switch model.layout {
		case LayoutNormal:
			return LayoutMsg(LayoutCompact)
		case LayoutCompact:
			return LayoutMsg(LayoutNormal)
		default:
			return LayoutMsg(LayoutNormal)
		}
	}
}
