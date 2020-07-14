package drivers

import elastic "gopkg.in/olivere/elastic.v7"

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	//fmt.Println("ES initialized...")
	return client, err
}
