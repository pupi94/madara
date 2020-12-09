package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pupi94/madara/config"
	"github.com/sirupsen/logrus"
	"net/url"
)

type Image struct {
	StoreID string `json:"store_id"`
	Host    string `json:"host"`
	Src     string `json:"src"`
}

func main() {
	config.InitKafkaProducer(context.Background())

	list := [][]string{
		{"11000", "https://www.gudaoimages.com/YK03800-OR-1.jpg"},
		{"23005", "https://www.gudaoimages.com/YK03800-OR-2.jpg"},
		{"23005", "https://www.gudaoimages.com/YK03800-OR-2.jpg"},
	}

	for _, item := range list {
		img := &Image{StoreID: item[0], Src: item[1]}
		uri, err := url.ParseRequestURI(item[1])
		if err != nil {
			logrus.Panic("ParseRequestURI Fail: ", err)
		}
		img.Host = uri.Host

		bs, err := json.Marshal(img)
		if err != nil {
			logrus.Panic("Marshal Fail: ", err)
		}

		err = config.AsyncProducer.Produce("image", bs, fmt.Sprintf("%s:%s", img.StoreID, img.Host))
		if err != nil {
			logrus.Panic("AsyncProducer Fail: ", err)
		}
	}
}
