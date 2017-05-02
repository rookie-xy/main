package types

type Array struct {
    data    []interface{}
    length  int
}

func NewArray(length int) Array {
    return Array{
        data: make([]interface{}, length),
        length: length,
    }
}

func (r Array) SetData(i int, data interface{}) int {
    if data == nil || i < 0 {
        return Error
    }

    r.data[i] = data

    return Ok
}

func (r Array) GetData(i int) interface{} {
    return r.data[i]
}

func (r Array) SetLength(length int) int {
    if length < 0 {
        return Error
    }

    r.length = length

    return Ok
}

func (r Array) GetLength() int {
    return r.length
}
