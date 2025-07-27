package platform

import (
	"context"
	"kumparan/internal/repository"
	"kumparan/internal/shared"
	"kumparan/internal/shared/common"
	"kumparan/internal/shared/dto"
)

type (
	Service interface {
		HealthCheck(ctx context.Context) (*dto.HealthCheckResponse, error)
	}

	implPlatform struct {
		deps shared.Deps

		repo repository.Holder
	}
)

func NewPlatformService(deps shared.Deps, repo repository.Holder) Service {
	return &implPlatform{deps: deps, repo: repo}
}

func (i *implPlatform) HealthCheck(ctx context.Context) (*dto.HealthCheckResponse, error) {
	var (
		redisStatus = common.ERROR
		dbStatus    = common.ERROR
	)

	_, err := i.repo.CacheRepository.CheckHealth(ctx)
	if err == nil {
		redisStatus = common.OK
	}

	//Get Generic Database object sql.DB
	err = i.deps.DB.Ping()
	if err != nil {
		dbStatus = common.ERROR
	}

	resp := dto.HealthCheckResponse{
		Status:      common.OK,
		RedisStatus: redisStatus,
		DbStatus:    dbStatus,
	}
	return &resp, nil
}
