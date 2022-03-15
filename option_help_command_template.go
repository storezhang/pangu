package pangu

var (
	_        = CommandHelpTemplate
	_ option = (*optionCommandHelpTemplate)(nil)
)

type optionCommandHelpTemplate struct {
	template string
}

// CommandHelpTemplate 配置命令帮助信息
func CommandHelpTemplate(template string) *optionCommandHelpTemplate {
	return &optionCommandHelpTemplate{
		template: template,
	}
}

func (cht *optionCommandHelpTemplate) apply(options *options) {
	options.helpCommandTemplate = cht.template
}
