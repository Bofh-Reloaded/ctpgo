package wiring

import (
	"github.com/golobby/container/v2"
	"log"
)

func Wire(resolver interface{}) {
	err := container.Singleton(resolver)
	if (err != nil) {
		log.Fatalf("Wiring failure for %v with error %v", resolver, err)
	}
}

func Fill(obj interface{}) {
	err := container.Fill(obj)
	if (err != nil) {
		log.Fatalf("Fill failure for %v with error %v", obj, err)
	}
}

func Bind(obj interface{}) {
	err := container.Bind(obj)
	if (err != nil) {
		log.Fatalf("Bind failure for %v with error %v", obj, err)
	}
}
