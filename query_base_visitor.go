// Code generated from Query.g4 by ANTLR 4.10.1. DO NOT EDIT.

package main // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitPredicateExpr(ctx *PredicateExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNestedExpr(ctx *NestedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitPrimaryPred(ctx *PrimaryPredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitIsPred(ctx *IsPredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitOpPred(ctx *OpPredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitCmpPred(ctx *CmpPredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitApply(ctx *ApplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNumberValue(ctx *NumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitValueTuple(ctx *ValueTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitFieldTuple(ctx *FieldTupleContext) interface{} {
	return v.VisitChildren(ctx)
}
