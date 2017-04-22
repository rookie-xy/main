/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "os"

type File struct {
    *Log
    *os.File

     name     string
     size     int64
     bytes  []byte

     action   IO
}

type IO interface {
    Open(name string) int
    Closer() int
    Reader() int
    Writer() int
}

func NewFile(log *Log) *File {
    return &File{
        Log  : log,
        File : os.Stdout,
    }
}

func (f *File) SetName(name string) int {
    if name == "" {
        return Error
    }

    f.name = name

    return Ok
}

func (f *File) GetName() string {
    return f.name
}

func (f *File) SetSize(size int64) int {
    if size < 0 {
        return Error
    }

    f.size = size

    return Ok
}

func (f *File) GetSize() int64 {
    return f.size
}

func (f *File) SetBytes(bytes []byte) int {
    if bytes == nil {
        return Error
    }

    f.bytes = bytes

    return Ok
}

func (f *File) GetBytes() []byte {
    return f.bytes
}

func (f *File) Open(name string) int {
    var error error

    f.File, error = os.OpenFile(name, os.O_RDWR, 0777)
    if error != nil {
        f.Info("open file error: %s", error)
        return Error
    }

    stat, error := f.Stat()
    if error != nil {
        f.Info("stat file error: %s", error)
        return Error
    }

    f.size = stat.Size()

    return Ok
}

func (f *File) Closer() int {
    if error := f.Close(); error != nil {
        f.Info("close file error: %s", error)
        return Error
    }

    return Ok
}

func (f *File) Reader() int {
    var char []byte

    if size := f.size; size <= 0 {
        f.Error("file size is: %d\n", size)
        return Error
    } else {
        char = make([]byte, size)
    }

    _, error := f.Read(char)
    if error != nil {
        f.Error("file read error: %s", error)
        return Error
    }

    f.bytes = char

    return Ok
}

func (f *File) Writer() int {
    return Ok
}
