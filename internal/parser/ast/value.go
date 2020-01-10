package ast

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

// 6.27 JSON value function
type (
	JSONValueFunction struct {
		Node

		JsonValue               token.Token
		LeftParen               token.Token
		JSONAPICommonSyntax     *JSONAPICommonSyntax
		JSONReturningClause     *JSONReturningClause
		JSONValueEmptyBehaviour *JSONValueEmptyBehaviour
		JSONValueErrorBehaviour *JSONValueErrorBehaviour
		On                      token.Token
		EmptyOrError            token.Token
	}

	JSONReturningClause struct {
		Node

		Returning token.Token
		DataType  *DataType
	}

	JSONValueEmptyBehaviour struct {
		Node

		ErrorOrNullOrDefault token.Token
		ValueExpression      *ValueExpression
	}

	JSONValueErrorBehaviour struct {
		Node

		ErrorOrNullOrDefault token.Token
		ValueExpression      *ValueExpression
	}
)

// 6.28 value expression
type (
	ValueExpression struct {
		Node

		CommonValueExpression  *CommonValueExpression
		BooleanValueExpression *BooleanValueExpression
		RowValueExpression     *RowValueExpression
	}

	CommonValueExpression struct {
		Node

		NumericValueExpression         *NumericValueExpression
		StringValueExpression          *StringValueExpression
		DatetimeValueExpression        *DatetimeValueExpression
		IntervalValueExpression        *IntervalValueExpression
		UserDefinedTypeValueExpression *UserDefinedTypeValueExpression
		ReferenceValueExpression       *ReferenceValueExpression
		CollectionValueExpression      *CollectionValueExpression
	}

	UserDefinedTypeValueExpression struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
	}

	ReferenceValueExpression struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
	}

	CollectionValueExpression struct {
		Node

		ArrayValueExpression    *ArrayValueExpression
		MultisetValueExpression *MultisetValueExpression
	}
)

// 6.29 numeric value expression
type (
	NumericValueExpression struct {
		Node

		NumericValueExpression *NumericValueExpression
		PlusSignOrMinusSign    token.Token
		Term                   *Term
	}

	Term struct {
		Node

		Term              *Term
		AsteriskOrSolidus token.Token
		Factor            *Factor
	}

	Factor struct {
		Node

		Sign           token.Token
		NumericPrimary *NumericPrimary
	}

	NumericPrimary struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
		NumericValueFunction   *NumericValueFunction
	}
)

