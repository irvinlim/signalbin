package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	timeout          time.Duration
	secondSignalQuit bool
	exitCode         int
	signalExitCode   bool

	command = &cobra.Command{
		Use:   "signalbin SIGNALS",
		Short: "Start signalbin and trap on a list of signals.",
		Example: strings.Join([]string{
			"  # Trap on SIGINT and SIGTERM and sleep for 10s after signal",
			"  signalbin SIGINT SIGTERM",
			"",
			"  # Trap on SIGINT and sleep for 60s",
			"  signalbin -t 60s",
			"",
			"  # Interrupt graceful termination sleep on second signal onwards",
			"  signalbin -t 60s --second-signal-quit",
		}, "\n"),
		Long: "This utility traps on a list of signals to aid with testing of signal handling routines.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(args)
		},
	}
)

func init() {
	command.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second,
		"Time to sleep for graceful termination.")
	command.Flags().BoolVarP(&secondSignalQuit, "second-signal-quit", "q", false,
		"Whether to immediately quit graceful termination on the second signal.")
	command.Flags().IntVarP(&exitCode, "exit-code", "e", 0,
		"Specify an explicit exit code to use when exiting.")
	command.Flags().BoolVarP(&signalExitCode, "signal-exit-code", "s", false,
		"If true, will propagate the signal via the exit code using 128+signalnum.")
}

func main() {
	defer klog.Flush()

	if err := command.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(126)
	}
}
