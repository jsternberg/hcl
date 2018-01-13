package hcl

// ExprList tests if the given expression is a static list construct and,
// if so, extracts the expressions that represent the list elements.
// If the given expression is not a static list, error diagnostics are
// returned.
//
// A particular Expression implementation can support this function by
// offering a method called ExprList that takes no arguments and returns
// []Expression. This method should return nil if a static list cannot
// be extracted.
func ExprList(expr Expression) ([]Expression, Diagnostics) {
	type exprList interface {
		ExprList() []Expression
	}

	if exL, supported := expr.(exprList); supported {
		if list := exL.ExprList(); list != nil {
			return list, nil
		}
	}
	return nil, Diagnostics{
		&Diagnostic{
			Severity: DiagError,
			Summary:  "Invalid expression",
			Detail:   "A static list expression is required.",
			Subject:  expr.StartRange().Ptr(),
		},
	}
}
