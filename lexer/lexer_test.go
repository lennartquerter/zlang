package lexer

import (
	"testing"

	"zLang/token"
)

func TestLexer_NextToken_Basic(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testStringSequence(input, tests, t)
}

func TestLexer_NextToken_Keywords(t *testing.T) {
	input := `
if (x > y) {
	return true;
} else {
	return false;
}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},

		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},

		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	testStringSequence(input, tests, t)
}

func TestLexer_NextToken_Identifiers(t *testing.T) {
	input := `
let five = 5;
let ten = 10;

let add = fn(x,y) {
	x+y;
};

let result = add(five, ten);

/*!

5 > 6
10< 20

"foobar"
"foo bar"
[1, 2];
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},

		{token.INT, "5"},
		{token.GT, ">"},
		{token.INT, "6"},

		{token.INT, "10"},
		{token.LT, "<"},
		{token.INT, "20"},

		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.RBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.LBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},


	}

	testStringSequence(input, tests, t)

}

func TestLexer_NextToken_DoubleChar(t *testing.T) {
	input := `
y == e

x != y
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{

		{token.IDENT, "y"},
		{token.EQ, "=="},
		{token.IDENT, "e"},
		{token.IDENT, "x"},
		{token.NOT_EQ, "!="},
		{token.IDENT, "y"},
	}

	testStringSequence(input, tests, t)
}

func testStringSequence(input string, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}, t *testing.T) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
