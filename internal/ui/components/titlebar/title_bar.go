package titlebar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vadimklimov/cpi-navigator/internal"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
)

type Model struct {
	common common.Common
	title  string
}

func New() *Model {
	return &Model{
		common: common.New(),
		title:  internal.AppLongName,
	}
}

func (*Model) Init() tea.Cmd {
	return nil
}

func (model *Model) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model *Model) View() string {
	return model.common.Styles.TitleBar.Title.Render(model.title)
}
