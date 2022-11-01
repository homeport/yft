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
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/homeport/yft/internal/cmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "YAML File Tool command line tests")
}

func asset(path ...string) string {
	return filepath.Join(append([]string{"..", "..", "assets"}, path...)...)
}

func yft(args ...string) (out string, err error) {
	r, w, err := os.Pipe()
	Expect(err).ToNot(HaveOccurred())

	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
	}()

	os.Stdout = w
	os.Args = append([]string{"yft"}, args...)
	err = Execute()
	w.Close()

	var buf bytes.Buffer
	_, copyErr := io.Copy(&buf, r)
	Expect(copyErr).ToNot(HaveOccurred())

	return strings.TrimRight(buf.String(), "\n"), err
}
