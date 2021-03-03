package tipeee

import (
	"strings"
	"time"
)

// ListtipsResponse represents a response from
// Ulule's API to a GET /v2.0/partners/tips request.
type ListTipsResponse struct {
	Pager *Pager `json:"pager"`
	Tips  []*Tip `json:"tips"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse(time.RFC3339Nano, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// User represents a Tipeee user.
type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email,omitempty"`
	Pseudo string `json:"pseudo,omitempty"`
}

// Tip represents a Tipeee tip.
type Tip struct {
	ID                int                 `json:"id"`
	Amount            int                 `json:"amount"`
	QuantifiedRewards []*QuantifiedReward `json:"rewards"`
	StartAt           *Time               `json:"start_at,omitempty"`
	EndAt             *Time               `json:"end_at,omitempty"`
	DonationType      string              `json:"donation_type,omitempty"`
}

// QuantifiedReward associates a reward and a quantity.
type QuantifiedReward struct {
	Quantity int     `json:"quantity"`
	Reward   *Reward `json:"reward"`
}

// Reward describes one (among possibly several) reward for a Tip.
type Reward struct {
	ID    int `json:"id"`
	Price int `json:"price"`
	// TODO: translations
}

// Pager is used for pagination.
type Pager struct {
	NbItems     int `json:"item_nbr"`
	NbPages     int `json:"page_nbr"`
	CurrentPage int `json:"current_page"`
}
