package runtime

import "fmt"

type Value interface {
    Dump() string
    Type() ValueType
}

type ValueType uint8
const (
    _ ValueType = iota
    ValueType_Array
    ValueType_Boolean
    ValueType_Closure
    ValueType_Null
    ValueType_Number
    ValueType_String
)

var (
    valueType_StringMappings = map[ValueType]string{
        ValueType_Array:   "Array",
        ValueType_Boolean: "Boolean",
        ValueType_Closure: "Closure",
        ValueType_Number:  "Number",
        ValueType_Null:    "NullValue_Instance",
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