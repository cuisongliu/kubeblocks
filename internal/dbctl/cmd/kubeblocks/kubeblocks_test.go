/*
Copyright ApeCloud Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeblocks

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"

	"github.com/apecloud/kubeblocks/internal/dbctl/util/helm"
	"github.com/apecloud/kubeblocks/version"
)

var _ = Describe("kubeblocks", func() {
	var cmd *cobra.Command
	var streams genericclioptions.IOStreams

	BeforeEach(func() {
		streams, _, _, _ = genericclioptions.NewTestIOStreams()
	})

	It("kubeblocks", func() {
		tf := cmdtesting.NewTestFactory().WithNamespace("test")
		defer tf.Cleanup()

		cmd = NewKubeBlocksCmd(tf, streams)
		Expect(cmd).ShouldNot(BeNil())
		Expect(cmd.HasSubCommands()).Should(BeTrue())
	})

	It("check install", func() {
		tf := cmdtesting.NewTestFactory().WithNamespace("test")
		defer tf.Cleanup()

		var cfg string
		cmd = newInstallCmd(tf, streams)
		cmd.Flags().StringVar(&cfg, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")

		Expect(cmd).ShouldNot(BeNil())
		Expect(cmd.HasSubCommands()).Should(BeFalse())

		o := &installOptions{
			options: options{
				IOStreams: streams,
			},
		}
		Expect(o.complete(tf, cmd)).To(Succeed())
		Expect(o.Namespace).To(Equal("test"))
	})

	It("run install", func() {
		o := &installOptions{
			options: options{
				IOStreams: streams,
				cfg:       helm.FakeActionConfig(),
				Namespace: "default",
			},
			Version: version.DefaultKubeBlocksVersion,
			Monitor: true,
		}
		Expect(o.run()).To(Or(Succeed(), HaveOccurred()))
		Expect(len(o.Sets)).To(Equal(1))
		Expect(o.Sets[0]).To(Equal(kMonitorParam))
	})

	It("check uninstall", func() {
		tf := cmdtesting.NewTestFactory().WithNamespace("test")
		defer tf.Cleanup()

		var cfg string
		cmd = newUninstallCmd(tf, streams)
		cmd.Flags().StringVar(&cfg, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")

		Expect(cmd).ShouldNot(BeNil())
		Expect(cmd.HasSubCommands()).Should(BeFalse())

		o := &options{
			IOStreams: streams,
		}
		Expect(o.complete(tf, cmd)).To(Succeed())
		Expect(o.Namespace).To(Equal("test"))
	})

	It("run uninstall", func() {
		o := &options{
			IOStreams: streams,
			cfg:       helm.FakeActionConfig(),
			Namespace: "default",
		}

		Expect(o.run()).To(MatchError(MatchRegexp("Failed to uninstall KubeBlocks")))
	})
})
