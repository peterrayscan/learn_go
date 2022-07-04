package creat_factory

import (
	"log"
	"testing"
)

func Test_SimpleFactory(t *testing.T) {
	fruit, err := newFruit("apple seed")
	if err != nil {
		log.Fatal(err)
	}

	fruit.grow()
}
