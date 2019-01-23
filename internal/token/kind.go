package token

import "fmt"

type Kind uint8
const (
    Kind_Unsupported Kind = iota
    Kind_Whitespace
    Kind_Identifier
    Kind_Number
    Kind_String // `"`, and everything between them

    Kind_Keyword_True
    Kind_Keyword_False
    Kind_Keyword_If
    Kind_Keyword_Else
    Kind_Keyword_While

    Kind_Operator_Assign               // `=`
    Kind_Operator_Negate               // `!`
    Kind_Operator_Equal                // `==`
    Kind_Operator_NotEqual             // `!=`
    Kind_Operator_GreaterThan          // `>`
    Kind_Operator_LessThan             // `<`
    Kind_Operator_GreaterThanOrEqualTo // `>=`
    Kind_Operator_LessThanOrEqualTo    // `<=`
    Kind_Operator_And                  // `&&`
    Kind_Operator_Or                   // `||`
    Kind_Operator_Add                  // `+`
    Kind_Operator_Subtract             // `-`
    Kind_Operator_Multiply             // `*`
    Kind_Operator_Divide               // `/`

    Kind_Syntax_ValueSeparator // `,`
    Kind_Syntax_Comment        // `#`

    Kind_Brackets_Parentheses_Open   // `(`
    Kind_Brackets_Parentheses_Close  // `)`
    Kind_Brackets_Curly_Open         // `{`
    Kind_Brackets_Curly_Close        // `}`
)

var (
    kindStrings = map[Kind]string{
        Kind_Unsupported: "unsupported",
        Kind_Whitespace:  "whitespace",
        Kind_Identifier:  "identifier",
        Kind_Number:      "number",
        Kind_String:      "string",

        Kind_Keyword_True:  "keyword:true",
        Kind_Keyword_False: "keyword:false",
        Kind_Keyword_If:    "keyword:if",
        Kind_Keyword_Else:  "keyword:else",
        Kind_Keyword_While: "keyword:while",

        Kind_Operator_Assign:               "operator:assign",
        Kind_Operator_Negate:               "operator:negate",
        Kind_Operator_Equal:                "operator:equal",
        Kind_Operator_NotEqual:             "operator:not-equal",
        Kind_Operator_GreaterThan:          "operator:greater-than",
        Kind_Operator_LessThan:             "operator:less-than",
        Kind_Operator_GreaterThanOrEqualTo: "operator:greater-than-or-equal-to",
        Kind_Operator_LessThanOrEqualTo:    "operator:less-than-or-equal-to",
        Kind_Operator_And:                  "operator:and",
        Kind_Operator_Or:                   "operator:or",
        Kind_Operator_Add:                  "operator:add",
        Kind_Operator_Subtract:             "operator:subtract",
        Kind_Operator_Multiply:             "operator:multiply",
        Kind_Operator_Divide:               "operator:divide",

        Kind_Syntax_ValueSeparator:      "syntax:value-separator",
        Kind_Syntax_Comment:             "syntax:comment",

        Kind_Brackets_Parentheses_Open:  "brackets:parentheses:open",
        Kind_Brackets_Parentheses_Close: "brackets:parentheses:close",
        Kind_Brackets_Curly_Open:        "brackets:curly:open",
        Kind_Brackets_Curly_Close:       "brackets:curly:close",
    }
)

func (k Kind) String() string {
    kindString := kindStrings[k]
    if kindString == "" {
        panic(fmt.Sprintf("undefined string for Kind %d", k))
    }
    return kindString
}