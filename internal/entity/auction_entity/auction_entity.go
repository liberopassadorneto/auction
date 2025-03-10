package auction_entity

import "time"

type Auction struct {
	ID               string
	ProductName      string
	Category         string
	Description      string
	ProductCondition ProductCondition
	Status           AuctionStatus
	Timestamp        time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota
	Used
	Refurbished
)
