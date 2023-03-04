https://liwenzhou.com/posts/Go/install/
# Go Simple Tutorial

> Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言 [wiki](https://zh.wikipedia.org/wiki/Go)

## 1. Go 安装

[golang_download](https://golang.org/dl/)

```shell
$ wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
$ tar -zxvf go1.20.1.linux-amd64.tar.gz -C /usr/local 
```

配置环境变量： Linux下有两个文件可以配置环境变量，其中 ` /etc/profile ` 是对所有用户生效的；` $HOME/.profile ` 是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。
```shell
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```
修改 `/etc/profile` 后要重启生效，修改 `$HOME/.profile` 后使用 source 命令加载 `$HOME/.profile` 文件即可生效。 检查：
```shell
$ go version
```

从 Go `1.11` 版本开始，Go 提供了 [Go Modules](https://github.com/golang/go/wiki/Modules) 的机制，推荐设置以下环境变量，第三方包的下载将通过国内镜像，避免出现官方网址被屏蔽的问题。

```shell
$ go env -w GOPROXY=https://goproxy.cn,direct
```
`GOROOT` 和 `GOPATH` 都是环境变量，其中 `GOROOT` 是我们安装 go 开发包的路径，而从 Go 1.8 版本开始，Go开发包在安装完成后会为 `GOPATH` 设置一个默认目录，并且在Go1.14及之后的版本中启用了Go Module模式之后，不一定非要将代码写到GOPATH目录下，所以也就不需要我们再自己配置GOPATH了，使用默认的即可。

想要查看你电脑上的 `GOPATH` 路径，只需要打开终端输入以下命令并回车：
```shell
go env
```
`GOPROXY` 非常重要
Go1.14 版本之后，都推荐使用 `go mod` 模式来管理依赖环境了，也不再强制我们把代码必须写在 `GOPATH` 下面的src目录了，你可以在你电脑的任意位置编写go代码。（网上有些教程适用于1.11版本之前。）

默认GoPROXY配置是：`GOPROXY=https://proxy.golang.org,direct`，由于国内访问不到 `https://proxy.golang.org` ，所以我们需要换一个 `PROXY` ，这里推荐使用 `https://goproxy.io或https://goproxy.cn` 。

可以执行下面的命令修改GOPROXY：
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```

## 2. Hello World

现在我们来创建第一个 Go 项目 —— `hello`。在我们桌面创建一个 `hello` 目录。

**go mod init**


使用 `go module` 模式新建项目时，我们需要通过 `go mod init` 项目名命令对项目进行初始化，该命令会在项目根目录下生成 `go.mod ` 文件。例如，我们使用 `hello` 作为我们第一个Go项目的名称，执行如下命令。
```shell
go mod int hello
```

**编写代码**

接下来在该目录中创建一个 `main.go` 文件:
```go
package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 包

func main(){  // main函数，是程序执行的入口
	fmt.Println("Hello World!")  // 在终端打印 Hello World!
}
```

**编译**

`go build` 命令表示将源代码编译成可执行文件。

在 hello 目录下执行：
```shell
go build
```


我们还可以使用 `-o` 参数来指定编译后得到的可执行文件的名字。
```shell
go build -o hello
```

**go run**

`go run main.go` 也可以执行程序，该命令本质上是先在临时目录编译程序然后再执行。

如果你不清楚上方关于` go run ` 执行机制的描述，那么你最好今后都使用 `go build` 编译再执行。

**go install**

`go instal`l表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的`bin`目录下。因为我们把`GOPATH`下的`bin`目录添加到了环境变量中，所以我们就可以在任意地方直接执行可执行文件了

**跨平台编译**

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，Go语言支持跨平台编译——在当前平台（例如Windows）下编译其他平台（例如Linux）的可执行文件。

**Windows编译Linux可执行文件**

如果我想在Windows下编译一个Linux下可执行文件，那需要怎么做呢？只需要在编译时指定目标操作系统的平台和处理器架构即可。

注意：无论你在Windows电脑上使用VsCode编辑器还是Goland编辑器，都要注意你使用的终端类型，因为不同的终端下命令不一样！！！目前的Windows通常默认使用的是PowerShell终端。

如果你的Windows使用的是cmd，那么按如下方式指定环境变量。
```shell
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

如果你的Windows使用的是PowerShell终端，那么设置环境变量的语法为

```shell
$ENV:CGO_ENABLED=0
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
```

在你的Windows终端下执行完上述命令后，再执行下面的命令，得到的就是能够在Linux平台运行的可执行文件了。
```shell
go build
```

**Windows编译Mac可执行文件**

Windows下编译Mac平台64位可执行程序：

cmd终端下执行：
```shell
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

PowerShell终端下执行：
```shell
$ENV:CGO_ENABLED=0
$ENV:GOOS="darwin"
$ENV:GOARCH="amd64"
go build
```


**Mac编译Linux可执行文件**

Mac电脑编译得到Linux平台64位可执行程序：
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

**Mac编译Windows可执行文件**

Mac电脑编译得到Windows平台64位可执行程序：
```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

**Linux编译Mac可执行文件**

Linux平台下编译Mac平台64位可执行程序：
```shell
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
```

**Linux编译Windows可执行文件**

Linux平台下编译Windows平台64位可执行程序：
```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```
