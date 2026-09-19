package domain

type Statistics struct {
	PostsCreated      int
	PostsAverageViews *float64
}

func NewStatistics(
	postsCreated int,
	postsAverageViews *float64,
) Statistics {
	return Statistics{
		PostsCreated:      postsCreated,
		PostsAverageViews: postsAverageViews,
	}
}
