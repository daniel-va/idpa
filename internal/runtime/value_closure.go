package runtime

import "strconv"

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

type TypeAnnotation int
func (t TypeAnnotation) Allow(types ...ValueType) TypeAnnotation {
    result := t
    for _, currentType := range types {
        result |= 1 << currentType
    }
    return result
}

func (t TypeAnnotation) Allows(aType ValueType) bool {
    return (int(t) & int(aType)) != 0
}