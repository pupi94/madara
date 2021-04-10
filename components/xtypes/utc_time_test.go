package xtypes

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gitlab.shoplazza.site/shoplaza/samoyed/config"

	sq "github.com/Masterminds/squirrel"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"gitlab.shoplazza.site/shoplaza/samoyed/utils/structs"
)

type Store struct {
	ID                  UUIDBinary `db:"id" json:"id" faker:"uuid_hyphenated"`               //  not null
	OriginID            *string    `db:"origin_id" json:"origin_id"`                         //
	HistoryProductTags  *string    `db:"history_product_tags" json:"history_product_tags"`   //
	CreatedAt           *UTCTime   `db:"created_at" json:"created_at"`                       //  not null
	UpdatedAt           *UTCTime   `db:"updated_at" json:"updated_at"`                       //  not null
	CollectionNumberSeq *int64     `db:"collection_number_seq" json:"collection_number_seq"` //  default:1 专辑编号序列
	Config              *string    `db:"config" json:"config"`                               //
}

func (Store) TableName() string {
	return "stores"
}

func TestLocalTime_Value(t *testing.T) {
	config.InitDB()
	db := config.DB
	query, args, err := sq.Delete("stores").ToSql()
	if err != nil {
		panic("clear db failed!")
	}
	db.MustExec(query, args...)
	sID := UUIDBinary(uuid.New().String())
	createdAt := UTCTime(time.Now())
	//createdAt := LocalTimex(time.Now().Unix())
	updatedAt := UTCTime(time.Now())

	store := &Store{
		ID:                  sID,
		OriginID:            aws.String("123"),
		HistoryProductTags:  nil,
		CreatedAt:           &createdAt,
		UpdatedAt:           &updatedAt,
		CollectionNumberSeq: nil,
		Config:              nil,
	}
	query, args, err = sq.Insert("stores").SetMap(structs.SqMap(store)).ToSql()
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(query, args...); err != nil {
		panic(err)
	}
}

func TestJSON_MarshalJSON(t *testing.T) {
	sID := UUIDBinary(uuid.New().String())
	createdAt := UTCTime(time.Now())
	//createdAt := LocalTimex(time.Now().Unix())
	updatedAt := UTCTime(time.Now())

	store := &Store{
		ID:                  sID,
		OriginID:            aws.String("123"),
		HistoryProductTags:  nil,
		CreatedAt:           &createdAt,
		UpdatedAt:           &updatedAt,
		CollectionNumberSeq: nil,
		Config:              nil,
	}
	st, err := json.Marshal(store)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(st))

	sts := Store{}
	if err := json.Unmarshal(st, &sts); err != nil {
		panic(err)
	}
}
