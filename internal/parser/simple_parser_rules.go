package parser

import (
	"github.com/tomarrell/lbadd/internal/parser/ast"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

func (p *simpleParser) parseSQLStatement(r reporter) (stmt *ast.SQLStmt) {
	stmt = &ast.SQLStmt{}

	if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordExplain {
		stmt.Explain = next
		p.consumeToken()

		if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordQuery {
			stmt.Query = next
			p.consumeToken()

			if next, ok := p.lookahead(r); ok && next.Type() == token.KeywordPlan {
				stmt.Plan = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.KeywordPlan)
				// At this point, just assume that 'QUERY' was a mistake. Don't
				// abort. It's very unlikely that 'PLAN' occurs somewhere, so
				// assume that the user meant to input 'EXPLAIN <statement>'
				// instead of 'EXPLAIN QUERY PLAN <statement>'.
			}
		}
	}

	// according to the grammar, these are the tokens that initiate a statement
	p.searchNext(r, token.StatementSeparator, token.EOF, token.KeywordAlter, token.KeywordAnalyze, token.KeywordAttach, token.KeywordBegin, token.KeywordCommit, token.KeywordCreate, token.KeywordDelete, token.KeywordDetach, token.KeywordDrop, token.KeywordEnd, token.KeywordInsert, token.KeywordPragma, token.KeywordReindex, token.KeywordRelease, token.KeywordRollback, token.KeywordSavepoint, token.KeywordSelect, token.KeywordUpdate, token.KeywordVacuum)

	next, ok := p.unsafeLowLevelLookahead()
	if !ok {
		r.incompleteStatement()
		return
	}

	// lookahead processing to check what the statement ahead is
	switch next.Type() {
	case token.KeywordAlter:
		stmt.AlterTableStmt = p.parseAlterTableStmt(r)
	case token.KeywordAnalyze:
		stmt.AnalyzeStmt = p.parseAnalyzeStmt(r)
	case token.KeywordAttach:
		stmt.AttachStmt = p.parseAttachDatabaseStmt(r)
	case token.KeywordBegin:
		stmt.BeginStmt = p.parseBeginStmt(r)
	case token.KeywordCommit:
		stmt.CommitStmt = p.parseCommitStmt(r)
	case token.KeywordDetach:
		stmt.DetachStmt = p.parseDetachDatabaseStmt(r)
	case token.KeywordEnd:
		stmt.CommitStmt = p.parseCommitStmt(r)
	case token.KeywordRollback:
		stmt.RollbackStmt = p.parseRollbackStmt(r)
	case token.KeywordVacuum:
		stmt.VacuumStmt = p.parseVacuumStmt(r)
	case token.StatementSeparator:
		r.incompleteStatement()
		p.consumeToken()
	case token.KeywordPragma:
		// we don't support pragmas, as we don't need them yet
		r.unsupportedConstruct(next)
		p.skipUntil(token.StatementSeparator, token.EOF)
	default:
		r.unsupportedConstruct(next)
		p.skipUntil(token.StatementSeparator, token.EOF)
	}

	p.searchNext(r, token.StatementSeparator, token.EOF)
	next, ok = p.unsafeLowLevelLookahead()
	if !ok {
		return
	}
	if next.Type() == token.StatementSeparator {
		// if there's a statement separator, consume this token and get the next
		// token, so that it can be checked if that next token is an EOF
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
	}
	if next.Type() == token.EOF {
		p.consumeToken()
		return
	}
	return
}

