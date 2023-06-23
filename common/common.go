package common

// Assuming this structure for a bid
type Bid struct {
	ID     string  `json:"id"`
	Bid    float64 `json:"bid"`
	AdHTML string  `json:"adhtml"`
}

// BidRequest structure
type BidRequest struct {
	ID   string `json:"id"`
	Imp  string `json:"imp"` // for example, this could be the HTML of the ad to be shown
	Site string `json:"site"`
}

// BidResponse structure
type BidResponse struct {
	ID  string `json:"id"`
	Bid Bid    `json:"bid"`
}
