package pangu

var (
	_        = AppHelpTemplate
	_ option = (*optionAppHelpTemplate)(nil)
)

type optionAppHelpTemplate struct {
	template string
}

// AppHelpTemplate 配置应用帮助信息
func AppHelpTemplate(template string) *optionAppHelpTemplate {
	return &optionAppHelpTemplate{
		template: template,
	}
}

func (aht *optionAppHelpTemplate) apply(options *options) {
	options.helpAppTemplate = aht.template
}
