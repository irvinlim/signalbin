package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

// Run will run the command and return errors if any.
func Run(args []string) error {
	klog.Info("starting signalbin")

	// Parse signals to trap
	signals, err := ParseSignalsFromArgs(args)
	if err != nil {
		return errors.Wrapf(err, "cannot parse signals from args")
	}
	if len(signals) == 0 {
		return fmt.Errorf("no signals specified")
	}

	// Trap signals
	done := make(chan os.Signal, 2)
	for _, sig := range signals {
		signal.Notify(done, sig)
	}
	klog.Infof("trapping on the following signals: %v", signals)

	HandleSignals(done)
	return nil
}

// HandleSignals will handle signals as configured.
func HandleSignals(done <-chan os.Signal) {
	sig := <-done
	startTime := time.Now()
	klog.Infof("received initial signal: %v", sig)

	if timeout > 0 {
		klog.Infof("graceful termination starting, sleep for %v", timeout)
		t := time.NewTimer(timeout)

	gracefulTermination:
		for {
			select {
			case newSig := <-done:
				klog.Infof("received subsequent signal: %v", newSig)
				if secondSignalQuit {
					break gracefulTermination
				}
			case <-t.C:
				klog.Infof("graceful termination finished")
				break gracefulTermination
			}
		}
	}

	klog.Infof("signal handler finished after %v", time.Since(startTime))
	code := exitCode

	// Exit with signal exit code.
	if signalExitCode {
		if sig, ok := sig.(syscall.Signal); ok {
			klog.Infof("using signal for exit code: %v (sig = %v)", unix.SignalName(sig), int(sig))
			code = int(sig) + 128
		}
	}

	// Otherwise, use the exit code defined via flags.
	klog.Infof("signalbin exiting with exit code %v", code)
	os.Exit(code)
}
