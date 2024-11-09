package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/MikelGV/jsonParser/lexer/tokens"
)

type Lexer struct {
    // Here should go something but i don't know what right now
    Name string
    Input string
    Tokens chan tokens.Token
    State LexFn
    
    Start int
    Pos int
    Width int
}

func (t *Lexer) Emit(tokenType tokens.TokenType)  {
    t.Tokens <- tokens.Token{Type: tokenType, Value: t.Input[t.Start:t.Pos]}
    t.Start = t.Pos
}

func (t *Lexer) InputToEnd() string {
    return t.Input[t.Pos:]
}

func (t *Lexer) Inc() {
    t.Pos++
    if t.Pos >= utf8.RuneCountInString(t.Input) {
        t.Emit(tokens.TOKEN_EOF)
    }
}

func (t *Lexer) SkipWhiteSpace() {
    for {
        ch := t.Next()

        if !unicode.IsSpace(ch) {
            t.Dec()
            break
        }

        if ch == tokens.EOF {
            t.Emit(tokens.TOKEN_EOF)
            break
        }
    }
}

func (t *Lexer) Next() {}

func (t *Lexer) Dec() {}

func BeginLexing(name, input string) *Lexer {
    l := &Lexer{
        Name: name,
        Input: input,
        State: LexingBein,
        Tokens: make(chan tokens.Token, 3),
    }

    return l
}

func LexingBein(lexer *Lexer) LexFn {
    lexer.SkipWhiteSpace()

    if strings.HasPrefix(lexer.InputToEnd(), tokens.LEFT_KEY) {
        return LexLeftKey
    } else {
        return LexerKey
    }
}

func isEOF(lexer *Lexer) {}
