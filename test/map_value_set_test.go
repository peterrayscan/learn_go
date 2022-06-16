package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_MapValueSet(t *testing.T) {
	testx.RunFunc(func() {
		type set map[uint32]struct{}
		mapWithSet := map[uint32]set{}

		type temp struct {
			first  uint32
			second uint32
		}

		temps := []temp{
			{1, 2},
			{1, 3},
			{1, 3},
			{3, 1},
			{3, 1},
			{3, 2},
		}

		for _, temp := range temps {
			s, exist := mapWithSet[temp.first]
			if !exist {
				mapWithSet[temp.first] = set{temp.second: struct{}{}}
				continue
			}
			s[temp.second] = struct{}{}
		}
		fmt.Println("map is:\n", mapWithSet)
	})
}
