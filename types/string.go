/*
 * Copyright (C) 2016 Meng Shi
 */

package types

type String struct {
    Len int
    Data interface{} 
}

var NilString = String{ 0, nil }
