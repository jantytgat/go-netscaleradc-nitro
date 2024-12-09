package nitro

const (
	ResourcePackagePath = "github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/"
)

type ResourceReader interface {
	GetTypeName() string
}
