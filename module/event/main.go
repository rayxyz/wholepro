package main

import "fmt"

type Publisher interface {
	Publish(value interface{})
}

type Observer interface {
	Notify(value interface{})
}

type ObserverFunc func(value interface{})

func (fn ObserverFunc) Notify(value interface{}) {
	fn(value)
}

type Observerable []Observer

func (obs *Observerable) AddObserver(ob Observer) {
	*obs = append(*obs, ob)
}

func (obs *Observerable) Publish(value interface{}) {
	for _, obsx := range *obs {
		obsx.Notify(value)
	}
}

type Field struct {
	Value int64
	Observerable
}

func (f *Field) Set(v int64) {
	f.Value = v
	f.Publish(v)
}

func listen0(value interface{}) {
	fmt.Printf("listen0 value: %#v\n", value)
}

func listen1(value interface{}) {
	fmt.Printf("listen1 value: %#v\n", value)
}

func main() {
	fmt.Println("Initiating event mechanism module...")
	v := &Field{}
	v.AddObserver(ObserverFunc(listen0))
	v.AddObserver(ObserverFunc(listen1))
	v.Set(222)
}
