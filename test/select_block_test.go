package test

import (
	"learn/tools/testx"
	"testing"
)

func Test_SelectBlock(t *testing.T) {
	testx.RunFunc(func() {
		// this will block until timeout
		select {}
	})
}
