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
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/gonvenience/bunt"
	"github.com/gonvenience/neat"
	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
	yamlv3 "gopkg.in/yaml.v3"
)

var restructureCmdSettings struct {
	inplace bool
}

// restructureCmd represents the paths command
var restructureCmd = &cobra.Command{
	Use:   "restructure <file>",
	Args:  cobra.ExactArgs(1),
	Short: "Restructure the order of keys",
	Long: func() string {
		exampleYAML := `---
releases:
- sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393
  name: binary-buildpack
  version: 1.0.32
  url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32
`

		var data yamlv3.Node
		_ = yamlv3.Unmarshal([]byte(exampleYAML), &data)

		before, _ := neat.ToYAMLString(data)
		ytbx.RestructureObject(&data)
		after, _ := neat.ToYAMLString(data)

		return bunt.Sprintf(`Restructure the order of keys in YAML maps.

The restructure logic tries to put human friendly and identifying keys such as
the name, or id key before other entries.

Example:
%s

Result:
%s
`, before, after)
	}(),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(_ *cobra.Command, args []string) error {
		location := args[0]

		input, err := ytbx.LoadFile(location)
		if err != nil {
			return err
		}

		for i := range input.Documents {
			ytbx.RestructureObject(input.Documents[i])
		}

		if restructureCmdSettings.inplace {
			info, err := os.Stat(location)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			writer := bufio.NewWriter(&buf)
			for _, document := range input.Documents {
				out, marshalErr := yamlv3.Marshal(document)
				if marshalErr != nil {
					return marshalErr
				}

				fmt.Fprint(writer, string(out))
			}

			writer.Flush()
			err = os.WriteFile(location, buf.Bytes(), info.Mode())
			if err != nil {
				return err
			}

		} else {
			for _, document := range input.Documents {
				out, err := neat.ToYAMLString(document)
				if err != nil {
					return err
				}

				fmt.Print(out)
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(restructureCmd)
	restructureCmd.Flags().SortFlags = false
	restructureCmd.Flags().BoolVarP(&restructureCmdSettings.inplace, "in-place", "i", false, "overwrite input file with output of this command")
	restructureCmd.Flags().BoolVarP(&ytbx.DisableRemainingKeySort, "disable-remaining-key-sort", "s", false, "disables that all unknown keys are sorted to improve the readability")
}
