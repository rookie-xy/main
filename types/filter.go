package types

type Filter_t struct {
    name string
    Filter
}

func NewFilter(f Filter) *Filter_t {
    return &Filter_t{
        "filter",
        f,
    }
}
