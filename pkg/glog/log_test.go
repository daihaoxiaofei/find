package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	Error(`Error`)
	Info(`Info`)
	Debug(`Debug`)
	Ok(`Ok`)
}
