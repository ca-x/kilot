package templates

import _ "embed"

var (
	// OdmBizTpl provides the default templates for biz using odm to generate.
	//go:embed odm_biz.tpl
	OdmBizTpl string

	// OdmDataTpl provides the default templates for data using odm to generate.
	//go:embed odm_data.tpl
	OdmDataTpl string
)
