package main

import "fmt"

type car struct {
	color         string
	engineType    string
	hasSunroof    bool
	hasNavigation bool
}

type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(et string) CarBuilder
	SetHasSunroof(flag bool) CarBuilder
	SetHasNavigation(flag bool) CarBuilder
	Build() *car
}

func NewCarBuilder() CarBuilder {
	return &carBuilderStruct{
		car: &car{},
	}
}

type carBuilderStruct struct {
	car *car
}

func (cb *carBuilderStruct) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilderStruct) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

func (cb *carBuilderStruct) SetHasNavigation(flag bool) CarBuilder {
	cb.car.hasNavigation = flag
	return cb
}

func (cb *carBuilderStruct) SetHasSunroof(flag bool) CarBuilder {
	cb.car.hasSunroof = flag
	return cb
}

func (cb *carBuilderStruct) Build() *car {
	return cb.car
}

type Director struct {
	builder CarBuilder
}

func (db *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *car {
	db.builder.SetColor(color).
		SetEngineType(engineType).
		SetHasNavigation(hasNavigation).
		SetHasSunroof(hasSunroof)

	return db.builder.Build()
}

func main() {
	builder := NewCarBuilder()

	director := Director{builder: builder}

	myCar := director.ConstructCar("blue", "v12", true, true)

	fmt.Printf("%+v", myCar)
}
