package types

import (
    "unsafe"
    "errors"
)

type Codec_t struct {
    name string
    Codec
}

func NewCodec(c Codec) Codec_t {
    return Codec_t{
        name: "codec",
        Codec: c,
    }
}

var Codecs []Codec

func SetCodec(cfg *Configure_t, cmd *Command_t, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*Codec_t)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))
    if field == nil {
        return Error
    }

    value := cfg.GetValue()
    if value == nil {
        return Error
    }

    values := value.(map[interface{}]interface{})

    var code Codec_t

    for k, v := range values {
        name := k.(string)

        for _, codec := range Codecs {
            if codec.Type(name) == Ignore {
                continue
            }

            codec.Init(v)

            code = NewCodec(codec)
            code.name = name
        }
    }

    *field = code

    return Ok
}

func (r Codec_t) Encode(in interface{}) (interface{}, error) {
    if handler := r.Codec; handler != nil {
        return handler.Encode(in)
    }

    return nil, errors.New("No found handler")
}

func (r Codec_t) Decode(in []byte) (interface{}, error) {
    if handler := r.Codec; handler != nil {
        return handler.Decode(in)
    }

    return nil, errors.New("No found handler")
}
