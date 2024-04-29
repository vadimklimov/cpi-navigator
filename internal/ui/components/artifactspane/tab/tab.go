package tab

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vadimklimov/cpi-navigator/internal/cpi/api"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
)

type Model struct {
	common      common.Common
	tabs        []Tab
	activeTabID int
}

type Tab struct {
	ArtifactType, Label string
}

type ActiveTabMsg string

func New() *Model {
	supportedArtifactTypes := api.SupportedArtifactTypes()

	tabs := []Tab{
		{supportedArtifactTypes.Designtime.IntegrationFlow.Name, "Integration flows"},
		{supportedArtifactTypes.Designtime.ValueMapping.Name, "Value mappings"},
		{supportedArtifactTypes.Designtime.MessageMapping.Name, "Message mappings"},
		{supportedArtifactTypes.Designtime.ScriptCollection.Name, "Script collections"},
	}

	return &Model{
		common:      common.New(),
		tabs:        tabs,
		activeTabID: 0,
	}
}

func (model *Model) Init() tea.Cmd {
	model.activeTabID = 0

	return nil
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.common.KeyMap.Left):
			model.activeTabID = (model.activeTabID - 1 + len(model.tabs)) % len(model.tabs)
			cmds = append(cmds, model.ActiveTabCmd)
		case key.Matches(msg, model.common.KeyMap.Right):
			model.activeTabID = (model.activeTabID + 1) % len(model.tabs)
			cmds = append(cmds, model.ActiveTabCmd)
		}

	case ActiveTabMsg:
		model.activeTabID = model.tabIDByArtifactType(string(msg))
	}

	return model, tea.Batch(cmds...)
}

func (model *Model) View() string {
	var (
		style   lipgloss.Style
		builder = strings.Builder{}
	)

	for idx, tab := range model.tabs {
		if idx == model.activeTabID {
			style = model.common.Styles.IntegrationArtifactsPane.Tabs.Tab.Active.Copy()
		} else {
			style = model.common.Styles.IntegrationArtifactsPane.Tabs.Tab.Inactive.Copy()
		}

		builder.WriteString(style.Render(tab.Label))

		if idx != (len(model.tabs) - 1) {
			builder.WriteString(model.common.Styles.IntegrationArtifactsPane.Tabs.Tab.Separator.String())
		}
	}

	return lipgloss.NewStyle().Render(builder.String())
}

func (model *Model) ActiveTabCmd() tea.Msg {
	return ActiveTabMsg(model.artifactTypeByActiveTab())
}

func (model *Model) artifactTypeByActiveTab() string {
	return model.tabs[model.activeTabID].ArtifactType
}

func (model *Model) tabIDByArtifactType(artifactType string) int {
	return slices.IndexFunc(model.tabs, func(tab Tab) bool {
		return tab.ArtifactType == artifactType
	})
}
