package queries

import "github.com/olivere/elastic"

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}
type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQueries...)
	return query
}
