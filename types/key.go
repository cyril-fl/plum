package types

// keys
type Key string

const (
	// Navigation keys
	KEY_ENTER     Key = "enter"
	KEY_ESC       Key = "esc"
	KEY_BACKSPACE Key = "backspace"
	KEY_TAB       Key = "tab"
	KEY_SHIFT_TAB Key = "shift+tab"

	// Arrow keys
	KEY_UP      Key = "up"
	KEY_DOWN    Key = "down"
	KEY_LEFT    Key = "left"
	KEY_RIGHT   Key = "right"
	KEY_PG_UP   Key = "pgup"
	KEY_PG_DOWN Key = "pgdown"

	// Shortcut
	SHORT_CTRL_C    Key = "ctrl+c"
	SHORT_CTRL_Q    Key = "ctrl+q"
	SHORT_CTRL_N    Key = "ctrl+n"
	SHORT_SHIFT_TAB Key = "shift+tab"

	// Command keys
	CMD_SHIFT_ENTER Key = "shift+enter"
	CMD_ALT_ENTER   Key = "alt+enter"
	CMD_CTRL_S      Key = "ctrl+s"

	// Other keys
	KEY_EMPTY Key = ""
)
