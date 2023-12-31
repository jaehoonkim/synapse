// Code generated by Ice-cream-maker DO NOT EDIT.
package template
 
func (Template) ColumnNames() []string {
	return []string{
 		"uuid",
 		"name",
 		"summary",
 		"flow",
 		"inputs",
 		"origin",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (TemplateCommand) ColumnNames() []string {
	return []string{
 		"name",
 		"summary",
 		"inputs",
 		"outputs",
 		"client_version",
 		"category",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (row Template) Values() []interface{} {
	return []interface{}{
		row.Uuid,
		row.Name,
		row.Summary,
		row.Flow,
		row.Inputs,
		row.Origin,
		row.Created,
		row.Updated,
		row.Deleted,
	}
}
 
func (row TemplateCommand) Values() []interface{} {
	return []interface{}{
		row.Name,
		row.Summary,
		row.Inputs,
		row.Outputs,
		row.ClientVersion,
		row.Category,
		row.Created,
		row.Updated,
		row.Deleted,
	}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
 
func (row *Template) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Uuid,
		&row.Name,
		&row.Summary,
		&row.Flow,
		&row.Inputs,
		&row.Origin,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	)
}
 
func (row *TemplateCommand) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Name,
		&row.Summary,
		&row.Inputs,
		&row.Outputs,
		&row.ClientVersion,
		&row.Category,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	)
}
 
func (row *Template) Ptrs() []interface{} {
	return []interface{}{
		&row.Uuid,
		&row.Name,
		&row.Summary,
		&row.Flow,
		&row.Inputs,
		&row.Origin,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	}
}
 
func (row *TemplateCommand) Ptrs() []interface{} {
	return []interface{}{
		&row.Name,
		&row.Summary,
		&row.Inputs,
		&row.Outputs,
		&row.ClientVersion,
		&row.Category,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	}
}
