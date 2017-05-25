package types

type Array_t struct {
    data    []interface{}
    length  int
}

func NewArray(length int) Array_t {
    return Array_t{
        data: make([]interface{}, length),
        length: length,
    }
}

func (r Array_t) SetData(i int, data interface{}) int {
    if data == nil || i < 0 {
        return Error
    }

    r.data[i] = data

    return Ok
}

func (r Array_t) GetData(i int) interface{} {
    return r.data[i]
}

func (r Array_t) SetLength(length int) int {
    if length < 0 {
        return Error
    }

    r.length = length

    return Ok
}

func (r Array_t) GetLength() int {
    return r.length
}
