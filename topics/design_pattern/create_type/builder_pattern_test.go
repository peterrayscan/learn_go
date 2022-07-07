package createtype

import (
	"fmt"
	"testing"
)

/*
	builder pattern
	建造者模式：
*/

/*
	separate the construction of a complex object from representation
	设计思想：
		* Builder interface (包含1.biuld_XX method 返回的是biulder接口，2.get_XX 返回对象)
		* 父struct
		* Director struct, 属性为Builder, 实现Construct()和SetBuilder()方法
		* 不同的子struct组合，实现接口builder
*/

//定义父struct
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

//builder interface
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Vehicle
}

//Director
type Director struct {
	builder Builder
}

func (director *Director) Construct() {
	director.builder.SetWheels().SetSeats().SetStructure() //链式调用
}

func (director *Director) SetBuilder(builder Builder) {
	director.builder = builder
}

//car struct
type Car struct {
	vehicle Vehicle
}

//实现继承Builder
func (car *Car) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}
func (car *Car) SetSeats() Builder {
	car.vehicle.Seats = 4
	return car
}
func (car *Car) SetStructure() Builder {
	car.vehicle.Structure = "Car"
	return car
}
func (car *Car) GetVehicle() Vehicle {
	return car.vehicle
}

func Test_BuilderPattern(t *testing.T) {
	director := Director{}

	//car
	car := &Car{}
	director.SetBuilder(car)
	director.Construct()
	vehicle := director.builder.GetVehicle()
	//vehicle = car.GetVehicle()
	fmt.Println(vehicle)
	if vehicle.Wheels != 4 {
		t.Errorf("car wheels must be 4, but get %d\n", vehicle.Wheels)
	}
	if vehicle.Seats != 4 {
		t.Errorf("car seats must be 4, but get %d\n", vehicle.Seats)
	}
	if vehicle.Structure != "Car" {
		t.Errorf("vehicle structure must be Car, but get %s\n", vehicle.Structure)
	}
}
