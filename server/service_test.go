package server

import (
	"testing"
	"fmt"
)

func TestNewService(t *testing.T) {

}

func TestService_NextIP(t *testing.T) {
	svr := NewService()
	addr, err := svr.NextIP()
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}
