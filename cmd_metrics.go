package main

import (
	"compress/flate"
	"regexp"
	"runtime"

	"github.com/spf13/cobra"
)

func newMetricsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Metrics utilities",
		RunE:  func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	cmd.AddCommand(newMetricsCopy())
	cmd.AddCommand(newMetricsDump())
	cmd.AddCommand(newMetricsLoad())
	return cmd
}

func newMetricsCopy() *cobra.Command {
	var (
		opts   PromDumpOptions
		store  VictoriaMetricsStore
		filter string
		silent bool
	)
	cmd := &cobra.Command{
		Use:   "copy <prom-endpoint> <vm-endpoint>",
		Short: "copy metrics data from prometheus to victoria-metrics",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Endpoint = args[0]
			store.Endpoint = args[1]
			if len(filter) > 0 {
				re, err := regexp.Compile(filter)
				if err != nil {
					return err
				}
				opts.Filter = re
			}
			if !silent {
				opts.Progress = new(ConsoleProgressListener)
			}
			return PromDumpMetrics(&store, opts.SetDefaults())
		},
	}
	cmd.Flags().Int64Var(&opts.Start, "start", 0, "start timestamp")
	cmd.Flags().Int64Var(&opts.End, "end", 0, "end timestamp")
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
	cmd.Flags().BoolVar(&silent, "silent", false, "silent mode")

	cmd.Flags().MarkHidden("vm.headers")
	return cmd
}

func newMetricsDump() *cobra.Command {
	var (
		opts   PromDumpOptions
		filter string
		dir    string
		gz     int
		rebase uint64
		silent bool
	)
	cmd := &cobra.Command{
		Use:   "dump <endpoint>",
		Short: "Dump metrics data from prometheus",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Endpoint = args[0]
			if len(filter) > 0 {
				re, err := regexp.Compile(filter)
				if err != nil {
					return err
				}
				opts.Filter = re
			}
			store, err := NewFileMetricsStore(dir, gz, rebase)
			if err != nil {
				return err
			}
			if !silent {
				opts.Progress = new(ConsoleProgressListener)
			}
			return PromDumpMetrics(store, opts.SetDefaults())
		},
	}
	cmd.Flags().Int64Var(&opts.Start, "start", 0, "start timestamp")
	cmd.Flags().Int64Var(&opts.End, "end", 0, "end timestamp")
	cmd.Flags().Int64Var(&opts.Step, "step", 15, "resolution step")
	cmd.Flags().IntVar(&opts.MaxSteps, "steps", 720, "max steps per query")
	cmd.Flags().IntVar(&opts.MaxSamples, "samples", 50000000, "max samples per query")
	cmd.Flags().IntVar(&opts.Concurrency, "threads", runtime.NumCPU()*2, "number of worker threads")
	cmd.Flags().StringToStringVar(&opts.Selector, "selector", map[string]string{}, "series selector")
	cmd.Flags().StringToStringVar(&opts.Headers, "headers", map[string]string{}, "additional http headers for prometheus requests")
	cmd.Flags().StringVar(&filter, "filter", "", "metrics filter regexp")
	cmd.Flags().StringVar(&dir, "dir", "/tmp/metrics.d", "output directory")
	cmd.Flags().IntVar(&gz, "gzip", flate.DefaultCompression, "gzip compression level (-1:DefaultCompression 0:NoCompression 1:BestSpeed 9:BestCompression)")
	cmd.Flags().Uint64Var(&rebase, "rebase", 0, "rebase time of metrics (in epoch second)")
	cmd.Flags().BoolVar(&silent, "silent", false, "silent mode")

	// keeping data as it is recommanded
	cmd.Flags().MarkHidden("rebase")
	return cmd
}

func newMetricsLoad() *cobra.Command {
	var (
		dir     string
		gz      int
		threads int
		rebase  uint64
		headers map[string]string
		labels  map[string]string
		silent  bool
	)
	cmd := &cobra.Command{
		Use:   "load <endpoint>",
		Short: "Load metrics data to victoria-metrics",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := NewFileMetricsStore(dir, gz, rebase)
			if err != nil {
				return err
			}
			if !silent {
				store.ShowProgress(true)
			}
			return store.UploadToVM(args[0], headers, labels, threads)
		},
	}
	cmd.Flags().StringVar(&dir, "dir", "/tmp/metrics.d", "input directory")
	cmd.Flags().IntVar(&threads, "threads", runtime.NumCPU()*2, "number of worker threads")
	cmd.Flags().IntVar(&gz, "gzip", flate.DefaultCompression, "gzip compression level (-1:DefaultCompression 0:NoCompression 1:BestSpeed 9:BestCompression)")
	cmd.Flags().Uint64Var(&rebase, "rebase", 0, "rebase time of metrics (in epoch second)")
	cmd.Flags().StringToStringVar(&headers, "vm.headers", map[string]string{}, "additional http headers for victoria-metrics requests")
	cmd.Flags().StringToStringVar(&labels, "labels", map[string]string{}, "extra labels")
	cmd.Flags().BoolVar(&silent, "silent", false, "silent mode")
	cmd.Flags().MarkHidden("vm.headers")
	return cmd
}
