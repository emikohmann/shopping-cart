package search

type Result struct {
	Paging  Paging      `json:"paging"`
	Results interface{} `json:"results"`
}

type Paging struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Query struct {
	Limit  int
	Offset int
	Data   Data
}

type Data map[string][]interface{}
