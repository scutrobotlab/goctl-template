# goctl-template

## 概述

[goctl template](https://go-zero.dev/docs/tutorials/cli/template)

模板（Template）是数据驱动生成的基础，所有的代码（rest api、rpc、model、docker、kube）生成都会依赖模板，
默认情况下，模板生成器会选择内存中的模板进行生成，而对于有模板修改需求的开发者来讲，则需要将模板进行落盘，
从而进行模板修改，在下次代码生成时会加载指定路径下的模板进行生成。

## 生成

```shell
goctl template init --home .
```

## 使用

Go Get
```shell
go get github.com/scutrobotlab/goctl-template
```

Remote

```shell
goctl api go --api user.api --remote https://github.com/scutrobotlab/goctl-template --dir .
```

Home

```shell
git clone https://github.com/scutrobotlab/goctl-template.git
```

```shell
goctl api go --api user.api --home goctl-template --dir .
```
