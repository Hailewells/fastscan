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
```

**!Disclamer**
## ⚠️ Important Disclaimer: Next Upgrades
# 🔥 8. Next Upgrades (Tell me which you want)

We can level this up quickly:

- ✅ CIDR / subnet scanning
- ✅ banner grabbing
- ✅ HTTP detection (like httpx)
- ✅ output to JSON
- ✅ progress bar
- ✅ SYN scan (advanced, raw sockets)
- ✅ integrate with tools like nuclei-style workflows
