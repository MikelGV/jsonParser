package lexer

import "github.com/MikelGV/jsonParser/lexer/tokens"


func LexRightKey(lexer *Lexer) LexFn {
    lexer.Pos += len(tokens.RIGHT_KEY)
    lexer.Emit(tokens.TOKEN_RIGHT_KEY)
    return LexingBein
}
