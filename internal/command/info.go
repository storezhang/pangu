package command

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	*Default
}

func NewInfo() *Info {
	return &Info{
		Default: New(constant.CommandInfo).Usage(`打印应用程序信息`).Aliases(`i`, `information`).Build(),
	}
}

func (i *Info) Run(_ *runtime.Context) (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{constant.HeaderName, constant.HeaderValue})
	table.Append([]string{constant.ColumnName, internal.Name})
	table.Append([]string{constant.ColumnVersion, internal.Version})
	table.Append([]string{constant.ColumnBuild, internal.Build})
	table.Append([]string{constant.ColumnComplied, internal.Compiled})
	table.Append([]string{constant.ColumnRevision, internal.Revision})
	table.Append([]string{constant.ColumnBranch, internal.Branch})
	table.Append([]string{constant.ColumnRuntime, internal.Runtime})
	table.Render()

	return
}