package library

import "github.com/daniel-va/idpa/internal/runtime"

func initTypes() {
    Entries["to_string"] = runtime.ClosureValue{
       Arity: 1,
       Call: func(params []runtime.Value) (runtime.Value, *runtime.Error) {
           return runtime.StringValue(params[0].Dump()), nil
       },
    }
}
