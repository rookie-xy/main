/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import (
    "unsafe"
)

type Context struct {
    Name   String
    Data   [128]*unsafe.Pointer
}

func NewContext() *Context {
    return &Context{
        Name: String{ len("context"), "context" },
    }
}

func (r *Context) Get() *Context {
    return r
}

func (r *Context) Set(c *Context) int {
    /*
    if c == nil {
        return Error
    }

    r = c
    */
    return Ok
}

func (c *Context) SetData(index int, p *unsafe.Pointer) int {
    c.Data[index] = p
    return Ok
}

func (c *Context) GetData(index int) *unsafe.Pointer {
    return c.Data[index]
}

func (c *Context) GetDatas() []*unsafe.Pointer {
    return c.Data[1:]
}

func Block(c *Configure, m []Moduler, modType int64, cfgType int) int {
    for i := 0; m[i] != nil; i++ {
        module := m[i].Type()

        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if this := handle.Create(); this != nil {
                if context := handle.Contexts(); context != nil {
                    module.CtxIndex++
                    if context.SetData(module.CtxIndex, &this) == Error {
                        return Error
                    }
                }

            } else {
                return Error
            }

        } else {
            continue
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

    /*
    for i := 0; m[i] != nil; i++ {
        module := m[i].Type()
        if module.Type != modType {
            continue
        }

        if handle := module.Context; handle != nil {
            if context := handle.Contexts(); context != nil {
                if this := context.GetData(module.CtxIndex); this != nil {
                    if handle.Insert(this) == Error {
                        return Error
                    }
                }
            }

        } else {
            continue
        }
    }
    */

    return Ok
}
