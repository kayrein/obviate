# obviate

Pipe the results of `go test` or `encore test` through this tool to filter out anything but the test failures.

You must specify -json as one of the arguments to the test program, so that it outputs JSON that this tool can read
(you may send it whatever other arguments you like).

## Usage
`encore test -json ./... | obviate`

## Installation
`go install github.com/kayrein/obviate`

## Arguments
Pass `-w` to obviate to provide the output for an entire failing test, rather than individual subtests.

Example: `encore test -json ./... | obviate -w`

Pass `-v` to obviate to print the version.

## See also
This project primarily intended to be used for [Encore](https://encore.dev)'s quite noisy output.