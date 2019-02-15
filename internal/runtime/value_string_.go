package runtime

type StringValue string

func (v StringValue) Dump() string {
    return string(v)
}

func (v StringValue) Type() ValueType {
    return ValueType_String
}