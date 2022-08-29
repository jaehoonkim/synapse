// Code generated by Ice-cream-maker DO NOT EDIT.
package v3
 
func (Service_create) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"timestamp",
 		"pdate",
 		"name",
 		"summary",
 		"template_uuid",
 		"step_count",
 		"subscribed_channel",
 		"step_position",
 		"status",
 		"created",
	}
}
 
func (Service) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"timestamp",
 		"pdate",
 		"name",
 		"summary",
 		"template_uuid",
 		"step_count",
 		"subscribed_channel",
 		"assigned_client_uuid",
 		"step_position",
 		"status",
 		"message",
 		"created",
	}
}
 
func (ServiceResult_create) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"timestamp",
 		"pdate",
 		"result_type",
 		"result",
	}
}
 
func (ServiceResult) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"timestamp",
 		"pdate",
 		"result_type",
 		"result",
	}
}
 
func (ServiceStep_create) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"seq",
 		"timestamp",
 		"pdate",
 		"name",
 		"summary",
 		"method",
 		"args",
 		"result_filter",
 		"status",
 		"created",
	}
}
 
func (ServiceStep) ColumnNames() []string {
	return []string{
 		"cluster_uuid",
 		"uuid",
 		"seq",
 		"timestamp",
 		"pdate",
 		"name",
 		"summary",
 		"method",
 		"args",
 		"result_filter",
 		"status",
 		"started",
 		"ended",
 		"created",
	}
}
 
func (Service_polling) ColumnNames() []string {
	return []string{
 		"uuid",
 		"timestamp",
 		"status",
 		"created",
	}
}
 
func (row Service_create) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Timestamp,
		row.PartitionDate,
		row.Name,
		row.Summary,
		row.TemplateUuid,
		row.StepCount,
		row.SubscribedChannel,
		row.StepPosition,
		row.Status,
		row.Created,
	}
}
 
func (row Service) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Timestamp,
		row.PartitionDate,
		row.Name,
		row.Summary,
		row.TemplateUuid,
		row.StepCount,
		row.SubscribedChannel,
		row.AssignedClientUuid,
		row.StepPosition,
		row.Status,
		row.Message,
		row.Created,
	}
}
 
func (row ServiceResult_create) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Timestamp,
		row.PartitionDate,
		row.ResultSaveType,
		row.Result,
	}
}
 
func (row ServiceResult) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Timestamp,
		row.PartitionDate,
		row.ResultSaveType,
		row.Result,
	}
}
 
func (row ServiceStep_create) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Sequence,
		row.Timestamp,
		row.PartitionDate,
		row.Name,
		row.Summary,
		row.Method,
		row.Args,
		row.ResultFilter,
		row.Status,
		row.Created,
	}
}
 
func (row ServiceStep) Values() []interface{} {
	return []interface{}{
		row.ClusterUuid,
		row.Uuid,
		row.Sequence,
		row.Timestamp,
		row.PartitionDate,
		row.Name,
		row.Summary,
		row.Method,
		row.Args,
		row.ResultFilter,
		row.Status,
		row.Started,
		row.Ended,
		row.Created,
	}
}
 
func (row Service_polling) Values() []interface{} {
	return []interface{}{
		row.Uuid,
		row.Timestamp,
		row.Status,
		row.Created,
	}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
 
func (row *Service_create) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.TemplateUuid,
		&row.StepCount,
		&row.SubscribedChannel,
		&row.StepPosition,
		&row.Status,
		&row.Created,
	)
}
 
func (row *Service) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.TemplateUuid,
		&row.StepCount,
		&row.SubscribedChannel,
		&row.AssignedClientUuid,
		&row.StepPosition,
		&row.Status,
		&row.Message,
		&row.Created,
	)
}
 
func (row *ServiceResult_create) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.ResultSaveType,
		&row.Result,
	)
}
 
func (row *ServiceResult) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.ResultSaveType,
		&row.Result,
	)
}
 
func (row *ServiceStep_create) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Sequence,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.ResultFilter,
		&row.Status,
		&row.Created,
	)
}
 
func (row *ServiceStep) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ClusterUuid,
		&row.Uuid,
		&row.Sequence,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.ResultFilter,
		&row.Status,
		&row.Started,
		&row.Ended,
		&row.Created,
	)
}
 
func (row *Service_polling) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Uuid,
		&row.Timestamp,
		&row.Status,
		&row.Created,
	)
}
 
func (row *Service_create) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.TemplateUuid,
		&row.StepCount,
		&row.SubscribedChannel,
		&row.StepPosition,
		&row.Status,
		&row.Created,
	}
}
 
func (row *Service) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.TemplateUuid,
		&row.StepCount,
		&row.SubscribedChannel,
		&row.AssignedClientUuid,
		&row.StepPosition,
		&row.Status,
		&row.Message,
		&row.Created,
	}
}
 
func (row *ServiceResult_create) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.ResultSaveType,
		&row.Result,
	}
}
 
func (row *ServiceResult) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Timestamp,
		&row.PartitionDate,
		&row.ResultSaveType,
		&row.Result,
	}
}
 
func (row *ServiceStep_create) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Sequence,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.ResultFilter,
		&row.Status,
		&row.Created,
	}
}
 
func (row *ServiceStep) Ptrs() []interface{} {
	return []interface{}{
		&row.ClusterUuid,
		&row.Uuid,
		&row.Sequence,
		&row.Timestamp,
		&row.PartitionDate,
		&row.Name,
		&row.Summary,
		&row.Method,
		&row.Args,
		&row.ResultFilter,
		&row.Status,
		&row.Started,
		&row.Ended,
		&row.Created,
	}
}
 
func (row *Service_polling) Ptrs() []interface{} {
	return []interface{}{
		&row.Uuid,
		&row.Timestamp,
		&row.Status,
		&row.Created,
	}
}
