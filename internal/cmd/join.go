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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gonvenience/ytbx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var joinCmdSettings struct {
	filename string
	ending   bool
}

// joinCmd represents the get command
var joinCmd = &cobra.Command{
	Use:           "join <file> [<file>] [...]",
	Aliases:       []string{"combine"},
	Short:         "Joins multiple YAML files into a multi-document file",
	Long:          "Joins multiple YAML files into a multi-document file.\n",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var buf bytes.Buffer

		for _, location := range args {
			location := filepath.Clean(location)

			inputfile, err := ytbx.LoadFile(location)
			if err != nil {
				return err
			}

			for _, document := range inputfile.Documents {
				out, err := yaml.Marshal(document)
				if err != nil {
					return err
				}

				fmt.Fprintf(&buf, "--- # file %s\n", location)
				buf.Write(out)

				if joinCmdSettings.ending {
					fmt.Fprint(&buf, "...\n")
				}
			}
		}

		switch joinCmdSettings.filename {
		case "-":
			fmt.Print(buf.String())

		default:
			return ioutil.WriteFile(joinCmdSettings.filename, buf.Bytes(), os.FileMode(0644))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)

	joinCmd.PersistentFlags().StringVarP(&joinCmdSettings.filename, "name", "f", "", "Name of the target file")
	joinCmd.PersistentFlags().BoolVarP(&joinCmdSettings.ending, "document-ending", "e", false, "Write YAML document ending ('...') after each document")
}
