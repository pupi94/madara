package es_mapping

const Product = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"tags":{ "type":"keyword" },
			"location":{ "type":"geo_point" },
			"suggest_field":{ "type":"completion" }
		}
	}
}`
