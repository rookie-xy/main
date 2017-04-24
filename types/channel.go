package types


type Channel struct {
    name string

    Channelable
}

type Channelable interface {
    Push(e *Event) int
    Pull() *Event
}

func NewChannel() *Channel {
    return &Channel{}
}

func (c *Channel) Push(e *Event) int {
    if method := c.Channelable; method != nil {
        return method.Push(e)
    }

    return Ok
}

func (c *Channel) Pull() *Event {
    if method := c.Channelable; method != nil {
        return method.Pull()
    }

    return nil
}
