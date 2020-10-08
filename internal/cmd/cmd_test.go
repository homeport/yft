// Copyright © 2020 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("program sub-commands", func() {
	Context("get command", func() {
		It("should get the value using Dot-Style path from local file", func() {
			output, err := yft("get", asset("examples", "simple.yml"), "list.one.somekey")
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEquivalentTo("foobar"))
		})
	})

	Context("compare command", func() {
		It("should return the common paths in two files", func() {
			output, err := yft("compare", asset("examples", "simple.yml"), asset("examples", "another.yml"))
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEquivalentTo("/list/name=one/somekey"))
		})
	})

	Context("paths command", func() {
		It("should return all paths in the given YAML file", func() {
			output, err := yft("paths", asset("examples", "simple.yml"))
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEquivalentTo("/list/name=one/somekey"))
		})
	})

	Context("restructure command", func() {
		It("should restructure the keys in the YAML file to be more human friendly to read", func() {
			output, err := yft("restructure", asset("examples", "restructure.yml"))
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEquivalentTo(`---
releases:
- name: binary-buildpack
  url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32
  version: 1.0.32
  sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393`))
		})

		It("should restructure the keys in the YAML file in-place to be more human friendly to read", func() {
			exampleYAML := []byte(`---
releases:
- sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393
  name: binary-buildpack
  version: 1.0.32
  url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32
`)
			tmpfile, err := ioutil.TempFile("", "exampleYAML")
			Expect(err).ToNot(HaveOccurred())
			defer os.Remove(tmpfile.Name())

			_, err = tmpfile.Write(exampleYAML)
			Expect(err).ToNot(HaveOccurred())

			err = tmpfile.Close()
			Expect(err).ToNot(HaveOccurred())

			output, err := yft("restructure", "--in-place", tmpfile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEmpty())

			actualBytes, err := ioutil.ReadFile(tmpfile.Name())
			Expect(err).ToNot(HaveOccurred())

			Expect(string(actualBytes)).To(BeEquivalentTo(`releases:
    - name: binary-buildpack
      url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32
      version: 1.0.32
      sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393
`))
		})
	})

	Context("version command", func() {
		It("should get the internal development version result", func() {
			output, err := yft("version")
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(BeEquivalentTo("version development"))
		})
	})
})
