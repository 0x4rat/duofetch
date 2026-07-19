# duofetch

A fast, cross-platform (Linux + Windows) system information tool that shows a quick summary by default or a full detailed breakdown with `-d`.

---

## Features

- **Compact summary** (default) — logo + info side by side; returns instantly with no sampling
- **Full detailed view** (`-d`) — per-core CPU usage, all disks with live I/O throughput, all network interfaces with live upload/download rates, swap, load averages, and process count
- **Single static binary** — no Python, no libc worries, no runtime dependencies; `scp` it onto any server and run it immediately
- **Cross-platform** — compiles and runs on Linux and Windows from one codebase
- **Distro-aware ASCII logos** — Ubuntu, Debian, Arch, Fedora, CentOS, RHEL, openSUSE, Alpine, Windows, and a generic Linux (Tux) fallback
- **ANSI color output** — auto-enabled on Linux; uses `ENABLE_VIRTUAL_TERMINAL_PROCESSING` on Windows Terminal and modern Windows consoles
- **Configurable** — TOML config file to toggle individual fields independently for each view, pick accent color, disable or replace the logo
- **`--gen-config`** — writes a fully-commented default config file so you can see and edit every option

---

## Screenshots

![summary](docs/summary.png)
![detailed](docs/detailed.png)

---

## Installation

### Using `go install`

```sh
go install github.com/0x4rat/duofetch@latest
```

### Building from source

```sh
git clone https://github.com/0x4rat/duofetch
cd duofetch
CGO_ENABLED=0 go build -ldflags="-s -w" -o duofetch .
```

The result is a single fully-static binary you can copy to any Linux or Windows server and run with no dependencies whatsoever.

### Cross-compiling

**Linux (amd64):**

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o duofetch-linux-amd64 .
```

**Linux (arm64):**

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o duofetch-linux-arm64 .
```

**Windows (amd64):**

```sh
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o duofetch-windows-amd64.exe .
```

Or use the Makefile:

```sh
make all          # build for Linux amd64/arm64 and Windows amd64 → dist/
make build        # build for the current platform
```

---

## Usage

| Command                               | Description                                 |
| ------------------------------------- | ------------------------------------------- |
| `duofetch`                            | Compact summary (instant)                   |
| `duofetch -d` / `duofetch --detailed` | Full detailed breakdown (~1 second)         |
| `duofetch --no-color`                 | Disable ANSI color output                   |
| `duofetch --no-logo`                  | Hide the ASCII logo                         |
| `duofetch --gen-config`               | Write default config file                   |
| `duofetch --gen-config --force`       | Overwrite existing config without prompting |
| `duofetch --version`                  | Print version                               |
| `duofetch --help`                     | Print usage                                 |

### Examples

```sh
# Quick summary — what did I just get?
duofetch

# Full breakdown with live CPU, disk, and network rates
duofetch -d

# Summary without color (for logs or scripts)
duofetch --no-color

# Generate a config file to customize the output
duofetch --gen-config
```

---

## What each view shows

### Summary (default, instant)

Prints the distro ASCII logo on the left and a quick info block on the right:

| Field   | Example                |
| ------- | ---------------------- |
| OS      | Ubuntu 22.04           |
| Kernel  | 5.15.0-76-generic      |
| Host    | myserver               |
| Uptime  | 2 days, 5 hours        |
| Shell   | bash                   |
| CPU     | AMD Ryzen 7 5800X (16) |
| RAM     | 4.2 GB / 16.0 GB       |
| Disk    | 45.2 GB / 200.0 GB     |
| Network | eth0 @ 192.168.1.100   |

All values are read instantaneously — no sampling, no delay.

### Detailed (`-d`, ~1 second)

Takes one ~500 ms sample to compute live rates, then prints full sections:

