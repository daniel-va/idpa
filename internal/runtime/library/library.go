package library

import (
    "github.com/daniel-va/idpa/internal/runtime"
)

var Entries = map[string]runtime.Value{}

func init() {
    initIO()
    initTypes()
    initArray()
}