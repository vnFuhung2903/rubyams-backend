package repo

import (
	"github.com/vnFuhung2903/rubyams-backend/internal/entity"
	"gorm.io/gorm"
)

type VoteRepository interface {
	FindAll() ([]*entity.Vote, error)
	FindByTxHash(txHash string) (*entity.Vote, error)
	FindByVoter(voter string) (*entity.Vote, error)
	CreateVote(voter string, txHash string, course uint, amount uint) (*entity.Vote, error)
}

type voteRepository struct {
	Db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepository{Db: db}
}

func (vr *voteRepository) FindAll() ([]*entity.Vote, error) {
	var votes []*entity.Vote
	res := vr.Db.Find(&votes)
	if res.Error != nil {
		return nil, res.Error
	}
	return votes, nil
}

func (vr *voteRepository) FindByTxHash(txHash string) (*entity.Vote, error) {
	var vote entity.Vote
	res := vr.Db.First(&vote, entity.Vote{TxHash: txHash})
	if res.Error != nil {
		return nil, res.Error
	}
	return &vote, nil
}

func (vr *voteRepository) FindByVoter(voter string) (*entity.Vote, error) {
	var vote entity.Vote
	res := vr.Db.First(&vote, entity.Vote{Voter: voter})
	if res.Error != nil {
		return nil, res.Error
	}
	return &vote, nil
}

func (vr *voteRepository) CreateVote(voter string, txHash string, course uint, score uint) (*entity.Vote, error) {
	newVote := &entity.Vote{
		Voter:  voter,
		TxHash: txHash,
		Course: course,
		Score:  score,
	}
	res := vr.Db.Create(newVote)
	if res.Error != nil {
		return nil, res.Error
	}
	return newVote, nil
}
