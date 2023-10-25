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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:           "get <file> <path>",
	Args:          cobra.ExactArgs(2),
	Short:         "Get the value at a given path",
	Long:          "Get the value at a given path in the file.\n" + getPathHelp(),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(_ *cobra.Command, args []string) error {
		location := args[0]
		pathString := args[1]

		inputfile, err := ytbx.LoadFile(location)
		if err != nil {
			return err
		}

		obj, err := ytbx.Grab(inputfile.Documents[0], pathString)
		if err != nil {
			return err
		}

		boldKeys := true
		useIndentLines := true
		outputProcessor := neat.NewOutputProcessor(useIndentLines, boldKeys, &neat.DefaultColorSchema)

		output, err := outputProcessor.ToYAML(obj)
		if err != nil {
			return fmt.Errorf("failed to render gathered data: %w", err)
		}

		fmt.Print(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
