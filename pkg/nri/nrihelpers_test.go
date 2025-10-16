package nri

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/containerd/nri/pkg/api"
)

var _ = Describe("NRI Helpers", func() {
	Context("getNetworkNamespace", func() {
		It("extracts network namespace path from pod sandbox", func() {
			pod := &api.PodSandbox{
				Linux: &api.LinuxPodSandbox{Namespaces: []*api.LinuxNamespace{{Type: "network", Path: "/proc/1/ns/net"}}},
			}
			Expect(getNetworkNamespace(pod)).To(Equal("/proc/1/ns/net"))
		})

		It("returns empty string when network namespace is missing", func() {
			pod := &api.PodSandbox{Linux: &api.LinuxPodSandbox{Namespaces: []*api.LinuxNamespace{{Type: "uts", Path: "/proc/1/ns/uts"}}}}
			Expect(getNetworkNamespace(pod)).To(Equal(""))
		})
	})
})
