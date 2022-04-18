package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"

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
