package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	DefaultPage  = 1
	DefaultLimit = 20
	MaxLimit     = 100
)

type Pagination struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"-"`
}

func NewPagination(c *fiber.Ctx) Pagination {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	if page < 1 {
		page = DefaultPage
	}
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

type SortParams struct {
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

func NewSortParams(c *fiber.Ctx, defaultSortBy, defaultSortOrder string) SortParams {
	sortBy := c.Query("sort_by", defaultSortBy)
	sortOrder := c.Query("sort_order", defaultSortOrder)

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = defaultSortOrder
	}

	return SortParams{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
}
