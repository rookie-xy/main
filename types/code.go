package types

type Code struct {
    name string
    Codec
}

func NewCode(c Codec) *Code {
    return &Code{
        "code",
        c,
    }
}
