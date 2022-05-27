package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
	"unicode/utf8"
)

func Test_StringRuneByte(t *testing.T) {
	testx.RunFunc(func() {
		englishStr := "hello world"
		fmt.Println("englishStr:", englishStr)
		fmt.Println("len of englishStr:", len(englishStr))
		fmt.Println("[]byte is:", []byte(englishStr))
		for i := 0; i < len(englishStr); i++ {
			fmt.Printf("%v: %v,   ", i, englishStr[i])
		}
		fmt.Println()
		fmt.Println("rune count of englishStr:", utf8.RuneCountInString(englishStr))
		for index, runeValue := range englishStr {
			fmt.Printf("%#U starts at %d, value: %#q\n", runeValue, index, runeValue)
		}
		fmt.Println("Using DecodeRuneInString")
		runeV, size := utf8.DecodeRuneInString(englishStr)
		fmt.Printf("runeV:%#U, size: %d\n", runeV, size)

		fmt.Println()

		chineseStr := "哈喽沃尔德"
		fmt.Println("chineseStr:", chineseStr)
		fmt.Println("len of chineseStr:", len(chineseStr))
		fmt.Println("[]byte is:", []byte(chineseStr))
		for i := 0; i < len(chineseStr); i++ {
			fmt.Printf("%v: %v,   ", i, chineseStr[i])
		}
		fmt.Println()
		fmt.Println("rune count of chineseStr:", utf8.RuneCountInString(chineseStr))
		for index, runeValue := range chineseStr {
			fmt.Printf("%#U starts at %d, value: %#q\n", runeValue, index, runeValue)
		}
		runeV, size = utf8.DecodeRuneInString(chineseStr)
		fmt.Printf("runeV:%#U, size: %d\n", runeV, size)
	})
}
