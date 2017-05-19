package elasticsearch

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

// Contextの初期化とelasticsearchのコネクションを確立します。
func Elastic() (*elastic.Client, context.Context) {
	ctx := context.Background()

	client, err := elastic.NewClient(
		// elasticsearchのdocker内IPアドレス
		elastic.SetURL("127.0.0.1:9200"),
		// これつけないとつながらない?
		// https://github.com/olivere/elastic/issues/58#issuecomment-156052782
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	return client, ctx
}
