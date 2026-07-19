// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package sysinfo

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"os"

	"github.com/shirou/gopsutil/v3/host"
	psnet "github.com/shirou/gopsutil/v3/net"
)

// FormatUptime converts uptime seconds to a human-readable string.
func FormatUptime(secs uint64) string {
	d := secs / 86400
	h := (secs % 86400) / 3600
	m := (secs % 3600) / 60

	var parts []string
	if d == 1 {
		parts = append(parts, "1 day")
	} else if d > 1 {
		parts = append(parts, fmt.Sprintf("%d days", d))
	}
	if h == 1 {
		parts = append(parts, "1 hour")
	} else if h > 1 {
		parts = append(parts, fmt.Sprintf("%d hours", h))
	}
	if m == 1 {
		parts = append(parts, "1 min")
	} else if m > 1 || len(parts) == 0 {
		parts = append(parts, fmt.Sprintf("%d mins", m))
	}
	return strings.Join(parts, ", ")
}

// FormatBytes converts a byte count to a human-readable string.
func FormatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// FormatBps formats bytes-per-second into a readable rate string.
func FormatBps(bps float64) string {
	if bps < 0 {
		return "N/A"
	}
	return FormatBytes(uint64(bps)) + "/s"
}

// RootPath returns the root/system drive path for the current OS.
func RootPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\"
	}
	return "/"
}

// GetShell returns the current user's shell name.
func GetShell() string {
	if runtime.GOOS == "windows" {
		if os.Getenv("PSModulePath") != "" {
			return "powershell"
		}
		if cs := os.Getenv("ComSpec"); cs != "" {
			return strings.ToLower(filepath.Base(cs))
		}
		return "cmd"
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "sh"
	}
	return filepath.Base(shell)
}

// CleanCPUModel removes marketing noise from CPU model strings.
func CleanCPUModel(model string) string {
	r := strings.NewReplacer("(R)", "", "(TM)", "", "(tm)", "")
	model = r.Replace(model)
	for strings.Contains(model, "  ") {
		model = strings.ReplaceAll(model, "  ", " ")
	}
	return strings.TrimSpace(model)
}

// FormatOSName produces a display-friendly OS name from host info.
func FormatOSName(h *host.InfoStat) string {
	name := strings.Title(h.Platform)
	switch strings.ToLower(h.Platform) {
	case "ubuntu":
		name = "Ubuntu"
	case "debian":
		name = "Debian GNU/Linux"
	case "arch", "archlinux":
		name = "Arch Linux"
	case "fedora":
		name = "Fedora Linux"
	case "centos":
		name = "CentOS Linux"
	case "rhel", "redhat":
		name = "Red Hat Enterprise Linux"
	case "opensuse", "opensuse-leap", "opensuse-tumbleweed":
		name = "openSUSE"
	case "alpine":
		name = "Alpine Linux"
	case "windows":
		name = "Windows"
	case "darwin":
		name = "macOS"
	}
	if h.PlatformVersion != "" && h.PlatformVersion != "rolling" {
		name += " " + h.PlatformVersion
	}
	return name
}

// IsActiveIface returns true if the interface is up and not loopback.
func IsActiveIface(iface psnet.InterfaceStat) bool {
	hasUp, hasLoop := false, false
	for _, f := range iface.Flags {
		switch f {
		case "up":
			hasUp = true
		case "loopback":
			hasLoop = true
		}
	}
	return hasUp && !hasLoop
}

// IsIPv4 returns true if s looks like a dotted-decimal IPv4 address.
func IsIPv4(s string) bool {
	return strings.Count(s, ".") == 3 && !strings.Contains(s, ":")
}

// SplitAddr strips the CIDR prefix from an addr string like "192.168.1.1/24".
func SplitAddr(addr string) string {
	return strings.Split(addr, "/")[0]
}

// PseudoFS is the set of filesystem types to skip in disk listings.
var PseudoFS = map[string]bool{
	"tmpfs": true, "devtmpfs": true, "sysfs": true, "proc": true,
	"devpts": true, "cgroup": true, "cgroup2": true, "pstore": true,
	"bpf": true, "securityfs": true, "hugetlbfs": true, "mqueue": true,
	"debugfs": true, "tracefs": true, "fusectl": true, "configfs": true,
	"autofs": true, "squashfs": true, "overlay": true, "nsfs": true,
	"efivarfs": true, "ramfs": true,
}
