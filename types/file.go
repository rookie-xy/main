/*
 * Copyright (C) 2016 Meng Shi
 */

package types

import "os"

type File_t struct {
    *Log_t
    *os.File

     name     string
     size     int64
     bytes  []byte

     Filer
}

func NewFile(log *Log_t) *File_t {
    return &File_t{
        Log_t: log,
        File : os.Stdout,
    }
}

func (f *File_t) SetName(name string) int {
    if name == "" {
        return Error
    }

    f.name = name

    return Ok
}

func (f *File_t) GetName() string {
    return f.name
}

func (f *File_t) SetSize(size int64) int {
    if size < 0 {
        return Error
    }

    f.size = size

    return Ok
}

func (f *File_t) GetSize() int64 {
    return f.size
}

func (f *File_t) SetBytes(bytes []byte) int {
    if bytes == nil {
        return Error
    }

    f.bytes = bytes

    return Ok
}

func (f *File_t) GetBytes() []byte {
    return f.bytes
}

func (f *File_t) Open(name string) int {
    var error error

    f.File, error = os.OpenFile(name, os.O_RDWR, 0777)
    f.File.Stat()
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

func (f *File_t) Closer() int {
    if error := f.Close(); error != nil {
        f.Info("close file error: %s", error)
        return Error
    }

    return Ok
}

func (f *File_t) Reader() int {
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

func (f *File_t) Writer() int {
    return Ok
}
