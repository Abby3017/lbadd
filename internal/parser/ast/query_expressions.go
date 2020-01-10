package ast

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

// 7 Query expressions

// 7.1 row value constructor
type (
	RowValueConstructor struct {
		Node

		CommonValueExpression       *CommonValueExpression
		BooleanValueExpression      *BooleanValueExpression
		ExplicitRowValueConstructor *ExplicitRowValueConstructor
	}

	ExplicitRowValueConstructor struct {
		Node

		LeftParen                      token.Token
		RowValueConstructorElement     *RowValueConstructorElement
		Comma                          token.Token
		RowValueConstructorElementList *RowValueConstructorElementList
		RightParen                     token.Token
		Row                            token.Token
		RowSubquery                    *RowSubquery
	}

	RowValueConstructorElementList struct {
		Node

		RowValueConstructorElement []*RowValueConstructorElement
	}

	RowValueConstructorElement struct {
		Node

		ValueExpression *ValueExpression
	}

	ContextuallyTypedRowValueConstructorElementList struct {
		Node

		ContextuallyTypedRowValueConstructorElement []*ContextuallyTypedRowValueConstructorElement
	}

	ContextuallyTypedRowValueConstructorElement struct {
		Node

		ValueExpression                     *ValueExpression
		ContextuallyTypedValueSpecification *ContextuallyTypedValueSpecification
	}

	RowValueConstructorPredicand struct {
		Node

		CommonValueExpression       *CommonValueExpression
		BooleanPredicand            *BooleanPredicand
		ExplicitRowValueConstructor *ExplicitRowValueConstructor
	}
)

// 7.2 row value expression
type (
	RowValueExpression struct {
		Node

		RowValueSpecialCase         *RowValueSpecialCase
		ExplicitRowValueConstructor *ExplicitRowValueConstructor
	}

	TableRowValueExpression struct {
		Node

		RowValueSpecialCase *RowValueSpecialCase
		RowValueConstructor *RowValueConstructor
	}

	ContextuallyTypedRowValueExpression struct {
		Node

		RowValueSpecialCase                 *RowValueSpecialCase
		ContextuallyTypedRowValueExpression *ContextuallyTypedRowValueExpression
	}

	RowValuePredicand struct {
		Node

		RowValueSpecialCase          *RowValueSpecialCase
		RowValueConstructorPredicand *RowValueConstructorPredicand
	}

	RowValueSpecialCase struct {
		Node

		NonparenthesizedValueExpressionPrimary *NonparenthesizedValueExpressionPrimary
	}
)

// 7.3 table value constructor
type (
	TableValueConstructor struct {
		Node

		Values                 token.Token
		RowValueExpressionList *RowValueExpressionList
	}

	RowValueExpressionList struct {
		Node

		TableRowValueExpression []*TableRowValueExpression
	}

	ContextuallyTypedTableValueConstructor struct {
		Node

		Values                                  token.Token
		ContextuallyTypedRowValueExpressionList *ContextuallyTypedRowValueExpressionList
	}

	ContextuallyTypedRowValueExpressionList struct {
		Node

		ContextuallyTypedRowValueExpression []*ContextuallyTypedRowValueExpression
	}
)

// 7.4 table expression
type (
	TableExpression struct {
		Node

		FromClause    *FromClause
		WhereClause   *WhereClause
		GroupByClause *GroupByClause
		HavingClause  *HavingClause
		WindowClause  *WindowClause
	}
)

// 7.5 from clause
type (
	FromClause struct {
		Node

		From               token.Token
		TableReferenceList *TableReferenceList
	}

	TableReferenceList struct {
		Node

		TableReference []*TableReference
	}
)

