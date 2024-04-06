package main

import (
	"github.com/XMLHexagram/tdl_ffi/service"
)

import "C"

func init() {

}

func main() {

}

//export Test
func Test() string {
	return "This is a test message from Go"
}

//export StartGrpcServer
func StartGrpcServer(port int) {
	service.S.Start(port)
}
