package elasticsearch

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

// Contextの初期化とelasticsearchのコネクションを確立します。
func Elastic() (*elastic.Client, context.Context) {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}
	return client, ctx
}
