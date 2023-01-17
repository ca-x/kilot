package templates

import _ "embed"

var (
	// OdmBizLayerLogicTpl provides the default templates for biz using odm to generate.
	//go:embed odm_biz_layer_logic.tpl
	OdmBizLayerLogicTpl string

	// OdmDataLayerLogicTpl provides the default templates for data using odm to generate.
	//go:embed odm_data_layer_logic.tpl
	OdmDataLayerLogicTpl string

	// OdmUseCaseLogicTpl  provides the default templates for table using odm to generate.
	//go:embed odm_biz_use_case_logic.tpl
	OdmUseCaseLogicTpl string

	// NormalBizLayerLogicTpl provides the default template for biz using std mongo driver
	//go:embed normal_biz_layer_logic.tpl
	NormalBizLayerLogicTpl string

	// NormalDataLayerLogicTpl provides the default template for data using std mongo driver
	//go:embed normal_data_layer_logic.tpl
	NormalDataLayerLogicTpl string

	// NormalUseCaseLogicTpl  provides the default template for table using std mongo driver
	//go:embed normal_biz_use_case_logic.tpl
	NormalUseCaseLogicTpl string
)
