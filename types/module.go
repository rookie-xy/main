package types


type Module struct {
    CtxIndex   int
    Index      int
    Context    Contextable
    Commands   []Command
    Type       int64
}

type Moduleable interface {
    Init(o *Option) int
    Main(c *Configure) int
    Exit() int
    Type() *Module
}

var Modables []Moduleable

func (m *Module) Self() *Module {
    return m
}

func Load(modules []Moduleable, module Moduleable) []Moduleable {
    if modules == nil && module == nil {
        return nil
    }

    modules = append(modules, module)

    return modules
}


func (m *Module) SetIndex(i int) {
    m.Index = i
}

func GetSomeModules(modable []Moduleable, modType int64) []Moduleable {
    var modules []Moduleable

    for i := 0; modable[i] != nil; i++ {
        module := modable[i].Type()

        if module.Type == modType {
            modules = Load(modules, modable[i])
        }
    }

    modules = Load(modules, nil)

    return modules
}

func GetSpacModules(modables []Moduleable) []Moduleable {
    var modules []Moduleable

    for i := 0; modables[i] != nil; i++ {
        module := modables[i].Type()

        if module.Type == SYSTEM_MODULE ||
           module.Type == CONFIG_MODULE {
            continue
        }

        modules = Load(modules, modables[i])
    }

    modules = Load(modules, nil)

    return modules
}

func GetPartModules(modables []Moduleable, modType int64) []Moduleable {
    if modables == nil || len(modables) <= 0 {
        return nil
    }

    switch modType {

    case SYSTEM_MODULE:
        modules := GetSomeModules(modables, modType)
        if modules != nil {
            return modules
        }

    case CONFIG_MODULE:
        modules := GetSomeModules(modables, modType)
        if modules != nil {
            return modules
        }
    }

    var modules []Moduleable

    modType = modType >> 28

    for i := 0; modables[i] != nil; i++ {
        module := modables[i].Type()
        moduleType := module.Type >> 28

        if moduleType == modType {
            modules = Load(modules, modables[i])
        }
    }

    modules = Load(modules, nil)

    return modules
}
