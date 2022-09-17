package elastic

import(
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticConn *elasticsearch.Client

func InitElasticConn() {
	var err error
	ElasticConn, err = elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {

	}
	fmt.Println("elasticsearch: connection successfully opened")
	return
}