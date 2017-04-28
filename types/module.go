package types

type Module struct {
    Index      int
    CtxIndex   int
    Context    Contexter
    Commands   []Command
    Type       int64
}

var Modulers []Moduler

func (m *Module) Self() *Module {
    return m
}

func Load(modulers []Moduler, moduler Moduler) []Moduler {
    if modulers == nil && moduler == nil {
        return nil
    }

    modulers = append(modulers, moduler)

    return modulers
}

func (m *Module) SetIndex(i int) {
    m.Index = i
}

func GetSomeModules(m []Moduler, modType int64) []Moduler {
    var modulers []Moduler

    for i := 0; m[i] != nil; i++ {
        module := m[i].Type()

        if module.Type == modType {
            modulers = Load(modulers, m[i])
        }
    }

    modulers = Load(modulers, nil)

    return modulers
}

func GetSpacModules(m []Moduler) []Moduler {
    var modulers []Moduler

    for i := 0; m[i] != nil; i++ {
        module := m[i].Type()

        if module.Type == SYSTEM_MODULE ||
           module.Type == CONFIG_MODULE {
            continue
        }

        modulers = Load(modulers, m[i])
    }

    modulers = Load(modulers, nil)

    return modulers
}

func GetPartModules(m []Moduler, modType int64) []Moduler {
    if m == nil || len(m) <= 0 {
        return nil
    }

    switch modType {

    case SYSTEM_MODULE:
        modulers := GetSomeModules(m, modType)
        if modulers != nil {
            return modulers
        }

    case CONFIG_MODULE:
        modulers := GetSomeModules(m, modType)
        if modulers != nil {
            return modulers
        }
    }

    var modulers []Moduler

    modType = modType >> 28

    for i := 0; m[i] != nil; i++ {
        module := m[i].Type()
        moduleType := module.Type >> 28

        if moduleType == modType {
            modulers = Load(modulers, m[i])
        }
    }

    modulers = Load(modulers, nil)

    return modulers
}
