# signalbin

![license](https://img.shields.io/github/license/irvinlim/signalbin)
[![docker pulls](https://img.shields.io/docker/pulls/irvinlim/signalbin.svg)](https://hub.docker.com/r/irvinlim/signalbin)
[![image size](https://img.shields.io/docker/image-size/irvinlim/signalbin?sort=date)](https://hub.docker.com/r/irvinlim/signalbin/tags)

Like httpbin, but for UNIX signals.

`signalbin` is a small utility written in Go to test interactions with signal handlers.

## Docker

```shell
docker run --rm -it irvinlim/signalbin SIGINT,SIGTERM -sq -t=30s
```

## Usage

```shell
This utility traps on a list of signals to aid with testing of signal handling routines.

Usage:
  signalbin SIGNALS [flags]

Examples:
  # Trap on SIGINT and SIGTERM and sleep for 10s after signal
  signalbin SIGINT SIGTERM

  # Trap on SIGINT and sleep for 60s
  signalbin -t 60s

  # Interrupt graceful termination sleep on second signal onwards
  signalbin -t 60s --second-signal-quit

Flags:
  -e, --exit-code int        Specify an explicit exit code to use when exiting.
  -h, --help                 help for signalbin
  -q, --second-signal-quit   Whether to immediately quit graceful termination on the second signal.
  -s, --signal-exit-code     If true, will propagate the signal via the exit code using 128+signalnum.
  -t, --timeout duration     Time to sleep for graceful termination. (default 10s)
```

## License

MIT
