package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func QueryBenchResults(ctx context.Context, db *sql.DB, query string, args []interface{}) ([]*BenchResult, error) {
	var (
		rs   = []*BenchResult{}
		q    = new(strings.Builder)
		tags = map[string][]string{}
	)
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// query results
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	q.WriteString("select id, tag from bench_result_tag where id in (''")
	for rows.Next() {
		var (
			r   BenchResult
			cmd []byte
		)
		err = rows.Scan(&r.ID, &r.Name, &cmd, &r.Started, &r.Finished, &r.Exit, &r.Error, &r.Stdout, &r.Stderr)
		if err != nil {
			rows.Close()
			return nil, err
		}
		json.Unmarshal(cmd, &r.Cmd)
		rs = append(rs, &r)
		q.WriteString(", '" + r.ID + "'")
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}
	q.WriteString(")")
	if len(rs) == 0 {
		return rs, nil
	}
	// query tags
	rows, err = tx.Query(q.String())
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id  string
			tag string
		)
		err = rows.Scan(&id, &tag)
		if err != nil {
			rows.Close()
			return nil, err
		}
		tags[id] = append(tags[id], tag)
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for _, r := range rs {
		r.Tags = tags[r.ID]
	}
	return rs, nil
}

func BuildQuerySQL(query string, includeOutput bool, limit int, offset int) (string, []interface{}, error) {
	buf := new(strings.Builder)
	buf.WriteString("select `id`, `name`, `cmd`, `started`, `finished`, `exit`, `error`")
	if includeOutput {
		buf.WriteString(", `stdout`, `stderr`")
	} else {
		buf.WriteString(", '' `stdout`, '' `stderr`")
	}
	buf.WriteString(" from bench_result")
	var (
		args []interface{}
		err  error
	)
	if len(query) > 0 {
		buf.WriteString(" where ")
		args, err = buildWhereClause(buf, query)
		if err != nil {
			return "", nil, err
		}
	}
	buf.WriteString(" order by started desc")
	if limit > 0 {
		buf.WriteString(" limit " + strconv.Itoa(limit))
		if offset > 0 {
			buf.WriteString(" offset " + strconv.Itoa(offset))
		}
	}
	return buf.String(), args, nil
}

func buildWhereClause(out *strings.Builder, query string) ([]interface{}, error) {
	collector := new(ParseErrorCollector)

	lexer := NewQueryLexer(antlr.NewInputStream(query))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	parser := NewQueryParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)

	expr := parser.Query()
	if len(collector.errors) > 0 {
		return nil, collector.errors[0]
	}

	visitor := &ResultQueryBuilder{query: out}
	res := visitor.Visit(expr)
	if res != nil {
		err, ok := res.(error)
		if ok {
			return nil, err
		}
		return nil, fmt.Errorf("unknown build result: %v", res)
	}
	return visitor.args, nil
}

