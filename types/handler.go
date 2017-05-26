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
    Set() unsafe.Pointer
    Get() []*unsafe.Pointer
}

type Channel interface {
    New() Channel
    Init(configure interface{}) int

    Push() int
    Pull() int

    //Publish(e *Event) int
    //Subscribe() *Event
    //Filter
    //Codec
}

type Input interface {

}

type Output interface {

}

type Codec interface {
    New() Codec
    Init(configure interface{}) int
    Encode(in interface{}) (interface{}, error)
    Decode(in []byte) (interface{}, error)
    Type(name string) int
}

type Filter interface {
    New() Filter
    Init(configure interface{}) int
    Washing(in []byte) (interface{}, error)
    Type(name string) int
}

type Configure interface {
    Get() int
    Set() int
}

type Cycle interface {
    Start(c *Configure_t) int
    Stop() int
}

type Event interface {
    JsonEncode()
    JsonDecode()
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
