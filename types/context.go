/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type ContextCreateFunc func(c *Configure) unsafe.Pointer
type ContextInitFunc func(c *Configure, configure *unsafe.Pointer) string

type Context struct {
    Name    String
    Create  ContextCreateFunc
    Init    ContextInitFunc
}
/*
type Contextable interface {
    Create() int
    Insert() int
}
*/

func Block(c *Configure, m []*Moduleable, modType int64, cfgType int) int {
    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != modType {
            continue
        }

        context := (*Context)(unsafe.Pointer(module.Context))
        if context == nil {
            continue
        }

        if handle := context.Create; handle != nil {
            this := handle(c)
            if cycle.SetContext(module.Index, &this) == Error {
                return Error
            }
        }
    }
/*
    configure := cycle.GetConfigure()
    if configure == nil {
        return Error
    }
    */

    if c.SetModuleType(modType) == Error {
        return Error
    }

    if c.SetCommandType(cfgType) == Error {
        return Error
    }

    if c.Materialized(modules) == Error {
        return Error
    }

    for m := 0; modules[m] != nil; m++ {
        module := modules[m]
        if module.Type != modType {
            continue
        }

        this := (*Context)(unsafe.Pointer(module.Context))
        if this == nil {
            continue
        }

        context := cycle.GetContext(module.Index)
        if context == nil {
            continue
        }

        if init := this.Init; init != nil {
            if init(cycle, context) == "-1" {
                return Error
            }
        }
    }

    return Ok
}
