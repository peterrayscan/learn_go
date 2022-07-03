package creat_factory

import (
	"log"
	"testing"
)

func Test_SimpleFatory(t *testing.T) {
	fruit, err := newFruit("apple seed")
	if err != nil {
		log.Fatal(err)
	}

	fruit.grow()
}
