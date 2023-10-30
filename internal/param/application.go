package param

import (
	"time"

	"github.com/pangum/pangu/internal/internal/app"
)

type Application struct {
	*Config
	*Help
	*Banner
	*Code

	// 合法性验证，包括
	// 启动器构造方法合法性验证
	// 构造方法合法性验证
	Verify bool
	// 应用描述
	Description string
	// 使用方式
	Usage string

	// 版权
	Copyright string
	// 作者
	Authors app.Authors
	// 元数据
	Metadata map[string]any

	// 超时
	Timeout time.Duration
}

func NewApplication() *Application {
	return &Application{
		Config: newConfig(),
		Help:   newHelp(),
		Banner: newBanner(),
		Code:   newCode(),

		Verify:      true,
		Description: "一个使用github.com/pangum构建的应用程序",
		Authors: app.Authors{{
			Name:  "storezhang",
			Email: "storezhang@gmail.com",
		}},
		Metadata: make(map[string]any),

		Timeout: 15 * time.Second,
	}
}

func (a *Application) Override(latest *Application) {
	if latest.Config.Set {
		a.Config = latest.Config
	}
	if latest.Help.Set {
		a.Help = latest.Help
	}
	if latest.Banner.Set {
		a.Banner = latest.Banner
	}
	if latest.Code.Set {
		a.Code = latest.Code
	}
}
