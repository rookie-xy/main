/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
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
/*
     Event        chan *Event

     Handle
     */
     Parser
}

type Parser interface {
    Marshal(in interface{}) ([]byte, error)
    Unmarshal(in []byte, out interface{}) int
}

func NewConfigure(log *Log) *Configure {
    return &Configure{
        File:  NewFile(log),
        //Event: make(chan *Event),
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
/*
func (c *Configure) SetHandle(handle Handle) int {
    if handle == nil {
       return Error
    }

    c.Handle = handle

    return Ok
}

func (c *Configure) GetHandle() Handle {
    if c.Handle == nil {
        return nil
    }

    return c.Handle
}
*/

func (c *Configure) SetParser(parser Parser) int {
    if parser == nil {
       return Error
    }

    c.Parser = parser

    return Ok
}

func (c *Configure) GetParser() Parser {
    if c.Parser == nil {
        return nil
    }

    return c.Parser
}

/*
func (c *Configure) Materialized(cycle *Cycle, modules []*Module) int {
    if c.value == nil {
        content := c.GetBytes()
        if content == nil {
            log.Error("configure content: %s, filename: %s, size: %d\n",
                      content, c.GetFileName(), c.GetSize())
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
            c.Materialized(cycle, modules)
        }

    case map[interface{}]interface{}:
        if c.doParse(v, cycle, modules) == Error {
            return Error
        }

    default:
        c.Warn("unknown")
    }

    return Ok
}

func (c *Configure) doParse(materialized map[interface{}]interface{}, cycle *Cycle, m []*Module) int {
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
            module := modules[m]

            commands := module.Commands;
            if commands == nil {
                continue;
            }

            for i := 0; commands[i].Name.Len != 0; i++ {

                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                				found = true;

                    context := cycle.GetContext(module.Index)

                    c.value = value

																    if cycle.SetConfigure(c) == Error {
                        flag = Error
																				    break
                    }

                    command.Set(cycle, &command, context)
                }
            }
        }
    }

    if flag == Error {
        return ConfigError
    }

    return ConfigOk
}
*/

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

/* default impl */
func (c *Configure) Set() int {
    c.Warn("configure handle set")
    return Ok
}

func (c *Configure) Get() int {
    c.Warn("configure content set")
    return Ok
}

func (c *Configure) Marshal(in interface{}) ([]byte, error) {
    c.Warn("configure Marshal")
    return nil, nil
}

func (c *Configure) Unmarshal(in []byte, out interface{}) int {
    c.Warn("configure Unmarshal")
    return Ok
}
