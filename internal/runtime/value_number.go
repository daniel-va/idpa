package runtime

import "strconv"

type NumberValue float64

func (v NumberValue) Dump() string {
    return strconv.FormatFloat(float64(v), 'f', -1, 64)
}

func (v NumberValue) Type() ValueType {
    return ValueType_Number
}
