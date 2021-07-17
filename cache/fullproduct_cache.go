package cache

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pupi94/madara/component/redis"
	"github.com/pupi94/madara/component/syncx"
	"github.com/pupi94/madara/component/xtype"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/model"
	"golang.org/x/sync/errgroup"
)

type FullProductCache struct {
	redis *redis.Client

	descCache      *DescriptionCache
	productCache   *ProductCache
	inventoryCache *InventoryCache

	sharedCall syncx.SharedCalls
}

func NewFullProductCache(redisClient *redis.Client) *FullProductCache {
	db := config.DB
	return &FullProductCache{
		redis:          redisClient,
		productCache:   NewProductCache(redisClient, db),
		descCache:      NewDescCache(redisClient, db),
		inventoryCache: NewInventoryCache(redisClient, db),
		sharedCall:     syncx.NewSharedCalls(),
	}
}

func (pc *FullProductCache) Save(ctx context.Context, storeId xtype.Uuid, product *model.FullProduct) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return pc.productCache.Save(ctx, storeId, product)
	})

	eg.Go(func() error {
		return pc.descCache.Save(ctx, product.StoreID, product.ID, aws.StringValue(product.Description))
	})

	eg.Go(func() error {
		var inventoryDate = make(map[xtype.Uuid]int64)
		inventoryDate[product.ID] = aws.Int64Value(product.InventoryQuantity)
		for _, v := range product.Variants {
			inventoryDate[v.ID] = aws.Int64Value(v.InventoryQuantity)
		}
		return pc.inventoryCache.BatchSave(ctx, product.StoreID, inventoryDate)
	})
	return eg.Wait()
}

func (pc *FullProductCache) BatchSave(ctx context.Context, storeId xtype.Uuid, products []*model.FullProduct) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return pc.productCache.BatchSave(ctx, storeId, products)
	})

	eg.Go(func() error {
		var data = make(map[xtype.Uuid]string)
		for _, p := range products {
			data[p.ID] = aws.StringValue(p.Description)
		}
		return pc.descCache.BatchSave(ctx, storeId, data)
	})

	eg.Go(func() error {
		var data = make(map[xtype.Uuid]int64)
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

func (pc *FullProductCache) Delete(ctx context.Context, storeId, productId xtype.Uuid, variantIds []xtype.Uuid) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return pc.productCache.Delete(ctx, storeId, productId)
	})

	eg.Go(func() error {
		return pc.descCache.Delete(ctx, storeId, productId)
	})

	eg.Go(func() error {
		var ids = []xtype.Uuid{productId}
		ids = append(ids, variantIds...)
		return pc.inventoryCache.Delete(ctx, storeId, ids...)
	})
	return eg.Wait()
}

func (pc *FullProductCache) BatchDelete(ctx context.Context, storeId xtype.Uuid, productIds, variantsIds []xtype.Uuid) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		var ids []xtype.Uuid
		for _, id := range productIds {
			ids = append(ids, id)
		}
		return pc.productCache.BatchDelete(ctx, storeId, ids)
	})

	eg.Go(func() error {
		return pc.descCache.Delete(ctx, storeId, productIds...)
	})

	eg.Go(func() error {
		var ids = make([]xtype.Uuid, len(productIds))
		copy(ids, variantsIds)
		return pc.inventoryCache.Delete(ctx, storeId, ids...)
	})
	return eg.Wait()
}

func (pc *FullProductCache) Get(ctx context.Context, storeId, productId xtype.Uuid) (*model.FullProduct, error) {
	return pc.doSharedGet(ctx, storeId, productId)
}

func (pc *FullProductCache) MultiGet(ctx context.Context, storeId xtype.Uuid, productIds []xtype.Uuid) ([]*model.FullProduct, error) {
	var products = make([]*model.FullProduct, 0)
	for _, productId := range productIds {
		p, err := pc.Get(ctx, storeId, productId)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (pc *FullProductCache) doSharedGet(ctx context.Context, storeId, productId xtype.Uuid) (*model.FullProduct, error) {
	val, err := pc.sharedCall.Do(fmt.Sprintf("%s:product:%s", storeId, productId), func() (interface{}, error) {
		return pc.doGet(ctx, storeId, productId)
	})
	return val.(*model.FullProduct), err
}

func (pc *FullProductCache) doGet(ctx context.Context, storeId, productId xtype.Uuid) (*model.FullProduct, error) {
	eg, ctx := errgroup.WithContext(ctx)

	var (
		product     *model.FullProduct
		description string
		inventory   int64
	)
	eg.Go(func() error {
		var err error

		product, err = pc.productCache.Get(ctx, storeId, productId)
		if err != nil {
			return err
		}
		var variantIds = make([]xtype.Uuid, 0)
		for _, v := range product.Variants {
			variantIds = append(variantIds, v.ID)
		}
		invMap, err := pc.inventoryCache.MultiGet(ctx, storeId, variantIds)
		if err != nil {
			return err
		}
		for _, v := range product.Variants {
			if inv, ok := invMap[v.ID]; ok {
				v.InventoryQuantity = &inv
			} else {
				v.InventoryQuantity = aws.Int64(0)
			}
		}
		return nil
	})

	eg.Go(func() error {
		var err error
		description, err = pc.descCache.Get(ctx, storeId, productId)
		return err
	})

	eg.Go(func() error {
		var err error
		inventory, err = pc.inventoryCache.Get(ctx, storeId, productId)
		return err
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	product.Description = &description
	product.InventoryQuantity = &inventory
	return product, nil
}
