package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

// ParseSignal parses a signal by name or number.
func ParseSignal(s string) (syscall.Signal, bool) {
	// Look up signal by name, e.g. SIGINT
	if signal := unix.SignalNum(s); signal != 0 {
		return signal, true
	}

	// Lookup by signal number
	if sigNum, err := strconv.Atoi(s); err == nil {
		signal := syscall.Signal(sigNum)
		if unix.SignalName(signal) != "" {
			return signal, true
		}
	}

	return 0, false
}

// ParseSignals parses a string of signal numbers or names, joined with comma.
func ParseSignals(s string) ([]syscall.Signal, error) {
	split := strings.Split(s, ",")
	signals := make([]syscall.Signal, 0, len(split))
	for _, s := range split {
		signal, ok := ParseSignal(s)
		if !ok {
			return nil, fmt.Errorf("cannot parse signal: %v", s)
		}
		signals = append(signals, signal)
	}
	return signals, nil
}

// ParseSignalsFromArgs parses a list of arguments into a list of signals.
func ParseSignalsFromArgs(args []string) ([]syscall.Signal, error) {
	sigMap := make(map[int]syscall.Signal)
	for _, arg := range args {
		parsed, err := ParseSignals(arg)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse arg: %v", arg)
		}
		for _, sig := range parsed {
			sigMap[int(sig)] = sig
		}
	}

	signals := make([]syscall.Signal, 0, len(sigMap))
	for _, sig := range sigMap {
		signals = append(signals, sig)
	}
	sort.Slice(signals, func(i, j int) bool {
		return signals[i] < signals[j]
	})

	return signals, nil
}
