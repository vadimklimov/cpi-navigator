package browser

import (
	"net/url"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/err"
)

func OpenURLCmd(url *url.URL) tea.Cmd {
	return func() tea.Msg {
		if url != nil {
			if e := browser.OpenURL(url.String()); e != nil {
				return err.ErrorMsg(e)
			}
		}

		return nil
	}
}
