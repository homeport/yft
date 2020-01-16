// Copyright © 2018 The Homeport Team
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

package cmd

import (
	"fmt"

	"github.com/gonvenience/neat"
	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yft",
	Short: "YAML file tool",
	Long:  `YAML file tool provides a set of commands to work with the content of a given YAML file`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func getPathHelp() string {
	exampleYAML := `---
list:
- name: one
  somekey: foobar`

	examplePath := ytbx.Path{PathElements: []ytbx.PathElement{
		{Name: "list"},
		{Key: "name", Name: "one"},
		{Name: "somekey"},
	}}

	return fmt.Sprintf(`
There are two supported path syntaxes available:
- Dot-style (used by Spruce) or
- GoPatch-style (used by BOSH).

Example:
%s
Path in Dot-style would be: %s
The same path in GoPatch is: %s

`,
		neatYAML(exampleYAML),
		examplePath.ToDotStyle(),
		examplePath.ToGoPatchStyle())
}

func neatYAML(data string) string {
	node, err := ytbx.LoadYAMLDocuments([]byte(data))
	if err != nil {
		panic(err)
	}

	output, err := neat.ToYAMLString(node[0])
	if err != nil {
		panic(err)
	}

	return output
}
