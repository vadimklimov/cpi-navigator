package integrationartifact

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/golang-module/carbon/v2"
	"github.com/vadimklimov/cpi-navigator/internal/cpi/api"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/err"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/artifactspane/tab"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/attributespane/attribute"
)

type Model struct {
	common               common.Common
	integrationflows     list.Model
	valuemappings        list.Model
	messagemappings      list.Model
	scriptcollections    list.Model
	selectedArtifactType string
}

type (
	IntegrationFlowsMsg  []api.IntegrationArtifact
	ValueMappingsMsg     []api.IntegrationArtifact
	MessageMappingsMsg   []api.IntegrationArtifact
	ScriptCollectionsMsg []api.IntegrationArtifact
)

var supportedArtifactTypes = api.SupportedArtifactTypes()

func New() *Model {
	common := common.New()

	init := func() list.Model {
		width := common.Styles.IntegrationArtifactsPane.Dataset.Area.GetWidth()
		height := common.Styles.IntegrationArtifactsPane.Dataset.Area.GetHeight()

		list := list.New(make([]list.Item, 0), NewIntegrationArtifactItemDelegate(), width, height)
		list.DisableQuitKeybindings()
		list.SetShowHelp(false)
		list.SetShowTitle(false)
		list.SetFilteringEnabled(false)
		list.SetShowPagination(false)
		list.SetShowStatusBar(false)
		list.SetStatusBarItemName("artifact", "artifacts")
		list.InfiniteScrolling = true
		list.Styles.NoItems = common.Styles.IntegrationArtifactsPane.Dataset.NoItems

		return list
	}

	return &Model{
		common:               common,
		integrationflows:     init(),
		valuemappings:        init(),
		messagemappings:      init(),
		scriptcollections:    init(),
		selectedArtifactType: supportedArtifactTypes.Designtime.IntegrationFlow.Name,
	}
}

func (model *Model) Init() tea.Cmd {
	model.selectedArtifactType = supportedArtifactTypes.Designtime.IntegrationFlow.Name

	return tea.Batch(
		model.IntegrationFlowsInitCmd,
		model.ValueMappingsInitCmd,
		model.MessageMappingsInitCmd,
		model.ScriptCollectionsInitCmd,
	)
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	artifactsToList := func(artifacts []api.IntegrationArtifact) []list.Item {
		items := make([]list.Item, 0, len(artifacts))
		for _, artifact := range artifacts {
			items = append(items, Item(artifact))
		}

		return items
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.common.KeyMap.Up), key.Matches(msg, model.common.KeyMap.Down):
			switch model.selectedArtifactType {
			case supportedArtifactTypes.Designtime.IntegrationFlow.Name:
				model.integrationflows, cmd = model.integrationflows.Update(msg)
			case supportedArtifactTypes.Designtime.ValueMapping.Name:
				model.valuemappings, cmd = model.valuemappings.Update(msg)
			case supportedArtifactTypes.Designtime.MessageMapping.Name:
				model.messagemappings, cmd = model.messagemappings.Update(msg)
			case supportedArtifactTypes.Designtime.ScriptCollection.Name:
				model.scriptcollections, cmd = model.scriptcollections.Update(msg)
			}

			cmds = append(cmds, cmd)
		}

	case tab.ActiveTabMsg:
		model.selectedArtifactType = string(msg)

	case IntegrationFlowsMsg:
		model.integrationflows.SetItems(artifactsToList(msg))

	case ValueMappingsMsg:
		model.valuemappings.SetItems(artifactsToList(msg))

	case MessageMappingsMsg:
		model.messagemappings.SetItems(artifactsToList(msg))

	case ScriptCollectionsMsg:
		model.scriptcollections.SetItems(artifactsToList(msg))
	}

	return model, tea.Batch(cmds...)
}

func (model *Model) View() string {
	switch model.selectedArtifactType {
	case supportedArtifactTypes.Designtime.IntegrationFlow.Name:
		return model.integrationflows.View()
	case supportedArtifactTypes.Designtime.ValueMapping.Name:
		return model.valuemappings.View()
	case supportedArtifactTypes.Designtime.MessageMapping.Name:
		return model.messagemappings.View()
	case supportedArtifactTypes.Designtime.ScriptCollection.Name:
		return model.scriptcollections.View()
	default:
		return model.integrationflows.View()
	}
}

func (model *Model) IntegrationFlowsInitCmd() tea.Msg {
	model.integrationflows.ResetSelected()

	return IntegrationFlowsMsg(make([]api.IntegrationArtifact, 0))
}

func (model *Model) ValueMappingsInitCmd() tea.Msg {
	model.valuemappings.ResetSelected()

	return ValueMappingsMsg(make([]api.IntegrationArtifact, 0))
}

