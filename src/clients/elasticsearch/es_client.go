package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/prosline/pl_logger/logger"
	"time"
)
var (
	Client esClientInterface = &esClient{}
)
type esClientInterface interface{
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct{
	client *elastic.Client
}

func Init(){
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}
func (c *esClient) setClient(client *elastic.Client){
	c.client = client
}
func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}
