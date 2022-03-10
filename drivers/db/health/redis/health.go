package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	dbHealth "github.com/burntcarrot/apollo/drivers/db/health"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const MAX_FETCH_ROWS = 9 * 100000

type HealthRepo struct {
	Conn   *redis.Client
	logger *zap.SugaredLogger
}

func NewHealthRepo(conn *redis.Client, logger *zap.SugaredLogger) *HealthRepo {
	return &HealthRepo{
		Conn:   conn,
		logger: logger,
	}
}

func (h *HealthRepo) CreateService(ctx context.Context, srv health.Domain, id uint) (health.Domain, error) {
	createdService := dbHealth.Service{
		Name: srv.Name,
		URI:  srv.URI,
	}

	raw, err := json.Marshal(createdService)
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to marshal: %v\n", err)
		return health.Domain{}, errors.New("failed to marshal")
	}

	// remember, a service ID can be used by other components in the service
	key := fmt.Sprintf("apollo:%d:services", id)
	err = h.Conn.RPush(ctx, key, raw).Err()
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to push to %s: %v\n", key, err)
		return health.Domain{}, errors.New("failed to create service")
	}

	// also, push the created service to key (possible data duplication)
	// but, we get the birds eye view of all services in the system
	newkey := "apollo:services"
	err = h.Conn.RPush(ctx, newkey, raw).Err()
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to push to %s: %v\n", newkey, err)
		return health.Domain{}, errors.New("failed to create service")
	}

	return createdService.ToDomain(), nil
}

func (h *HealthRepo) GetServicesByID(ctx context.Context, id uint) ([]health.Domain, error) {
	key := fmt.Sprintf("apollo:%d:services", id)
	raw, err := h.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		h.logger.Errorf("[health/redis/GetServicesByID] failed to fetch %s: %v\n", key, err)
		return []health.Domain{}, errors.New("failed to fetch services")
	}

	srv := new(dbHealth.Service)
	var services []health.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), srv); err != nil {
			h.logger.Errorf("[health/redis/GetServicesByID] failed to unmarshal: %v\n", err)
			return []health.Domain{}, errors.New("failed to unmarshal")
		}

		services = append(services, srv.ToDomain())
	}

	return services, nil
}

func (h *HealthRepo) GetAllServices(ctx context.Context) ([]health.Domain, error) {
	key := "apollo:services"
	raw, err := h.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		h.logger.Errorf("[health/redis/GetAllServices] failed to fetch %s: %v\n", key, err)
		return []health.Domain{}, errors.New("failed to fetch services")
	}

	srv := new(dbHealth.Service)
	var services []health.Domain

	for _, j := range raw {
		if err := json.Unmarshal([]byte(j), srv); err != nil {
			h.logger.Errorf("[health/redis/GetAllServices] failed to unmarshal: %v\n", err)
			return []health.Domain{}, errors.New("failed to unmarshal")
		}

		services = append(services, srv.ToDomain())
	}

	return services, nil
}
