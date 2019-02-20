package library

import "github.com/daniel-va/idpa/internal/runtime"

func initArray() {
    Entries["array"] = runtime.ClosureValue{
        Arity: -1,
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            return runtime.ArrayValue(params), nil
        },
    }

    Entries["array_length"] = runtime.ClosureValue{
        Arity: 1,
        ParamTypes: []runtime.TypeAnnotation{
            runtime.AllowTypes(runtime.ValueType_Array),
        },
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            array := params[0].(runtime.ArrayValue)
            return runtime.NumberValue(len(array)), nil
        },
    }

    Entries["array_get"] = runtime.ClosureValue{
        Arity: 2,
        ParamTypes: []runtime.TypeAnnotation{
          runtime.AllowTypes(runtime.ValueType_Array),
          runtime.AllowTypes(runtime.ValueType_Number),
        },
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            array      := params[0].(runtime.ArrayValue)
            index, err := array_asIndex(params[1])
            if err != nil {
                return nil, err
            }

            err = array_checkIndex(array, index)
            if err != nil {
                return nil, err
            }

            return array[index], nil
        },
    }

    Entries["array_set"] = runtime.ClosureValue{
        Arity: 3,
        ParamTypes: []runtime.TypeAnnotation{
            runtime.AllowTypes(runtime.ValueType_Array),
            runtime.AllowTypes(runtime.ValueType_Number),
            runtime.AllowTypes(),
        },
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            array := params[0].(runtime.ArrayValue)
            index, err := array_asIndex(params[1])
            if err != nil {
                return nil, err
            }
            value := params[2]

            err = array_checkIndex(array, index)
            if err != nil {
                return nil, err
            }

            array[index] = value
            return runtime.NullValue{}, nil
        },
    }

    Entries["array_append"] = runtime.ClosureValue{
        Arity: 2,
        ParamTypes: []runtime.TypeAnnotation{
            runtime.AllowTypes(runtime.ValueType_Array),
            runtime.AllowTypes(),
        },
        Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
            array := params[0].(runtime.ArrayValue)
            value := params[1]

            array = append(array, value)
            return array, nil
        },
    }
}

func array_asIndex(value runtime.Value) (int64, *runtime.Error) {
    indexValue := value.(runtime.NumberValue)
    if indexValue < 0 {
        return 0, runtime.Err("index can't be negative")
    }
    return int64(indexValue), nil
}

func array_checkIndex(array runtime.ArrayValue, index int64) *runtime.Error {
    if int64(len(array)) <= index {
        return runtime.Err("index %d not contained in array of size %d", index, len(array))
    }
    return nil
}