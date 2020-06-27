module knstart

go 1.13

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.1-0.20191009090205-6c0755d89d1e // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	go.uber.org/zap v1.14.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.17.6
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/pkg v0.0.0-20200624210428-eb05e8dd5b5b
	knative.dev/test-infra v0.0.0-20200624184627-7eae2aae38e4 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
)
