package repository

import (
	"context"
	"github.com/arandich/marketplace-orders/internal/config"
	"github.com/arandich/marketplace-orders/internal/model"
	"github.com/arandich/marketplace-orders/pkg/metrics"
	pb "github.com/arandich/marketplace-proto/api/proto/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type OrderRepository struct {
	pgPool      *pgxpool.Pool
	promMetrics metrics.Metrics
	logger      *zerolog.Logger
	cfg         config.Config
	clients     model.Clients
}

func NewOrderRepository(ctx context.Context, pgPool *pgxpool.Pool, promMetrics metrics.Metrics, cfg config.Config, clients model.Clients) *OrderRepository {
	return &OrderRepository{
		pgPool:      pgPool,
		promMetrics: promMetrics,
		logger:      zerolog.Ctx(ctx),
		cfg:         cfg,
		clients:     clients,
	}
}

func (o OrderRepository) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return nil, nil
}

func (o OrderRepository) InitOrder(ctx context.Context, req *pb.InitOrderRequest) (*pb.InitOrderResponse, error) {
	user, err := o.clients.IdService.GetUser(ctx, nil)
	if err != nil {
		return nil, err
	}

	o.logger.Info().Str("user", user.UserId).Msg("init order")
	return nil, nil
}
