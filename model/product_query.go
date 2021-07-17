package model

import (
	"context"
	"github.com/pupi94/madara/component/xtype"
	"gorm.io/gorm"
)

var ProductColumns = []string{
	"id", "store_id", "title", "description", "published_at", "inventory_quantity", "created_at", "updated_at",
	"published",
}

type PQueryOptions struct {
	Variants    bool
	Images      bool
	Description bool
	Inventory   bool
}

type PQueryOption func(*PQueryOptions)

type PFilterOptions struct {
	SortBy        string
	SortDirection string
	Page          *uint64
	PageSize      *uint64
}

type PFilterOption func(*PFilterOptions)

type ProductQuery struct {
	Options *PQueryOptions
}

func NewProductQuery(setters ...PQueryOption) *ProductQuery {
	opts := &PQueryOptions{
		Variants:    true,
		Images:      true,
		Description: true,
	}

	for _, setter := range setters {
		setter(opts)
	}
	return &ProductQuery{Options: opts}
}

func (q *ProductQuery) SelectFullProducts(ctx context.Context, db *gorm.DB, opts ...PFilterOption) ([]*FullProduct, error) {
	return nil, nil
}

func (q *ProductQuery) GetFullProduct(ctx context.Context, db *gorm.DB, storeId, id xtype.Uuid) (*FullProduct, error) {
	return nil, nil
}

func PQNoVariants() PQueryOption {
	return func(args *PQueryOptions) {
		args.Variants = false
	}
}

func PQNoImages() PQueryOption {
	return func(args *PQueryOptions) {
		args.Images = false
	}
}

func PQNoDesc() PQueryOption {
	return func(args *PQueryOptions) {
		args.Description = false
	}
}

func PQNoInventory() PQueryOption {
	return func(args *PQueryOptions) {
		args.Description = false
	}
}

func PQWithSortBy(s string) PFilterOption {
	return func(args *PFilterOptions) {
		args.SortBy = s
	}
}

func PQWithSortDirection(s string) PFilterOption {
	return func(args *PFilterOptions) {
		args.SortDirection = s
	}
}

func PQWithPage(n uint64) PFilterOption {
	return func(args *PFilterOptions) {
		args.Page = &n
	}
}

func PQWithPageSize(n uint64) PFilterOption {
	return func(args *PFilterOptions) {
		args.PageSize = &n
	}
}
