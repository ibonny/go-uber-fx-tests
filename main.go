package main

import "go.uber.org/fx"

type MyStruct bool

func (MyStruct) DoSomething() {
	println("Yo.")
}

func NewStruct() *MyStruct {
	return new(MyStruct)
}

func testout(ms *MyStruct) {
	ms.DoSomething()
}

func main() {
	println("Yo.")

	fx.New(
		fx.NopLogger,
		fx.Provide(NewStruct),
		fx.Invoke(testout),
	)
}
