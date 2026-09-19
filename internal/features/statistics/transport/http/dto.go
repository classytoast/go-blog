package statistics_transport_http

import "github.com/classytoast/go-blog/internal/core/domain"

type GetStatisticsResponse struct {
	PostsCreated      int      `json:"post_created"`
	PostsAverageViews *float64 `json:"post_average_views"`
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	return GetStatisticsResponse{
		PostsCreated:      statistics.PostsCreated,
		PostsAverageViews: statistics.PostsAverageViews,
	}
}
