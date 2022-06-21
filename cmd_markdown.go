package main

import (
	"io/ioutil"
	"os"
	"text/template"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed template.html
var page string

func newMarkdownRenderCmd() *cobra.Command {
	var (
		out  string
		opts struct {
			Title    string
			Markdown string
		}
	)
	cmd := &cobra.Command{
		Use:   "render <markdown>",
		Short: "Render a markdown document",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := template.New("page").Parse(page)
			if err != nil {
				return err
			}
			o := os.Stdout
			if len(out) > 0 {
				f, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					return err
				}
				defer f.Close()
				o = f
			}
			raw, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			opts.Markdown = string(raw)
			return t.Execute(o, opts)
		},
	}
	cmd.Flags().StringVarP(&out, "output", "o", "", "output file")
	cmd.Flags().StringVar(&opts.Title, "title", "Document", "page title")
	return cmd
}