func (model *Model) MessageMappingsInitCmd() tea.Msg {
	model.messagemappings.ResetSelected()

	return MessageMappingsMsg(make([]api.IntegrationArtifact, 0))
}

func (model *Model) ScriptCollectionsInitCmd() tea.Msg {
	model.scriptcollections.ResetSelected()

	return ScriptCollectionsMsg(make([]api.IntegrationArtifact, 0))
}

func (model *Model) IntegrationArtifactsByPackageCmd(packageID string) tea.Cmd {
	return tea.Batch(
		model.IntegrationFlowsByPackageCmd(packageID),
		model.ValueMappingsByPackageCmd(packageID),
		model.MessageMappingsByPackageCmd(packageID),
		model.ScriptCollectionsByPackageCmd(packageID),
	)
}

func (*Model) IntegrationFlowsByPackageCmd(packageID string) tea.Cmd {
	return func() tea.Msg {
		integrationflows, e := api.IntegrationArtifactsByPackageAndType(packageID,
			supportedArtifactTypes.Designtime.IntegrationFlow.Name)
		if e != nil {
			return err.ErrorMsg(e)
		}

		return IntegrationFlowsMsg(integrationflows)
	}
}

func (*Model) ValueMappingsByPackageCmd(packageID string) tea.Cmd {
	return func() tea.Msg {
		valuemappings, e := api.IntegrationArtifactsByPackageAndType(packageID,
			supportedArtifactTypes.Designtime.ValueMapping.Name)
		if e != nil {
			return err.ErrorMsg(e)
		}

		return ValueMappingsMsg(valuemappings)
	}
}

func (*Model) MessageMappingsByPackageCmd(packageID string) tea.Cmd {
	return func() tea.Msg {
		messagemappings, e := api.IntegrationArtifactsByPackageAndType(packageID,
			supportedArtifactTypes.Designtime.MessageMapping.Name)
		if e != nil {
			return err.ErrorMsg(e)
		}

		return MessageMappingsMsg(messagemappings)
	}
}

func (*Model) ScriptCollectionsByPackageCmd(packageID string) tea.Cmd {
	return func() tea.Msg {
		scriptcollections, e := api.IntegrationArtifactsByPackageAndType(packageID,
			supportedArtifactTypes.Designtime.ScriptCollection.Name)
		if e != nil {
			return err.ErrorMsg(e)
		}

		return ScriptCollectionsMsg(scriptcollections)
	}
}

func (model *Model) selectedArtifactItem() list.Item {
	switch model.selectedArtifactType {
	case supportedArtifactTypes.Designtime.IntegrationFlow.Name:
		return model.integrationflows.SelectedItem()
	case supportedArtifactTypes.Designtime.ValueMapping.Name:
		return model.valuemappings.SelectedItem()
	case supportedArtifactTypes.Designtime.MessageMapping.Name:
		return model.messagemappings.SelectedItem()
	case supportedArtifactTypes.Designtime.ScriptCollection.Name:
		return model.scriptcollections.SelectedItem()
	default:
		return model.integrationflows.SelectedItem()
	}
}

func (model *Model) SelectedArtifactAttributes() []attribute.Attribute {
	selectedArtifactItem := model.selectedArtifactItem()
	if selectedArtifactItem == nil {
		return nil
	}

	artifact := selectedArtifactItem.(Item)

	switch model.selectedArtifactType {
	case supportedArtifactTypes.Designtime.IntegrationFlow.Name:
		return []attribute.Attribute{
			{Key: "ID", Value: artifact.ID},
			{Key: "Version", Value: artifact.Version},
			{Key: "Name", Value: artifact.Name},
			{Key: "Description", Value: artifact.Description},
			{Key: "Created by", Value: artifact.CreatedBy},
			{Key: "Created at", Value: carbon.CreateFromTimestampMilli(artifact.CreatedAt).ToIso8601ZuluString()},
			{Key: "Modified by", Value: artifact.ModifiedBy},
			{Key: "Modified at", Value: carbon.CreateFromTimestampMilli(artifact.ModifiedAt).ToIso8601ZuluString()},
		}

	case supportedArtifactTypes.Designtime.ValueMapping.Name,
		supportedArtifactTypes.Designtime.MessageMapping.Name,
		supportedArtifactTypes.Designtime.ScriptCollection.Name:
		return []attribute.Attribute{
			{Key: "ID", Value: artifact.ID},
			{Key: "Version", Value: artifact.Version},
			{Key: "Name", Value: artifact.Name},
			{Key: "Description", Value: artifact.Description},
		}

	default:
		return []attribute.Attribute{}
	}
}
