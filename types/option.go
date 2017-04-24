/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "fmt"

type Option struct {
    *Log
    *Configure

     argc   int
     argv   []string
     items  map[string]interface{}

     option Optionable
}

type Optionable interface {
    Parser() int
}

func NewOption(log *Log) *Option {
    return &Option{
        Log:   log,
        items: make(map[string]interface{}),
    }
}

func (o *Option) GetArgc() int {
    return o.argc
}

func (o *Option) GetArgv() []string {
    return o.argv
}

func (o *Option) SetArgs(argc int, argv []string) int {
    if argc <= 0 || argv == nil {
        return Error
    }

    o.argc = argc
    o.argv = argv

    return Ok
}

func (o *Option) SetItem(k string, v interface{}) {
    o.items[k] = v
}

func (o *Option) GetItem(k string) interface{} {
    return o.items[k]
}

func (o *Option) Set(option Optionable) int {
    if option == nil {
        return Error
    }

    o.option = option

    return Ok
}

func (o *Option) Get() Optionable {
    return o.option
}

func (o *Option) Materialized() int {
    if option := o.Get(); option != nil {
        if option.Parser() == Error {
            return Error
        }

        return Ok
    }

    // default method
    if o.Parser() == Error {
        return Error
    }

    return Ok
}

func (o *Option) Parser() int {
    fmt.Println("option type")
    return Ok
}
