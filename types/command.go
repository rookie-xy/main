/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

type SetFunc func(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int

type Command_t struct {
    Name    String_t
    Type    int
    Set     SetFunc
    Conf    int
    Offset  uintptr
    Post    interface{}
}

var NilCommand = Command_t{ NilString, 0, nil, 0, 0, nil }
