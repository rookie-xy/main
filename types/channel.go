package types

type Channel_t struct {
    name string

    Channel
}

func NewChannel(ch Channel) *Channel_t {
    return &Channel_t{
        "channel",
        ch,
    }
}
/*
func (c *Channel_t) Push(e *Event) int {
    if handler := c.Channel; handler != nil {
        return handler.Push()
    }

    return Ok
}

func (c *Channel_t) Pull() *Event {
    if handler := c.Channel; handler != nil {
        return handler.Pull()
    }

    return nil
}
*/
