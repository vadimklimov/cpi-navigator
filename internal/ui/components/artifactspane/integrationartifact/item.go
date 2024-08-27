package integrationartifact

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/vadimklimov/cpi-navigator/internal/cpi/api"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
)

type Item api.IntegrationArtifact

type ItemDelegate struct {
	common common.Common
}

func (item Item) FilterValue() string {
	return item.Name
}

func NewIntegrationArtifactItemDelegate() ItemDelegate {
	return ItemDelegate{
		common: common.New(),
	}
}

func (ItemDelegate) Height() int {
	return 1
}

func (ItemDelegate) Spacing() int {
	return 0
}

func (ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (itemDelegate ItemDelegate) Render(writer io.Writer, model list.Model, index int, listItem list.Item) {
	item := listItem.(Item)

	var style lipgloss.Style
	if index == model.Index() {
		style = itemDelegate.common.Styles.IntegrationArtifactsPane.Dataset.Item.Selected
	} else {
		style = itemDelegate.common.Styles.IntegrationArtifactsPane.Dataset.Item.Normal
	}

	width := model.Width() - style.GetHorizontalFrameSize()
	content := truncate.StringWithTail(item.Name, uint(width), "…")
	fmt.Fprint(writer, style.Render(content))
}
