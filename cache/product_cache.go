package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/models"
	"github.com/pupi94/madara/tools/redis"
	"golang.org/x/sync/errgroup"
)

type ProductCache struct {
	redis          *redis.Client
	descCache      *DescriptionCache
	inventoryCache *InventoryCache
}

func NewProductCache(redisClient *redis.Client) *ProductCache {
	return &ProductCache{
		redis:          redisClient,
		descCache:      NewDescCache(redisClient),
		inventoryCache: NewInventoryCache(redisClient),
	}
}

func (pc *ProductCache) Save(ctx context.Context, product *models.FullProduct) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := pc.saveMainInfo(ctx, product); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		if err := pc.descCache.Save(ctx, product.StoreID, product.ID, aws.StringValue(product.Description)); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		var inventoryDate = make(map[xtypes.Uuid]int64)
		inventoryDate[product.ID] = aws.Int64Value(product.InventoryQuantity)
		for _, v := range product.Variants {
			inventoryDate[v.ID] = aws.Int64Value(v.InventoryQuantity)
		}

		if err := pc.inventoryCache.BatchSave(ctx, product.StoreID, inventoryDate); err != nil {
			return err
		}
		return nil
	})
	return eg.Wait()
}

func (pc *ProductCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, products []*models.FullProduct) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := pc.batchSaveMainInfo(ctx, storeId, products); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		var data = make(map[xtypes.Uuid]string)
		for _, p := range products {
			data[p.ID] = aws.StringValue(p.Description)
		}
		return pc.descCache.BatchSave(ctx, storeId, data)
	})

	eg.Go(func() error {
		var data = make(map[xtypes.Uuid]int64)
		for _, p := range products {
			data[p.ID] = aws.Int64Value(p.InventoryQuantity)

			for _, v := range p.Variants {
				data[v.ID] = aws.Int64Value(v.InventoryQuantity)
			}
		}
		return pc.inventoryCache.BatchSave(ctx, storeId, data)
	})
	return eg.Wait()
}

func (pc *ProductCache) Delete(ctx context.Context, storeId, productId xtypes.Uuid, variantIds []xtypes.Uuid) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		res := pc.redis.Delete(pc.cacheKey(storeId, productId))
		return res.Error()
	})

	eg.Go(func() error {
		return pc.descCache.Delete(ctx, storeId, productId)
	})

	eg.Go(func() error {
		var ids = []xtypes.Uuid{productId}
		ids = append(ids, variantIds...)
		return pc.inventoryCache.Delete(ctx, storeId, ids...)
	})
	return eg.Wait()
}

func (pc *ProductCache) BatchDelete(ctx context.Context, storeId xtypes.Uuid, productIds, variantsIds []xtypes.Uuid) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		var ids []string
		for _, id := range productIds {
			ids = append(ids, pc.cacheKey(storeId, id))
		}
		res := pc.redis.Delete(ids...)
		return res.Error()
	})

	eg.Go(func() error {
		return pc.descCache.Delete(ctx, storeId, productIds...)
	})

	eg.Go(func() error {
		var ids = make([]xtypes.Uuid, len(productIds))
		copy(ids, variantsIds)
		return pc.inventoryCache.Delete(ctx, storeId, ids...)
	})
	return eg.Wait()
}

func (pc *ProductCache) Get(ctx context.Context, storeId, productId xtypes.Uuid) error {
	return nil
}

func (pc *ProductCache) Select(ctx context.Context, storeId xtypes.Uuid, productIds []xtypes.Uuid) error {
	return nil
}

func (pc *ProductCache) saveMainInfo(ctx context.Context, product *models.FullProduct) error {
	temp := product.Description
	defer func() {
		product.Description = temp
	}()

	product.Description = nil
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	res := pc.redis.Set(pc.cacheKey(product.StoreID, product.ID), data)
	return res.Error()
}

func (pc *ProductCache) batchSaveMainInfo(ctx context.Context, storeId xtypes.Uuid, products []*models.FullProduct) error {
	var dataMap = make(map[string]interface{})

	for _, product := range products {
		temp := product.Description
		product.Description = nil

		data, err := json.Marshal(product)
		if err != nil {
			return err
		}
		dataMap[pc.cacheKey(storeId, product.ID)] = data
		product.Description = temp
	}
	res := pc.redis.MultiSet(dataMap)
	return res.Error()
}

func (pc *ProductCache) cacheKey(storeId, id xtypes.Uuid) string {
	return fmt.Sprintf("%s:product:%s", storeId, id)
}