// 6.30 numeric value function
type (
	NumericValueFunction struct {
		Node

		PositionExpression       *PositionExpression
		RegexOccurrencesFunction *RegexOccurrencesFunction
		RegexPositionExpression  *RegexPositionExpression
		ExtractExpression        *ExtractExpression
		LengthExpression         *LengthExpression
		CardinalityExpression    *CardinalityExpression
		MaxCardinalityExpression *MaxCardinalityExpression
		AbsoluteValueExpression  *AbsoluteValueExpression
		ModulusExpression        *ModulusExpression
		TrigonometricFunction    *TrigonometricFunction
		GeneralLogarithmFunction *GeneralLogarithmFunction
		CommonLogarithm          *CommonLogarithm
		NaturalLogarithm         *NaturalLogarithm
		ExponentialFunction      *ExponentialFunction
		PowerFunction            *PowerFunction
		SquareRoot               *SquareRoot
		FloorFunction            *FloorFunction
		CeilingFunction          *CeilingFunction
		WidthBucketFunction      *WidthBucketFunction
		MatchNumberFunction      *MatchNumberFunction
	}

	PositionExpression struct {
		Node

		CharacterPositionExpression *CharacterPositionExpression
		BinaryPositionExpression    *BinaryPositionExpression
	}

	RegexOccurrencesFunction struct {
		Node

		OccurrencesRegex token.Token
		LeftParen        token.Token
		XQueryPattern    *XQueryPattern
		Flag             token.Token
		XQueryOptionFlag *XQueryOptionFlag
		In               token.Token
		From             token.Token
		StartPosition    *StartPosition
		Using            token.Token
		CharLengthUnits  *CharLengthUnits
		RightParen       token.Token
	}

	XQueryPattern struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	XQueryOptionFlag struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	RegexSubjectString struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	RegexPositionExpression struct {
		Node

		PositionRegex             token.Token
		LeftParen                 token.Token
		RegexPositionStartOrAfter *RegexPositionStartOrAfter
		XQueryPattern             *XQueryPattern
		Flag                      token.Token
		XQueryOptionFlag          *XQueryOptionFlag
		In                        token.Token
		RegexSubjectString        *RegexSubjectString
		From                      token.Token
		StartPosition             *StartPosition
		Using                     token.Token
		CharLengthUnits           *CharLengthUnits
		Occurrence                token.Token
		RegexOccurrence           *RegexOccurrence
	}

	RegexPositionStartOrAfter struct {
		Node

		StartOrAfter token.Token
	}

	RegexOccurrence struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	RegexCaptureGroup struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	CharacterPositionExpression struct {
		Node

		Position                  token.Token
		LeftParen                 token.Token
		CharacterValueExpression1 *CharacterValueExpression
		In                        token.Token
		CharacterValueExpression2 *CharacterValueExpression
		Using                     token.Token
		CharLengthUnits           *CharLengthUnits
		RightParen                token.Token
	}

	BinaryPositionExpression struct {
		Node

		Position               token.Token
		LeftParen              token.Token
		BinaryValueExpression1 *BinaryValueExpression
		In                     token.Token
		BinaryValueExpression2 *BinaryValueExpression
		RightParen             token.Token
	}

	LengthExpression struct {
		Node

		CharLengthExpression  *CharLengthExpression
		OctetLengthExpression *OctetLengthExpression
	}

	CharLengthExpression struct {
		Node

		CharLengthOrCharacterLength token.Token
		LeftParen                   token.Token
		CharacterValueExpression    *CharacterValueExpression
		Using                       token.Token
		CharLengthUnits             *CharLengthUnits
		RightParen                  token.Token
	}

	OctetLengthExpression struct {
		Node

		OctetLength           token.Token
		LeftParen             token.Token
		StringValueExpression *StringValueExpression
		RightParen            token.Token
	}

	ExtractExpression struct {
		Node

		Extract       token.Token
		LeftParen     token.Token
		ExtractField  *ExtractField
		From          token.Token
		ExtractSource *ExtractSource
		RightParen    token.Token
	}

	ExtractField struct {
		Node

		PrimaryDatetimeField *PrimaryDatetimeField
		TimeZoneField        *TimeZoneField
	}

	TimeZoneField struct {
		Node

		TimezoneHourOrTimezoneMinute token.Token
	}

	ExtractSource struct {
		Node

		DatetimeValueExpression *DatetimeValueExpression
		IntervalValueExpression *IntervalValueExpression
	}

	CardinalityExpression struct {
		Node

		Cardinality               token.Token
		LeftParen                 token.Token
		CollectionValueExpression *CollectionValueExpression
		RightParen                token.Token
	}

	MaxCardinalityExpression struct {
		Node

		ArrayMaxCardinality  token.Token
		LeftParen            token.Token
		ArrayValueExpression *ArrayValueExpression
		RightParen           token.Token
	}

	AbsoluteValueExpression struct {
		Node

		Abs                    token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	ModulusExpression struct {
		Node

		Mod                            token.Token
		LeftParen                      token.Token
		NumericValueExpressionDividend *NumericValueExpressionDividend
		Comma                          token.Token
		NumericValueExpressionDivisor  *NumericValueExpressionDivisor
		RightParen                     token.Token
	}

	NumericValueExpressionDividend struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	NumericValueExpressionDivisor struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	TrigonometricFunctionName struct {
		Node

		Sin  token.Token
		Cos  token.Token
		Tan  token.Token
		Sinh token.Token
		Cosh token.Token
		Tanh token.Token
		Asin token.Token
		Acos token.Token
		Atan token.Token
	}

	GeneralLogarithmFunction struct {
		Node

		Log                      token.Token
		LeftParen                token.Token
		GeneralLogarithmBase     *GeneralLogarithmBase
		Comma                    token.Token
		GeneralLogarithmArgument *GeneralLogarithmArgument
		RightParen               token.Token
	}

	GeneralLogarithmBase struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	GeneralLogarithmArgument struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	CommonLogarithm struct {
		Node

		Log10                  token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	NaturalLogarithm struct {
		Node

		Ln                     token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	ExponentialLogarithm struct {
		Node

		Exp                    token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	PowerLogarithm struct {
		Node

		Power                          token.Token
		LeftParen                      token.Token
		NumericValueExpressionBase     *NumericValueExpressionBase
		Comma                          token.Token
		NumericValueExpressionExponent *NumericValueExpressionExponent
		RightParen                     token.Token
	}

	NumericValueExpressionBase struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	NumericValueExpressionExponent struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	SquareRoot struct {
		Node

		Sqrt                   token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	FloorFunction struct {
		Node

		Floor                  token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	CeilingFunction struct {
		Node

		CeilOrCeiling          token.Token
		LeftParen              token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}

	WidthBucketFunction struct {
		Node

		WidthBucket        token.Token
		LeftParen          token.Token
		WidthBucketOperand *WidthBucketOperand
		Comma1             token.Token
		WidthBucketBound1  *WidthBucketBound1
		Comma2             token.Token
		WidthBucketBound2  *WidthBucketBound2
		Comma3             token.Token
		WidthBucketCount   *WidthBucketCount
		RightParen         token.Token
	}

	WidthBucketOperand struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	WidthBucketBound1 struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	WidthBucketBound2 struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	WidthBucketCount struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	MatchNumberFunction struct {
		Node

		MatchNumber token.Token
		LeftParen   token.Token
		RightParen  token.Token
	}
)

// 6.31 string value expression
type (
	StringValueExpression struct {
		Node

		CharacterValueExpression *CharacterValueExpression
		BinaryValueExpression    *BinaryValueExpression
	}

	CharacterValueExpression struct {
		Node

		Concatenation   *Concatenation
		CharacterFactor *CharacterFactor
	}

	Concatenation struct {
		Node

		CharacterValueExpression *CharacterValueExpression
		ConcatenationOperator    token.Token
		CharacterFactor          *CharacterFactor
	}

	CharacterFactor struct {
		Node

		CharacterPrimary *CharacterPrimary
		CollateClause    *CollateClause
	}

	CharacterPriamry struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
		StringValueFunction    *StringValueFunction
	}

	BinaryValueExpression struct {
		Node

		BinaryConcatenation *BinaryConcatenation
		BinaryFactor        *BinaryFactor
	}

	BinaryFactor struct {
		Node

		BinaryPrimary *BinaryPrimary
	}

	BinaryPrimary struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
		StringValueFunction    *StringValueFunction
	}

	BinaryConcatenation struct {
		Node

		BinaryValueExpression *BinaryValueExpression
		ConcatenationOperator token.Token
		BinaryFactor          *BinaryFactor
	}
)

// 6.32 string value function
type (
	StringValueFunction struct {
		Node

		CharacterValueFunction *CharacterValueFunction
		BinaryValueFunction    *BinaryValueFunction
		JSONValueConstructor   *JSONValueConstructor
		JSONQuery              *JSONQuery
	}

	CharacterValueFunction struct {
		Node

		CharacterSubstringFunction         *CharacterSubstringFunction
		RegularExpressionSubstringFunction *RegularExpressionSubstringFunction
		RegexSubstringFunction             *RegexSubstringFunction
		Fold                               *Fold
		Transcoding                        *Transcoding
		CharacterTransliteration           *CharacterTransliteration
		RegexTransliteration               *RegexTransliteration
		TrimFunction                       *TrimFunction
		CharacterOverlayFunction           *CharacterOverlayFunction
		NormalizeFunction                  *NormalizeFunction
		SpecificTypeMethod                 *SpecificTypeMethod
		ClassifierMethod                   *ClassifierMethod
	}

	CharacterSubstringFunction struct {
		Node

		Substring                token.Token
		LeftParen                token.Token
		CharacterValueExpression *CharacterValueExpression
		From                     token.Token
		StartPosition            *StartPosition
		For                      token.Token
		StringLength             *StringLength
		Using                    token.Token
		CharLengthUnits          *CharLengthUnits
		RightParen               token.Token
	}

	RegularExpressionSubstringFunction struct {
		Node

		Substring                 token.Token
		LeftParen                 token.Token
		CharacterValueExpression1 *CharacterValueExpression
		Similar                   token.Token
		CharacterValueExpression2 *CharacterValueExpression
		Escape                    token.Token
		EscapeCharacter           *EscapeCharacter
		RightParen                token.Token
	}

	RegexSubstringFunction struct {
		Node

		SubstringRegex     token.Token
		LeftParen          token.Token
		XQueryPattern      *XQueryPattern
		Flag               token.Token
		XQueryOptionFlag   *XQueryOptionFlag
		In                 token.Token
		RegexSubjectString *RegexSubjectString
		From               token.Token
		StartPosition      *StartPosition
		Using              token.Token
		CharLengthUnits    *CharLengthUnits
		Occurrence         token.Token
		RegexOccurrence    *RegexOccurrence
		Group              token.Token
		RegexCaptureGroup  *RegexCaptureGroup
		RightParen         token.Token
	}

	Fold struct {
		Node

		UpperOrLower             token.Token
		LeftParen                token.Token
		CharacterValueExpression *CharacterValueExpression
		RightParen               token.Token
	}

	Transcoding struct {
		Node

		Convert                  token.Token
		LeftParen                token.Token
		CharacterValueExpression *CharacterValueExpression
		Using                    token.Token
		TranscodingName          *TranscodingName
		RightParen               token.Token
	}

	CharacterTransliteration struct {
		Node

		Translate                token.Token
		LeftParen                token.Token
		CharacterValueExpression *CharacterValueExpression
		Using                    token.Token
		TransliterationName      *TransliterationName
		RightParen               token.Token
	}

	RegexTransliteration struct {
		Node

		TranslateRegex                 token.Token
		LeftParen                      token.Token
		XQueryPattern                  *XQueryPattern
		Flag                           token.Token
		XQueryOptionFlag               *XQueryOptionFlag
		In                             token.Token
		RegexSubjectString             *RegexSubjectString
		With                           token.Token
		XQueryReplacementString        *XQueryReplacementString
		From                           token.Token
		StartPosition                  *StartPosition
		Using                          token.Token
		CharLengthUnits                *CharLengthUnits
		Occurrence                     token.Token
		RegexTransliterationOccurrence *RegexTransliterationOccurrence
		RightParen                     token.Token
	}

	XQueryReplacementString struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	RegexTransliterationOccurrence struct {
		Node

		RegexOccurrence *RegexOccurrence
		All             token.Token
	}

	TrimFunction struct {
		Node

		Trim         token.Token
		LeftParen    token.Token
		TrimOperands *TrimOperands
		RightParen   token.Token
	}

	TrimOperands struct {
		Node

		TrimSpecification *TrimSpecification
		TrimCharacter     *TrimCharacter
		From              token.Token
		TrimSource        *TrimSource
	}

	TrimSource struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	TrimSpecification struct {
		Node

		LeadingOrTrailingOrBoth token.Token
	}

	TrimCharacter struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	CharacterOverlayFunction struct {
		Node

		Overlay                   token.Token
		LeftParen                 token.Token
		CharacterValueExpression1 *CharacterValueExpression
		Placing                   token.Token
		CharacterValueExpression2 *CharacterValueExpression
		From                      token.Token
		StartPosition             *StartPosition
		For                       token.Token
		StringLength              *StringLength
		Using                     token.Token
		CharLengthUnits           *CharLengthUnits
		RightParen                token.Token
	}

	NormalizeFunction struct {
		Node

		Normalize                     token.Token
		LeftParen                     token.Token
		CharacterValueExpression      *CharacterValueExpression
		Comma1                        token.Token
		NormalForm                    *NormalForm
		Comma2                        token.Token
		NormalizeFunctionResultLength *NormalizeFunctionResultLength
		RightParen                    token.Token
	}

	NormalForm struct {
		Node

		Nfc  token.Token
		Nfd  token.Token
		Nfkc token.Token
		Nfkd token.Token
	}

	NormalizeFunctionResultLength struct {
		Node

		CharacterLength            *CharacterLength
		CharacterLargeObjectLength *CharacterLargeObjectLength
	}

	SpecificTypeMethod struct {
		Node

		UserDefinedTypeValueExpression *UserDefinedTypeValueExpression
		Period                         token.Token
		Specifictype                   token.Token
		LeftParen                      token.Token
		RightParen                     token.Token
	}

	BinaryValueFunction struct {
		Node

		BinarySubstringFunction *BinarySubstringFunction
		BinaryTrimFunction      *BinaryTrimFunction
		BinaryOverlayFunction   *BinaryOverlayFunction
	}

	BinarySubstringFunction struct {
		Node

		Substring             token.Token
		LeftParen             token.Token
		BinaryValueExpression *BinaryValueExpression
		From                  token.Token
		StartPosition         *StartPosition
		For                   token.Token
		StringLength          *StringLength
		RightParen            token.Token
	}

	BinaryTrimFunction struct {
		Node

		Trim               token.Token
		LeftParen          token.Token
		BinaryTrimOperands *BinaryTrimOperands
		RighParen          token.Token
	}

	BinaryTrimOperands struct {
		Node

		TrimSpecification *TrimSpecification
		TrimOctet         *TrimOctet
		From              token.Token
		BinaryTrimSource  *BinaryTrimSource
	}

	BinaryTrimSource struct {
		Node

		BinaryValueExpression *BinaryValueExpression
	}

	TrimOctet struct {
		Node

		BinaryValueExpression *BinaryValueExpression
	}

	BinaryOverlayFunction struct {
		Node

		Overlay                token.Token
		LeftParen              token.Token
		BinaryValueExpression1 *BinaryValueExpression
		Placing                token.Token
		BinaryValueExpression2 *BinaryValueExpression
		From                   token.Token
		StartPosition          *StartPosition
		For                    token.Token
		StringLength           *StringLength
		RightParen             token.Token
	}

	StartPosition struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	StringLength struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	ClassifierFunction struct {
		Node

		Classifier             token.Token
		LeftParen              token.Token
		RowPatternVariableName *RowPatternVariableName
		RightParen             token.Token
	}
)

// 6.33 JSON value constructor
type (
	JSONValueConstructor struct {
		Node

		JSONObjectConstructor *JSONObjectConstructor
		JSONArrayConstructor  *JSONArrayConstructor
	}

	JSONObjectConstructor struct {
		Node

		JSONObject                  token.Token
		LeftParen                   token.Token
		JSONNameAndValue            []*JSONNameAndValue
		JSONConstructorNullClause   *JSONConstructorNullClause
		JSONKeyUniquenessConstraint *JSONKeyUniquenessConstraint
		JSONOutputClause            *JSONOutputClause
		RightParen                  token.Token
	}

	JSONNameAndValue struct {
		Node

		Key                 token.Token
		JSONName            *JSONName
		Value               token.Token
		Colon               token.Token
		JSONValueExpression *JSONValueExpression
	}

	JSONName struct {
		Node

		CharacterValueExpression *CharacterValueExpression
	}

	JSONConstructorNullClause struct {
		Node

		NullOrAbsent token.Token
		On           token.Token
		Null         token.Token
	}

	JSONArrayConstructor struct {
		Node

		JSONArrayConstructorByEnumeration *JSONArrayConstructorByEnumeration
		JSONArrayConstructorByQuery       *JSONArrayConstructorByQuery
	}

	JSONArrayConstructorByEnumeration struct {
		Node

		JSONArray                 token.Token
		LeftParen                 token.Token
		JSONValueExpression       []*JSONValueExpression
		JSONConstructorNullClause *JSONConstructorNullClause
		JSONOutputClause          *JSONOutputClause
		RightParen                token.Token
	}

	JSONArrayConstructorByQuery struct {
		Node

		JSONArray                 token.Token
		LeftParen                 token.Token
		QueryExpression           *QueryExpression
		JSONInputClause           *JSONInputClause
		JSONConstructorNullClause *JSONConstructorNullClause
		JSONOutputClause          *JSONOutputClause
		RightParen                token.Token
	}
)

// 6.34 JSON query
type (
	JSONQuery struct {
		Node

		LeftParen                 token.Token
		JSONAPICommonSyntax       *JSONAPICommonSyntax
		JSONOutputClause          *JSONOutputClause
		JSONQueryWrapperBehaviour *JSONQueryWrapperBehaviour
		Wrapper                   token.Token
		JSONQueryQuotesBehaviour  *JSONQueryQuotesBehaviour
		Quotes                    token.Token
		On                        token.Token
		Scalar                    token.Token
		String                    token.Token
		JSONQueryEmptyBehaviour   *JSONQueryEmptyBehaviour
		Empty                     token.Token
		JSONQueryErrorBehaviour   *JSONQueryErrorBehaviour
		Error                     token.Token
		RightParen                token.Token
	}

	JSONQueryWrapperBehaviour struct {
		Node

		Without                    token.Token
		With                       token.Token
		ConditionalOrUnconditional token.Token
		Array                      token.Token
	}

	JSONQueryQuotesBehaviour struct {
		Node

		KeepOrOmit token.Token
	}

	JSONQueryEmptyBehaviour struct {
		Node

		Error  token.Token
		Null   token.Token
		Empty  token.Token
		Array  token.Token
		Object token.Token
	}

	JSONQueryErrorBehaviour struct {
		Node

		Error  token.Token
		Null   token.Token
		Empty  token.Token
		Array  token.Token
		Object token.Token
	}
)

// 6.35 datetime value expression
type (
	DatetimeValueExpression struct {
		Node

		DatetimeTerm            *DatetimeTerm
		IntervalValueExpression *IntervalValueExpression
		PlusSign                token.Token
		MinusSign               token.Token
		DatetiemValueExpression *DatetiemValueExpression
		IntervalTerm            *IntervalTerm
	}

	DatetimeTerm struct {
		Node

		DatetimeFactor *DatetimeFactor
	}

	DatetimeFactor struct {
		Node

		DatetimePrimary *DatetimePrimary
		TimeZone        *TimeZone
	}

	DatetimePrimary struct {
		Node

		ValueExpressionPrimary *ValueExpressionPrimary
		DatetimeValueFunction  *DatetimeValueFunction
	}

	TimeZone struct {
		Node

		At                token.Token
		TimeZoneSpecifier *TimeZoneSpecifier
	}

	TimeZoneSpecifier struct {
		Node

		Local           token.Token
		Time            token.Token
		Zone            token.Token
		IntervalPrimary *IntervalPrimary
	}
)

// 6.36 datetime value function
type (
	DatetimeValueFunction struct {
		Node

		CurrentDateValueFunction           *CurrentDateValueFunction
		CurrentTimeValueFunction           *CurrentTimeValueFunction
		CurrentTimestampFunction           *CurrentTimestampFunction
		CurrentLocalTimeValueFunction      *CurrentLocalTimeValueFunction
		CurrentLocalTimestampValueFunction *CurrentLocalTimestampValueFunction
	}

	CurrentDateValueFunction struct {
		Node

		CurrentDate token.Token
	}

	CurrentTimeValueFunction struct {
		Node

		CurrentTime   token.Token
		LeftParen     token.Token
		TimePrecision *TimePrecision
		RightParen    token.Token
	}

	CurrentLocalTimeValueFunction struct {
		Node

		Localtime     token.Token
		LeftParen     token.Token
		TimePrecision *TimePrecision
		RightParen    token.Token
	}

	CurrentTimestampValueFunction struct {
		Node

		CurrentTimestamp   token.Token
		LeftParen          token.Token
		TimestampPrecision *TimestampPrecision
		RightParen         token.Token
	}

	CurrentLocalTimestampValueFunction struct {
		Node

		Localtimestamp     token.Token
		LeftParen          token.Token
		TimestampPrecision *TimestampPrecision
		RightParen         token.Token
	}
)

// 6.37 interval value expression
type (
	IntervalValueExpression struct {
		Node

		IntervalTerm             *IntervalTerm
		IntervalValueExpression1 *IntervalValueExpression
		PlusSignOrMinusSign      token.Token
		IntervalTerm1            *IntervalTerm
		LeftParen                token.Token
		DatetimeValueExpression  *DatetimeValueExpression
		DatetimeTerm             *DatetimeTerm
		RightParen               token.Token
		IntervalQualifier        *IntervalQualifier
	}

	IntervalTerm struct {
		Node

		IntervalFactor    *IntervalFactor
		IntervalTerm2     *IntervalTerm
		AsteriskOrSolidus token.Token
		Factor            *Factor
		Term              *Term
	}

	IntervalFactor struct {
		Node

		Sign            token.Token
		IntervalPrimary *IntervalPrimary
	}

	IntervalPrimary struct {
		Node

		ValueExpressionPrimary  *ValueExpressionPrimary
		IntervalQualifier       *IntervalQualifier
		IntervalValueExpression *IntervalValueExpression
	}
)

// 6.38 interval value function
type (
	IntervalValueFunction struct {
		Node

		IntervalAbsoluteValueFunction *IntervalAbsoluteValueFunction
	}

	IntervalAbsoluteValueFunction struct {
		Node

		Abs                     token.Token
		LeftParen               token.Token
		IntervalValueExpression *IntervalValueExpression
		RightParen              token.Token
	}
)

// 6.39 boolean value expression
type (
	BooleanValueExpression struct {
		Node

		BooleanValueExpression *BooleanValueExpression
		Or                     token.Token
		BooleanTerm            *BooleanTerm
	}

	BooleanTerm struct {
		Node

		BooleanTerm   *BooleanTerm
		And           token.Token
		BooleanFactor *BooleanFactor
	}

	BooleanFactor struct {
		Node

		Not         token.Token
		BooleanTest *BooleanTest
	}

	BooleanTest struct {
		Node

		BooleanPrimary *BooleanPrimary
		Is             token.Token
		Not            token.Token
		TruthValue     *TruthValue
	}

	TruthValue struct {
		Node

		TrueOrFalseOrUnknokwn token.Token
	}

	BooleanPrimary struct {
		Node

		Predicate        *Predicate
		BooleanPredicand *BooleanPredicand
	}

	BooleanPredicand struct {
		Node

		ParenthesizedBooleanValueExpression    *ParenthesizedBooleanValueExpression
		NonparenthesizedValueExpressionPrimary *NonparenthesizedValueExpressionPrimary
	}

	ParenthesizedBooleanValueExpression struct {
		Node

		LeftParen              token.Token
		BooleanValueExpression *BooleanValueExpression
		RightParen             token.Token
	}
)

// 6.40 array value expression
type (
	ArrayValueExpression struct {
		Node

		ArrayConcatenation *ArrayConcatenation
		ArrayPrimary       *ArrayPrimary
	}

	ArrayConcatenation struct {
		Node

		ArrayValueExpression1 *ArrayValueExpression
		ConcatenationOperator *ConcatenationOperator
		ArrayPrimary          *ArrayPrimary
	}

	ArrayPrimary struct {
		Node

		ArrayValueFunction     *ArrayValueFunction
		ValueExpressionPrimary *ValueExpressionPrimary
	}
)

// 6.41 array value function
type (
	ArrayValueFunction struct {
		Node

		TrimArrayFunction *TrimArrayFunction
	}

	TrimArrayFunction struct {
		Node

		TrimArray              token.Token
		LeftParen              token.Token
		ArrayValueExpression   *ArrayValueExpression
		Comma                  token.Token
		NumericValueExpression *NumericValueExpression
		RightParen             token.Token
	}
)

// 6.42 array value constructor
type (
	ArrayValueConstructor struct {
		Node

		ArrayValueConstructorByEnumeration *ArrayValueConstructorByEnumeration
		ArrayValueConstructorByQuery       *ArrayValueConstructorByQuery
	}

	ArrayValueConstructorByEnumeration struct {
		Node

		Array                  token.Token
		LeftBracketOrTrigraph  token.Token
		ArrayElementList       *ArrayElementList
		RightBracketOrTrigraph token.Token
	}

	ArrayElementList struct {
		Node

		ArrayElement []*ArrayElement
	}

	ArrayElement struct {
		Node

		ValueExpression *ValueExpression
	}

	ArrayValueConstructorByQuery struct {
		Node

		Array         token.Token
		TableSubquery *TableSubquery
	}
)

// 6.43 multiset value expression
type (
	MultisetValueExpression struct {
		Node

		MultisetValueExpression *MultisetValueExpression
		Multiset                token.Token
		UnionOrExcept           token.Token
		AllOrDistinct           token.Token
		MultisetTerm            *MultisetTerm
	}

	MultisetTerm struct {
		Node

		MultisetTerm    *MultisetTerm
		Multiset        token.Token
		Intersect       token.Token
		AllOrDistinct   token.Token
		MultisetPrimary *MultisetPrimary
	}

	MultisetPrimary struct {
		Node

		MultisetValueFunction  *MultisetValueFunction
		ValueExpressionPrimary *ValueExpressionPrimary
	}
)

// 6.44 multiset value function
type (
	MultisetValueFunction struct {
		Node

		MultisetSetFunction *MultisetSetFunction
	}

	MultisetSetFunction struct {
		Node

		Set                     token.Token
		LeftParen               token.Token
		MultisetValueExpression *MultisetValueExpression
		RightParen              token.Token
	}
)

// 6.45 multiset value constructor
type (
	MultisetValueConstrutor struct {
		Node

		MultisetValueConstructorByEnumeration *MultisetValueConstructorByEnumeration
		MultisetValueConstrutorByQuery        *MultisetValueConstrutorByQuery
		TableValueConstructorByQuery          *TableValueConstructorByQuery
	}

	MultisetValueConstructorByEnumeration struct {
		Node

		Multiset               token.Token
		LeftBracketOrTrigraph  token.Token
		MultisetElementList    *MultisetElementList
		RightBracketOrTrigraph token.Token
	}

	MultisetElementList struct {
		Node

		MultisetElement []*MultisetElement
	}

	MultisetElement struct {
		Node

		ValueExpression *ValueExpression
	}

	MultisetValueConstrutorByQuery struct {
		Node

		Multiset      token.Token
		TableSubquery *TableSubquery
	}

	TableValueConstructorByQuery struct {
		Node

		Table         token.Token
		TableSubquery *TableSubquery
	}
)
