package nitro

const (
	ResourcePackagePath = "github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/"
)

type ResourceReader interface {
	GetTypeName() string
}
