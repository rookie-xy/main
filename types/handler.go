package types

import "unsafe"

type Moduler interface {
    Init(o *Option) int
    Main(c *Configure) int
    Exit() int
    Self() *Module
}

type Context interface {
    Create() unsafe.Pointer
    GetDatas() []*unsafe.Pointer
}

type Channeler interface {
    Push(e *Event) int
    Pull() *Event
}

type Configurer interface {
    GetConfigure() int
    SetConfigure() int
}

type Cycler interface {
    Start(c *Configure, name string) int
    Stop() int
}

type Eventer interface {
    JsonEncode()
    JsonDecode()
}

type Filer interface {
    Open(name string) int
    Closer() int
    Reader() int
    Writer() int
}

type Loger interface {
    Dump() int
}

type Optioner interface {
    Parser() int
}

type Stringer interface {

}

type Filter interface {

}

type Codec interface {
    New() Codec
    Init(configure interface{}) int
    Encode(in interface{}) (interface{}, error)
    Decode(in []byte) (interface{}, error)
    Type(name string) int
}