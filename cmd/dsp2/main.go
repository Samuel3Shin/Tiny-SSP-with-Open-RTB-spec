package main

import (
	"net/http"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/dsp2"
)

func main() {
	http.HandleFunc("/get-bid", dsp2.GetBidHandler)
	http.ListenAndServe(":8082", nil)
}
