package types

import "time"

type HoldersResponse []struct {
	Token   string   `json:"token"`
	Holders []Holder `json:"holders"`
}

type Holder struct {
	ProxyWallet           string    `json:"proxyWallet" db:"proxy_wallet"`
	Bio                   string    `json:"bio" db:"bio"`
	Asset                 string    `json:"asset" db:"asset"`
	Pseudonym             string    `json:"pseudonym" db:"pseudonym"`
	Amount                float64   `json:"amount" db:"amount"`
	DisplayUsernamePublic bool      `json:"displayUsernamePublic" db:"display_username_public"`
	OutcomeIndex          int       `json:"outcomeIndex" db:"outcome_index"`
	Name                  string    `json:"name" db:"name"`
	ProfileImage          string    `json:"profileImage" db:"profile_image"`
	ProfileImageOptimized string    `json:"profileImageOptimized" db:"profile_image_optimized"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	MarketID              string    `json:"market_id" db:"market_id"`
}
