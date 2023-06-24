package ssp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
)

type BidGetter interface {
	GetBidFromDSP(bidRequest common.BidRequest, url string) common.BidResponse
}

type SSP struct {
	BidGetter
}

func NewSSP(bg BidGetter) *SSP {
	return &SSP{BidGetter: bg}
}

func (s *SSP) GetBidFromDSPs(bidRequest common.BidRequest) (maxBid common.BidResponse) {
	bidResponses := make(chan common.BidResponse, 2)

	go func() {
		bidResponses <- s.GetBidFromDSP(bidRequest, "http://localhost:8081/get-bid")
	}()

	go func() {
		bidResponses <- s.GetBidFromDSP(bidRequest, "http://localhost:8082/get-bid")
	}()

	response1 := <-bidResponses
	response2 := <-bidResponses

	if response1.SeatBid[0].Bid[0].Price > response2.SeatBid[0].Bid[0].Price {
		return response1
	}

	return response2
}

func (s *SSP) BidRequestHandler(w http.ResponseWriter, r *http.Request) {
	var bidRequest common.BidRequest
	if err := json.NewDecoder(r.Body).Decode(&bidRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxBid := s.GetBidFromDSPs(bidRequest)
	bid := maxBid.SeatBid[0].Bid[0]

	s.fireImpressionPixel(bid)

	if err := json.NewEncoder(w).Encode(maxBid); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *SSP) fireImpressionPixel(bid common.Bid) {
	logMessage := fmt.Sprintf("ID: %s, Bid: %f, AdID: %s", bid.ID, bid.Price, bid.AdID)

	_, err := http.Post("http://localhost:8083", "text/plain", strings.NewReader(logMessage))
	if err != nil {
		log.Printf("Failed to fire impression pixel: %v", err)
	}
}

func (s *SSP) StartServer() {
	http.HandleFunc("/bid", s.BidRequestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