- **CPU** — model, physical/logical core counts, per-core usage % with a bar graph, load averages (Linux only)
- **Memory** — RAM used/free/total/cached with usage bar, swap used/total
- **Disks** — every real partition: mountpoint, filesystem type, used/total with bar, device name, live read/write throughput
- **Network** — every interface: IPv4/IPv6 addresses, MAC, up/down status, total bytes sent/received, live upload/download rate
- **Host** — hostname, OS, kernel + architecture, uptime, boot time, process count

The ~1 second wait is the cost of accurate live rates; the default summary has no wait.

---

## Customization

### Config file location

| Platform | Path                             |
| -------- | -------------------------------- |
| Linux    | `~/.config/duofetch/config.toml` |
| Windows  | `%AppData%\duofetch\config.toml` |

### Generate the default config

```sh
duofetch --gen-config
```

This writes a fully-commented config file. If the file already exists you will be prompted before it is overwritten. Pass `--force` to skip the prompt.

### Full example config.toml

```toml
# duofetch configuration

[display]
# Disable all ANSI color output.
no_color = false

# Hide the ASCII logo in the summary view.
no_logo = false

# Accent color for labels and section headers.
# Options: red, green, yellow, blue, magenta, cyan, white, orange
accent_color = "cyan"

# Path to a custom ASCII art file to use instead of the built-in logo.
# custom_logo = "/home/user/.config/duofetch/mylogo.txt"

# ── Summary view (default, instant) ──────────────────────────────────────────
[summary]
show_os       = true   # OS name and version
show_kernel   = true   # kernel version
show_hostname = true   # machine hostname
show_uptime   = true   # system uptime
show_shell    = true   # current shell
show_cpu      = true   # CPU model and core count
show_ram      = true   # RAM used / total
show_disks    = true   # root disk used / total
show_network  = true   # primary interface and IP

# ── Detailed view (-d) ────────────────────────────────────────────────────────
[detailed]
show_os              = true   # OS name and version
show_kernel          = true   # kernel version and architecture
show_hostname        = true   # machine hostname
show_uptime          = true   # system uptime
show_cpu             = true   # CPU model, core counts, per-core usage
show_ram             = true   # RAM used/free/total/cached
show_swap            = true   # swap used / total
show_disks           = true   # all real partitions with usage
show_disk_io         = true   # live disk read/write throughput
show_network         = true   # all network interfaces with addresses
show_net_throughput  = true   # live per-interface upload/download rate
show_load            = true   # load averages (Linux only)
show_processes       = true   # number of running processes
```

### CLI flags override the config

`--no-color` and `--no-logo` always take effect regardless of what the config says.

---

## Supported platforms and logos

| Platform | Logos available                                                                         |
| -------- | --------------------------------------------------------------------------------------- |
| Linux    | Ubuntu, Debian, Arch Linux, Fedora, CentOS, RHEL, openSUSE, Alpine, generic Linux (Tux) |
| Windows  | Windows                                                                                 |

Distro detection on Linux reads the `ID=` field from `/etc/os-release`. On Windows the Windows logo is always used.

### Adding a new logo

1. Open `internal/logos/logos.go`.
2. Add a new entry to the `All` map. The key must match the `ID=` value in `/etc/os-release` for that distro (lowercase).

```go
"myDistro": {
    Color: "green",  // any color name supported by the render package
    Lines: []string{
        "  _____  ",
        " / My  \\ ",
        "| Distro|",
        " \\_____/ ",
    },
},
```

3. Rebuild. The logo will be picked up automatically on any machine running that distro.

---

## License

duofetch is free software licensed under the **GNU General Public License v3.0**.

This program is distributed in the hope that it will be useful, but **WITHOUT ANY WARRANTY** — without even the implied warranty of **MERCHANTABILITY** or **FITNESS FOR A PARTICULAR PURPOSE**. The author accepts no liability for any damage arising from its use. See the [LICENSE](LICENSE) file for the full license text.

---

## Contributing

Bug reports, logo contributions, and pull requests are welcome. Please open an issue or PR on GitHub.

---

## Contact

Questions or feedback? Reach out at **0x4rat@protonmail.com**.
