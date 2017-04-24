/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "time"

type Event struct {
    id         int64
    name       string

    magic      uint8   /* message/notice/datagram */
    flag       uint8   /* 有无偏移， 需不需应答 */
    option     uint8   /* data package/control package */
    opcode     uint8   /* 操作 */
    offset     uint32
    length     uint64

    serial     int32   /* 一个event分布两个里面 */
    timestamp  int64   /* 合并event */

    data       []byte

    This       chan *Event
}

func NewEvent() *Event {
    return &Event{
        name:      "event",
        timestamp: time.Now().Unix(),
        This:      make(chan *Event),
    }
}

func (e *Event) SetName(name string) int {
    if name == "" {
        return Error
    }

    e.name = name

    return Ok
}

func (e *Event) GetName() string {
    return e.name
}

func (e *Event) SetMagic(magic uint8) {
    e.magic = magic
}

func (e *Event) GetMagic() uint8 {
    return e.magic
}

func (e *Event) SetFlag(flag uint8) {
    e.flag = flag
}

func (e *Event) GetFlag() uint8 {
    return e.flag
}

func (e *Event) SetOption(option uint8) {
    e.option = option
}

func (e *Event) GetOption() uint8 {
    return e.option
}

func (e *Event) SetOffset(offset uint32) {
    e.offset = offset
}

func (e *Event) GetOffset() uint32 {
    return e.offset
}

func (e *Event) SetLength(length uint64) {
    e.length = length
}

func (e *Event) GetLength() uint64 {
    return e.length
}

func (e *Event) SetOpcode(opcode uint8) {
    e.opcode = opcode
}

func (e *Event) GetOpcode() uint8 {
    return e.opcode
}