// 7.6 table reference
type (
	TableReference struct {
		Node

		TableFactor *TableFactor
		JoinedTable *JoinedTable
	}

	TableFactor struct {
		Node

		TablePrimary *TablePrimary
		SampleClause *SampleClause
	}

	SampleClause struct {
		Node

		Tablesample      token.Token
		SampleMethod     *SampleMethod
		LeftParen        token.Token
		SamplePercentage *SamplePercentage
		RightParen       token.Token
		RepeatableClause *RepeatableClause
	}

	SampleMethod struct {
		Node

		Bernoulli token.Token
		System    token.Token
	}

	RepeatableClause struct {
		Node

		Repeatable     token.Token
		LeftParen      token.Token
		RepeatArgument *RepeatArgument
		RightParen     token.Token
	}

	SamplePercentage struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	RepeatArgument struct {
		Node

		NumericValueExpression *NumericValueExpression
	}

	TablePrimary struct {
		Node

		TableOrQueryName                   *TableOrQueryName
		QuerySystemTimePeriodSpecification *QuerySystemTimePeriodSpecification
		DerivedTable                       *DerivedTable
		LateralDerivedTable                *LateralDerivedTable
		CollectionDerivedTable             *CollectionDerivedTable
		TableFunctionDerivedTable          *TableFunctionDerivedTable
		PTFDerivedTable                    *PTFDerivedTable
		OnlySpec                           *OnlySpec
		DataChangeDeltaTable               *DataChangeDeltaTable
		JSONTable                          *JSONTable
		JSONTablePrimitive                 *JSONTablePrimitive
		CorrelationOrRecognition           *CorrelationOrRecognition
		CorrelationName                    *CorrelationName
		ParenthesizedJoinedTable           *ParenthesizedJoinedTable
	}

	CorrelationOrRecognition struct {
		Node

		As                                 token.Token
		CorrelationName                    *CorrelationName
		ParenthesizedDerivedColumnList     *ParenthesizedDerivedColumnList
		RowPatternRecognitionClauseAndName *RowPatternRecognitionClauseAndName
	}

	QuerySystemTimePeriodSpecification struct {
		Node

		For                   token.Token
		SystemTime            token.Token
		As                    token.Token
		Of                    token.Token
		Between               token.Token
		AsymmetricOrSymmetric token.Token
		PointInTime1          *PointInTime
		PointInTime2          *PointInTime
		From                  token.Token
		To                    token.Token
	}

	PointInTime struct {
		Node

		DatetimeValueExpression *DatetimeValueExpression
	}

	OnlySpec struct {
		Node

		Only             token.Token
		LeftParen        token.Token
		TableOrQueryName *TableOrQueryName
		RightParen       token.Token
	}

	LateralDerivedTable struct {
		Node

		Lateral       token.Token
		TableSubquery *TableSubquery
	}

	CollectionDerivedTable struct {
		Node

		Unnest                    token.Token
		LeftParen                 token.Token
		CollectionValueExpression []*CollectionValueExpression
		RightParen                token.Token
		With                      token.Token
		Ordinality                token.Token
	}

	TableFunctionDerivedTable struct {
		Node

		Table                     token.Token
		LeftParen                 token.Token
		CollectionValueExpression *CollectionValueExpression
		RightParen                token.Token
	}

	DerivedTable struct {
		Node

		TableSubquery *TableSubquery
	}

	PTFDerivedTable struct {
		Node

		Table             token.Token
		LeftParen         token.Token
		RoutineInvocation *RoutineInvocation
		RightParen        token.Token
	}

	TableOrQueryName struct {
		Node

		TableName           *TableName
		TransitionTableName *TransitionTableName
		QueryName           *QueryName
	}

	DerivedColumnList struct {
		Node

		ColumnNameList *ColumnNameList
	}

	ColumnNameList struct {
		Node

		ColumnName []*ColumnName
	}

	DataChangeDeltaTable struct {
		Node

		ResultOption        *ResultOption
		Table               token.Token
		LeftParen           token.Token
		DataChangeStatement *DataChangeStatement
		RightParen          token.Token
	}

	DataChangeStatement struct {
		Node

		DeleteStatementSearched *DeleteStatementSearched
		InsertStatement         *InsertStatement
		MergeStatement          *MergeStatement
		UpdateStatementSearched *UpdateStatementSearched
	}

	ResultOption struct {
		Node

		FinalOrNewOrOld token.Token
	}

	ParenthesizedJoinedTable struct {
		Node

		LeftParen                token.Token
		ParenthesizedJoinedTable *ParenthesizedJoinedTable
		JoinedTable              *JoinedTable
		RightParen               token.Token
	}

	RowPatternRecognitionClauseAndName struct {
		Node

		As1                               token.Token
		RowPatternInputName               *RowPatternInputName
		RowPatternInputDerivedColumnList  *RowPatternInputDerivedColumnList
		RowPatternRecognitionClause       *RowPatternRecognitionClause
		As2                               token.Token
		RowPatternOutputName              *RowPatternOutputName
		RowPatternOutputDerivedColumnList *RowPatternOutputDerivedColumnList
	}

	RowPatternInputName struct {
		Node

		CorrelationName *CorrelationName
	}

	RowPatternOutputName struct {
		Node

		CorrelationName *CorrelationName
	}

	RowPatternInputDerivedColumnList struct {
		Node

		ParenthesizedDerivedColumnList *ParenthesizedDerivedColumnList
	}

	RowPatternOutputDerivedColumnList struct {
		Node

		ParenthesizedDerivedColumnList *ParenthesizedDerivedColumnList
	}

	ParenthesizedDerivedColumnList struct {
		Node

		LeftParen         token.Token
		DerivedColumnList *DerivedColumnList
		RightParen        token.Token
	}
)

