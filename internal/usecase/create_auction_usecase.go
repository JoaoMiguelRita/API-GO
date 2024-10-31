package create_auction_use_case

import "auction_golang/internal/entity/auction_entity"

type AuctionInputDto struct {
	ProductName string,
	Condition auction_entity.AuctionCondition,
}

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
}

func NewAuctionUseCase(repository auction_entity.AuctionRepositoryInterface) *AuctionUseCase{
	return &AuctionUseCase{
		AuctionRepository: repository,
	}
}

func (u AuctionUseCase) CreateAuction(input *AuctionInputDto) (*AuctionOutputDto, error){
	a := auction_entity.CreateAuction(input.ProductName, input.Condition)

	if err := u.AuctionRepository.CreateAuction(a); err != nil {
		return nil
	}

	return &AuctionOutputDto{
		Id: a.Id,
		ProductName: a.ProductName,
		Condition: a.Condition,
	}, nil
}