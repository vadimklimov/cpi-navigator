package common

import (
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/keymap"
	"github.com/vadimklimov/cpi-navigator/internal/ui/common/styles"
)

type Common struct {
	KeyMap *keymap.KeyMap
	Styles *styles.Styles
}

func New() Common {
	return Common{
		KeyMap: keymap.DefaultKeyMap(),
		Styles: styles.DefaultStyles(),
	}
}
