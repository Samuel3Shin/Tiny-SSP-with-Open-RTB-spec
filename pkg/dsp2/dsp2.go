package dsp2

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
)

var bannerAds = map[string][][]string{
	"1660x260": {
		{
			"https://wooltariusa.com/collections/snacks?utm_source=naver&utm_medium=display&utm_campaign=PC_%EB%A9%94%EC%9D%B8_%ED%83%80%EC%9E%84%EB%B3%B4%EB%93%9C&utm_content=%EC%9A%B8%ED%83%80%EB%A6%AC%EB%AA%B0_230619&utm_term=k%EA%B0%84%EC%8B%9D",
			"https://dsp-ad-objects.s3.amazonaws.com/wooltari.jpeg"},
		{
			"https://www.hyundaiusa.com/us/ko/vehicles/ioniq-6?cmpid=tnad_ko_hotplace_nv_830x130&utm_source=tnad_nv&utm_medium=tnad_ko_banner",
			"https://dsp-ad-objects.s3.amazonaws.com/hyundai_ioniq6.jpeg"},
	},
	"840x480": {
		{
			"https://patron.naver.com/ntv/c/intro/globalcontents",
			"https://dsp-ad-objects.s3.amazonaws.com/samchongsa.jpeg"},
		{"https://happybean.naver.com/fundings/detail/F944",
			"https://dsp-ad-objects.s3.amazonaws.com/happybean.jpeg"},
		{"https://tv.naver.com/ktstudiogenie", "https://dsp-ad-objects.s3.amazonaws.com/madang.jpeg"},
		{"https://myprofile.naver.com/main", "https://dsp-ad-objects.s3.amazonaws.com/people.png"},
	},
}

func GetBidHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	var bidRequest common.BidRequest
	err := json.NewDecoder(r.Body).Decode(&bidRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a bid
	bidResponse := GenerateBid(bidRequest)

	// Send the bid
	json.NewEncoder(w).Encode(bidResponse)
}

func GenerateBid(bidRequest common.BidRequest) common.BidResponse {
	cfg := common.GetConfig()
	rand.Seed(time.Now().UnixNano())
	seatBids := []common.SeatBid{}
	bannerSize := fmt.Sprintf("%dx%d", bidRequest.Imp[0].Banner.W, bidRequest.Imp[0].Banner.H)
	sampleAdms := bannerAds[bannerSize]

	for i := 0; i < 3; i++ {
		bids := []common.Bid{}
		for j := 0; j < len(sampleAdms); j++ {
			bidAmount := rand.Float64() * 100
			bids = append(bids, common.Bid{
				ID:    fmt.Sprintf("bid%d", j+4),
				ImpID: bidRequest.Imp[0].ID,
				Price: bidAmount,
				AdID:  fmt.Sprintf("ad%d", j+4),
				AdM:   fmt.Sprintf("<a href='%s' target='_blank'><img src='%s' /><img src='%s' style='height:1px; width:1px;' /></a>", sampleAdms[j][0], sampleAdms[j][1], fmt.Sprintf("%s?adID=%s", cfg.LOGSERVER_URL, fmt.Sprintf("ad%d", j+4))),
				NURL:  fmt.Sprintf("%s?adID=%s", cfg.LOGSERVER_URL, fmt.Sprintf("ad%d", j+4)),
			})
		}

		seatBids = append(seatBids, common.SeatBid{
			Bid: bids,
		})
	}

	return common.BidResponse{
		ID:      bidRequest.ID,
		SeatBid: seatBids,
	}
}
