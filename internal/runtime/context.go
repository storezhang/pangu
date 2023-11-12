package runtime

import (
	"context"
	"time"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/urfave/cli/v2"
)

var _ context.Context = (*Context)(nil)

// Context 描述上下文
type Context struct {
	context *cli.Context
}

// NewContext 创建上下文
func NewContext(ctx *cli.Context) *Context {
	return &Context{
		context: ctx,
	}
}

func (c *Context) Deadline() (time.Time, bool) {
	return c.context.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *Context) Err() error {
	return c.context.Err()
}

func (c *Context) String(name string) string {
	return c.context.String(name)
}

func (c *Context) Set(name string, value string) error {
	return c.context.Set(name, value)
}

func (c *Context) Arguments() app.Arguments {
	return c.context.Args()
}

func (c *Context) Value(key any) any {
	return c.context.Value(gox.ToString(key))
}
