/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

type Module interface {
    Init(o *Option_t) int
    Main(c *Configure_t) int
    Exit() int
    Self() *Module_t
}

type Context interface {
    Builder() unsafe.Pointer
    Configure() []*unsafe.Pointer
    Self() *Context_t
}

type Channel interface {
    New() Channel
    Register(publish, subscribe string) int
    Push(in *Event) int
    Intercept() int
    Pull(out *Event) int
    Type(name string) int
}

type Input interface {
    Listen() int
    Reader() int
}

type Output interface {
    Writer(events Event) int
}

type Codec interface {
    New() Codec
    Init(configure interface{}) int
    Encode(in interface{}) (interface{}, error)
    Decode(out []byte) (interface{}, error)
    Type(name string) int
}

type Filter interface {
    New() Filter
    Init(configure interface{}) int
    Washing(in Event) (interface{}, error)
    Type(name string) int
}

type Configure interface {
    Get() int
    Set() int
}

type Cycle interface {
    New() Cycle
    Init(configure interface{}) int
    Start(p *unsafe.Pointer, c *Context_t) int
    Stop() int
    Type(name string) int
}

type Event interface {
    New() Event
    Init(configure interface{}) int
    Type(name string) int
}

type Filer interface {
    Open(name string) int
    Closer() int
    Reader() int
    Writer() int
}

type Log interface {
    Dump() int
}

type String interface {
    Len() int
}
