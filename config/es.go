package config

func InitEsClient() {
	//client, err := elastic.NewClient(elastic.SetURL(Cfg.EsURL), elastic.SetSniff(false), elastic.SetHttpClient(&http.Client{
	//	Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
	//}))
	//if err != nil {
	//	fmt.Printf("Failed to create es_mapping client: %s \n", err)
	//	panic(err)
	//}
	//ESClient = client
}
