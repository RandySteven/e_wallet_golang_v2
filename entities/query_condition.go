package entities

type QueryCondition struct {
	SortedBy string `form:"sortBy"`
	Sort     string `form:"sort"`
	Limit    string `form:"limit"`
	Page     string `form:"page"`
}