// 7.7 row pattern recognition clause
type (
	RowPatternRecognitionClause struct {
		Node

		MatchRecognize         token.Token
		LeftParen              token.Token
		RowPatternPartitionBy  *RowPatternPartitionBy
		RowPatternOrderBy      *RowPatternOrderBy
		RowPatternMeasures     *RowPatternMeasures
		RowPatternRowsPerMatch *RowPatternRowsPerMatch
		RowPatternCommonSyntax *RowPatternCommonSyntax
		RightParen             token.Token
	}

	RowPatternPartitionBy struct {
		Node

		Partition               token.Token
		By                      token.Token
		RowPatternPartitionList *RowPatternPartitionList
	}

	RowPatternPartitionList struct {
		Node

		RowPatternPartitionColumn []*RowPatternPartitionColumn
	}

	RowPatternPartitionColumn struct {
		Node

		ColumnReference *ColumnReference
		CollateClause   *CollateClause
	}

	RowPatternOrderBy struct {
		Node

		Order                 token.Token
		By                    token.Token
		SortSpecificationList *SortSpecificationList
	}

	RowPatternRowsPerMatch struct {
		Node

		One                          token.Token
		Row                          token.Token
		All                          token.Token
		Rows                         token.Token
		Per                          token.Token
		Match                        token.Token
		RowPatternEmptyMatchHandling *RowPatternEmptyMatchHandling
	}

	RowPatternEmptyMatchHandling struct {
		Node

		ShowOrOmit token.Token
		Empty      token.Token
		Matches    token.Token
		With       token.Token
		Unmatched  token.Token
		Rows       token.Token
	}
)

// 7.8 row pattern measures
type (
	RowPatternMeasures struct {
		Node

		Measures              token.Token
		RowPatternMeasureList *RowPatternMeasureList
	}

	RowPatternMeasureList struct {
		Node

		RowPatternMeasureDefinition []*RowPatternMeasureDefinition
	}

	RowPatternMeasureDefinition struct {
		Node

		RowPatternMeasureExpression *RowPatternMeasureExpression
		As                          token.Token
		MeasureName                 *MeasureName
	}

	RowPatternMeasureExpression struct {
		Node

		ValueExpression *ValueExpression
	}
)

