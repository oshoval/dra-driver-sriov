package cni

import (
	"net"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	cni100 "github.com/containernetworking/cni/pkg/types/100"
	resourcev1 "k8s.io/api/resource/v1"
)

var _ = Describe("CNI Status Conversion", func() {
	Context("cniResultToNetworkData", func() {
		It("converts CNI result to NetworkDeviceData correctly", func() {
			res := &cni100.Result{
				CNIVersion: "1.0.0",
				Interfaces: []*cni100.Interface{
					{Name: "eth0", Mac: "aa:bb:cc:dd:ee:ff", Sandbox: "/proc/1/ns/net"},
				},
				IPs: []*cni100.IPConfig{
					{Address: mustParseCIDR("10.1.2.3/24")},
				},
			}

			nd, err := cniResultToNetworkData(res)
			Expect(err).ToNot(HaveOccurred())
			Expect(nd).To(Equal(&resourcev1.NetworkDeviceData{
				InterfaceName:   "eth0",
				HardwareAddress: "aa:bb:cc:dd:ee:ff",
				IPs:             []string{"10.1.2.0/24"},
			}))
		})
	})
})

func mustParseCIDR(s string) (out net.IPNet) {
	_, ipn, _ := net.ParseCIDR(s)
	out = *ipn
	return
}
