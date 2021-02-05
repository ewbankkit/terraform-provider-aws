package namevaluesfilters

import (
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/experimental/playground/notfound/service"
)

type NameValuesFilters map[string][]string

func New(i interface{}) NameValuesFilters {
	return make(NameValuesFilters).Add(i)
}

func (filters NameValuesFilters) Add(i interface{}) NameValuesFilters {
	return filters
}

func (filters NameValuesFilters) ServiceFilters() []*service.Filter {
	return []*service.Filter{}
}
