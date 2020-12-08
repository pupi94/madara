package main

import (
	"context"
	"encoding/json"
	"github.com/pupi94/madara/config"
)

func main() {
	config.InitKafkaProducer(context.Background())
	m := map[string]string{"a": "1", "b": "2"}
	bs, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		err := config.AsyncProducer.Produce("image", bs, 1)
		if err != nil {
			panic(err)
		}
	}
}
