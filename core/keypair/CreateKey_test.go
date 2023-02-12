package keypair

import (
	"fmt"
	"testing"
)

func TestCreateKey(t *testing.T) {
	ch := CreateKey()
	select {
	case p := <-ch:
		fmt.Println(p)
	}
	// {3ffdbe6a43dab8404013ee4c370d8a1b1b1c447817983e2cc1f29022ab13f375 0x53572674504e32349dd8535833a9c4679536eddd}
}
