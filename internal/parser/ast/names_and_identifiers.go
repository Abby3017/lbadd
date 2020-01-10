package ast

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

// 5.4 names and identifiers
type (
	Identifier struct {
		Node

		ActualIdentifier *ActualIdentifier
	}

	ActualIdentifier struct {
		Node

		RegularIdentifier          *RegularIdentifier
		DelimitedIdentifier        *DelimitedIdentifier
		UnicodeDelimitedIdentifier *UnicodeDelimitedIdentifier
	}

	SQLLanguageIdentifier struct {
		Node

		SQLLanguageIdentifierStart *SQLLanguageIdentifierStart
		SQLLanguageIdentifierPart  []*SQLLanguageIdentifierPart
	}

	SQLLanguageIdentifierStart struct {
		Node

		SimpleLatinLetter token.Token
	}

	SQLLanguageIdentifierPart struct {
		Node

		SimpleLatinLetter token.Token
		Digit             token.Token
		Underscore        token.Token
	}

	AuthorizationIdentifier struct {
		Node

		RoleName       *RoleName
		UserIdentifier *UserIdentifier
	}

	TableName struct {
		Node

		LocalOrSchemaQualifiedName *LocalOrSchemaQualifiedName
	}

	DomainName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	SchemaName struct {
		Node

		CatalogName           *CatalogName
		Period                token.Token
		UnqualifiedSchemaName *UnqualifiedSchemaName
	}

	UnqualifiedSchemaName struct {
		Node

		Identifier *Identifier
	}

	CatalogName struct {
		Node

		Identifier *Identifier
	}

	SchemaQualifiedName struct {
		Node

		SchemaName          *SchemaName
		Period              token.Token
		QualifiedIdentifier *QualifiedIdentifier
	}

	LocalOrSchemaQualifiedName struct {
		Node

		LocalOrSchemaQualifier *LocalOrSchemaQualifier
		Period                 token.Token
		QualifiedIdentifier    *QualifiedIdentifier
	}

	LocalOrSchemaQualifier struct {
		Node

		SchemaName     *SchemaName
		LocalQualifier *LocalQualifier
	}

	QualifiedIdentifier struct {
		Node

		Identifier *Identifier
	}

	ColumnName struct {
		Node

		Identifier *Identifier
	}

	CorrelationName struct {
		Node

		Identifier *Identifier
	}

	QueryName struct {
		Node

		Identifier *Identifier
	}

	SQLClientModuleName struct {
		Node

		Identifier *Identifier
	}

	ProcedureName struct {
		Node

		Identifier *Identifier
	}

	SchemaQualifiedRoutineName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	MethodName struct {
		Node

		Identifier *Identifier
	}

	SpecificName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	CursorName struct {
		Node

		LocalQualifiedName *LocalQualifiedName
	}

	LocalQualifiedName struct {
		Node

		LocalQualifier      *LocalQualifier
		Period              token.Token
		QualifiedIdentifier *QualifiedIdentifier
	}

	LocalQualifier struct {
		Node

		Module token.Token
	}

	HostParameterName struct {
		Node

		Colon      token.Token
		Identifier *Identifier
	}

	SQLParameterName struct {
		Node

		Identifier *Identifier
	}

	ConstraintName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	ExternalRoutineName struct {
		Node

		Identifier             *Identifier
		CharacterStringLiteral token.Token
	}

	TriggerName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	CharacterSetName struct {
		Node

		SchemaName            *SchemaName
		Period                token.Token
		SQLLanguageIdentifier *SQLLanguageIdentifier
	}

	TransliterationName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	TranscodingName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	SchemaResolvedUserDefinedTypeName struct {
		Node

		UserDefinedTypeName *UserDefinedTypeName
	}

	UserDefinedTypeName struct {
		Node

		SchemaName          *SchemaName
		Period              token.Token
		QualifiedIdentifier *QualifiedIdentifier
	}

	AttributeName struct {
		Node

		Identifier *Identifier
	}

	FieldName struct {
		Node

		Identifier *Identifier
	}

	SavepointName struct {
		Node

		Identifier *Identifier
	}

	SequenceGeneratorName struct {
		Node

		SchemaQualifiedName *SchemaQualifiedName
	}

	RoleName struct {
		Node

		Identifier *Identifier
	}

	UserIdentifier struct {
		Node

		Identifier *Identifier
	}

	ConnectionName struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
	}

	SQLServerName struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
	}

	ConnectionUserName struct {
		Node

		SimpleValueSpecification *SimpleValueSpecification
	}

	SQLStatementName struct {
		Node

		StatementName         *StatementName
		ExtendedStatementName *ExtendedStatementName
	}

	StatementName struct {
		Node

		Identifier *Identifier
	}

	ExtendedStatementName struct {
		Node

		ScopeOption              *ScopeOption
		SimpleValueSpecification *SimpleValueSpecification
	}

	DynamicCursorName struct {
		Node

		ConventionalDynamicCursorName *ConventionalDynamicCursorName
		PTFCursorName                 *PTFCursorName
	}

	ConventionalDynamicCursorName struct {
		Node

		CursorName         *CursorName
		ExtendedCursorName *ExtendedCursorName
	}

	ExtendedCursorName struct {
		Node

		ScopeOption              *ScopeOption
		SimpleValueSpecification *SimpleValueSpecification
	}

	PTFCursorName struct {
		Node

		Ptf                      token.Token
		SimpleValueSpecification *SimpleValueSpecification
	}

	DescriptorName struct {
		Node

		ConventionalDescriptorName *ConventionalDescriptorName
		PTFDescriptorName          *PTFDescriptorName
	}

	ConventionalDescriptorName struct {
		Node

		NonExtendedDescriptorName *NonExtendedDescriptorName
		ExtendedDescriptorName    *ExtendedDescriptorName
	}

	NonExtendedDescriptorName struct {
		Node

		Identifier *Identifier
	}

	ExtendedDescriptorName struct {
		Node

		ScopeOption              *ScopeOption
		SimpleValueSpecification *SimpleValueSpecification
	}

	ScopeOption struct {
		Node

		GlobalOrLocal token.Token
	}

	PTFDescriptorName struct {
		Node

		Ptf token.Token
	}

	WindowName struct {
		Node

		Identifier *Identifier
	}

	RowPatternVariableName struct {
		Node

		CorrelationName *CorrelationName
	}

	MeasureName struct {
		Node

		Identifier *Identifier
	}
)
