package grpc

import (
	"context"
	"basket-api/model"
	"basket-api/proto/gen"
	"basket-api/repositories"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type BasketService struct {
	repository repositories.BasketRepository
	gen.UnimplementedBasketServer
}

func NewBasketService(s *grpc.Server, repository repositories.BasketRepository) *BasketService {
	srv := BasketService{
		repository: repository,
	}

	gen.RegisterBasketServer(s, &srv)
	reflection.Register(s)

	return &srv

}

func (bs *BasketService) GetBasket(ctx context.Context, request *gen.GetBasketRequest) (*gen.CustomerBasketResponse, error) {
	userId, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("Begin GetBasketById call for basket id %s", userId)

	data, err := bs.repository.GetBasket(ctx, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve basket")
	}
	if data == nil {
		return &gen.CustomerBasketResponse{}, nil
	}

	return mapToCustomerBasketResponse(data), nil

}

func (bs *BasketService) UpdateBasket(ctx context.Context, request *gen.UpdateBasketRequest) (*gen.CustomerBasketResponse, error) {
	userId, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to retrieve user ID from context: %v", err)
	}

	log.Debug().Msgf("Begin UpdateBasket call for basket id %s", userId)

	customerBasket := mapToCustomerBasket(userId, request)
	err = bs.repository.UpdateBasket(ctx, customerBasket)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update basket: %v", err)
	}

	return mapToCustomerBasketResponse(customerBasket), nil
}

func (bs *BasketService) DeleteBasket(ctx context.Context, request *gen.DeleteBasketRequest) (*gen.DeleteBasketResponse, error) {
	userId, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to retrieve user ID from context: %v", err)
	}

	err = bs.repository.DeleteBasket(ctx, userId)
	return nil, err
}

func getUserIDFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Failed to retrieve metadata from context")
	}

	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "Missing user ID in context metadata")
	}

	userID := userIDs[0]
	if userID == "" {
		return "", status.Errorf(codes.Unauthenticated, "Empty user ID in context metadata")
	}

	return userID, nil
}

func mapToCustomerBasketResponse(customerBasket *model.CustomerBasket) *gen.CustomerBasketResponse {
	response := &gen.CustomerBasketResponse{}
	for _, item := range customerBasket.Items {
		response.Items = append(response.Items, &gen.BasketItem{
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}

	return response
}

func mapToCustomerBasket(userID string, request *gen.UpdateBasketRequest) *model.CustomerBasket {
	response := &model.CustomerBasket{
		BuyerID: userID,
	}

	for _, item := range request.Items {
		response.Items = append(response.Items, model.BasketItem{
			ProductID: int(item.ProductId),
			Quantity:  int(item.Quantity),
		})
	}

	return response
}
