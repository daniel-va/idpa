package runtime

type BooleanValue bool

func (v BooleanValue) Dump() string {
    if v {
        return "true"
    } else {
        return "false"
    }
}

func (v BooleanValue) Type() ValueType {
    return ValueType_Boolean
}