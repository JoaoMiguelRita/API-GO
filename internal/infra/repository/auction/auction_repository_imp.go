package auction_repository_imp

import "auction_golang/internal/entity/auction_entity"

var auctions map[string]*auction_entity.Auction

type AuctionRepositoryImp struct{}

func (r *AuctionRepositoryImp) CreateAuction(auction *auction_entity.Auction) error {
	auctions[auction.Id] = auction

	return nil
}
