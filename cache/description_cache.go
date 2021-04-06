package cache

import (
	"context"
)

type DescriptionCache struct {
}

func (pc *DescriptionCache) Save(ctx context.Context, productId uint, desc string) {

}

func (pc *DescriptionCache) BatchSave(ctx context.Context, list map[uint]string) {

}
