package main

import (
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/ssp"
)

func main() {
	s := ssp.NewSSP(&ssp.SSP{})
	s.StartServer()
}
