# 介绍

盘古是一个以Golang语言写成的应用快速开发框架，是在工作中的实际经验而写成的框架
::: danger 特别强调
<!--@formatter:off-->
当你代码量少的时候，很难体会到`盘古`框架的优势
<!--@formatter:on-->
:::

## 它是如何工作的

盘古的设计理念是：在系统复杂的时候代码依然简单！实现方式也很简单，通过内置的依赖注入来管理系统间的各种依赖而不是依赖于编写各种`init`初始化代码。
同时在系统中增加内置的命令行支持，方便添加自定义的命令。而在用户侧，只需要关注自己的业务代码，添加到盘古内部后，盘古负责将这些代码管理起来；用 一句话来概括就是：只关注自己的业务逻辑，而将应用程序启动以及各种配置项完全交给盘古来管理

## 功能特性

**内置依赖注入框架**

- 运行时依赖注入
- 强大的类型系统

**内置配置文件加载**

- 支持`Yaml`
- 支持`Toml`
- 支持`Json`
- 支持`XML`
- 非常容易扩展支持其它配置格式

**内置支持数据验证**

- 越来越多的内置验证器
- 国际化

**丰富的内置功能支持**

- 支持输出用户良好的版本信息
- 支持添加服务
- 支持添加命令
- 支持添加参数
- 支持数据迁移

**非常强大的配置功能**

- 所有配置都是接口实现（易于找到全部的配置项）
- 支持默认配置（零配置，开箱即用）
- 支持配置覆盖（可以在程序的任何地方覆盖原来的配置）

## 设计原则

力求保持简单是`盘古`在代码设计中最大的原则，此外还有

- 语义化（比如`Name`可以配置的项有：`应用程序`名称、`添加依赖`时的名称以及`获取依赖`时的名称）
- 尽量不暴露内部实现
- 线程安全