package auctions

import (
	"context"
	"github.com/liberopassadorneto/auction/internal/entity/auction_entity"
	"github.com/liberopassadorneto/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionMongo struct {
	ID               string                          `bson:"_id"`
	ProductName      string                          `bson:"product_name"`
	Category         string                          `bson:"category"`
	Description      string                          `bson:"description"`
	ProductCondition auction_entity.ProductCondition `bson:"product_condition"`
	Status           auction_entity.AuctionStatus    `bson:"status"`
	Timestamp        int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (r *AuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) (
	*internal_error.InternalError,
) {
	auctionMongo := AuctionMongo{
		ID:               auction.ID,
		ProductName:      auction.ProductName,
		Category:         auction.Category,
		Description:      auction.Description,
		ProductCondition: auction.ProductCondition,
		Status:           auction.Status,
		Timestamp:        auction.Timestamp.Unix(),
	}

	_, err := r.Collection.InsertOne(ctx, auctionMongo)
	if err != nil {
		return internal_error.NewInternalServerError("error creating auction in mongodb")
	}

	return nil
}
