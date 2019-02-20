package library

import (
    "fmt"
    "github.com/daniel-va/idpa/internal/runtime"
    "strconv"
)

func initTypes() {
    Entries["null"] = runtime.NullValue{}

    Entries["to_string"] = runtime.ClosureValue{
        Arity: 1,
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            return runtime.StringValue(params[0].Dump()), nil
        },
    }

    Entries["to_number"] = runtime.ClosureValue{
        Arity: 1,
        ParamTypes: []runtime.TypeAnnotation{
            runtime.AllowTypes(
                runtime.ValueType_Null,
                runtime.ValueType_Number,
                runtime.ValueType_String,
                runtime.ValueType_Boolean,
            ),
        },
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            value := params[0]
            switch value.Type() {
            case runtime.ValueType_Null:
                return runtime.NumberValue(0), nil
            case runtime.ValueType_Number:
                return value.(runtime.NumberValue), nil
            case runtime.ValueType_String:
                numberValue, err := strconv.ParseFloat(string(value.(runtime.StringValue)), 64)
                if err != nil {
                    return nil, runtime.Err("string does not contain a valid number")
                }
                return runtime.NumberValue(numberValue), nil
            case runtime.ValueType_Boolean:
                if value.(runtime.BooleanValue) {
                    return runtime.NumberValue(1), nil
                } else {
                    return runtime.NumberValue(0), nil
                }
            default:
                panic(fmt.Sprintf("unchecked type: %T", value.Type()))
            }
        },
    }
}
