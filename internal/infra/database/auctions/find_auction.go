package auctions

import (
	"context"
	"errors"
	"github.com/liberopassadorneto/auction/configuration/logger"
	"github.com/liberopassadorneto/auction/internal/entity/auction_entity"
	"github.com/liberopassadorneto/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (r *AuctionRepository) FindAuctionByID(ctx context.Context, id string) (
	*auction_entity.Auction,
	*internal_error.InternalError,
) {
	var auction AuctionMongo
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&auction)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("auction not found in mongodb", err)
			return nil, internal_error.NewNotFoundError("auction not found in mongodb")
		}
		logger.Error("error finding auction in mongodb", err)
		return nil, internal_error.NewInternalServerError("error finding auction in mongodb")
	}

	return &auction_entity.Auction{
		ID:               auction.ID,
		ProductName:      auction.ProductName,
		Category:         auction.Category,
		Description:      auction.Description,
		ProductCondition: auction.ProductCondition,
		Status:           auction.Status,
		Timestamp:        time.Unix(auction.Timestamp, 0),
	}, nil
}

func (r *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category, productName string,
) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("error finding auctions in mongodb", err)
		return nil, internal_error.NewInternalServerError("error finding auctions in mongodb")
	}

	defer cursor.Close(ctx)

	var auctions []AuctionMongo
	if err := cursor.All(ctx, &auctions); err != nil {
		logger.Error("error decoding auctions from mongodb", err)
		return nil, internal_error.NewInternalServerError("error decoding auctions from mongodb")
	}

	var auctionsEntity []auction_entity.Auction
	for _, auction := range auctions {
		auctionsEntity = append(auctionsEntity, auction_entity.Auction{
			ID:               auction.ID,
			ProductName:      auction.ProductName,
			Category:         auction.Category,
			Description:      auction.Description,
			ProductCondition: auction.ProductCondition,
			Status:           auction.Status,
			Timestamp:        time.Unix(auction.Timestamp, 0),
		})
	}
	return auctionsEntity, nil
}
