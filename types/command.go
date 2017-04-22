/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "unsafe"

type SetFunc func(cfg *Configure, cmd *Command, p *unsafe.Pointer) int

type Command struct {
    Name    String
    Type    int
    Set     SetFunc
    Conf    int
    Offset  uintptr
    Post    interface{}
}

var NilCommand = Command{ NilString, 0, nil, 0, 0, nil }
