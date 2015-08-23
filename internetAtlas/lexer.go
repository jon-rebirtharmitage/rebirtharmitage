package main

import (
	"log"
	"reflect"
)

func parseWired(input interface{}){
	log.Print(reflect.TypeOf(input))
}