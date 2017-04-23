package types


type Channel struct {
    name string
}

type Channelable interface {
    Push() int
    Pull() int
}

func NewChannel() *Channel {
    return &Channel{}
}
