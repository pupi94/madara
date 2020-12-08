package producers

import (
	"encoding/json"
	"github.com/pupi94/madara/config"
)

func SyncImage() error {
	m := map[string]string{"src": "https://google.com/image.jpg", "alt": "test"}
	bs, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return config.AsyncProducer.Produce(config.Env.ImageTopicName, bs, 1)
}
