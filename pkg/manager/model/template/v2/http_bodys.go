package template

type HttpRsp_Template struct {
	Template `json:",inline"`
	Commands []TemplateCommand `json:"commands,omitempty"`
}

type HttpRsp_TemplateCommand = TemplateCommand

type HttpReqTemplate_Update = Template_update
