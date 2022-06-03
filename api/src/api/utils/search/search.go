package search

import (
	"github.com/emikohmann/shopping-cart/api/src/api/domain/search"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const (
	keywordLimit    = "limit"
	keywordOffset   = "offset"
	valuesSeparator = ","
)

func ParseQuery(ctx *gin.Context) search.Query {
	data := make(search.Data)
	for key, values := range ctx.Request.URL.Query() {
		if key == keywordLimit || key == keywordOffset {
			continue
		}
		for _, value := range values {
			if strings.Contains(value, valuesSeparator) {
				for _, subValue := range strings.Split(value, valuesSeparator) {
					data[key] = append(data[key], subValue)
				}
			} else {
				data[key] = append(data[key], value)
			}
		}
	}
	limit, offset := 20, 0
	queryLimit, err := strconv.Atoi(ctx.Query(keywordLimit))
	if err == nil {
		limit = queryLimit
	}
	queryOffset, err := strconv.Atoi(ctx.Query(keywordOffset))
	if err == nil {
		offset = queryOffset
	}
	return search.Query{
		Limit:  limit,
		Offset: offset,
		Data:   data,
	}
}
