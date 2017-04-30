package types

type Filtration struct {
    name string
    Filter
}

func NewFiltration(f Filter) *Filtration {
    return &Filtration{
        "filtration",
        f,
    }
}
