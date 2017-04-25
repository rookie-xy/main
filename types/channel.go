package types


type Channel struct {
    name string

    Channeler
}

func NewChannel(ch Channeler) *Channel {
    return &Channel{
        "channel",
        ch,
    }
}

func (c *Channel) Push(e *Event) int {
    if handler := c.Channeler; handler != nil {
        return handler.Push(e)
    }

    return Ok
}

func (c *Channel) Pull() *Event {
    if handler := c.Channeler; handler != nil {
        return handler.Pull()
    }

    return nil
}
