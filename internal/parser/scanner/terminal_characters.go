package scanner

import (
	"unicode"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
)

// A terminal character that is specified in the SQL specification.
var (
	SQLTerminalCharacter = SQLLanguageCharacter
	SQLLanguageCharacter = matcher.Merge(
		SimpleLatinCharacter,
		Digit,
		SQLSpecialCharacter,
	)
	SimpleLatinCharacter = matcher.Merge(
		SimpleLatinUpperCaseLetter,
		SimpleLatinLowerCaseLetter,
	)
	SimpleLatinUpperCaseLetter = matcher.String("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	SimpleLatinLowerCaseLetter = matcher.String("abcdefghijklmnopqrstuvwxyz")
	Digit                      = matcher.String("0123456789")
	SQLSpecialCharacter        = matcher.Merge(
		Space,
		DoubleQuote,
		Percent,
		Ampersand,
		Quote,
		LeftParen,
		RightParen,
		Asterisk,
		PlusSign,
		Comma,
		MinusSign,
		Period,
		Solidus,
		Colon,
		Semicolon,
		LessThanOperator,
		EqualsOperator,
		GreaterThanOperator,
		QuestionMark,
		LeftBracket,
		RightBracket,
		Circumflex,
		Underscore,
		VerticalBar,
		LeftBrace,
		RightBrace,
		DollarSign,
		Apostrophe,
	)
	Space               = matcher.RuneWithDesc("Space", '\u0020')
	DoubleQuote         = matcher.RuneWithDesc("Double Quote", '"')
	Percent             = matcher.RuneWithDesc("Percent", '%')
	Ampersand           = matcher.RuneWithDesc("Ampersand", '&')
	Quote               = matcher.RuneWithDesc("Quote", '\'')
	LeftParen           = matcher.RuneWithDesc("Left Paren", '(')
	RightParen          = matcher.RuneWithDesc("Right Paren", ')')
	Asterisk            = matcher.RuneWithDesc("Asterisk", '*')
	PlusSign            = matcher.RuneWithDesc("Plus Sign", '+')
	Comma               = matcher.RuneWithDesc("Comma", ',')
	MinusSign           = matcher.RuneWithDesc("Minus Sign", '-')
	Period              = matcher.RuneWithDesc("Period", '.')
	Solidus             = matcher.RuneWithDesc("Solidus", '/')
	ReverseSolidus      = matcher.RuneWithDesc("Reverse Solidus", '/')
	Colon               = matcher.RuneWithDesc("Colon", ':')
	Semicolon           = matcher.RuneWithDesc("Semicolon", ';')
	LessThanOperator    = matcher.RuneWithDesc("Less Than Operator", '<')
	EqualsOperator      = matcher.RuneWithDesc("Equals Operator", '=')
	GreaterThanOperator = matcher.RuneWithDesc("Greater Than Operator", '>')
	QuestionMark        = matcher.RuneWithDesc("Question Mark", '?')
	// Left and Right bracket trigraph must be lexed by the scanner, as these
	// are sequences of three runes
	LeftBracket  = matcher.RuneWithDesc("Left Bracket", '[')
	RightBracket = matcher.RuneWithDesc("Right Bracket", ']')
	Circumflex   = matcher.RuneWithDesc("Circumflex", '^')
	Underscore   = matcher.RuneWithDesc("Underscore", '_')
	VerticalBar  = matcher.RuneWithDesc("Vertical Bar", '|')
	LeftBrace    = matcher.RuneWithDesc("Left Brace", '{')
	RightBrace   = matcher.RuneWithDesc("Right Brace", '}')
	DollarSign   = matcher.RuneWithDesc("Dollar Sign", '$')
	Apostrophe   = matcher.RuneWithDesc("Apostrophe", '\'')
)

// A fragment of a lexical unit.
var (
	IdentifierPart = matcher.Merge(
		IdentifierStart,
		IdentifierExtend,
	)
	IdentifierStart = matcher.Merge(
		matcher.New("Lu", unicode.Lu),
		matcher.New("Ll", unicode.Ll),
		matcher.New("Lt", unicode.Lt),
		matcher.New("Lm", unicode.Lm),
		matcher.New("Lo", unicode.Lo),
		matcher.New("Nl", unicode.Nl),
	)
	IdentifierExtend = matcher.Merge(
		matcher.RuneWithDesc("Middle Dot", '\u00b7'),
		matcher.New("Mn", unicode.Mn),
		matcher.New("Mc", unicode.Mc),
		matcher.New("Nd", unicode.Nd),
		matcher.New("Pc", unicode.Pc),
		matcher.New("Cf", unicode.Cf),
	)
	Multiplier = matcher.String("KMGTP") // Kilo, Mega, Giga, Tera, Peta
)
