package v1

import (
	"github.com/NexClipper/sudory/pkg/server/model"
	metav1 "github.com/NexClipper/sudory/pkg/server/model/meta/v1"
	"github.com/NexClipper/sudory/pkg/server/model/orm"
)

//TemplateCommand Property
type TemplateCommandProperty struct {
	//템플릿 UUID
	TemplateUuid string `json:"template_uuid" xorm:"char(32) notnull index 'template_uuid' comment('templates uuid')"`
	//순서
	Sequence int32 `json:"sequence,omitempty" xorm:"int notnull 'sequence' comment('sequence')"`
	//메소드
	//@example: "kubernetes.deployment.get.v1", "kubernetes.pod.list.v1"
	Method string `json:"method,omitempty" xorm:"varchar(255) notnull 'method' comment('method')"`
	//arguments
	Args map[string]string `json:"args,omitempty" xorm:"text null 'args' comment('args')"`
}

//MODEL: TEMPLATE_COMMAND
type TemplateCommand struct {
	metav1.LabelMeta        `json:",inline" xorm:"extends"` //inline labelmeta
	TemplateCommandProperty `json:",inline" xorm:"extends"` //inline property
}

//DATABASE SCHEMA: TEMPLATE_COMMAND
type DbSchemaTemplateCommand struct {
	metav1.DbMeta   `xorm:"extends"`
	TemplateCommand `xorm:"extends"`
}

var _ orm.TableName = (*DbSchemaTemplateCommand)(nil)

func (DbSchemaTemplateCommand) TableName() string {
	return "template_command"
}

//HTTP REQUEST BODY: TEMPLATE
type HttpReqTemplateCommand struct {
	TemplateCommand `json:",inline"`
}

//HTTP RESPONSE BODY: TEMPLATE
type HttpRspTemplateCommand struct {
	TemplateCommand `json:",inline"`
}

var _ model.Modeler = (*HttpRspTemplateCommand)(nil)

func (HttpRspTemplateCommand) GetType() string {
	return "HTTP RSP TEMPLATE_COMMAND"
}

//HTTP RESPONSE BODY: MANY TEMPLATE
type HttpRspTemplateCommands []HttpRspTemplateCommand

var _ model.Modeler = (*HttpRspTemplateCommands)(nil)

func (HttpRspTemplateCommands) GetType() string {
	return "HTTP RSP []TEMPLATE_COMMAND"
}

//변환 TemplateCommand -> DbSchema*
func TransToDbSchema(s []TemplateCommand) []DbSchemaTemplateCommand {
	var out = make([]DbSchemaTemplateCommand, len(s))
	for n, it := range s {
		out[n] = DbSchemaTemplateCommand{TemplateCommand: it}
	}
	return out
}

//변환 HttpReq* -> TemplateCommand
func TransFormHttpReqTemplate(s []HttpReqTemplateCommand) []TemplateCommand {
	var out = make([]TemplateCommand, len(s))
	for n, it := range s {
		out[n] = it.TemplateCommand
	}
	return out
}
