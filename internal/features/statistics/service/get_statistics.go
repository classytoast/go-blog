package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	posts, err := s.repository.GetPosts(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			" get posts from repository: %w", err,
		)
	}

	return calcStatistics(posts), nil
}

func calcStatistics(posts []domain.Post) domain.Statistics {
	statistics := domain.NewStatistics(0, nil)

	if len(posts) == 0 {
		return statistics
	}

	statistics.PostsCreated = len(posts)

	totalViews := 0
	for _, post := range posts {
		totalViews += int(post.ViewsCount)
	}

	if statistics.PostsCreated > 0 && totalViews > 0 {
		postsAverageViews := float64(totalViews) / float64(statistics.PostsCreated)
		statistics.PostsAverageViews = &postsAverageViews
	}

	return statistics
}
