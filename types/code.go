package types

import (
    "unsafe"
)

type Code struct {
    name string
    Codec
}

func NewCode(c Codec) Code {
    return Code{
        name: "code",
        Codec: c,
    }
}

var Codecs []Codec

func Setup(codecs []Codec, codec Codec) []Codec {
    if codecs == nil && codec == nil {
        return nil
    }

    codecs = append(codecs, codec)

    return codecs
}

func SetCodec(cfg *Configure, cmd *Command, ptr *unsafe.Pointer) int {
    if cfg == nil || cmd == nil || ptr == nil {
        return Error
    }

    field := (*Code)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))
    if field == nil {
        return Error
    }

    value := cfg.GetValue()
    if value == nil {
        return Error
    }

    values := value.(map[interface{}]interface{})

    var code Code

    for k, v := range values {
        name := k.(string)

        for _, codec := range Codecs {
            if codec.Type(name) == Ignore {
                continue
            }

            codec.Init(v)

            code = NewCode(codec)
            code.name = name
        }
    }

    *field = code

    return Ok
}

func (r Code) Encode() int {
    if handler := r.Codec; handler != nil {
        return handler.Encode()
    }

    return Ok
}

func (r Code) Decode() int {
    if handler := r.Codec; handler != nil {
        return handler.Decode()
    }

    return Ok
}
