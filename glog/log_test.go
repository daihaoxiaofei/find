package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	LogPath = `../log`

	set(`err`, ModeDiscard)
	Error(`test1`)
	set(`err`, ModeStdout)
	Error(`test2`)
	set(`err`, ModeSF)
	Error(`test3`)
}
