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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// splitCmd represents the get command
var splitCmd = &cobra.Command{
	Use:           "split <file> [<file>] [...]",
	Aliases:       []string{"unravel", "break-down", "segment"},
	Short:         "Splits a multi-document file into separate files",
	Long:          "Splits a multi-document file into separate files.\n",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, location := range args {
			stat, err := os.Stat(location)
			if err != nil {
				return err
			}

			inputfile, err := ytbx.LoadFile(location)
			if err != nil {
				return err
			}

			extension := filepath.Ext(location)
			base := strings.TrimSuffix(location, extension)

			for i, document := range inputfile.Documents {
				filename := fmt.Sprintf("%s-%d%s", base, i, extension)

				bytes, err := yaml.Marshal(document)
				if err != nil {
					return err
				}

				if err := ioutil.WriteFile(filename, bytes, stat.Mode()); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
