package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alephao/nftool/cmd"
	"github.com/spf13/cobra"
)

func printOptions(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("**Options:**\n\n```\n")
		flags.PrintDefaults()

		parentFlags := cmd.InheritedFlags()
		parentFlags.SetOutput(buf)

		if parentFlags.HasAvailableFlags() {
			buf.WriteString("\n")
			parentFlags.PrintDefaults()
		}

		buf.WriteString("```\n\n")
	}
	return nil
}

func genMarkdown(cmd *cobra.Command, w io.Writer) error {
	// cmd.InitDefaultHelpCmd()
	// cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	name := cmd.CommandPath()

	buf.WriteString("### `" + name + "`\n\n")
	buf.WriteString(cmd.Short + "\n\n")
	if len(cmd.Long) > 0 {
		buf.WriteString(cmd.Long + "\n\n")
	}

	if cmd.Runnable() {
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))
	}

	if len(cmd.Example) > 0 {
		buf.WriteString("**Example**:\n\n")
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.Example))
	}

	if err := printOptions(buf, cmd, name); err != nil {
		return err
	}

	_, err := buf.WriteTo(w)

	return err
}

func genMarkdownTree(cmd *cobra.Command, file *os.File) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := genMarkdownTree(c, file); err != nil {
			return err
		}
	}

	if len(cmd.Commands()) > 0 {
		return nil
	}

	if err := genMarkdown(cmd, file); err != nil {
		return err
	}
	return nil
}

func main() {
	rootCmd := cmd.NewRootCmd()

	filename := "docs.md"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	err = genMarkdownTree(rootCmd, f)

	if err != nil {
		log.Fatal(err)
	}
}
