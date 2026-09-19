package statistics_service

import (
	"context"
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
)

type StatisticsRepository interface {
	GetPosts(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Post, error)
}

type StatisticsService struct {
	repository StatisticsRepository
}

func NewStatisticsService(
	repository StatisticsRepository,
) *StatisticsService {
	return &StatisticsService{
		repository: repository,
	}
}
