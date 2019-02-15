package token

import (
    "github.com/daniel-va/idpa/internal/source"
    "unicode"
)

func newParser(readCh <-chan rune) *parser {
    return &parser{
        readCh: readCh,
        pos:    source.Pos{ Row: 1, },
    }
}

type parser struct {
    readCh <-chan rune
    pos    source.Pos

    buffer *Token
}

func (p *parser) Parse() (Token, bool) {
    tk, ok := p.next()
    if !ok {
        return unparsed()
    }

    switch tk.Kind {
    case Kind_Syntax_Comment:
        p.parseComment()
        return p.Parse()
    case Kind_String:
        return p.parseString(tk), true
    default:
        return p.parseDefault(tk), true
    }
}


func (p *parser) parseDefault(tk Token) Token {
    nextTk, ok := p.next()
    if !ok {
        return tk
    }

    if joinKind, ok := p.join(tk, nextTk); ok {
        tk.Value += nextTk.Value
        tk.Kind   = joinKind
        return p.parseDefault(tk)
    } else {
        p.buffer = &nextTk
        return tk
    }
}

func (p *parser) parseComment() {
    tk, ok := p.next()
    if !ok {
        return
    }
    if tk.Kind == Kind_Whitespace && tk.Value == "\n" {
        return
    }
    p.parseComment()
}

func (p *parser) parseString(entry Token) Token {
    tk, ok := p.next()
    if !ok {
        return entry
    }
    var done bool

    switch tk.Value {
    case "\"":
        done = true
    case "\n":
        done = true
    case "\\":
        escapedTk, ok := p.next()
        if !ok {
            done = true
        }
        switch escapedTk.Value {
        case "n":
            tk.Value = "\n"
        case "r":
            tk.Value = "\r"
        case "t":
            tk.Value = "\t"
        case "\"":
            tk.Value = "\""
        case "\\":
            tk.Value = "\\"
        default:
            tk.Value = escapedTk.Value
        }
    }

    entry.Value += tk.Value
    if done {
        return entry
    }
    return p.parseString(entry)
}

func (p *parser) next() (Token, bool) {
    if p.buffer != nil {
        bufferedChar := *p.buffer
        p.buffer = nil
        return bufferedChar, true
    }

    char, ok := <-p.readCh
    if !ok {
        return Token{}, ok
    }
    switch char {
    case '\n':
        p.pos.Row += 1
        p.pos.Col  = 0
    default:
        p.pos.Col += 1
    }

    return Token{
        Value: string(char),
        Kind:  p.kindOf(char),
        Pos:   p.pos,
    }, true
}

func (p *parser) kindOf(char rune) Kind {
    switch char {
    case '_':
        return Kind_Identifier
    case ' ', '\n', '\r', '\t':
        return Kind_Whitespace
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return Kind_Number
    case '"':
        return Kind_String

    // operators
    case '=':
        return Kind_Operator_Assign
    case '!':
        return Kind_Operator_Negate
    case '>':
        return Kind_Operator_GreaterThan
    case '<':
        return Kind_Operator_LessThan
    case '+':
        return Kind_Operator_Add
    case '-':
        return Kind_Operator_Subtract
    case '*':
        return Kind_Operator_Multiply
    case '/':
        return Kind_Operator_Divide

    // syntax
    case ',':
        return Kind_Syntax_ValueSeparator
    case '#':
        return Kind_Syntax_Comment
    case '.':
       return Kind_Syntax_DecimalPoint

    // brackets
    case '(':
        return Kind_Brackets_Parentheses_Open
    case ')':
        return Kind_Brackets_Parentheses_Close
    case '{':
        return Kind_Brackets_Curly_Open
    case '}':
        return Kind_Brackets_Curly_Close
    }

    if unicode.IsLetter(char) {
        return Kind_Identifier
    }

    return Kind_Unsupported
}

func (p *parser) join(a, b Token) (Kind, bool) {
    switch a.Kind {
    case Kind_Unsupported:
        if kind, ok := p.joinUnsupported(a.Value, b.Value); ok {
            return kind, true
        } else {
            return Kind_Unsupported, true
        }
    case b.Kind:
        switch a.Kind {
        case Kind_Identifier:
            return p.joinIdentifier(a.Value + b.Value), true
        default:
            return p.joinKind(a.Kind)
        }
    }

    switch a.Kind {
    case Kind_Identifier:
        return a.Kind, b.Kind == Kind_Number

    case Kind_Keyword_False,
        Kind_Keyword_True,
        Kind_Keyword_If,
        Kind_Keyword_Else,
        Kind_Keyword_While,
        Kind_Keyword_Return:
        return b.Kind, b.Kind == Kind_Identifier

    // operators
    case Kind_Operator_Negate:
        return Kind_Operator_NotEqual, b.Kind == Kind_Operator_Assign
    case Kind_Operator_GreaterThan:
        return Kind_Operator_GreaterThanOrEqualTo, b.Kind == Kind_Operator_Equal
    case Kind_Operator_LessThan:
        return Kind_Operator_LessThanOrEqualTo, b.Kind == Kind_Operator_Equal
    }

    return Kind_Unsupported, false
}

func (p *parser) joinKind(kind Kind) (Kind, bool) {
    switch kind {
    case Kind_Identifier,
         Kind_Number:
        return kind, true

    // operators
    case Kind_Operator_Assign:
        return Kind_Operator_Equal, true

    default:
        return kind, false
    }
}

func (p *parser) joinIdentifier(value string) Kind {
    switch value {

    // keywords
    // Don't forget to add them in `parser.join` too.
    // Otherwise, keywords would mark the end of an identifier.
    case "true":
        return Kind_Keyword_True
    case "false":
        return Kind_Keyword_False
    case "if":
        return Kind_Keyword_If
    case "else":
        return Kind_Keyword_Else
    case "while":
        return Kind_Keyword_While
    case "return":
        return Kind_Keyword_Return

    default:
        return Kind_Identifier
    }
}

func (p *parser) joinUnsupported(a, b string) (Kind, bool) {
    switch a {
    case "&":
        return Kind_Operator_And, b == "&"
    case "|":
        return Kind_Operator_Or, b == "|"
    default:
        return Kind_Unsupported, false
    }
}

func unparsed() (Token, bool) {
    return Token{}, false
}