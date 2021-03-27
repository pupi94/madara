package testing

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pupi94/madara/testing/es_mapping"
)

func InitEs() {

}

func ClearEs() {

}

func ClearIndex() {

}

func CreateIndex(client *elastic.Client) {
	ctx := context.Background()
	_, err := client.CreateIndex("products").BodyString(es_mapping.Product).Do(ctx)
	if err != nil {
		panic(err)
	}
}

func IndexDoc() {

}

func SearchDocs(docs interface{}, query elastic.Query) (exist bool) {

	return true
}

func GetDoc(id int64) (exist bool) {
	return true
}