// 7.9 row pattern common syntax
type (
	RowPatternCommonSyntax struct {
		Node

		After                    token.Token
		Match                    token.Token
		RowPatternSkipTo         *RowPatternSkipTo
		RowPatternInitialOrSeek  *RowPatternInitialOrSeek
		Pattern                  token.Token
		LeftParen                token.Token
		RowPattern               *RowPattern
		RightParen               token.Token
		RowPatternSubsetClause   *RowPatternSubsetClause
		Definex                  token.Token
		RowPatternDefinitionList *RowPatternDefinitionList
	}

	RowPatternSkipTo struct {
		Node

		Skip                         token.Token
		To                           token.Token
		Past                         token.Token
		Next                         token.Token
		Last                         token.Token
		First                        token.Token
		Row                          token.Token
		RowPatternSkipToVariableName *RowPatternSkipToVariableName
	}

	RowPatternSkipToVariableName struct {
		Node

		RowPatternVariableName *RowPatternVariableName
	}

	RowPatternInitialOrSeek struct {
		Node

		InitialOrSeek token.Token
	}

	RowPattern struct {
		Node

		RowPatternTerm        *RowPatternTerm
		RowPatternAlternation *RowPatternAlternation
	}

	RowPatternAlternation struct {
		Node

		RowPattern     *RowPattern
		VerticalBar    token.Token
		RowPatternTerm *RowPatternTerm
	}

	RowPatternTerm struct {
		Node

		RowPatternTerm   *RowPatternTerm
		RowPatternFactor *RowPatternFactor
	}

	RowPatternFactor struct {
		Node

		RowPatternPrimary    *RowPatternPrimary
		RowPatternQuantifier *RowPatternQuantifier
	}

	RowPatternQuantifier struct {
		Node

		Asterisk             token.Token
		PlusSign             token.Token
		QuestionMark1        token.Token
		LeftBrace            token.Token
		UnsignedInteger1     *UnsignedInteger
		Comma                token.Token
		UnsignedInteger2     *UnsignedInteger
		RightBrace           token.Token
		QuestionMarkTrailing token.Token
	}

	RowPatternPrimary struct {
		Node

		RowPatternPrimaryVariableName *RowPatternPrimaryVariableName
		DollarSign                    token.Token
		Circumflex                    token.Token
		LeftParen                     token.Token
		RowPattern                    *RowPattern
		RightParen                    token.Token
		LeftBraceMinus                token.Token
		RowPattern                    *RowPattern
		RightMinusBrace               token.Token
		RowPatternPermute             *RowPatternPermute
	}

	RowPatternPrimaryVariableName struct {
		Node

		RowPatternVariableName *RowPatternVariableName
	}

	RowPatternPermute struct {
		Node

		Permute    token.Token
		LeftParen  token.Token
		RowPattern []*RowPattern
		RightParen token.Token
	}

	RowPatternSubsetClause struct {
		Node

		Subset               token.Token
		RowPatternSubsetList *RowPatternSubsetList
	}

	RowPatternSubsetList struct {
		Node

		RowPatternSubsetItem []*RowPatternSubsetItem
	}

	RowPatternSubsetItem struct {
		Node

		RowPatternSubsetItemVariableName *RowPatternSubsetItemVariableName
		EqualsOperator                   token.Token
		LeftParen                        token.Token
		RowPatternSubsetRhs              *RowPatternSubsetRhs
		RightParen                       token.Token
	}

	RowPatternSubsetItemVariableName struct {
		Node

		RowPatternVariableName *RowPatternVariableName
	}

	RowPatternSubsetRhs struct {
		Node

		RowPatternSubsetRhsVariableName []*RowPatternSubsetRhsVariableName
	}

	RowPatternSubsetRhsVariableName struct {
		Node

		RowPatternVariableName *RowPatternVariableName
	}
)
