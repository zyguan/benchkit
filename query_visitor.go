// Code generated from Query.g4 by ANTLR 4.10.1. DO NOT EDIT.

package main // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by QueryParser#PredicateExpr.
	VisitPredicateExpr(ctx *PredicateExprContext) interface{}

	// Visit a parse tree produced by QueryParser#NestedExpr.
	VisitNestedExpr(ctx *NestedExprContext) interface{}

	// Visit a parse tree produced by QueryParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by QueryParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by QueryParser#PrimaryPred.
	VisitPrimaryPred(ctx *PrimaryPredContext) interface{}

	// Visit a parse tree produced by QueryParser#IsPred.
	VisitIsPred(ctx *IsPredContext) interface{}

	// Visit a parse tree produced by QueryParser#OpPred.
	VisitOpPred(ctx *OpPredContext) interface{}

	// Visit a parse tree produced by QueryParser#CmpPred.
	VisitCmpPred(ctx *CmpPredContext) interface{}

	// Visit a parse tree produced by QueryParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by QueryParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by QueryParser#apply.
	VisitApply(ctx *ApplyContext) interface{}

	// Visit a parse tree produced by QueryParser#StringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by QueryParser#NumberValue.
	VisitNumberValue(ctx *NumberValueContext) interface{}

	// Visit a parse tree produced by QueryParser#BooleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by QueryParser#NullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by QueryParser#ValueTuple.
	VisitValueTuple(ctx *ValueTupleContext) interface{}

	// Visit a parse tree produced by QueryParser#FieldTuple.
	VisitFieldTuple(ctx *FieldTupleContext) interface{}
}
