package micro

import (
	"github.com/awalterschulze/gominikanren/sexpr/ast"
)

type Substitution = *ast.List

func EmptySubstitution() Substitution {
	return ast.NewList(true).List
}

type Goal func(Substitution) StreamOfSubstitutions

/*
(define (run-goal n g) 
	(takeInf n (g empty-s))
)
*/
func RunGoal(n int, g Goal) []Substitution {
	ss := g(EmptySubstitution())
	return takeStream(n, ss)
}

func SuccessO() Goal {
	return func(s Substitution) StreamOfSubstitutions {
		return NewSingletonStream(s)
	}
}

func FailureO() Goal {
	return func(s Substitution) StreamOfSubstitutions {
		return nil
	}
}

func EqualO(u, v *ast.SExpr) Goal {
	return func(s Substitution) StreamOfSubstitutions {
		ss, sok := unify(u, v, s)
		if sok {
			return NewSingletonStream(ss)
		}
		return nil
	}
}

/*
(define (nevero) 
	(lambda (s)
		(lambda () 
			((nevero) s)
		)
	)
)
*/
func NeverO() Goal {
	return func(s Substitution) StreamOfSubstitutions {
		return Suspension(func() StreamOfSubstitutions {
			return NeverO()(s)
		})
	}
}

/*
(define (alwayso) 
	(lambda (s)
		(lambda ()
			(
				(disj 
					S 
					(alwayso)
				)
				s
			)
		)
	)
)
*/
func AlwaysO() Goal {
	return func(s Substitution) StreamOfSubstitutions {
		return Suspension(func() StreamOfSubstitutions {
			return DisjointO(
				SuccessO(),
				AlwaysO(),
			)(s)
		})
	}
}