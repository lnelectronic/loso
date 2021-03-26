// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 25/3/2564 1:21
// ---------------------------------------------------------------------------
package models

// Filter Optine Model
type Filter struct {
	Skip      *int64
	Limit     *int64
	SortKey   string
	SortVal   int
	Condition interface{}
}

//New Instance Farm
func (f *Filter) New() *Filter {
	return &Filter{

		Skip:      f.Skip,
		Limit:     f.Limit,
		SortKey:   f.SortKey,
		SortVal:   f.SortVal,
		Condition: f.Condition,
	}
}
