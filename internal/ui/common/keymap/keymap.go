package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Enter   key.Binding
	Tab     key.Binding
	Quit    key.Binding
	Layout  key.Binding
	Refresh key.Binding
	Open    key.Binding
}

func DefaultKeyMap() *KeyMap {
	keymap := new(KeyMap)

	keymap.Up = key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	)

	keymap.Down = key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	)

	keymap.Left = key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "left"),
	)

	keymap.Right = key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "right"),
	)

	keymap.Enter = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	)

	keymap.Tab = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next"),
	)

	keymap.Quit = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)

	keymap.Layout = key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "layout"),
	)

	keymap.Refresh = key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	)

	keymap.Open = key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open"),
	)

	return keymap
}