func (p *simpleParser) parseAlterTableStmt(r reporter) (stmt *ast.AlterTableStmt) {
	stmt = &ast.AlterTableStmt{}

	p.searchNext(r, token.KeywordAlter)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Alter = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordTable {
		stmt.Table = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordTable)
		// do not consume anything, report that there is no 'TABLE' keyword, and
		// proceed as if we would have received it
	}

	schemaOrTableName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaOrTableName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	p.consumeToken() // consume the literal

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "." {
		// schemaOrTableName is schema name
		stmt.SchemaName = schemaOrTableName
		stmt.Period = next
		p.consumeToken()

		tableName, ok := p.lookahead(r)
		if !ok {
			return
		}
		if tableName.Type() == token.Literal {
			stmt.TableName = tableName
			p.consumeToken()
		}
	} else {
		stmt.TableName = schemaOrTableName
	}
	switch next.Type() {
	case token.KeywordRename:
		stmt.Rename = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordTo:
			stmt.To = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}
			stmt.NewTableName = next
			p.consumeToken()
		case token.KeywordColumn:
			stmt.Column = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}

			fallthrough
		case token.Literal:
			stmt.ColumnName = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.KeywordTo {
				r.unexpectedToken(token.KeywordTo)
				p.consumeToken()
				return
			}

			stmt.To = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}

			stmt.NewColumnName = next
			p.consumeToken()
		default:
			r.unexpectedToken(token.KeywordTo, token.KeywordColumn, token.Literal)
		}
	case token.KeywordAdd:
		stmt.Add = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		switch next.Type() {
		case token.KeywordColumn:
			stmt.Column = next
			p.consumeToken()

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Literal {
				r.unexpectedToken(token.Literal)
				p.consumeToken()
				return
			}
			fallthrough
		case token.Literal:
			stmt.ColumnDef = p.parseColumnDef(r)
		default:
			r.unexpectedToken(token.KeywordColumn, token.Literal)
		}
	}

	return
}

func (p *simpleParser) parseColumnDef(r reporter) (def *ast.ColumnDef) {
	def = &ast.ColumnDef{}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
		def.ColumnName = next
		p.consumeToken()

		if next, ok = p.lookahead(r); ok && next.Type() == token.Literal {
			def.TypeName = p.parseTypeName(r)
		}

		for {
			next, ok = p.optionalLookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.KeywordConstraint ||
				next.Type() == token.KeywordPrimary ||
				next.Type() == token.KeywordNot ||
				next.Type() == token.KeywordUnique ||
				next.Type() == token.KeywordCheck ||
				next.Type() == token.KeywordDefault ||
				next.Type() == token.KeywordCollate ||
				next.Type() == token.KeywordGenerated ||
				next.Type() == token.KeywordReferences {
				def.ColumnConstraint = append(def.ColumnConstraint, p.parseColumnConstraint(r))
			} else {
				break
			}
		}
	}
	return
}

func (p *simpleParser) parseTypeName(r reporter) (name *ast.TypeName) {
	name = &ast.TypeName{}

	// one or more name
	if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
		name.Name = append(name.Name, next)
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
	}
	for {
		if next, ok := p.lookahead(r); ok && next.Type() == token.Literal {
			name.Name = append(name.Name, next)
			p.consumeToken()
		} else {
			break
		}
	}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Delimiter {
		if next.Value() == "(" {
			name.LeftParen = next
			p.consumeToken()

			name.SignedNumber1 = p.parseSignedNumber(r)
		} else {
			r.unexpectedToken(token.Delimiter)
		}
	} else {
		return
	}

	if next, ok := p.lookahead(r); ok && next.Type() == token.Delimiter {
		switch next.Value() {
		case ",":
			name.Comma = next
			p.consumeToken()

			name.SignedNumber2 = p.parseSignedNumber(r)
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() != token.Delimiter {
				r.unexpectedToken(token.Delimiter)
			}
			fallthrough
		case ")":
			name.RightParen = next
			p.consumeToken()
		}
	} else {
		return
	}

	return
}

func (p *simpleParser) parseSignedNumber(r reporter) (num *ast.SignedNumber) {
	num = &ast.SignedNumber{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.UnaryOperator:
		num.Sign = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() != token.Literal {
			r.unexpectedToken(token.Literal)
			return
		}
		fallthrough
	case token.Literal:
		num.NumericLiteral = next
		p.consumeToken()
	default:
		r.unexpectedToken(token.UnaryOperator, token.Literal)
		return
	}
	return
}

