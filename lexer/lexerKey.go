package lexer

import (
	"strings"

	"github.com/MikelGV/jsonParser/lexer/errors"
	"github.com/MikelGV/jsonParser/lexer/tokens"
)

func LexerKey(lexer *Lexer) LexFn {
    
    for {
        if strings.HasPrefix(lexer.InputToEnd(), tokens.COLON) {
            lexer.Emit(tokens.TOKEN_KEY)
            return LexColon
        }

        lexer.Inc()


        if lexer.isEOF() {
            return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
        }
    }
}
