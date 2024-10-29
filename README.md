## 1.开启go modules功能
### 1)GO111MODULE=off，关闭go modules功能，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
### 2)GO111MODULE=on，开启go modules功能，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
### 3)GO111MODULE=auto，默认值，go命令会根据当前目录中是否有go.mod文件来决定是否启用module功能。这种情况下可以分为两种情形：
#### A.当项目路径在GOPATH目录外部时， 设置为GO111MODULE = on
#### B.当项目路径位于GOPATH内部时，即使存在go.mod, 设置为GO111MODULE = off
go env -w GO111MODULE=auto
## 2.生成go.mod文件
go mod init gin-demo
## 3.下载并安装 gin
go get -u github.com/gin-gonic/gin