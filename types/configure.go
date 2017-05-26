/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
    "sync"
)

var (
    ConfigOk    =  0
    ConfigError = -1
)

type Configure_t struct {
    *File_t
    *Codec_t

     sync.Mutex

     commandType  int
     moduleType   int64
     value        interface{}
     Event        chan *Event_t

     Channels      []Channel

     Configure
     Cycle
}

func NewConfigure(log *Log_t) *Configure_t {
    return &Configure_t{
        File_t:  NewFile(log),
        Event : make(chan *Event_t),
    }
}

func (c *Configure_t) SetModuleType(moduleType int64) int {
    if moduleType <= 0 {
        return Error
    }

    c.moduleType = moduleType

    return Ok
}

func (c *Configure_t) SetCommandType(commandType int) int {
    if commandType <= 0 {
        return Error
    }

    c.commandType = commandType

    return Ok
}

func (c *Configure_t) SetValue(value interface{}) int {
    if value == nil {
        return Error
    }

    return Ok
}

func (r *Configure_t) GetValue() interface{} {
    return r.value
}

func (r *Configure_t) NewConfigure(c Configure) int {
    if c == nil {
       return Error
    }

    r.Configure = c

    return Ok
}

func (r *Configure_t) SetConfigure() int {
    if handler := r.Configure; handler != nil {
        return handler.Set()
    }

    r.Warn("set configure not found")

    return Error
}

func (r *Configure_t) GetConfigure() int {
    if handler := r.Configure; handler != nil {
        return handler.Get()
    }

    r.Warn("get configure not found")

    return Error
}


func (r *Configure_t) Materialized(modules []Module) int {
    if r.value == nil {
        content := r.GetBytes()

        if content == nil {
            r.Error("configure content: %s, size: %d\n",
                      content, r.GetSize())

            return Error
        }

        var e error

        if r.value, e = r.Decode(content); e != nil {
            return Error
        }
    }

    switch v := r.value.(type) {

    case []interface{}:
        for _, value := range v {
            r.value = value
            r.Materialized(modules)
        }

    case map[interface{}]interface{}:
        if r.MapContext(v, modules) == Error {
            return Error
        }

    default:
        r.Warn("unknown")
    }

    return Ok
}

func (r *Configure_t) MapContext(materialized map[interface{}]interface{}, m []Module) int {
    flag := Ok

    modules := GetModules(m, r.moduleType)
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
                        data = handle.Get()[module.CtxIndex];
                        if data == nil {
                            return Error
                        }
                    }

                    r.value = value

                    command.Set(r, &command, data)
                }
            }
        }
    }

    if flag == Error {
        return ConfigError
    }

    return ConfigOk
}

func Block(c *Configure_t, m []Module, modType int64, cfgType int) int {
    for _, v := range m {
        if v != nil {
            if self := v.Self(); self != nil {
                if self.Type != modType {
                    continue
                }

                if handle := self.Context; handle != nil {
                    if this := handle.Set(); this != nil {
                        self.CtxIndex++
                        handle.Get()[self.CtxIndex] = &this
                    } else {
                        return Error
                    }

                } else {
                    continue
                }
            }
        }
    }

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

func SetFlag(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int {
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

func SetString(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int {
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

func SetNumber(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int {
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


func SetArray(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*Array_t)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

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

func SetChannel(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(c, Modules, CHANNEL_MODULE, flag)

    return Ok
}

func SetInput(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(c, Modules, INPUT_MODULE, flag)

    return Ok
}

func SetOutput(c *Configure_t, _ *Command_t, _ *unsafe.Pointer) int {
    if nil == c {
        return Error
    }

    flag := USER_CONFIG|CONFIG_ARRAY
    Block(c, Modules, OUTPUT_MODULE, flag)

    return Ok
}
