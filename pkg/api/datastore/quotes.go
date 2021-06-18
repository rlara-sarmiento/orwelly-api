package datastore

import "github.com/rlara-sarmiento/orwelly-api/pkg/api/model"

type Quotes interface {
	Create(id int, obj model.Quote) error
	Get(id int) (model.Quote, error)
}
