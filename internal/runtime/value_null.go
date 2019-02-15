package runtime

type NullValue struct {}

func (v NullValue) Dump() string {
    return v.Type().String()
}

func (NullValue) Type() ValueType {
    return ValueType_Null
}