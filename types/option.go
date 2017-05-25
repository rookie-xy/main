/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type Option_t struct {
    *Log_t
    *Configure_t

     argc   int
     argv   []string
     items  map[string]interface{}

     option Option
}

func NewOption(log *Log_t) *Option_t {
    return &Option_t{
        Log_t: log,
        items: make(map[string]interface{}),
    }
}

func (o *Option_t) GetArgc() int {
    return o.argc
}

func (o *Option_t) GetArgv() []string {
    return o.argv
}

func (o *Option_t) SetArgs(argc int, argv []string) int {
    if argc <= 0 || argv == nil {
        return Error
    }

    o.argc = argc
    o.argv = argv

    return Ok
}

func (o *Option_t) SetItem(k string, v interface{}) {
    o.items[k] = v
}

func (o *Option_t) GetItem(k string) interface{} {
    return o.items[k]
}
