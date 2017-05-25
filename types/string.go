/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type String_t struct {
    Len int
    Data interface{} 
}

var NilString = String_t{ 0, nil }