func (p *simpleParser) parseColumnConstraint(r reporter) (constr *ast.ColumnConstraint) {
	constr = &ast.ColumnConstraint{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordConstraint {
		constr.Constraint = next
		p.consumeToken()

		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			constr.Name = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
			// report that the token was unexpected, but continue as if the
			// missing literal token was present
		}
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordPrimary:
		// PRIMARY
		constr.Primary = next
		p.consumeToken()

		// KEY
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordKey {
			constr.Key = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordKey)
		}

		// ASC, DESC
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAsc {
			constr.Asc = next
			p.consumeToken()
		} else if next.Type() == token.KeywordDesc {
			constr.Desc = next
			p.consumeToken()
		}

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

		// AUTOINCREMENT
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordAutoincrement {
			constr.Autoincrement = next
			p.consumeToken()
		}

	case token.KeywordNot:
		// NOT
		constr.Not = next
		p.consumeToken()

		// NULL
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordNull {
			constr.Null = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordNull)
		}

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

	case token.KeywordUnique:
		// UNIQUE
		constr.Unique = next
		p.consumeToken()

		// conflict clause
		next, ok = p.optionalLookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordOn {
			constr.ConflictClause = p.parseConflictClause(r)
		}

	case token.KeywordCheck:
		// CHECK
		constr.Check = next
		p.consumeToken()

		// left paren
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == "(" {
			constr.LeftParen = next
			p.consumeToken()
		} else {
			r.unexpectedSingleRuneToken(token.Delimiter, '(')
			// assume that the opening paren has been omitted, report the
			// error but proceed as if it was found
		}

		// expr
		constr.Expr = p.parseExpression(r)

		// right paren
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Delimiter && next.Value() == ")" {
			constr.RightParen = next
			p.consumeToken()
		} else {
			r.unexpectedSingleRuneToken(token.Delimiter, ')')
			// assume that the opening paren has been omitted, report the
			// error but proceed as if it was found
		}

	case token.KeywordDefault:
		constr.Default = next
		p.consumeToken()

	case token.KeywordCollate:
		constr.Collate = next
		p.consumeToken()

	case token.KeywordGenerated:
		constr.Generated = next
		p.consumeToken()

	case token.KeywordReferences:
		constr.ForeignKeyClause = p.parseForeignKeyClause(r)
	default:
		r.unexpectedToken(token.KeywordPrimary, token.KeywordNot, token.KeywordUnique, token.KeywordCheck, token.KeywordDefault, token.KeywordCollate, token.KeywordGenerated, token.KeywordReferences)
	}

	return
}

// parseForeignKeyClause is not implemented yet and will always result in an
// unsupported construct error.
func (p *simpleParser) parseForeignKeyClause(r reporter) (clause *ast.ForeignKeyClause) {
	clause = &ast.ForeignKeyClause{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	r.unsupportedConstruct(next)
	p.searchNext(r, token.StatementSeparator, token.EOF)
	return
}

func (p *simpleParser) parseConflictClause(r reporter) (clause *ast.ConflictClause) {
	clause = &ast.ConflictClause{}

	next, ok := p.optionalLookahead(r)
	if !ok {
		return
	}

	// ON
	if next.Type() == token.KeywordOn {
		clause.On = next
		p.consumeToken()
	} else {
		// if there's no 'ON' token, the empty production is assumed, which is
		// why no error is reported here
		return
	}

	// CONFLICT
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordConflict {
		clause.Conflict = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordConflict)
		return
	}

	// ROLLBACK, ABORT, FAIL, IGNORE, REPLACE
	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	switch next.Type() {
	case token.KeywordRollback:
		clause.Rollback = next
		p.consumeToken()
	case token.KeywordAbort:
		clause.Abort = next
		p.consumeToken()
	case token.KeywordFail:
		clause.Fail = next
		p.consumeToken()
	case token.KeywordIgnore:
		clause.Ignore = next
		p.consumeToken()
	case token.KeywordReplace:
		clause.Replace = next
		p.consumeToken()
	default:
		r.unexpectedToken(token.KeywordRollback, token.KeywordAbort, token.KeywordFail, token.KeywordIgnore, token.KeywordReplace)
	}
	return
}

// parseExpression is not implemented yet and will always result in an
// unsupported construct error.
func (p *simpleParser) parseExpression(r reporter) (expr *ast.Expr) {
	expr = &ast.Expr{}

	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		expr.LiteralValue = next
		p.consumeToken()
	} else {
		r.unsupportedConstruct(next)
		p.searchNext(r, token.StatementSeparator, token.EOF)
	}
	return
}

// parseAttachDatabaseStmt parses a single ATTACH statement as defined in the spec:
// https://sqlite.org/lang_attach.html
func (p *simpleParser) parseAttachDatabaseStmt(r reporter) (stmt *ast.AttachStmt) {
	stmt = &ast.AttachStmt{}
	p.searchNext(r, token.KeywordAttach)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Attach = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordDatabase {
		stmt.Database = next
		p.consumeToken()
	}
	stmt.Expr = p.parseExpression(r)

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordAs {
		stmt.As = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordAs)
		return
	}

	schemaName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	stmt.SchemaName = schemaName
	p.consumeToken()
	return
}

