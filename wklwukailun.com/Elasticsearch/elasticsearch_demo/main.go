package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Person{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().Index("student").Type("_doc").BodyJson(p1).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

/*
使用postman查询
GET  127.0.0.1:9200/student/_search/

json
{
    "query":{
        "match":{
            "married":false
        }
    },
    "size":1
}
返回

{
    "took": 4,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 1,
            "relation": "eq"
        },
        "max_score": 0.2876821,
        "hits": [
            {
                "_index": "student",
                "_id": "_kIQF4IBKSVsm2-vl8Uy",
                "_score": 0.2876821,
                "_source": {
                    "name": "rion",
                    "age": 22,
                    "married": false
                }
            }
        ]
    }
}
*/
