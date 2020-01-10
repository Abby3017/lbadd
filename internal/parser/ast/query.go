package ast

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

type (
	// Query is the ast root of any SQL statement.
	Query struct {
		Node

		QueryExpression *QueryExpression
	}

	QueryExpression struct {
		Node

		WithClause          *WithClause
		QueryExpressionBody *QueryExpressionBody
		OrderByClause       *OrderByClause
		ResultOffsetClause  *ResultOffsetClause
		FetchFirstClause    *FetchFirstClause
	}

	QueryExpressionBody struct {
		Node

		QueryExpressionBody *QueryExpressionBody
		UnionOrExcept       token.Token
		AllOrDistinct       token.Token
		CorrespondingSpec   *CorrespondingSpec
		QueryTerm           *QueryTerm
	}

	QueryTerm struct {
		Node

		QueryTerm         *QueryTerm
		Intersect         token.Token
		AllOrDistinct     token.Token
		CorrespondingSpec *CorrespondingSpec
		QueryPrimary      *QueryPrimary
	}

	QueryPrimary struct {
		Node

		SimpleTable         *SimpleTable
		LeftParen           token.Token
		QueryExpressionBody *QueryExpressionBody
		OrderByClause       *OrderByClause
		ResultOffsetClause  *ResultOffsetClause
		FetchFirstClause    *FetchFirstClause
		RightParen          token.Token
	}

	SimpleTable struct {
		Node

		QuerySpecification    *QuerySpecification
		TableValueConstructor *TableValueConstructor
		ExplicitTable         *ExplicitTable
	}

	ExplicitTable struct {
		Node

		Table            token.Token
		TableOrQueryName *TableOrQueryName
	}

	CorrespondingSpec struct {
		Node

		Corresponding           token.Token
		By                      token.Token
		LeftParen               token.Token
		CorrespondingColumnList *CorrespondingColumnList
		RightParen              token.Token
	}

	CorrespondingColumnList struct {
		Node
	}

	OrderByClause struct {
		Node

		Order                 token.Token
		By                    token.Token
		SortSpecificationList *SortSpecificationList
	}

	ResultOffsetClause struct {
		Node

		Offset         token.Token
		OffsetRowCount *OffsetRowCount
		RowOrRows      token.Token
	}

	FetchFirstClause struct {
		Node

		Fetch              token.Token
		FirstOrNext        token.Token
		FetchFirstQuantity *FetchFirstQuantity
		RowOrRows          token.Token
		OnlyOrWithTies     token.Token
	}

	FetchFirstQuantity struct {
		Node

		FetchFirstRowCount   *FetchFirstRowCount
		FetchFirstPercentage *FetchFirstPercentage
	}

	OffsetRowCount struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
	}

	FetchFirstRowCount struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
	}

	FetchFirstPercentage struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
		Percent                  token.Token
	}

	QuerySpecification struct {
		Node

		Select          token.Token
		SetQuantifier   *SetQuantifier
		SelectList      *SelectList
		TableExpression *TableExpression
	}

	SelectList struct {
		Node

		Asterisk      token.Token
		SelectSublist []*SelectSublist
	}

	SelectSublist struct {
		Node

		DerivedColumn     *DerivedColumn
		QualifiedAsterisk *QualifiedAsterisk
	}

	QualifiedAsterisk struct {
		Node

		AsteriskedIdentifierChain *AsteriskedIdentifierChain
		Period                    token.Token
		Asterisk                  token.Token
		AllFieldsReference        *AllFieldsReference
	}

	AsteriskedIdentifierChain struct {
		Node

		AsteriskedIdentifier []*AsteriskedIdentifier
	}

	AsteriskedIdentifier struct {
		Node

		Identifier *Identifier
	}

	DerivedColumn struct {
		Node

		ValueExpression *ValueExpression
		AsClause        *AsClause
	}

	AsClause struct {
		Node

		As         token.Token
		ColumnName *ColumnName
	}

	AllFieldsReference struct {
		Node

		ValueExpressionPrimary  *ValueExpressionPrimary
		Period                  token.Token
		Asterisk                token.Token
		As                      token.Token
		LeftParen               token.Token
		AllFieldsColumnNameList *AllFieldsColumnNameList
		RightParen              token.Token
	}

	AllFieldsColumnNameList struct {
		Node

		ColumnNameList *ColumnNameList
	}
)

func (i *Identifier) String() string { return string(i) }
