package singleton

import "sync"

var once sync.Once

var foo Foo

type Foo struct {
}

func GetInstance() *Foo {
	once.Do(func() {
		foo = Foo{}
	})
	return &foo
}
