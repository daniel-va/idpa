package runtime

type ArrayValue []Value

func (v ArrayValue) Dump() string {
    dump := "["
    if len(v) > 0 {
        dump += " "
        for i, element := range v {
            if i > 0 {
                dump += ", "
            }
            dump += element.Dump()
        }
        dump += " "
    }
    return dump + "]"
}

func (v ArrayValue) Type() ValueType {
    return ValueType_Array
}
