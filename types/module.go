package types


type Module struct {
    CtxIndex   uint
    Index      uint
    Context    Contextable
    Commands   []Command
    Type       int64
}

type Moduleable interface {
    Init(c *Configure) int
    Main(c *Channel) int
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

func Init(modules []Moduleable, c *Configure) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Init(c)
    }
}

func Main(modules []Moduleable, c *Channel) {
    for i := 0; modules[i] != nil; i++ {
        module := modules[i]
        module.Main(c)
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
