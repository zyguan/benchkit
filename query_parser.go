// Code generated from Query.g4 by ANTLR 4.10.1. DO NOT EDIT.

package main // Query
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type QueryParser struct {
	*antlr.BaseParser
}

var queryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func queryParserInit() {
	staticData := &queryParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "','", "'<'", "'<='", "'>'", "'>='", "'='", "'!='",
		"'not'", "'and'", "'or'", "'is'", "'in'", "'like'", "'match'", "", "",
		"", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "LT", "LE", "GT", "GE", "EQ", "NE", "NOT", "AND", "OR",
		"IS", "IN", "LIKE", "MATCH", "STRING", "NUMBER", "BOOLEAN", "NULL",
		"IDENT", "WS",
	}
	staticData.ruleNames = []string{
		"query", "expr", "predicate", "primary", "field", "apply", "value",
		"tuple",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 117, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 45, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 66, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 78, 8, 5, 10, 5, 12, 5, 81, 9, 5, 1, 5, 1, 5, 3, 5, 85, 8, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 91, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 97,
		8, 7, 10, 7, 12, 7, 100, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7,
		108, 8, 7, 10, 7, 12, 7, 111, 9, 7, 1, 7, 1, 7, 3, 7, 115, 8, 7, 1, 7,
		0, 1, 2, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 2, 1, 0, 14, 16, 1, 0, 4, 9,
		128, 0, 16, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 65, 1,
		0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 84, 1, 0, 0, 0, 12, 90, 1, 0, 0, 0, 14,
		114, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0,
		0, 19, 20, 6, 1, -1, 0, 20, 28, 3, 4, 2, 0, 21, 22, 5, 1, 0, 0, 22, 23,
		3, 2, 1, 0, 23, 24, 5, 2, 0, 0, 24, 28, 1, 0, 0, 0, 25, 26, 5, 10, 0, 0,
		26, 28, 3, 2, 1, 3, 27, 19, 1, 0, 0, 0, 27, 21, 1, 0, 0, 0, 27, 25, 1,
		0, 0, 0, 28, 37, 1, 0, 0, 0, 29, 30, 10, 2, 0, 0, 30, 31, 5, 11, 0, 0,
		31, 36, 3, 2, 1, 3, 32, 33, 10, 1, 0, 0, 33, 34, 5, 12, 0, 0, 34, 36, 3,
		2, 1, 2, 35, 29, 1, 0, 0, 0, 35, 32, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37,
		35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 37, 1, 0, 0,
		0, 40, 60, 3, 6, 3, 0, 41, 42, 3, 6, 3, 0, 42, 44, 5, 13, 0, 0, 43, 45,
		5, 10, 0, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0,
		46, 47, 5, 20, 0, 0, 47, 60, 1, 0, 0, 0, 48, 50, 3, 6, 3, 0, 49, 51, 5,
		10, 0, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		53, 7, 0, 0, 0, 53, 54, 3, 6, 3, 0, 54, 60, 1, 0, 0, 0, 55, 56, 3, 6, 3,
		0, 56, 57, 7, 1, 0, 0, 57, 58, 3, 6, 3, 0, 58, 60, 1, 0, 0, 0, 59, 40,
		1, 0, 0, 0, 59, 41, 1, 0, 0, 0, 59, 48, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0,
		60, 5, 1, 0, 0, 0, 61, 66, 3, 8, 4, 0, 62, 66, 3, 10, 5, 0, 63, 66, 3,
		12, 6, 0, 64, 66, 3, 14, 7, 0, 65, 61, 1, 0, 0, 0, 65, 62, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 7, 1, 0, 0, 0, 67, 68, 5, 21,
		0, 0, 68, 9, 1, 0, 0, 0, 69, 70, 5, 21, 0, 0, 70, 71, 5, 1, 0, 0, 71, 85,
		5, 2, 0, 0, 72, 73, 5, 21, 0, 0, 73, 74, 5, 1, 0, 0, 74, 79, 3, 6, 3, 0,
		75, 76, 5, 3, 0, 0, 76, 78, 3, 6, 3, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 83, 5, 2, 0, 0, 83, 85, 1, 0, 0, 0, 84, 69, 1, 0, 0,
		0, 84, 72, 1, 0, 0, 0, 85, 11, 1, 0, 0, 0, 86, 91, 5, 17, 0, 0, 87, 91,
		5, 18, 0, 0, 88, 91, 5, 19, 0, 0, 89, 91, 5, 20, 0, 0, 90, 86, 1, 0, 0,
		0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 13,
		1, 0, 0, 0, 92, 93, 5, 1, 0, 0, 93, 98, 3, 12, 6, 0, 94, 95, 5, 3, 0, 0,
		95, 97, 3, 12, 6, 0, 96, 94, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0,
		101, 102, 5, 2, 0, 0, 102, 115, 1, 0, 0, 0, 103, 104, 5, 1, 0, 0, 104,
		109, 3, 8, 4, 0, 105, 106, 5, 3, 0, 0, 106, 108, 3, 8, 4, 0, 107, 105,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 5, 2, 0, 0,
		113, 115, 1, 0, 0, 0, 114, 92, 1, 0, 0, 0, 114, 103, 1, 0, 0, 0, 115, 15,
		1, 0, 0, 0, 13, 27, 35, 37, 44, 50, 59, 65, 79, 84, 90, 98, 109, 114,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// QueryParserInit initializes any static state used to implement QueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryParserInit() {
	staticData := &queryParserStaticData
	staticData.once.Do(queryParserInit)
}

// NewQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryParser(input antlr.TokenStream) *QueryParser {
	QueryParserInit()
	this := new(QueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &queryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Query.g4"

	return this
}

// QueryParser tokens.
const (
	QueryParserEOF     = antlr.TokenEOF
	QueryParserT__0    = 1
	QueryParserT__1    = 2
	QueryParserT__2    = 3
	QueryParserLT      = 4
	QueryParserLE      = 5
	QueryParserGT      = 6
	QueryParserGE      = 7
	QueryParserEQ      = 8
	QueryParserNE      = 9
	QueryParserNOT     = 10
	QueryParserAND     = 11
	QueryParserOR      = 12
	QueryParserIS      = 13
	QueryParserIN      = 14
	QueryParserLIKE    = 15
	QueryParserMATCH   = 16
	QueryParserSTRING  = 17
	QueryParserNUMBER  = 18
	QueryParserBOOLEAN = 19
	QueryParserNULL    = 20
	QueryParserIDENT   = 21
	QueryParserWS      = 22
)

// QueryParser rules.
const (
	QueryParserRULE_query     = 0
	QueryParserRULE_expr      = 1
	QueryParserRULE_predicate = 2
	QueryParserRULE_primary   = 3
	QueryParserRULE_field     = 4
	QueryParserRULE_apply     = 5
	QueryParserRULE_value     = 6
	QueryParserRULE_tuple     = 7
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryParserEOF, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Query() (localctx IQueryContext) {
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expr(0)
	}
	{
		p.SetState(17)
		p.Match(QueryParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(QueryParserAND, 0)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PredicateExprContext struct {
	*ExprContext
}

func NewPredicateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredicateExprContext {
	var p = new(PredicateExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PredicateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateExprContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitPredicateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExprContext struct {
	*ExprContext
}

func NewNestedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExprContext {
	var p = new(NestedExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NestedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NestedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNestedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(QueryParserOR, 0)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *QueryParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, QueryParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPredicateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(20)
			p.Predicate()
		}

	case 2:
		localctx = NewNestedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(22)
			p.expr(0)
		}
		{
			p.SetState(23)
			p.Match(QueryParserT__1)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(QueryParserNOT)
		}
		{
			p.SetState(26)
			p.expr(3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_expr)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(30)
					p.Match(QueryParserAND)
				}
				{
					p.SetState(31)
					p.expr(3)
				}

			case 2:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_expr)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(33)
					p.Match(QueryParserOR)
				}
				{
					p.SetState(34)
					p.expr(2)
				}

			}

		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) CopyFrom(ctx *PredicateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryPredContext struct {
	*PredicateContext
}

func NewPrimaryPredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryPredContext {
	var p = new(PrimaryPredContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *PrimaryPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryPredContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryPredContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitPrimaryPred(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmpPredContext struct {
	*PredicateContext
	cmp antlr.Token
}

func NewCmpPredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpPredContext {
	var p = new(CmpPredContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *CmpPredContext) GetCmp() antlr.Token { return s.cmp }

func (s *CmpPredContext) SetCmp(v antlr.Token) { s.cmp = v }

func (s *CmpPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpPredContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *CmpPredContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *CmpPredContext) LT() antlr.TerminalNode {
	return s.GetToken(QueryParserLT, 0)
}

func (s *CmpPredContext) LE() antlr.TerminalNode {
	return s.GetToken(QueryParserLE, 0)
}

func (s *CmpPredContext) GT() antlr.TerminalNode {
	return s.GetToken(QueryParserGT, 0)
}

func (s *CmpPredContext) GE() antlr.TerminalNode {
	return s.GetToken(QueryParserGE, 0)
}

func (s *CmpPredContext) EQ() antlr.TerminalNode {
	return s.GetToken(QueryParserEQ, 0)
}

func (s *CmpPredContext) NE() antlr.TerminalNode {
	return s.GetToken(QueryParserNE, 0)
}

func (s *CmpPredContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitCmpPred(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsPredContext struct {
	*PredicateContext
}

func NewIsPredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsPredContext {
	var p = new(IsPredContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *IsPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsPredContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *IsPredContext) IS() antlr.TerminalNode {
	return s.GetToken(QueryParserIS, 0)
}

func (s *IsPredContext) NULL() antlr.TerminalNode {
	return s.GetToken(QueryParserNULL, 0)
}

func (s *IsPredContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *IsPredContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitIsPred(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpPredContext struct {
	*PredicateContext
	op antlr.Token
}

func NewOpPredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpPredContext {
	var p = new(OpPredContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *OpPredContext) GetOp() antlr.Token { return s.op }

func (s *OpPredContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpPredContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *OpPredContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *OpPredContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryParserIN, 0)
}

func (s *OpPredContext) LIKE() antlr.TerminalNode {
	return s.GetToken(QueryParserLIKE, 0)
}

func (s *OpPredContext) MATCH() antlr.TerminalNode {
	return s.GetToken(QueryParserMATCH, 0)
}

func (s *OpPredContext) NOT() antlr.TerminalNode {
	return s.GetToken(QueryParserNOT, 0)
}

func (s *OpPredContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitOpPred(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Predicate() (localctx IPredicateContext) {
	this := p
	_ = this

	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_predicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrimaryPredContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Primary()
		}

	case 2:
		localctx = NewIsPredContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Primary()
		}
		{
			p.SetState(42)
			p.Match(QueryParserIS)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserNOT {
			{
				p.SetState(43)
				p.Match(QueryParserNOT)
			}

		}
		{
			p.SetState(46)
			p.Match(QueryParserNULL)
		}

	case 3:
		localctx = NewOpPredContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Primary()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserNOT {
			{
				p.SetState(49)
				p.Match(QueryParserNOT)
			}

		}
		{
			p.SetState(52)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OpPredContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserIN)|(1<<QueryParserLIKE)|(1<<QueryParserMATCH))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OpPredContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(53)
			p.Primary()
		}

	case 4:
		localctx = NewCmpPredContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Primary()
		}
		{
			p.SetState(56)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CmpPredContext).cmp = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserLT)|(1<<QueryParserLE)|(1<<QueryParserGT)|(1<<QueryParserGE)|(1<<QueryParserEQ)|(1<<QueryParserNE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CmpPredContext).cmp = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(57)
			p.Primary()
		}

	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *PrimaryContext) Apply() IApplyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApplyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApplyContext)
}

func (s *PrimaryContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PrimaryContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Apply()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Tuple()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(QueryParserIDENT, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(QueryParserIDENT)
	}

	return localctx
}

// IApplyContext is an interface to support dynamic dispatch.
type IApplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primary returns the _primary rule contexts.
	Get_primary() IPrimaryContext

	// Set_primary sets the _primary rule contexts.
	Set_primary(IPrimaryContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IPrimaryContext

	// SetArgs sets the args rule context list.
	SetArgs([]IPrimaryContext)

	// IsApplyContext differentiates from other interfaces.
	IsApplyContext()
}

type ApplyContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_primary IPrimaryContext
	args     []IPrimaryContext
}

func NewEmptyApplyContext() *ApplyContext {
	var p = new(ApplyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_apply
	return p
}

func (*ApplyContext) IsApplyContext() {}

func NewApplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplyContext {
	var p = new(ApplyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_apply

	return p
}

func (s *ApplyContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplyContext) Get_primary() IPrimaryContext { return s._primary }

func (s *ApplyContext) Set_primary(v IPrimaryContext) { s._primary = v }

func (s *ApplyContext) GetArgs() []IPrimaryContext { return s.args }

func (s *ApplyContext) SetArgs(v []IPrimaryContext) { s.args = v }

func (s *ApplyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(QueryParserIDENT, 0)
}

func (s *ApplyContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *ApplyContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitApply(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Apply() (localctx IApplyContext) {
	this := p
	_ = this

	localctx = NewApplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryParserRULE_apply)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(QueryParserIDENT)
		}
		{
			p.SetState(70)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(71)
			p.Match(QueryParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(QueryParserIDENT)
		}
		{
			p.SetState(73)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(74)

			var _x = p.Primary()

			localctx.(*ApplyContext)._primary = _x
		}
		localctx.(*ApplyContext).args = append(localctx.(*ApplyContext).args, localctx.(*ApplyContext)._primary)
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QueryParserT__2 {
			{
				p.SetState(75)
				p.Match(QueryParserT__2)
			}
			{
				p.SetState(76)

				var _x = p.Primary()

				localctx.(*ApplyContext)._primary = _x
			}
			localctx.(*ApplyContext).args = append(localctx.(*ApplyContext).args, localctx.(*ApplyContext)._primary)

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(82)
			p.Match(QueryParserT__1)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullValueContext struct {
	*ValueContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(QueryParserNULL, 0)
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberValueContext struct {
	*ValueContext
}

func NewNumberValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberValueContext {
	var p = new(NumberValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER, 0)
}

func (s *NumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanValueContext struct {
	*ValueContext
}

func NewBooleanValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(QueryParserBOOLEAN, 0)
}

func (s *BooleanValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitBooleanValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(QueryParserSTRING)
		}

	case QueryParserNUMBER:
		localctx = NewNumberValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(QueryParserNUMBER)
		}

	case QueryParserBOOLEAN:
		localctx = NewBooleanValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(QueryParserBOOLEAN)
		}

	case QueryParserNULL:
		localctx = NewNullValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(QueryParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_tuple
	return p
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) CopyFrom(ctx *TupleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueTupleContext struct {
	*TupleContext
	_value IValueContext
	items  []IValueContext
}

func NewValueTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueTupleContext {
	var p = new(ValueTupleContext)

	p.TupleContext = NewEmptyTupleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleContext))

	return p
}

func (s *ValueTupleContext) Get_value() IValueContext { return s._value }

func (s *ValueTupleContext) Set_value(v IValueContext) { s._value = v }

func (s *ValueTupleContext) GetItems() []IValueContext { return s.items }

func (s *ValueTupleContext) SetItems(v []IValueContext) { s.items = v }

func (s *ValueTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTupleContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValueTupleContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitValueTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldTupleContext struct {
	*TupleContext
	_field IFieldContext
	items  []IFieldContext
}

func NewFieldTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldTupleContext {
	var p = new(FieldTupleContext)

	p.TupleContext = NewEmptyTupleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleContext))

	return p
}

func (s *FieldTupleContext) Get_field() IFieldContext { return s._field }

func (s *FieldTupleContext) Set_field(v IFieldContext) { s._field = v }

func (s *FieldTupleContext) GetItems() []IFieldContext { return s.items }

func (s *FieldTupleContext) SetItems(v []IFieldContext) { s.items = v }

func (s *FieldTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTupleContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldTupleContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitFieldTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Tuple() (localctx ITupleContext) {
	this := p
	_ = this

	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryParserRULE_tuple)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(93)

			var _x = p.Value()

			localctx.(*ValueTupleContext)._value = _x
		}
		localctx.(*ValueTupleContext).items = append(localctx.(*ValueTupleContext).items, localctx.(*ValueTupleContext)._value)
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QueryParserT__2 {
			{
				p.SetState(94)
				p.Match(QueryParserT__2)
			}
			{
				p.SetState(95)

				var _x = p.Value()

				localctx.(*ValueTupleContext)._value = _x
			}
			localctx.(*ValueTupleContext).items = append(localctx.(*ValueTupleContext).items, localctx.(*ValueTupleContext)._value)

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(QueryParserT__1)
		}

	case 2:
		localctx = NewFieldTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(QueryParserT__0)
		}
		{
			p.SetState(104)

			var _x = p.Field()

			localctx.(*FieldTupleContext)._field = _x
		}
		localctx.(*FieldTupleContext).items = append(localctx.(*FieldTupleContext).items, localctx.(*FieldTupleContext)._field)
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QueryParserT__2 {
			{
				p.SetState(105)
				p.Match(QueryParserT__2)
			}
			{
				p.SetState(106)

				var _x = p.Field()

				localctx.(*FieldTupleContext)._field = _x
			}
			localctx.(*FieldTupleContext).items = append(localctx.(*FieldTupleContext).items, localctx.(*FieldTupleContext)._field)

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(QueryParserT__1)
		}

	}

	return localctx
}

func (p *QueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
