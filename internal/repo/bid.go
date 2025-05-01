package repo

import (
	"github.com/vnFuhung2903/rubyams-backend/internal/entity"
	"gorm.io/gorm"
)

type BidRepository interface {
	FindAll() ([]*entity.Bid, error)
	FindByTxHash(txHash string) (*entity.Bid, error)
	FindByBidder(bidder string) (*entity.Bid, error)
	CreateBid(bidder string, txHash string, course uint, amount uint) (*entity.Bid, error)
}

type bidRepository struct {
	Db *gorm.DB
}

func NewBidRepository(db *gorm.DB) BidRepository {
	return &bidRepository{Db: db}
}

func (br *bidRepository) FindAll() ([]*entity.Bid, error) {
	var bids []*entity.Bid
	res := br.Db.Find(&bids)
	if res.Error != nil {
		return nil, res.Error
	}
	return bids, nil
}

func (br *bidRepository) FindByTxHash(txHash string) (*entity.Bid, error) {
	var bid entity.Bid
	res := br.Db.First(&bid, entity.Bid{TxHash: txHash})
	if res.Error != nil {
		return nil, res.Error
	}
	return &bid, nil
}

func (br *bidRepository) FindByBidder(bidder string) (*entity.Bid, error) {
	var bid entity.Bid
	res := br.Db.First(&bid, entity.Bid{Bidder: bidder})
	if res.Error != nil {
		return nil, res.Error
	}
	return &bid, nil
}

func (br *bidRepository) CreateBid(bidder string, txHash string, course uint, amount uint) (*entity.Bid, error) {
	newBid := &entity.Bid{
		Bidder: bidder,
		TxHash: txHash,
		Course: course,
		Amount: amount,
		Result: true,
	}
	res := br.Db.Create(newBid)
	if res.Error != nil {
		return nil, res.Error
	}
	return newBid, nil
}
