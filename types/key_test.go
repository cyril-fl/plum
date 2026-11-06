package types

import "testing"

func TestKeySymbol(t *testing.T) {
	tests := map[Key]string{
		KEY_ENTER:      "↳",
		CMD_CTRL_ENTER: "^↳",
		KEY_SHIFT_TAB:  "⇧⇥",
		SHORT_CTRL_D:   "^D",
		SHORT_CTRL_O:   "^O",
		SHORT_CTRL_X:   "^X",
		SHORT_CTRL_V:   "^V",
		Key("f5"):      "f5",
	}

	for key, want := range tests {
		if got := key.Symbol(); got != want {
			t.Errorf("%q.Symbol() = %q, want %q", key, got, want)
		}
	}
}
