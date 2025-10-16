package cni

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RawExec", func() {
	Context("pluginErr", func() {
		It("formats error messages correctly", func() {
			e := &RawExec{}
			err := e.pluginErr(errors.New("boom"), nil, []byte("some-stderr"))
			Expect(err.Error()).To(ContainSubstring("netplugin failed"))
		})
	})
})
