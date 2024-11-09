package lexer

import (
	"strings"

	"github.com/MikelGV/jsonParser/lexer/errors"
	"github.com/MikelGV/jsonParser/lexer/tokens"
)


func LexLeftKey(lexer *Lexer) LexFn {
    lexer.Pos += len(tokens.LEFT_KEY)
    lexer.Emit(tokens.TOKEN_LEFT_KEY)
    return LexSection
}

func LexSection(lexer *Lexer) LexFn {
    
    for {
        if lexer.IsEOF() {
            return lexer.Errorf(errors.LEXER_ERROR_MISSING_RIGHT_KEY)
        }

        if strings.HasPrefix(lexer.InputToEnd(), tokens.RIGHT_KEY) {
            lexer.Emit(tokens.TOKEN_SECTION)
            return LexRightKey
        }

        lexer.Inc()
    }
}
