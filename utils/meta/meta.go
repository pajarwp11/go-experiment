package meta

import "math"

type Meta struct {
	Current  int `json:"current_page"`
	PerPage  int `json:"per_page"`
	LastPage int `json:"last_page"`
	Total    int `json:"total"`
	From     int `json:"from"`
	To       int `json:"to"`
}

func CreateMeta(page, limit, from, total int) Meta {
	to := from + limit
	if to > total {
		to = total
	}
	return Meta{
		Current:  page,
		PerPage:  limit,
		LastPage: int(math.Ceil(float64(total) / float64(limit))),
		Total:    total,
		From:     from + 1,
		To:       to,
	}
}
