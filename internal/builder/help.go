package builder

import (
	"github.com/pangum/pangu/internal/param"
)

type Help struct {
	params      *param.Help
	application *Application
}

func NewHelp(application *Application) (help *Help) {
	help = new(Help)
	help.params = application.params.Help
	help.application = application

	return
}

func (h *Help) App(tooltip string) (help *Help) {
	h.params.App = tooltip
	help = h

	return
}

func (h *Help) Command(tooltip string) (help *Help) {
	h.params.Command = tooltip
	help = h

	return
}

func (h *Help) Subcommand(tooltip string) (help *Help) {
	h.params.Subcommand = tooltip
	help = h

	return
}

func (h *Help) Build() (application *Application) {
	h.params.Set = true
	h.application.params.Help = h.params
	application = h.application

	return
}
