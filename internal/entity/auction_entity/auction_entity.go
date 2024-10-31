package auction_entity

func CreateAuction()(productName string, condition AuctionCondition) *Auction{
	return &Auction{
		Id: uuid.New().String(),
		ProductName: productName,
		Condition condition,
	}
}

type AuctionCondition int

type Auction struct {
	Id string
	ProductName string
	Condition AuctionCondition
}

const{
	New AuctionCondition = iota
	Old
}

type AuctionRepositoryInterface interface {
	CreateAuction(auction *Auction) error	
}