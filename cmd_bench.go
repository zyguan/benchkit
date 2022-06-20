package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

func newBenchRunCmd() *cobra.Command {
	var (
		tags   []string
		name   string
		report string
		prom   string
		vm     string
		filter string
		opts   PromDumpOptions
		store  VictoriaMetricsStore
	)
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a benchmark command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			if len(name) == 0 {
				name = args[0]
			}
			if len(filter) > 0 {
				re, err := regexp.Compile(filter)
				if err != nil {
					return err
				}
				opts.Filter = re
			}
			result := BenchRun(ctx, name, tags, args[0], args[1:]...)
			glog.Infof("bench result: id=%s exit=%d", result.ID, result.Exit)

			if len(prom) > 0 && len(vm) > 0 {
				glog.Infof("copy metrics from %s to %s", prom, vm)
				opts.Endpoint, store.Endpoint = prom, vm
				opts.Start, opts.End = result.Started, result.Finished
				err := PromDumpMetrics(&store, opts.SetDefaults())
				if err != nil {
					glog.Errorf("failed to copy metrics from %s to %s: %v", prom, vm, err)
				}
			}

			if len(report) > 0 {
				glog.Infof("report to %s", report)
				err := result.ReportTo(report)
				if err != nil {
					glog.Errorf("failed to report to %s: %v", report, err)
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringSliceVarP(&tags, "tag", "t", []string{}, "benchmark tags")
	cmd.Flags().StringVar(&name, "name", "", "benchmark name")
	cmd.Flags().StringVar(&report, "report-to", "", "endpoint to report to")
	cmd.Flags().StringVar(&prom, "prom", "", "prometheus endpoint")
	cmd.Flags().StringVar(&vm, "vm", "", "victoria-metrics endpoint")
	cmd.Flags().Int64Var(&opts.Step, "step", 15, "resolution step")
	cmd.Flags().IntVar(&opts.MaxSteps, "steps", 720, "max steps per query")
	cmd.Flags().IntVar(&opts.MaxSamples, "samples", 50000000, "max samples per query")
	cmd.Flags().IntVar(&opts.Concurrency, "threads", runtime.NumCPU()*2, "number of worker threads")
	cmd.Flags().StringToStringVar(&opts.Selector, "selector", map[string]string{}, "series selector")
	cmd.Flags().StringToStringVar(&opts.Headers, "headers", map[string]string{}, "additional http headers for prometheus requests")
	cmd.Flags().StringToStringVar(&store.Headers, "vm.headers", map[string]string{}, "additional http headers for victoria-metrics requests")
	cmd.Flags().StringToStringVar(&store.Labels, "labels", map[string]string{}, "extra labels")
	cmd.Flags().IntVar(&store.Batch, "batch", 50, "import batch size")
	cmd.Flags().Uint64Var(&store.Rebase, "rebase", 0, "rebase time of metrics (in epoch second)")
	cmd.Flags().StringVar(&filter, "filter", "", "metrics filter regexp")

	cmd.Flags().MarkHidden("vm.headers")
	return cmd
}

type serverOptions struct {
	dsn string
	db  *sql.DB
}

func newBenchServerCmd() *cobra.Command {
	opts := new(serverOptions)
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Benchmark server commands",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.db, err = sql.Open("mysql", opts.dsn)
			if err != nil {
				return err
			}
			return opts.db.Ping()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.db.Close()
		},
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	cmd.PersistentFlags().StringVar(&opts.dsn, "dsn", "", "store dsn")
	cmd.AddCommand(newBenchServerInitCmd(opts))
	cmd.AddCommand(newBenchServerStartCmd(opts))
	return cmd
}

func newBenchServerInitCmd(opts *serverOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init a benchmark server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := BenchServerInitDB(opts.db)
			if err == nil {
				fmt.Fprintln(os.Stdout, "done")
			}
			return err
		},
	}
	return cmd
}

func newBenchServerStartCmd(opts *serverOptions) *cobra.Command {
	var (
		addr string
	)
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start a benchmark server",
		RunE: func(cmd *cobra.Command, args []string) error {
			glog.Infof("listening on %s", addr)
			return http.ListenAndServe(addr, BenchServerHandler(opts.db))
		},
	}
	cmd.Flags().StringVar(&addr, "address", ":8080", "address to listen on")
	return cmd
}
