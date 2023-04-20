package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEDAL"
	EOF     = "EOF"
	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 123456
	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	// ()	丸括弧	パーレンシス(parenthesis)
	// []	角括弧	ブラケット(bracket)
	// {}	波括弧	ブレース(brace)
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
