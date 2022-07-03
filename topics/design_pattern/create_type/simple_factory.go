package creat_factory

import (
	"fmt"
)

type fruit interface {
	grow()
}

type apple struct{}

func (a *apple) grow() {
	fmt.Println("apple grows")
}

type banana struct{}

func (b *banana) grow() {
	fmt.Println("banana grows")
}

// newFruit new fruit
// 简单工厂模式，调用者只关注传入的参数，不用知道创建的逻辑
// 在 参数过多 or 逻辑复杂 的情况下，容易耦合度过高
// 因此只适用于简单逻辑的场景
func newFruit(seed string) (fruit, error) {

	switch seed {
	case "apple seed":
		return &apple{}, nil

	case "banana seed":
		return &banana{}, nil

	default:
		return nil, fmt.Errorf("new fruit fail with un-support seed")
	}
}
