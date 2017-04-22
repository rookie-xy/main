package types

import "unsafe"

type Module struct {
    CtxIndex   uint
    Index      uint
    Context    unsafe.Pointer
    Commands   []Command
    Type       int64
}

type Moduleable interface {
    Init() int
    Main() int
    Exit() int
    Type() *Module
}

var Modules []Moduleable

func Load(modules []Moduleable, module Moduleable) []Moduleable {
    if modules == nil && module == nil {
        return nil
    }

    modules = append(modules, module)

    return modules
}

func Init(modules []Moduleable) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Init()
    }
}

func Main(modules []Moduleable) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Main()
    }
}

func Exit(modules []Moduleable) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Exit()
    }
}

func (m *Module) SetIndex(i uint) {
    m.Index = i
}
