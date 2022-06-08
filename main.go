package main

import (
	"flag"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		glog.Errorf("execute error: %+v", err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	flag.CommandLine.Parse(nil)
	cmd := &cobra.Command{
		Use: "benchkit",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			glog.Flush()
		},
	}
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		if f.Name == "log_backtrace_at" {
			f.Name = "logbacktrace"
		}
		if f.Name == "log_dir" {
			f.Name = "logdir"
			f.Usage = strings.ToLower(f.Usage)
		}
		cmd.PersistentFlags().AddGoFlag(f)
	})
	cmd.AddCommand(newMetricsCmd())
	return cmd
}
