package expr

type Stmt interface {
	Accept(visitor StmtVisitor) (interface{}, error)
}

type StmtVisitor interface {
	VisitExpressionStmt(expr *Expression) (interface{}, error)
	VisitPrintStmt(expr *Print) (interface{}, error)
}

type Expression struct {
	expression Expr
}

func (e *Expression) Accept(visitor StmtVisitor) (interface{}, error) {
	val, err := visitor.VisitExpressionStmt(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Print struct {
	expression Expr
}

func (e *Print) Accept(visitor StmtVisitor) (interface{}, error) {
	val, err := visitor.VisitPrintStmt(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}
