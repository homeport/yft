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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var splitCmdSettings struct {
	directory string
}

// splitCmd represents the get command
var splitCmd = &cobra.Command{
	Use:           "split <file> [<file>] [...]",
	Aliases:       []string{"unravel", "break-down", "segment"},
	Short:         "Splits a multi-document file into separate files",
	Long:          "Splits a multi-document file into separate files.\n",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(_ *cobra.Command, args []string) error {
		for _, arg := range args {
			location := filepath.Clean(arg)

			stat, err := os.Stat(location)
			if err != nil {
				return err
			}

			inputfile, err := ytbx.LoadFile(location)
			if err != nil {
				return err
			}

			basename := filepath.Base(location)
			extension := filepath.Ext(basename)
			prefix := strings.TrimSuffix(basename, extension)

			for i, document := range inputfile.Documents {
				// Special case, ignore empty documents
				if len(document.Content) == 1 && document.Content[0].Tag == "!!null" {
					continue
				}

				var filename string = fmt.Sprintf("%s-%d%s", prefix, i, extension)
				if len(splitCmdSettings.directory) > 0 {
					filename = filepath.Join(splitCmdSettings.directory, filename)
				}

				bytes, err := yaml.Marshal(document)
				if err != nil {
					return err
				}

				if err := os.WriteFile(filename, append([]byte("---\n"), bytes...), stat.Mode()); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().SortFlags = false
	splitCmd.Flags().StringVarP(&splitCmdSettings.directory, "directory", "d", "", "Write files to directory rather than current working directory")
}
