package common

const (
	GoVersionRequire = "1.18.0"
)

const (
	AppName    = "snt"
	AppDesc    = "Sonata脚手架"
	AppGitRepo = "github.com/golang/protobuf/protoc-gen-go"
	AppVersion = "0.0.1"
)

type Dependency struct {
	Name    string
	Version string
	Fix     *Fix
}

const (
	FixMethodHint      = "hint"
	FixMethodGoInstall = "go_get"
)

type Fix struct {
	Type   string
	Method string
}

var Dependencies = []*Dependency{
	{
		Name:    "protoc",
		Version: "latest",
		Fix: &Fix{
			Type:   FixMethodHint,
			Method: "https://github.com/protocolbuffers/protobuf/releases 选择合适版本下载并将 protoc 放入 $PATH 中",
		},
	},
	{
		Name:    "protoc-gen-go",
		Version: "latest",
		Fix: &Fix{
			Type:   FixMethodGoInstall,
			Method: "go install github.com/golang/protobuf/protoc-gen-go@latest",
		},
	},
}
