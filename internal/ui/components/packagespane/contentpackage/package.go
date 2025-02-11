package contentpackage

import (
	"net/url"
	"slices"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vadimklimov/cpi-navigator/internal/config"
	"github.com/vadimklimov/cpi-navigator/internal/cpi/api"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/err"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/sort"
	"github.com/vadimklimov/cpi-navigator/internal/ui/components/attributespane/attribute"
)

type Model struct {
	common   common.Common
	packages list.Model
}

type ContentPackagesMsg []api.ContentPackage

func New() *Model {
	common := common.New()

	init := func() list.Model {
		width := common.Styles.ContentPackagesPane.Dataset.Area.GetWidth()
		height := common.Styles.ContentPackagesPane.Dataset.Area.GetHeight()

		list := list.New(make([]list.Item, 0), NewContentPackageItemDelegate(), width, height)
		list.DisableQuitKeybindings()
		list.SetShowHelp(false)
		list.SetShowTitle(false)
		list.SetFilteringEnabled(false)
		list.SetShowPagination(false)
		list.SetShowStatusBar(false)
		list.SetStatusBarItemName("package", "packages")
		list.InfiniteScrolling = true
		list.Styles.NoItems = common.Styles.ContentPackagesPane.Dataset.NoItems

		return list
	}

	return &Model{
		common:   common,
		packages: init(),
	}
}

func (model *Model) Init() tea.Cmd {
	return model.ContentPackagesCmd
}

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, model.common.KeyMap.Up), key.Matches(msg, model.common.KeyMap.Down):
			model.packages, cmd = model.packages.Update(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case ContentPackagesMsg:
		model.packages.SetItems(convertPackagesToListItems(msg))
		model.packages.ResetSelected()
	}

	return model, tea.Batch(cmds...)
}

func (model *Model) View() string {
	return model.packages.View()
}

func (*Model) ContentPackagesCmd() tea.Msg {
	packages, e := api.ContentPackages()
	if e != nil {
		return err.ErrorMsg(e)
	}

	return ContentPackagesMsg(packages)
}

func (model *Model) selectedPackageItem() list.Item {
	return model.packages.SelectedItem()
}

func (model *Model) SelectedPackageID() *string {
	selectedPackageItem := model.selectedPackageItem()
	if selectedPackageItem == nil {
		return nil
	}

	selectedPackage := selectedPackageItem.(Item)

	return &selectedPackage.ID
}

func (model *Model) SelectedPackageAttributes() []attribute.Attribute {
	selectedPackageItem := model.selectedPackageItem()
	if selectedPackageItem == nil {
		return nil
	}

	pkg := selectedPackageItem.(Item)

	return []attribute.Attribute{
		{Key: "ID", Value: pkg.ID},
		{Key: "Version", Value: pkg.Version},
		{Key: "Name", Value: pkg.Name},
		{Key: "Short text", Value: pkg.ShortText},
		{Key: "Vendor", Value: pkg.Vendor},
		{Key: "Mode", Value: pkg.Mode},
		{Key: "Created by", Value: pkg.CreatedBy},
		{Key: "Created at", Value: time.UnixMilli(pkg.CreationDate).UTC().Format(time.RFC3339)},
		{Key: "Modified by", Value: pkg.ModifiedBy},
		{Key: "Modified at", Value: time.UnixMilli(pkg.ModifiedDate).UTC().Format(time.RFC3339)},
	}
}

func (model *Model) SelectedPackageWebUIURL() *url.URL {
	tenantWebUIURL := config.TenantWebUIURL()
	if tenantWebUIURL == nil {
		return nil
	}

	tenantWorkspaceWebUIURL := tenantWebUIURL.JoinPath("shell/design")

	selectedPackageID := model.SelectedPackageID()
	if selectedPackageID == nil {
		return tenantWorkspaceWebUIURL
	}

	return tenantWorkspaceWebUIURL.JoinPath("contentpackage", *selectedPackageID)
}

func convertPackagesToListItems(packages []api.ContentPackage) []list.Item {
	sort.Sort(packages, sort.Options{
		Field: config.UIPackagesPaneSortField(),
		Order: config.UIPackagesPaneSortOrder(),
	})

	items := make([]list.Item, 0, len(packages))
	for pkg := range slices.Values(packages) {
		items = append(items, Item(pkg))
	}

	return items
}
