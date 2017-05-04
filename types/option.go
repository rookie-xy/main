/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type Option struct {
    *Log
    *Configure

     argc   int
     argv   []string
     items  map[string]interface{}

     option Optioner
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
