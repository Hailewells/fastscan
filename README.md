# fastscan
network scanner tool


# FastScan (Go)

Simple fast TCP port scanner with rate limiting.

## Features
- Concurrent scanning
- Timeout control
- Rate limiting (prevents overwhelming servers)

## Usage

```bash
go run . -host scanme.nmap.org -start 20 -end 100