// parseDetachDatabaseStmt parses a single DETACH statement as defined in spec:
// https://sqlite.org/lang_detach.html
func (p *simpleParser) parseDetachDatabaseStmt(r reporter) (stmt *ast.DetachStmt) {
	stmt = &ast.DetachStmt{}
	p.searchNext(r, token.KeywordDetach)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Detach = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordDatabase {
		stmt.Database = next
		p.consumeToken()
	}

	schemaName, ok := p.lookahead(r)
	if !ok {
		return
	}
	if schemaName.Type() != token.Literal {
		r.unexpectedToken(token.Literal)
		return
	}
	stmt.SchemaName = schemaName
	p.consumeToken()
	return
}

// parseVacuumStmt parses a single VACUUM statement as defined in the spec:
// https://sqlite.org/lang_vacuum.html
func (p *simpleParser) parseVacuumStmt(r reporter) (stmt *ast.VacuumStmt) {
	stmt = &ast.VacuumStmt{}
	p.searchNext(r, token.KeywordVacuum)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Vacuum = next
	p.consumeToken()

	// optionalLookahead is used because, the lookahead function
	// always looks for the next "real" token and not EOF.
	// Since Just "VACUUM" is a valid statement, we have to accept
	// the fact that there can be no tokens after the first keyword.
	// Same logic is applied for the next INTO keyword check too.
	next, ok = p.optionalLookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordInto {
		stmt.Into = next
		p.consumeToken()
		fileName, ok := p.lookahead(r)
		if !ok {
			return
		}
		if fileName.Type() != token.Literal {
			r.unexpectedToken(token.Literal)
			return
		}
		stmt.Filename = fileName
		p.consumeToken()
	}
	return
}

// parseAnalyzeStmt parses a single ANALYZE statement as defined in the spec:
// https://sqlite.org/lang_analyze.html
func (p *simpleParser) parseAnalyzeStmt(r reporter) (stmt *ast.AnalyzeStmt) {
	stmt = &ast.AnalyzeStmt{}
	p.searchNext(r, token.KeywordAnalyze)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Analyze = next
	p.consumeToken()

	// optionalLookahead is used, because ANALYZE alone is a valid statement
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		stmt.TableOrIndexName = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Literal)
		return
	}

	period, ok := p.optionalLookahead(r)
	if !ok || period.Type() == token.EOF {
		return
	}
	stmt.Period = period
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	stmt.TableOrIndexName = next
	p.consumeToken()
	return
}

// parseBeginStmt parses a single BEGIN statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseBeginStmt(r reporter) (stmt *ast.BeginStmt) {
	stmt = &ast.BeginStmt{}
	p.searchNext(r, token.KeywordBegin)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Begin = next
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordDeferred {
		stmt.Deferred = next
		p.consumeToken()
	} else if next.Type() == token.KeywordImmediate {
		stmt.Immediate = next
		p.consumeToken()
	} else if next.Type() == token.KeywordExclusive {
		stmt.Exclusive = next
		p.consumeToken()
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}
	return
}

// parseCommitStmt parses a single COMMIT statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseCommitStmt(r reporter) (stmt *ast.CommitStmt) {
	stmt = &ast.CommitStmt{}
	p.searchNext(r, token.KeywordCommit, token.KeywordEnd)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordCommit {
		stmt.Commit = next
	} else if next.Type() == token.KeywordEnd {
		stmt.End = next
	}
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}
	return
}

