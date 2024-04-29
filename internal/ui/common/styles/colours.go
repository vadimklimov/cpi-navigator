package styles

import "github.com/charmbracelet/lipgloss"

// Catppuccin Mocha.
type Colours struct {
	Rosewater lipgloss.Color
	Flamingo  lipgloss.Color
	Pink      lipgloss.Color
	Mauve     lipgloss.Color
	Red       lipgloss.Color
	Maroon    lipgloss.Color
	Peach     lipgloss.Color
	Yellow    lipgloss.Color
	Green     lipgloss.Color
	Teal      lipgloss.Color
	Sky       lipgloss.Color
	Sapphire  lipgloss.Color
	Blue      lipgloss.Color
	Lavender  lipgloss.Color
	Text      lipgloss.Color
	Subtext1  lipgloss.Color
	Subtext0  lipgloss.Color
	Overlay2  lipgloss.Color
	Overlay1  lipgloss.Color
	Overlay0  lipgloss.Color
	Surface2  lipgloss.Color
	Surface1  lipgloss.Color
	Surface0  lipgloss.Color
	Base      lipgloss.Color
	Mantle    lipgloss.Color
	Crust     lipgloss.Color
}

func DefaultColours() *Colours {
	colours := new(Colours)

	colours.Rosewater = lipgloss.Color("#f5e0dc")
	colours.Flamingo = lipgloss.Color("#f2cdcd")
	colours.Pink = lipgloss.Color("#f5c2e7")
	colours.Mauve = lipgloss.Color("#cba6f7")
	colours.Red = lipgloss.Color("#f38ba8")
	colours.Maroon = lipgloss.Color("#eba0ac")
	colours.Peach = lipgloss.Color("#fab387")
	colours.Yellow = lipgloss.Color("#f9e2af")
	colours.Green = lipgloss.Color("#a6e3a1")
	colours.Teal = lipgloss.Color("#94e2d5")
	colours.Sky = lipgloss.Color("#89dceb")
	colours.Sapphire = lipgloss.Color("#74c7ec")
	colours.Blue = lipgloss.Color("#89b4fa")
	colours.Lavender = lipgloss.Color("#b4befe")
	colours.Text = lipgloss.Color("#cdd6f4")
	colours.Subtext1 = lipgloss.Color("#bac2de")
	colours.Subtext0 = lipgloss.Color("#a6adc8")
	colours.Overlay2 = lipgloss.Color("#9399b2")
	colours.Overlay1 = lipgloss.Color("#7f849c")
	colours.Overlay0 = lipgloss.Color("#6c7086")
	colours.Surface2 = lipgloss.Color("#585b70")
	colours.Surface1 = lipgloss.Color("#45475a")
	colours.Surface0 = lipgloss.Color("#313244")
	colours.Base = lipgloss.Color("#1e1e2e")
	colours.Mantle = lipgloss.Color("#181825")
	colours.Crust = lipgloss.Color("#11111b")

	return colours
}
