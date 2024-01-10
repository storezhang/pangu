package get

import (
	"github.com/pangum/pangu/internal/command"
	"github.com/pangum/pangu/internal/container"
)

type Command struct {
	container.Put

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
