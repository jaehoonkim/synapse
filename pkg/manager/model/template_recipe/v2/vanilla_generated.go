// Code generated by Ice-cream-maker DO NOT EDIT.
package v2
 
func (TemplateRecipe) ColumnNames() []string {
	return []string{
 		"id",
 		"name",
 		"summary",
 		"method",
 		"args",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (row TemplateRecipe) Values() []interface{} {
	return []interface{}{
		row.ID,
		row.Name,
		row.Summary,
		row.Method,
		row.Args,
		row.Created,
		row.Updated,
		row.Deleted,
	}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
 
func (row *TemplateRecipe) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ID,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	)
}
 
func (row *TemplateRecipe) Ptrs() []interface{} {
	return []interface{}{
		&row.ID,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	}
}
