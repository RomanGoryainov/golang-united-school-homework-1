package solution

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func TestGetMessage(t *testing.T) {
	exp := emoji.Sprint("Hello :world_map:!")
	message := GetMessage()
	if message != emoji.Sprint("Hello :world_map:!") {
		t.Errorf("Expected:\n\t %q\n\t but got:\n\t %q", exp, message)
	}
}