type ParseErrorCollector struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (c *ParseErrorCollector) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errors = append(c.errors, errors.New("line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg))
}

type ResultQueryBuilder struct {
	BaseQueryVisitor

	query *strings.Builder
	args  []interface{}
}

func (v *ResultQueryBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ResultQueryBuilder) VisitQuery(ctx *QueryContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *ResultQueryBuilder) VisitPredicateExpr(ctx *PredicateExprContext) interface{} {
	return ctx.Predicate().Accept(v)
}

func (v *ResultQueryBuilder) VisitNestedExpr(ctx *NestedExprContext) interface{} {
	v.query.WriteString("(")
	if err := ctx.Expr().Accept(v); err != nil {
		return err
	}
	v.query.WriteString(")")
	return nil
}

func (v *ResultQueryBuilder) VisitNotExpr(ctx *NotExprContext) interface{} {
	v.query.WriteString("NOT ")
	return ctx.Expr().Accept(v)
}

func (v *ResultQueryBuilder) VisitAndExpr(ctx *AndExprContext) interface{} {
	if err := ctx.Expr(0).Accept(v); err != nil {
		return err
	}
	v.query.WriteString(" AND ")
	return ctx.Expr(1).Accept(v)
}

func (v *ResultQueryBuilder) VisitOrExpr(ctx *OrExprContext) interface{} {
	if err := ctx.Expr(0).Accept(v); err != nil {
		return err
	}
	v.query.WriteString(" OR ")
	return ctx.Expr(1).Accept(v)
}

func (v *ResultQueryBuilder) VisitPrimaryPred(ctx *PrimaryPredContext) interface{} {
	return ctx.Primary().Accept(v)
}

func (v *ResultQueryBuilder) VisitIsPred(ctx *IsPredContext) interface{} {
	if err := ctx.Primary().Accept(v); err != nil {
		return err
	}
	v.query.WriteString(" IS")
	if ctx.NOT() != nil {
		v.query.WriteString(" NOT")
	}
	v.query.WriteString(" NULL")
	return nil
}

func (v *ResultQueryBuilder) VisitOpPred(ctx *OpPredContext) interface{} {
	if err := ctx.Primary(0).Accept(v); err != nil {
		return err
	}
	if ctx.NOT() != nil {
		v.query.WriteString(" NOT")
	}
	op := strings.ToUpper(ctx.op.GetText())
	if op == "MATCH" {
		op = "REGEXP"
	}
	v.query.WriteString(" " + op + " ")
	return ctx.Primary(1).Accept(v)
}

func (v *ResultQueryBuilder) VisitCmpPred(ctx *CmpPredContext) interface{} {
	if err := ctx.Primary(0).Accept(v); err != nil {
		return err
	}
	v.query.WriteString(" " + strings.ToUpper(ctx.cmp.GetText()) + " ")
	return ctx.Primary(1).Accept(v)
}

func (v *ResultQueryBuilder) VisitPrimary(ctx *PrimaryContext) interface{} {
	child, ok := ctx.GetChild(0).(antlr.ParseTree)
	if !ok {
		return fmt.Errorf("unknown primary type: %T", ctx.GetChild(0))
	}
	return child.Accept(v)
}

func (v *ResultQueryBuilder) VisitField(ctx *FieldContext) interface{} {
	name := ctx.IDENT().GetText()
	if strings.ToLower(name) == "tags" {
		v.query.WriteString("(select `tag` from `bench_result_tag` where `id` = `bench_result`.`id`)")
	} else {
		v.query.WriteString("`" + name + "`")
	}
	return nil
}

func (v *ResultQueryBuilder) VisitApply(ctx *ApplyContext) interface{} {
	name := ctx.IDENT().GetText()
	v.query.WriteString(name + "(")
	for i, arg := range v.args {
		t, ok := arg.(antlr.ParseTree)
		if !ok {
			return fmt.Errorf("unknown argument type: %T", arg)
		}
		if i > 0 {
			v.query.WriteString(", ")
		}
		t.Accept(v)
	}
	v.query.WriteString(")")
	return nil
}

func (v *ResultQueryBuilder) VisitStringValue(ctx *StringValueContext) interface{} {
	s := ctx.GetText()
	v.query.WriteString("?")
	v.args = append(v.args, s[1:len(s)-1])
	return nil
}

func (v *ResultQueryBuilder) VisitNumberValue(ctx *NumberValueContext) interface{} {
	v.query.WriteString(ctx.GetText())
	return nil
}

func (v *ResultQueryBuilder) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	v.query.WriteString(ctx.GetText())
	return nil
}

func (v *ResultQueryBuilder) VisitNullValue(ctx *NullValueContext) interface{} {
	v.query.WriteString(ctx.GetText())
	return nil
}

func (v *ResultQueryBuilder) VisitValueTuple(ctx *ValueTupleContext) interface{} {
	v.query.WriteString("(")
	for i, item := range ctx.items {
		t, ok := item.(antlr.ParseTree)
		if !ok {
			return fmt.Errorf("unknown tuple element type: %T", item)
		}
		if i > 0 {
			v.query.WriteString(", ")
		}
		t.Accept(v)
	}
	v.query.WriteString(")")
	return nil
}

func (v *ResultQueryBuilder) VisitFieldTuple(ctx *FieldTupleContext) interface{} {
	v.query.WriteString("(")
	for i, item := range ctx.items {
		t, ok := item.(antlr.ParseTree)
		if !ok {
			return fmt.Errorf("unknown tuple element type: %T", item)
		}
		if i > 0 {
			v.query.WriteString(", ")
		}
		t.Accept(v)
	}
	v.query.WriteString(")")
	return nil
}
