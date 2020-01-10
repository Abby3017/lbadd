package ast

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

// 6 scalar expressions

// 6.1 data type
type (
	DataType struct {
		Node

		PredefinedType                  *PredefinedType
		RowType                         *RowType
		PathResolvedUserDefinedTypeName *PathResolvedUserDefinedTypeName
		ReferenceType                   *ReferenceType
		CollectionType                  *CollectionType
	}

	PredefinedType struct {
		Node

		CharacterStringType         *CharacterStringType
		Character                   token.Token
		Set                         token.Token
		CharacterSetSpecification   *CharacterSetSpecification
		NationalCharacterStringType *NationalCharacterStringType
		CollateClause               *CollateClause
		BinaryStringType            *BinaryStringType
		NumericType                 *NumericType
		BooleanType                 *BooleanType
		DatetimeType                *DatetimeType
		IntervalType                *IntervalType
	}

	CharacterStringType struct {
		Node

		Character                token.Token
		Char                     token.Token
		Varchar                  token.Token
		Varying                  token.Token
		LeftParen                token.Token
		CharacterLength          *CharacterLength
		RightParen               token.Token
		CharacterLargeObjectType *CharacterLargeObjectType
	}

	CharacterLargeObjectType struct {
		Node

		Character                  token.Token
		Char                       token.Token
		Clob                       token.Token
		Large                      token.Token
		Object                     token.Token
		LeftParen                  token.Token
		CharacterLargeObjectLength *CharacterLargeObjectLength
		RightParen                 token.Token
	}

	NationalCharacterStringType struct {
		Node

		National                         token.Token
		Nchar                            token.Token
		Character                        token.Token
		Char                             token.Token
		Varying                          token.Token
		LeftParen                        token.Token
		CharacterLength                  *CharacterLength
		RightParen                       token.Token
		NationalCharacterLargeObjectType *NationalCharacterLargeObjectType
	}

	NationalCharacterLargeObjectType struct {
		Node

		National                   token.Token
		Character                  token.Token
		Large                      token.Token
		Object                     token.Token
		Nchar                      token.Token
		Nclob                      token.Token
		LeftParen                  token.Token
		CharacterLargeObjectLength *CharacterLargeObjectLength
		RightParen                 token.Token
	}

	BinaryStringType struct {
		Node

		Binary                      token.Token
		Varying                     token.Token
		Varbinary                   token.Token
		LeftParen                   token.Token
		Length                      *Length
		RightParen                  token.Token
		BinaryLargeObjectStringType *BinaryLargeObjectStringType
	}

	BinaryLargeObjectStringType struct {
		Node

		Binary            token.Token
		Large             token.Token
		Object            token.Token
		Blob              token.Token
		LeftParen         token.Token
		LargeObjectLength *LargeObjectLength
		RightParen        token.Token
	}
)
