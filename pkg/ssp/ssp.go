package ssp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/rs/cors"
)

type BidGetter interface {
	GetBidFromDSP(bidRequest common.BidRequest, url string) (common.BidResponse, error)
}

type SSP struct {
	BidGetter
}

func NewSSP(bg BidGetter) *SSP {
	return &SSP{BidGetter: bg}
}

func (s *SSP) GetBidFromDSPs(bidRequest common.BidRequest) (maxBid common.BidResponse, err error) {
	cfg := common.GetConfig()
	bidResponses := make(chan common.BidResponse, 2)
	errResponses := make(chan error, 2)

	go func() {
		bidResponse, err := s.GetBidFromDSP(bidRequest, fmt.Sprintf("%s/get-bid", cfg.DSP1_URL))
		bidResponses <- bidResponse
		errResponses <- err
	}()

	go func() {
		bidResponse, err := s.GetBidFromDSP(bidRequest, fmt.Sprintf("%s/get-bid", cfg.DSP2_URL))
		bidResponses <- bidResponse
		errResponses <- err
	}()

	response1 := <-bidResponses
	err1 := <-errResponses
	response2 := <-bidResponses
	err2 := <-errResponses

	totalResponses := []common.BidResponse{}

	// Add to total list only if there's no error and response is of valid type
	if err1 == nil && reflect.TypeOf(response1) == reflect.TypeOf(common.BidResponse{}) {
		totalResponses = append(totalResponses, response1)
	}

	if err2 == nil && reflect.TypeOf(response2) == reflect.TypeOf(common.BidResponse{}) {
		totalResponses = append(totalResponses, response2)
	}

	var maxPrice float64
	var maxResponse common.BidResponse

	for _, response := range totalResponses {
		for _, seatBid := range response.SeatBid {
			for _, bid := range seatBid.Bid {
				if bid.Price > maxPrice {
					maxPrice = bid.Price
					maxResponse = response
				}
			}
		}
	}

	return maxResponse, nil
}

func (s *SSP) BidRequestHandler(w http.ResponseWriter, r *http.Request) {
	var bidRequest common.BidRequest
	if err := json.NewDecoder(r.Body).Decode(&bidRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxResponse, _ := s.GetBidFromDSPs(bidRequest)
	var maxBid common.Bid
	var maxPrice float64
	for _, seatBid := range maxResponse.SeatBid {
		for _, bid := range seatBid.Bid {
			if bid.Price > maxPrice {
				maxPrice = bid.Price
				maxBid = bid
			}
		}
	}

	// make sure to return only the max bid to the frontend
	maxResponse.SeatBid = []common.SeatBid{
		{
			Bid: []common.Bid{maxBid},
		},
	}

	if err := json.NewEncoder(w).Encode(maxResponse); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *SSP) StartServer() {
	handler := cors.Default().Handler(http.DefaultServeMux)
	http.HandleFunc("/bid", s.BidRequestHandler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