// parseRollbackStmt parses a single ROLLBACK statement as defined in the spec:
// https://sqlite.org/lang_transaction.html
func (p *simpleParser) parseRollbackStmt(r reporter) (stmt *ast.RollbackStmt) {
	stmt = &ast.RollbackStmt{}
	p.searchNext(r, token.KeywordRollback)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Rollback = next
	p.consumeToken()

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordTransaction {
		stmt.Transaction = next
		p.consumeToken()
	}

	// If the keyword TRANSACTION exists in the statement, we need to
	// check whether TO also exists. Out of TRANSACTION and TO, each not
	// existing and existing, we have the following logic.
	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordTo {
		stmt.To = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordSavepoint {
		stmt.Savepoint = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.Literal {
		stmt.SavepointName = next
	}
	p.consumeToken()
	return
}

func (p *simpleParser) parseDeleteStmt(r reporter) (stmt *ast.DeleteStmt) {
	stmt = &ast.DeleteStmt{}
	stmt.WithClause = p.parseWithClause(r)
	p.searchNext(r, token.KeywordDelete)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	stmt.Delete = next
	p.consumeToken()

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordFrom {
		stmt.From = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordFrom)
	}

	stmt.QualifiedTableName = p.parseQualifiedTableName(r)

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordWhere {
		stmt.Where = next
		p.consumeToken()
		stmt.Expr = p.parseExpression(r)
	}

	return
}

func (p *simpleParser) parseWithClause(r reporter) (withClause *ast.WithClause) {
	withClause = &ast.WithClause{}
	p.searchNext(r, token.KeywordWith)
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	withClause.With = next

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() == token.KeywordRecursive {
		withClause.Recursive = next
		p.consumeToken()
	}

	for {
		withClause.RecursiveCte = append(withClause.RecursiveCte, p.parseRecursiveCte(r))
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Value() == "," {
			p.consumeToken()
		} else {
			break
		}
	}
	return
}

func (p *simpleParser) parseRecursiveCte(r reporter) (recursiveCte *ast.RecursiveCte) {
	recursiveCte.CteTableName = p.parseCteTableName(r)
	next, ok := p.lookahead(r)
	if ok {
		return
	}
	if next.Type() == token.KeywordAs {
		recursiveCte.As = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.KeywordAs)
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Value() == "(" {
		recursiveCte.LeftParen = next
		p.consumeToken()
	} else {
		r.unexpectedToken(token.Delimiter)
	}
	recursiveCte.SelectStmt = p.parseSelectStmt(r)
	return
}

func (p *simpleParser) parseCteTableName(r reporter) (cteTableName *ast.CteTableName) {
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	cteTableName.TableName = next

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Value() == "(" {
		cteTableName.LeftParen = next
		p.consumeToken()
		for {
			columnName, ok := p.lookahead(r)
			if !ok {
				return
			}
			if columnName.Type() == token.Literal {
				cteTableName.ColumnName = append(cteTableName.ColumnName, columnName)
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == "," {
				p.consumeToken()
			}

			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Value() == ")" {
				cteTableName.RightParen = next
				p.consumeToken()
				break
			}
		}
	}
	return
}

func (p *simpleParser) parseSelectStmt(r reporter) (stmt *ast.SelectStmt) {
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	r.unsupportedConstruct(next)
	p.searchNext(r, token.StatementSeparator, token.EOF)
	return
}

func (p *simpleParser) parseQualifiedTableName(r reporter) (stmt *ast.QualifiedTableName) {
	next, ok := p.lookahead(r)
	if !ok {
		return
	}
	// We expect that the first literal can be either the schema name
	// or the table name. When we confirm the existance of a period, we
	// re-assign the table name and when we confirm that there is no
	// period, we reset the schema name value to nil.
	if next.Type() == token.Literal {
		stmt.SchemaName = next
		stmt.TableName = next
		p.consumeToken()
	}

	next, ok = p.lookahead(r)
	if !ok {
		return
	}
	if next.Type() != token.Literal {
		stmt.SchemaName = nil
	}
	if next.Value() == "." {
		stmt.Period = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.TableName = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordAs {
		stmt.As = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.Literal {
			stmt.Alias = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.Literal)
		}
	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordIndexed {
		stmt.Indexed = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordBy {
			stmt.By = next
			p.consumeToken()
			next, ok = p.lookahead(r)
			if !ok {
				return
			}
			if next.Type() == token.Literal {
				stmt.IndexName = next
				p.consumeToken()
			} else {
				r.unexpectedToken(token.Literal)
			}
		} else {
			r.unexpectedToken(token.KeywordBy)
		}

	}

	next, ok = p.optionalLookahead(r)
	if !ok || next.Type() == token.EOF {
		return
	}
	if next.Type() == token.KeywordNot {
		stmt.Not = next
		p.consumeToken()
		next, ok = p.lookahead(r)
		if !ok {
			return
		}
		if next.Type() == token.KeywordIndexed {
			stmt.Indexed = next
			p.consumeToken()
		} else {
			r.unexpectedToken(token.KeywordIndexed)
		}
	}
	return
}
