/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "errors"
)

var (
    ConfigOk    =  0
    ConfigError = -1
)

type Configure struct {
    *File

     commandType  int
     moduleType   int64
     value        interface{}

     Event        chan *Event

     Channeler

     Parser
     Configurer
}

func NewConfigure(log *Log) *Configure {
    return &Configure{
        File:  NewFile(log),
        Event: make(chan *Event),
    }
}

func (c *Configure) SetModuleType(moduleType int64) int {
    if moduleType <= 0 {
        return Error
    }

    c.moduleType = moduleType

    return Ok
}

func (c *Configure) SetCommandType(commandType int) int {
    if commandType <= 0 {
        return Error
    }

    c.commandType = commandType

    return Ok
}

func (c *Configure) SetValue(value interface{}) int {
    if value == nil {
        return Error
    }

    return Ok
}

func (c *Configure) GetValue() interface{} {
    return c.value
}

func (c *Configure) NewConfigurer(cr Configurer) int {
    if cr == nil {
       return Error
    }

    c.Configurer = cr

    return Ok
}

func (c *Configure) SetConfigure() int {
    if handler := c.Configurer; handler != nil {
        return handler.SetConfigure()
    }

    c.Warn("set configure not found")

    return Error
}

func (c *Configure) GetConfigure() int {
    if handler := c.Configurer; handler != nil {
        return handler.GetConfigure()
    }

    c.Warn("get configure not found")

    return Error
}

func (c *Configure) NewParser(parser Parser) int {
    if parser == nil {
       return Error
    }

    c.Parser = parser

    return Ok
}

func (c *Configure) Marshal(in interface{}) ([]byte, error) {
    if handler := c.Parser; handler != nil {
        return handler.Marshal(in)
    }

    return nil, errors.New("Marshal default not found")
}

func (c *Configure) Unmarshal(in []byte, out interface{}) int {
    if handler := c.Parser; handler != nil {
        return handler.Unmarshal(in, out)
    }

    c.Warn("Unmarshal default not found")

    return Error
}

func (c *Configure) Materialized(modules []Moduler) int {
    if c.value == nil {
        content := c.GetBytes()

        if content == nil {
            /*
            c.Error("configure content: %s, filename: %s, size: %d\n",
                      content, c.GetFileName(), c.GetSize())
                      */
            c.Error("configure content: %s, size: %d\n",
                      content, c.GetSize())

            return Error
        }

        if c.Parser.Unmarshal(content, &c.value) == Error {
            return Error
        }
    }

    switch v := c.value.(type) {

    case []interface{} :
        for _, value := range v {
            c.value = value
            c.Materialized(/*cycle, */modules)
        }

    case map[interface{}]interface{}:
        if c.doParse(v, /*cycle,*/ modules) == Error {
            return Error
        }

    default:
        c.Warn("unknown")
    }

    return Ok
}

func (c *Configure) doParse(materialized map[interface{}]interface{}, m []Moduler) int {
    flag := Ok

    modules := GetPartModules(m, c.moduleType)
    if modules == nil {
        return Error
    }

    for key, value := range materialized {
        if key != nil && value != nil {
            flag = Ok
        }

        name := key.(string)
        found := false

        for m := 0; flag != Error && !found && modules[m] != nil; m++ {
            module := modules[m].Type()

            commands := module.Commands
            if commands == nil {
                continue;
            }

            for i := 0; commands[i].Name.Len != 0; i++ {

                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                				found = true

                    //context := cycle.GetContext(module.Index)
                    var data *unsafe.Pointer
                    if handle := module.Context; handle != nil {
                        if context := handle.Contexts(); context != nil {
                            if this := context.GetData(module.Index); this != nil {
                                data = this
                            }
                        }
                    }

                    c.value = value

                    command.Set(c, &command, data)
                }
            }
        }
    }

    if flag == Error {
        return ConfigError
    }

    return ConfigOk
}

func (c *Configure) Block(module int64, config int) int {
    /*
    if Block(Modules, module, config) == Error {
        return Error
    }
    */

    return Ok
}

func SetFlag(cfg *Configure, cmd *Command, p *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || p == nil {
        return Error
    }

    field := (*bool)(unsafe.Pointer(uintptr(*p) + cmd.Offset))
    if field == nil {
        return Error
    }

    flag := cfg.GetValue()
    if flag == true {
        *field = true
    } else if flag == false {
        *field = false
    } else {
        return Error
    }

    /*
    if command.Post != nil {
        post := command.Post.(DvrConfPostType);
        post.Handler(cf, post, *p);
    }
    */

    return Ok
}

func SetString(cfg *Configure, cmd *Command, p *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || p == nil {
        return Error
    }

    field := (*string)(unsafe.Pointer(uintptr(*p) + cmd.Offset))
    if field == nil {
        return Error
    }

    strings := cfg.GetValue()
    if strings == nil {
        return Error
    }

    *field = strings.(string)

    return Ok
}

func SetNumber(cfg *Configure, cmd *Command, p *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || p == nil {
        return Error
    }

    field := (*int)(unsafe.Pointer(uintptr(*p) + cmd.Offset))
    if field == nil {
        return Error
    }

    number := cfg.GetValue()
    if number == nil {
        return Error
    }

    *field = number.(int)

    return Error
}


