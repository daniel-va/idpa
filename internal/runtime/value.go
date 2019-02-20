package runtime

import "fmt"

type Value interface {
    Dump() string
    Type() ValueType
}

type ValueType uint
const (
    ValueType_Array   = 1 << 0
    ValueType_Boolean = 1 << 1
    ValueType_Closure = 1 << 2
    ValueType_Null    = 1 << 3
    ValueType_Number  = 1 << 4
    ValueType_String  = 1 << 5
)

var (
    valueType_StringMappings = map[ValueType]string{
        ValueType_Array:   "Array",
        ValueType_Boolean: "Boolean",
        ValueType_Closure: "Closure",
        ValueType_Number:  "Number",
        ValueType_Null:    "null",
        ValueType_String:  "String",
    }
)

func (t ValueType) String() string {
    result, ok := valueType_StringMappings[t]
    if !ok {
        panic(fmt.Sprintf("no mapping for `%d` found", t))
    }
    return result
}