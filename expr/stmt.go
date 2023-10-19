package expr

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	VisitExpressionStmt(expr *Expression) interface{}
	VisitPrintStmt(expr *Print) interface{}
}

type Expression struct {
	expression Expr
}

func (e *Expression) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(e)
}

type Print struct {
	expression Expr
}

func (e *Print) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmt(e)
}
