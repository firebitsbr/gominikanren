// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(5),  /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(6),  /* backtick */
			shift(7),  /* symbol */
			shift(8),  /* int_lit */
			shift(9),  /* float_lit */
			shift(10), /* string_lit */
			shift(11), /* variable */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* space */
			nil,          /* backtick */
			nil,          /* symbol */
			nil,          /* int_lit */
			nil,          /* float_lit */
			nil,          /* string_lit */
			nil,          /* variable */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: SExpr */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: SExpr */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: SExpr */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			shift(17), /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(24), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(14), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(15), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(16), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(25), /* ) */
			shift(26), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(1), /* ), reduce: SExpr */
			reduce(1), /* space, reduce: SExpr */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(2), /* ), reduce: SExpr */
			reduce(2), /* space, reduce: SExpr */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(3), /* ), reduce: SExpr */
			reduce(3), /* space, reduce: SExpr */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			shift(28), /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: List */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(29), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(12), /* ), reduce: Atom */
			reduce(12), /* space, reduce: Atom */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(13), /* ), reduce: Atom */
			reduce(13), /* space, reduce: Atom */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(14), /* ), reduce: Atom */
			reduce(14), /* space, reduce: Atom */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(15), /* ), reduce: Atom */
			reduce(15), /* space, reduce: Atom */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(16), /* ), reduce: Atom */
			reduce(16), /* space, reduce: Atom */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			shift(31), /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: List */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(34), /* ) */
			shift(35), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(4), /* ), reduce: List */
			reduce(4), /* space, reduce: List */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			shift(37), /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(38), /* ) */
			shift(39), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: QuotedList */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(7), /* ), reduce: ContinueList */
			reduce(7), /* space, reduce: ContinueList */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(40), /* ) */
			shift(41), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(5), /* ), reduce: List */
			reduce(5), /* space, reduce: List */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(43), /* ) */
			shift(44), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(9), /* ), reduce: QuotedList */
			reduce(9), /* space, reduce: QuotedList */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: QuotedList */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: List */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(47), /* ) */
			shift(41), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(10), /* ), reduce: QuotedList */
			reduce(10), /* space, reduce: QuotedList */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(18), /* backtick */
			shift(19), /* symbol */
			shift(20), /* int_lit */
			shift(21), /* float_lit */
			shift(22), /* string_lit */
			shift(23), /* variable */
		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(49), /* ) */
			shift(41), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(8), /* ), reduce: ContinueList */
			reduce(8), /* space, reduce: ContinueList */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(6), /* ), reduce: List */
			reduce(6), /* space, reduce: List */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(50), /* ) */
			shift(41), /* space */
			nil,       /* backtick */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: QuotedList */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(11), /* ), reduce: QuotedList */
			reduce(11), /* space, reduce: QuotedList */
			nil,        /* backtick */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
}
