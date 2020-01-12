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

	NumericType struct {
		Node

		ExactNumericType         *ExactNumericType
		ApproximateNumericType   *ApproximateNumericType
		DecimalFloatingPointType *DecimalFloatingPointType
	}

	ExactNumericType struct {
		Node

		Numeric    token.Token
		Decimal    token.Token
		Dec        token.Token
		Smallint   token.Token
		Integer    token.Token
		Int        token.Token
		Bigint     token.Token
		LeftParen  token.Token
		Precision  *Precision
		Comma      token.Token
		Scale      *Scale
		RightParen token.Token
	}

	ApproximateNumericType struct {
		Node

		Float      token.Token
		Real       token.Token
		Double     token.Token
		Precision  token.Token
		LeftParen  token.Token
		Precision  *Precision
		RightParen token.Token
	}

	DecimalFloatingPointType struct {
		Node

		Decfloat   token.Token
		LeftParen  token.Token
		Precision  *Precision
		RightParen token.Token
	}

	Length struct {
		Node

		UnsignedInteger *UnsignedInteger
	}

	CharacterLength struct {
		Node

		Length          *Length
		CharLengthUnits *CharLengthUnits
	}

	LargeObjectLength struct {
		Node

		UnsignedInteger        *UnsignedInteger
		Multiplier             token.Token
		LargeObjectLengthToken token.Token
	}

	CharacterLargeObjectLength struct {
		Node

		LargeObjectLength *LargeObjectLength
		CharLengthUnits   *CharLengthUnits
	}

	CharLengthUnits struct {
		Node

		Characters token.Token
		Octets     token.Token
	}

	Precision struct {
		Node

		UnsignedInteger *UnsignedInteger
	}

	Scale struct {
		Node

		UnsignedInteger *UnsignedInteger
	}

	BooleanType struct {
		Node

		Boolean token.Token
	}

	DatetimeType struct {
		Node

		Date                  token.Token
		Time                  token.Token
		Timestamp             token.Token
		LeftParen             token.Token
		TimePrecision         *TimePrecision
		RightParen            token.Token
		WithOrWithoutTimeZone *WithOrWithoutTimeZone
		TimestampPrecision    *TimestampPrecision
	}

	WithOrWithoutTimeZone struct {
		Node

		With    token.Token
		Without token.Token
		Time    token.Token
		Zone    token.Token
	}

	TimePrecision struct {
		Node

		TimeFractionalSecondsPrecision *TimeFractionalSecondsPrecision
	}

	TimestampPrecision struct {
		Node

		TimeFractionalSecondsPrecision *TimeFractionalSecondsPrecision
	}

	TimeFractionalSecondsPrecision struct {
		Node

		UnsignedInteger *UnsignedInteger
	}

	IntervalType struct {
		Node

		Interval          token.Token
		IntervalQualifier *IntervalQualifier
	}

	RowType struct {
		Node

		Row         token.Token
		RowTypeBody *RowTypeBody
	}

	RowTypeBody struct {
		Node

		LeftParen       token.Token
		FieldDefinition []*FieldDefinition
		RightParen      token.Token
	}

	ReferenceType struct {
		Node

		Ref            token.Token
		LeftParen      token.Token
		ReferencedType *ReferencedType
		RightParen     token.Token
		ScopeClause    *ScopeClause
	}

	ScopeClause struct {
		Node

		Scope     token.Token
		TableName *TableName
	}

	ReferencedType struct {
		Node

		PathResolvedUserDefinedTypeName *PathResolvedUserDefinedTypeName
	}

	PathResolvedUserDefinedTypeName struct {
		Node

		UserDefinedTypeName *UserDefinedTypeName
	}

	CollectionType struct {
		Node

		ArrayType    *ArrayType
		MultisetType *MultisetType
	}

	ArrayType struct {
		Node

		DataType               *DataType
		Array                  token.Token
		LeftBracketOrTrigraph  token.Token
		MaximumCardinality     *MaximumCardinality
		RightBracketOrTrigraph token.Token
	}

	MaximumCardinality struct {
		Node

		UnsignedInteger *UnsignedInteger
	}

	MultisetType struct {
		Node

		DataType *DataType
		Multiset token.Token
	}
)

// 6.2 field definition
type (
	FieldDefinition struct {
		Node

		FieldName *FieldName
		DataType  *DataType
	}
)

// 6.3 value expression primary
type (
	ValueExpressionPrimary struct {
		Node

		ParenthesizedValueExpression           *ParenthesizedValueExpression
		NonparenthesizedValueExpressionPrimary *NonparenthesizedValueExpressionPrimary
	}

	ParenthesizedValueExpression struct {
		Node

		LeftParen       token.Token
		ValueExpression *ValueExpression
		RightParen      token.Token
	}

	NonparenthesizedValueExpressionPrimary struct {
		UnsignedValueSpecification    *UnsignedValueSpecification
		ColumnReference               *ColumnReference
		SetFunctionSpecification      *SetFunctionSpecification
		WindowFunction                *WindowFunction
		NestedWindowFunction          *NestedWindowFunction
		ScalarSubquery                *ScalarSubquery
		CaseExpression                *CaseExpression
		CastSpecification             *CastSpecification
		FieldReference                *FieldReference
		SubtypeTreatment              *SubtypeTreatment
		MethodInvocation              *MethodInvocation
		StaticMethodInvocation        *StaticMethodInvocation
		NewSpecification              *NewSpecification
		AttributeOrMethodReference    *AttributeOrMethodReference
		ReferenceResolution           *ReferenceResolution
		CollectionValueConstructor    *CollectionValueConstructor
		ArrayElementReference         *ArrayElementReference
		MultisetElementReference      *MultisetElementReference
		NextValueExpression           *NextValueExpression
		RoutineInvocation             *RoutineInvocation
		RowPatternNavigationOperation *RowPatternNavigationOperation
		JSONValueFunction             *JSONValueFunction
	}

	CollectionValueConstructor struct {
		Node

		ArrayValueConstructor    *ArrayValueConstructor
		MultisetValueConstructor *MultisetValueConstructor
	}
)

