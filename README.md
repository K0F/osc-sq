# osc-sq
Tiny osc sequencer.

[![Go](https://github.com/K0F/osc-sq/actions/workflows/go.yml/badge.svg)](https://github.com/K0F/osc-sq/actions/workflows/go.yml)


## build it
To build package, run simply:

```
git clone git@github.com:K0F/osc-sq.git
cd osc-sq
go mod tidy
go build
```

## run it

To run sequencer:

```
osc-sq -b 120 -p 10000
```

### Arguments
- `-b` for BPM
- `-p` for port on localhost
