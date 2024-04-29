package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/spf13/viper"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
)

type Model struct {
	common          common.Common
	tenant, message string
}

type StatusMsg string

func New() *Model {
	return &Model{
		common:  common.New(),
		tenant:  viper.GetString("tenant.name"),
		message: viper.GetString("tenant.webui_url"),
	}
}

func (*Model) Init() tea.Cmd {
	return nil
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatusMsg:
		model.message = string(msg)
	}

	return model, nil
}

func (model *Model) View() string {
	width := model.common.Styles.StatusBar.Area.GetWidth() -
		model.common.Styles.StatusBar.Area.GetHorizontalFrameSize() -
		model.common.Styles.StatusBar.Tenant.GetHorizontalFrameSize() -
		model.common.Styles.StatusBar.Message.GetHorizontalFrameSize() -
		len(model.tenant)
	message := truncate.StringWithTail(model.message, uint(width), "…")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		model.common.Styles.StatusBar.Tenant.Render(model.tenant),
		model.common.Styles.StatusBar.Message.Render(message),
	)
}

func (*Model) StatusMessageCmd(message string) tea.Cmd {
	return func() tea.Msg {
		return StatusMsg(message)
	}
}
