package exec

import . "github.com/daniel-va/idpa/internal/runtime"

type ContextScope struct {
    Parent *ContextScope

    Members map[string]Value
}

func NewScope() *ContextScope {
    return &ContextScope{
        Members: make(map[string]Value),
    }
}

func (s *ContextScope) Fetch(name string) (Value, bool) {
    value, ok := s.Members[name]
    if !ok && s.Parent != nil {
        return s.Parent.Fetch(name)
    }
    return value, ok
}

func (s *ContextScope) Assign(name string, value Value) {
    owningScope, ok := s.fetchOwningScope(name)
    if !ok {
        owningScope = s
    }
    owningScope.Members[name] = value
}

func (s *ContextScope) fetchOwningScope(name string) (*ContextScope, bool) {
    _, ok := s.Members[name]
    if !ok && s.Parent != nil {
        return s.Parent.fetchOwningScope(name)
    }
    return s, ok
}

func (s *ContextScope) Subscope() *ContextScope {
    subscope       := NewScope()
    subscope.Parent = s
    return subscope
}