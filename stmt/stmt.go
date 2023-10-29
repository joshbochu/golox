package stmt

import (
	"github.com/joshbochu/golox/expr"
	"github.com/joshbochu/golox/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) (interface{}, error)
}

type StmtVisitor interface {
	VisitExpressionStmt(expr *Expression) (interface{}, error)
	VisitPrintStmt(expr *Print) (interface{}, error)
	VisitVarStmt(expr *Var) (interface{}, error)
}

type Expression struct {
	Expression expr.Expr
}

func (e *Expression) Accept(visitor StmtVisitor) (interface{}, error) {
	val, err := visitor.VisitExpressionStmt(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Print struct {
	Expression expr.Expr
}

func (e *Print) Accept(visitor StmtVisitor) (interface{}, error) {
	val, err := visitor.VisitPrintStmt(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type Var struct {
	Name        token.Token
	Initializer expr.Expr
}

func (e *Var) Accept(visitor StmtVisitor) (interface{}, error) {
	val, err := visitor.VisitVarStmt(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}
