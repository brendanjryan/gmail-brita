# gmail-brita

A Go implementation of [gmail-britta](https://github.com/antifuchs/gmail-britta), a Ruby DSL for generating Gmail filters.

## Overview

This is a Go port of the original Ruby gem [gmail-britta](https://github.com/antifuchs/gmail-britta) by [Andreas Fuchs](https://github.com/antifuchs). While the original uses a Ruby DSL, this version uses YAML configuration files to define Gmail filters.

## Features

- YAML-based configuration for Gmail filters
- Support for multiple email addresses
- Complex filter conditions and actions
- XML output compatible with Gmail's filter import format
- Easy to use command-line interface

## Installation

```bash
go install github.com/brendanryan/gmail-brita@latest
```

## Usage

1. Create a YAML configuration file defining your filters (see examples directory)
2. Run the command:

```bash
gmail-brita -config filters.yaml -out gmail-filters.xml
```

3. Import the generated XML file into Gmail's filter settings

## Configuration

See the `examples` directory for sample filter configurations. The YAML format supports:

- Multiple email addresses
- Filter conditions (from, to, subject, has, etc.)
- Filter actions (archive, mark read, star, apply label, etc.)
- Complex combinations of conditions and actions

## Development

Requirements:

- Go 1.19 or later

```bash
# Clone the repository
git clone https://github.com/brendanryan/gmail-brita.git
cd gmail-brita

# Install dependencies
go mod download

# Run tests
make test

# Build
make build

# Run example
make run-example
```

## License

This project is licensed under the same terms as the original gmail-britta:

Copyright (c) 2012 Andreas Fuchs <asf@boinkor.net>, released under the MIT license.

## Attribution

This project is a Go port of [gmail-britta](https://github.com/antifuchs/gmail-britta) by Andreas Fuchs. The original project provided the inspiration and core concepts for this implementation.
