package items

import (
	"github.com/emikohmann/shopping-cart/api/src/api/domain/items"
	"github.com/emikohmann/shopping-cart/api/src/api/domain/search"
	"github.com/emikohmann/shopping-cart/api/src/api/utils/apierrors"
)

type Service interface {
	Create(item items.Item) (items.Item, apierrors.APIError)
	Get(id int64) (items.Item, apierrors.APIError)
	Search(query search.Query) (search.Result, apierrors.APIError)
}
