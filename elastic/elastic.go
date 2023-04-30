package elastic

import(
	"log"
	"github.com/elastic/go-elasticsearch/v8"
	"command-event-handler-service/models"
)

var ElasticConn *elasticsearch.Client

func InitElasticConn() {
	var err error
	ElasticConn, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: models.ElasticConf.Addrs,
		Username: models.ElasticConf.Username,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("elasticsearch: connection successfully opened")
	return
}