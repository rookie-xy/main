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
    *Code

     commandType  int
     moduleType   int64
     value        interface{}

     Event        chan *Event

     Channeler
     Filter

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


func (c *Configure) Materialized(modules []Moduler) int {
    if c.value == nil {
        content := c.GetBytes()

        if content == nil {
            c.Error("configure content: %s, size: %d\n",
                      content, c.GetSize())

            return Error
        }

        var e error

        if c.value, e = c.Decode(content); e != nil {
            return Error
        }
    }

    switch v := c.value.(type) {

    case []interface{}:
        for _, value := range v {
            c.value = value
            c.Materialized(modules)
        }

    case map[interface{}]interface{}:
        if c.doParse(v, modules) == Error {
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

        name  := key.(string)
        found := false

        for m := 0; flag != Error && !found && modules[m] != nil; m++ {
            module := modules[m].Self()

            commands := module.Commands
            if commands == nil {
                continue;
            }

            for i := 0; commands[i].Name.Len != 0; i++ {

                command := commands[i]

                if len(name) == command.Name.Len &&
                        name == command.Name.Data.(string) {

                				found = true

                    var data *unsafe.Pointer
                    if handle := module.Context; handle != nil {
                        data = handle.GetDatas()[module.CtxIndex];
                        if data == nil {
                            return Error
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

func Block(c *Configure, m []Moduler, modType int64, cfgType int) int {
    for _, v := range m {
        if v != nil {
            if self := v.Self(); self != nil {
                if self.Type != modType {
                    continue
                }

                if handle := self.Context; handle != nil {
                    if this := handle.Create(); this != nil {
                        self.CtxIndex++
                        handle.GetDatas()[self.CtxIndex] = &this

                    } else {
                        return Error
                    }

                } else {
                    continue
                }
            }
        }
    }

    /*
    for i := 0; m[i] != nil; i++ {
        module := m[i].Self()

        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if this := handle.Create(); this != nil {
                module.CtxIndex++
                handle.GetDatas()[module.CtxIndex] = &this

            } else {
                return Error
            }

        } else {
            continue
        }
    }
    */

    if c == nil {
        return Error
    }

    if c.SetModuleType(modType) == Error {
        return Error
    }

    if c.SetCommandType(cfgType) == Error {
        return Error
    }

    if c.Materialized(m) == Error {
        return Error
    }

    return Ok
}

func SetFlag(cfg *Configure, cmd *Command, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*bool)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

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

func SetString(cfg *Configure, cmd *Command, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*string)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

    value := cfg.GetValue()
    if value == nil {
        return Error
    }

    *field = value.(string)

    return Ok
}

func SetNumber(cfg *Configure, cmd *Command, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*int)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

    value := cfg.GetValue()
    if value == nil {
        return Error
    }

    *field = value.(int)

    return Error
}


func SetArray(cfg *Configure, cmd *Command, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*Array)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

    value := cfg.GetValue()
    if value == nil {
        return Error
    }

    values := value.([]interface{})
    length := len(values)

    array := NewArray(length)

    for k, v := range values {
        array.SetData(k, v)
    }

    *field = array

    return Ok
}

func SetChannel(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == cfg {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modulers, CHANNEL_MODULE, flag)

    return Ok
}

func SetInput(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == cfg {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modulers, INPUT_MODULE, flag)

    return Ok
}

func SetOutput(cfg *Configure, _ *Command, _ *unsafe.Pointer) int {
    if nil == cfg {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(cfg, Modulers, OUTPUT_MODULE, flag)

    return Ok
}
