package attribute

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
)

type Model struct {
	common     common.Common
	attributes table.Table
}

type Attribute struct {
	Key, Value string
}

type AttributesMsg []Attribute

func New() *Model {
	common := common.New()

	attributes := table.New().
		BorderTop(false).
		BorderBottom(false).
		BorderLeft(false).
		BorderRight(false).
		BorderRow(false).
		BorderColumn(false).
		StyleFunc(func(_, col int) lipgloss.Style {
			if col == 0 {
				return common.Styles.AttributesPane.Attribute.Key
			} else {
				return common.Styles.AttributesPane.Attribute.Value
			}
		})

	return &Model{
		common:     common,
		attributes: *attributes,
	}
}

func (model *Model) Init() tea.Cmd {
	return model.AttributesInitCmd
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case AttributesMsg:
		attributes := make([][]string, 0, len(msg))
		for _, attribute := range msg {
			attributes = append(attributes, []string{attribute.Key, attribute.Value})
		}

		model.attributes.ClearRows().Data(table.NewStringData()).Rows(attributes...)
	}

	return model, tea.Batch(cmds...)
}

func (model *Model) View() string {
	return model.attributes.Render()
}

func (*Model) AttributesInitCmd() tea.Msg {
	return AttributesMsg(make([]Attribute, 0))
}

func (*Model) AttributesCmd(attributes []Attribute) tea.Cmd {
	return func() tea.Msg {
		return AttributesMsg(attributes)
	}
}
