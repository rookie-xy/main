package types

import "context"

type Context_t struct {
    context.Context
    Kill context.CancelFunc
}

func NewContext() Context_t {
    return &Context_t{}
}

func (r *Context_t) Self() *Context_t {
    return r
}

func CreateBackgroundContext() context.Context {
    return context.Background()
}

func CreateTodoContext() context.Context {
    return context.TODO()
}

func (r *Context_t) WithCancel(c context.Context) int {
    if this, kill := context.WithCancel(c); this != nil {
        r.Context = this
        r.Kill = kill

    } else {
        return Error
    }

    return Ok
}

func (r *Context_t) Get() int {
    return Ok
}
