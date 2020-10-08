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

	"github.com/gonvenience/bunt"
	"github.com/gonvenience/wrap"
	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
)

// pathsCmd represents the paths command
var pathsCmd = &cobra.Command{
	Use:   "paths <file>",
	Args:  cobra.ExactArgs(1),
	Short: "List all paths to values",
	Long: func() string {
		exampleYAML := `---
yaml:
  structure:
    somekey: foobar
`

		examplePath := "/yaml/structure/somekey"

		return bunt.Sprintf(`Lists all paths to values in the given YAML file.

Example:
%s

would list you one path: %s
`,
			neatYAML(exampleYAML),
			examplePath,
		)
	}(),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := ytbx.ListPaths(args[0])
		if err != nil {
			return wrap.Error(err, "failed to get paths from file")
		}

		for _, entry := range list {
			fmt.Println(entry.ToGoPatchStyle())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pathsCmd)
}
