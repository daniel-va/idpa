package runtime

import (
    "strconv"
)

type ClosureValue struct {
    Arity int
    Call  func(params []Value) (Value, *Error)

    ParamTypes []TypeAnnotation
}

func (v ClosureValue) Dump() string {
    return string(v.Type()) + "/" + strconv.Itoa(v.Arity)
}

func (v ClosureValue) Type() ValueType {
    return ValueType_Closure
}


func AllowTypes(types ...ValueType) TypeAnnotation {
    return TypeAnnotation(0).Allow(types...)
}

type TypeAnnotation uint
func (t TypeAnnotation) Allow(types ...ValueType) TypeAnnotation {
    result := uint(t)
    for _, currentType := range types {
        result |= uint(currentType)
    }
    return TypeAnnotation(result)
}

func (t TypeAnnotation) Allows(aType ValueType) bool {
    return (uint(t) & uint(aType)) != 0
}