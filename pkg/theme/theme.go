package theme

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
)

var (
	// DefaultTextColor is the default text color
	DefaultTextColor = style.FgDefault

	// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
	GocuiDefaultTextColor = gocui.ColorDefault

	// ActiveBorderColor is the border color of the active frame
	ActiveBorderColor gocui.Attribute

	// InactiveBorderColor is the border color of the inactive active frames
	InactiveBorderColor gocui.Attribute

	// FilteredActiveBorderColor is the border color of the active frame, when it's being searched/filtered
	SearchingActiveBorderColor gocui.Attribute

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	GocuiSelectedLineBgColor gocui.Attribute

	OptionsColor gocui.Attribute

	// SelectedLineBgColor is the background color for the selected line
	SelectedLineBgColor = style.New()

	// SelectedRangeBgColor is the background color of the selected range of lines
	SelectedRangeBgColor = style.New()

	// CherryPickedCommitColor is the text style when cherry picking a commit
	CherryPickedCommitTextStyle = style.New()

	OptionsFgColor = style.New()

	DiffTerminalColor = style.FgMagenta

	UnstagedChangesColor = style.New()
)

// UpdateTheme updates all theme variables
func UpdateTheme(themeConfig config.ThemeConfig) {
	ActiveBorderColor = GetGocuiStyle(themeConfig.ActiveBorderColor)
	InactiveBorderColor = GetGocuiStyle(themeConfig.InactiveBorderColor)
	SearchingActiveBorderColor = GetGocuiStyle(themeConfig.SearchingActiveBorderColor)
	SelectedLineBgColor = GetTextStyle(themeConfig.SelectedLineBgColor, true)
	SelectedRangeBgColor = GetTextStyle(themeConfig.SelectedRangeBgColor, true)

	cherryPickedCommitBgTextStyle := GetTextStyle(themeConfig.CherryPickedCommitBgColor, true)
	cherryPickedCommitFgTextStyle := GetTextStyle(themeConfig.CherryPickedCommitFgColor, false)
	CherryPickedCommitTextStyle = cherryPickedCommitBgTextStyle.MergeStyle(cherryPickedCommitFgTextStyle)

	unstagedChangesTextStyle := GetTextStyle(themeConfig.UnstagedChangesColor, false)
	UnstagedChangesColor = unstagedChangesTextStyle

	GocuiSelectedLineBgColor = GetGocuiStyle(themeConfig.SelectedLineBgColor)
	OptionsColor = GetGocuiStyle(themeConfig.OptionsTextColor)
	OptionsFgColor = GetTextStyle(themeConfig.OptionsTextColor, false)

	DefaultTextColor = GetTextStyle(themeConfig.DefaultFgColor, false)
	GocuiDefaultTextColor = GetGocuiStyle(themeConfig.DefaultFgColor)
}
