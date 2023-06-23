package main

import (
	"net/http"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/dsp1"
)

func main() {
	http.HandleFunc("/get-bid", dsp1.GetBidHandler)
	http.ListenAndServe(":8081", nil)
}
