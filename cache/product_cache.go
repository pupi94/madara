package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pupi94/madara/component/redis"
	"github.com/pupi94/madara/component/xtype"
	"github.com/pupi94/madara/model"
	"gorm.io/gorm"
)

type ProductCache struct {
	redis *redis.Client
	db    *gorm.DB
}

func NewProductCache(redisClient *redis.Client, db *gorm.DB) *ProductCache {
	return &ProductCache{
		redis: redisClient,
		db:    db,
	}
}

func (pc *ProductCache) Save(ctx context.Context, storeId xtype.Uuid, product *model.FullProduct) error {
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

func (pc *ProductCache) BatchSave(ctx context.Context, storeId xtype.Uuid, products []*model.FullProduct) error {
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

func (pc *ProductCache) Delete(ctx context.Context, storeId, productId xtype.Uuid) error {
	res := pc.redis.Delete(pc.cacheKey(storeId, productId))
	return res.Error()
}

func (pc *ProductCache) BatchDelete(ctx context.Context, storeId xtype.Uuid, productIds []xtype.Uuid) error {
	var ids []string
	for _, id := range productIds {
		ids = append(ids, pc.cacheKey(storeId, id))
	}
	res := pc.redis.Delete(ids...)
	return res.Error()
}

func (pc *ProductCache) Get(ctx context.Context, storeId, productId xtype.Uuid) (*model.FullProduct, error) {
	return pc.doGet(ctx, storeId, productId)
}

func (pc *ProductCache) cacheKey(storeId, id xtype.Uuid) string {
	return fmt.Sprintf("%s:product:%s", storeId, id)
}

func (pc *ProductCache) doGet(ctx context.Context, storeId, id xtype.Uuid) (*model.FullProduct, error) {
	resp := pc.redis.Get(pc.cacheKey(storeId, id))

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	if resp.Hit() {
		data, err := resp.Bytes()
		if err != nil {
			return nil, err
		}
		var product model.FullProduct
		if err = json.Unmarshal(data, &product); err != nil {
			return nil, err
		}
		return &product, err
	}

	return pc.backfill(ctx, storeId, id)
}

func (pc *ProductCache) backfill(ctx context.Context, storeId, productId xtype.Uuid) (*model.FullProduct, error) {
	query := model.NewProductQuery(model.PQNoDesc(), model.PQNoInventory())
	product, err := query.GetFullProduct(ctx, pc.db, storeId, productId)
	if err != nil {
		return nil, err
	}
	if err := pc.Save(ctx, storeId, product); err != nil {
		return nil, err
	}
	return product, nil
}
