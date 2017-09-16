// Code generated by gocc; DO NOT EDIT.

package parser

import (
    . "github.com/awalterschulze/gominikanren/sexpr/ast"
    "github.com/awalterschulze/gominikanren/sexpr/token"
)

func getStr(v interface{}) string {
    t := v.(*token.Token)
    return string(t.Lit)
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : SExpr	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SExpr : Atom	<<  >>`,
		Id:         "SExpr",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SExpr : List	<<  >>`,
		Id:         "SExpr",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SExpr : QuotedList	<<  >>`,
		Id:         "SExpr",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `List : "(" ")"	<< NewList(false), nil >>`,
		Id:         "List",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewList(false), nil
		},
	},
	ProdTabEntry{
		String: `List : "(" SExpr ")"	<< NewList(false, X[1].(*SExpr)), nil >>`,
		Id:         "List",
		NTType:     2,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewList(false, X[1].(*SExpr)), nil
		},
	},
	ProdTabEntry{
		String: `List : "(" SExpr space ContinueList ")"	<< Prepend(false, X[1].(*SExpr), X[3].(*List)), nil >>`,
		Id:         "List",
		NTType:     2,
		Index:      6,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return Prepend(false, X[1].(*SExpr), X[3].(*List)), nil
		},
	},
	ProdTabEntry{
		String: `ContinueList : SExpr	<< &List{Items: []*SExpr{X[0].(*SExpr)}}, nil >>`,
		Id:         "ContinueList",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return &List{Items: []*SExpr{X[0].(*SExpr)}}, nil
		},
	},
	ProdTabEntry{
		String: `ContinueList : ContinueList space SExpr	<< Append(X[0].(*List), X[2].(*SExpr)), nil >>`,
		Id:         "ContinueList",
		NTType:     3,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return Append(X[0].(*List), X[2].(*SExpr)), nil
		},
	},
	ProdTabEntry{
		String: `QuotedList : backtick "(" ")"	<< NewList(true), nil >>`,
		Id:         "QuotedList",
		NTType:     4,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewList(true), nil
		},
	},
	ProdTabEntry{
		String: `QuotedList : backtick "(" SExpr ")"	<< NewList(true, X[2].(*SExpr)), nil >>`,
		Id:         "QuotedList",
		NTType:     4,
		Index:      10,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewList(true, X[2].(*SExpr)), nil
		},
	},
	ProdTabEntry{
		String: `QuotedList : backtick "(" SExpr space ContinueList ")"	<< Prepend(true, X[2].(*SExpr), X[4].(*List)), nil >>`,
		Id:         "QuotedList",
		NTType:     4,
		Index:      11,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return Prepend(true, X[2].(*SExpr), X[4].(*List)), nil
		},
	},
	ProdTabEntry{
		String: `Atom : symbol	<< NewSymbol(getStr(X[0])), nil >>`,
		Id:         "Atom",
		NTType:     5,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return NewSymbol(getStr(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Atom : int_lit	<< ParseInt(getStr(X[0])) >>`,
		Id:         "Atom",
		NTType:     5,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ParseInt(getStr(X[0]))
		},
	},
	ProdTabEntry{
		String: `Atom : float_lit	<< ParseFloat(getStr(X[0])) >>`,
		Id:         "Atom",
		NTType:     5,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ParseFloat(getStr(X[0]))
		},
	},
	ProdTabEntry{
		String: `Atom : string_lit	<< ParseString(getStr(X[0])) >>`,
		Id:         "Atom",
		NTType:     5,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ParseString(getStr(X[0]))
		},
	},
	ProdTabEntry{
		String: `Atom : variable	<< ParseVariable(getStr(X[0])) >>`,
		Id:         "Atom",
		NTType:     5,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ParseVariable(getStr(X[0]))
		},
	},
}