// 6.4 value specification and target specification
type (
	ValueSpecification struct {
		Node

		Literal                   *Literal
		GeneralValueSpecification *GeneralValueSpecification
	}

	UnsignedValueSpecification struct {
		Node

		UnsignedLiteral           *UnsignedLiteral
		GeneralValueSpecification *GeneralValueSpecification
	}

	GeneralValueSpecification struct {
		Node

		HostParameterSpecification      *HostParameterSpecification
		SQLParameterReference           *SQLParameterReference
		SynamicParameterSpecification   *SynamicParameterSpecification
		EmbeddedVariableSpecification   *EmbeddedVariableSpecification
		CurrentCollationSpecification   *CurrentCollationSpecification
		CurrentCatalog                  token.Token
		CurrentDefaultTransformGroup    token.Token
		CurrentPath                     token.Token
		CurrentRole                     token.Token
		CurrentSchema                   token.Token
		CurrentTransformGroupForType    token.Token
		PathResolvedUserDefinedTypeName *PathResolvedUserDefinedTypeName
		CurrentUser                     token.Token
		SessionUser                     token.Token
		SystemUser                      token.Token
		User                            token.Token
		Value                           token.Token
	}

	SimpleValueSpecification struct {
		Node

		Literal               *Literal
		HostParameterName     *HostParameterName
		SQLParameterReference *SQLParameterReference
		EmbeddedVariableName  *EmbeddedVariableName
	}

	TargetSpecification struct {
		Node

		HostParameterSpecification      *HostParameterSpecification
		SQLParameterReference           *SQLParameterReference
		ColumnReference                 *ColumnReference
		TargetArrayElementSpecification *TargetArrayElementSpecification
		DynamicParameterSpecification   *DynamicParameterSpecification
		EmbeddedVariableSpecification   *EmbeddedVariableSpecification
	}

	SimpleTargetSpecification struct {
		Node

		HostParameterName     *HostParameterName
		SQLParameterReference *SQLParameterReference
		ColumnReference       *ColumnReference
		EmbeddedVariableName  *EmbeddedVariableName
	}

	HostParameterSpecification struct {
		Node

		HostParameterName  *HostParameterName
		IndicatorParameter *IndicatorParameter
	}

	DynamicParameterSpecification struct {
		Node

		QuestionMark token.Token
	}

	EmbeddedVariableSpecification struct {
		Node

		EmbeddedVariableName *EmbeddedVariableName
		IndicatorVariable    *IndicatorVariable
	}

	IndicatorVariable struct {
		Node

		Indicator            token.Token
		EmbeddedVariableName *EmbeddedVariableName
	}

	IndicatorParameter struct {
		Node

		Indicator         token.Token
		HostParameterName *HostParameterName
	}

	TargetArrayElementSpecification struct {
		Node

		TargetArrayReference     *TargetArrayReference
		LeftBracketOrTrigraph    token.Token
		SimpleValueSpecification *SimpleValueSpecification
		RightBracketOrTrigraph   token.Token
	}

	TargetArrayReference struct {
		Node

		SQLParameterReference *SQLParameterReference
		ColumnReference       *ColumnReference
	}

	CurrentCollationSpecification struct {
		Node

		Collation             token.Token
		For                   token.Token
		LeftParen             token.Token
		StringValueExpression *StringValueExpression
		RightParen            token.Token
	}
)

// 6.5 contextually typed value specification
type (
	ContextuallyTypedValueSpecification struct {
		Node

		ImplicitlyTypedValueSpecification *ImplicitlyTypedValueSpecification
		DefaultSpecification              *DefaultSpecification
	}

	ImplicitlyTypedValueSpecification struct {
		Node

		NullSpecification  *NullSpecification
		EmptySpecification *EmptySpecification
	}

	NullSpecification struct {
		Node

		Null token.Token
	}

	EmptySpecification struct {
		Node

		Array                  token.Token
		Multiset               token.Token
		LeftBracketOrTrigraph  token.Token
		RightBracketOrTrigraph token.Token
	}
)

// 6.6 identifier chain
type (
	IdentifierChain struct {
		Node

		Identifier []*Identifier
	}

	BasicIdentifierChain struct {
		Node

		IdentifierChain *IdentifierChain
	}
)

// 6.7 column reference
type (
	ColumnReference struct {
		Node

		BasicIdentifierChain *BasicIdentifierChain
		Module               token.Token
		Period1              token.Token
		QualifiedIdentifier  *QualifiedIdentifier
		Period2              token.Token
		ColumnName           *ColumnName
	}
)

// 6.8 SQL parameter reference
type (
	SQLParameterReference struct {
		Node

		BasicIdentifierChain *BasicIdentifierChain
	}
)

// 6.9 set function specification
type (
	SetFunctionSpecification struct {
		Node

		RunningOrFinal    *RunningOrFinal
		AggregateFunction *AggregateFunction
		GroupingOperation *GroupingOperation
	}

	RunningOrFinal struct {
		Node

		Running token.Token
		Final   token.Token
	}

	GroupingOperation struct {
		Node

		Grouping        token.Token
		LeftParen       token.Token
		ColumnReference []*ColumnReference
		RightParen      token.Token
	}
)

// 6.10 window function
type (
	WindowFunction struct {
		Node

		WindowFunctionType        *WindowFunctionType
		Over                      token.Token
		WindowNameOrSpecification *WindowNameOrSpecification
	}
)

// TODO(TimSatke): continue
