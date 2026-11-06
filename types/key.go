// Package types defines shared values used across the application.
package types

// Key identifies a keyboard key or command.
type Key string

const (
	// Navigation keys.
	KEY_ENTER     Key = "enter"
	KEY_ESC       Key = "esc"
	KEY_BACKSPACE Key = "backspace"
	KEY_TAB       Key = "tab"
	KEY_SHIFT_TAB Key = "shift+tab"

	// Arrow keys.
	KEY_UP       Key = "up"
	KEY_DOWN     Key = "down"
	KEY_LEFT     Key = "left"
	KEY_RIGHT    Key = "right"
	KEY_PG_UP    Key = "pgup"
	KEY_PG_DOWN  Key = "pgdown"
	KEY_ALT_UP   Key = "alt+up"
	KEY_ALT_DOWN Key = "alt+down"

	// Shortcuts.
	SHORT_CTRL_C    Key = "ctrl+c"
	SHORT_CTRL_Q    Key = "ctrl+q"
	SHORT_CTRL_N    Key = "ctrl+n"
	SHORT_CTRL_E    Key = "ctrl+e"
	SHORT_CTRL_D    Key = "ctrl+d"
	SHORT_CTRL_O    Key = "ctrl+o"
	SHORT_CTRL_X    Key = "ctrl+x"
	SHORT_CTRL_V    Key = "ctrl+v"
	SHORT_CTRL_Y    Key = "ctrl+y"
	SHORT_CTRL_Z    Key = "ctrl+z"
	SHORT_SHIFT_TAB Key = "shift+tab"

	// Command keys.
	CMD_SHIFT_ENTER Key = "shift+enter"
	CMD_ALT_ENTER   Key = "alt+enter"
	CMD_CTRL_ENTER  Key = "ctrl+enter"
	CMD_CTRL_S      Key = "ctrl+s"

	// Other keys.
	KEY_EMPTY Key = ""
)

// symbols are the compact glyphs shown in shortcut hints.
var symbols = map[Key]string{
	KEY_ENTER:       "↳",
	KEY_ESC:         "⎋",
	KEY_TAB:         "⇥",
	KEY_SHIFT_TAB:   "⇧⇥",
	KEY_UP:          "↑",
	KEY_DOWN:        "↓",
	KEY_LEFT:        "←",
	KEY_RIGHT:       "→",
	KEY_ALT_UP:      "⌥↑",
	KEY_ALT_DOWN:    "⌥↓",
	SHORT_CTRL_C:    "^C",
	SHORT_CTRL_Q:    "^Q",
	SHORT_CTRL_N:    "^N",
	SHORT_CTRL_E:    "^E",
	SHORT_CTRL_D:    "^D",
	SHORT_CTRL_O:    "^O",
	SHORT_CTRL_X:    "^X",
	SHORT_CTRL_V:    "^V",
	SHORT_CTRL_Y:    "^Y",
	SHORT_CTRL_Z:    "^Z",
	CMD_SHIFT_ENTER: "⇧↳",
	CMD_ALT_ENTER:   "⌥↳",
	CMD_CTRL_ENTER:  "^↳",
	CMD_CTRL_S:      "^S",
}

// Symbol returns the glyph shown in shortcut hints, or the key itself when it has none.
func (k Key) Symbol() string {
	if symbol, ok := symbols[k]; ok {
		return symbol
	}
	return string(k)
}
