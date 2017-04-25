package types

import "unsafe"

type Moduler interface {
    Init(o *Option) int
    Main(c *Configure) int
    Exit() int
    Type() *Module
}

type Channeler interface {
    Push(e *Event) int
    Pull() *Event
}

type Configurer interface {
    GetConfigure() int
    SetConfigure() int
}

type Parser interface {
    Marshal(in interface{}) ([]byte, error)
    Unmarshal(in []byte, out interface{}) int
}

type Contexter interface {
    Create() unsafe.Pointer
    Insert(p *unsafe.Pointer) int

    Contexts() *Context
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

type Stringer struct {

}
