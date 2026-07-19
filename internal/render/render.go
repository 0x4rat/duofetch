// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package render

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/0x4rat/duofetch/internal/config"
	"github.com/0x4rat/duofetch/internal/logos"
	"github.com/0x4rat/duofetch/internal/sysinfo"
)

const logoInfoGap = "   " // spacing between logo and info block

// Summary renders the compact neofetch-style side-by-side view.
func Summary(info *sysinfo.SummaryInfo, cfg *config.Config) {
	var logoLines []string
	logoColor := "cyan"

	if !cfg.Display.NoLogo && !cfg.NoLogo {
		var logo logos.Logo
		if cfg.Display.CustomLogo != "" {
			logo = loadCustomLogo(cfg.Display.CustomLogo)
		} else {
			logo = logos.Get(info.DistroID)
		}
		logoColor = logo.Color
		for _, l := range logo.Lines {
			logoLines = append(logoLines, C(logoColor)+l+R())
		}
	}

	accent := cfg.Display.AccentColor
	if accent == "" {
		accent = logoColor
	}

	infoLines := buildSummaryLines(info, cfg, accent)

	// Print blank line before
	fmt.Println()
	printSideBySide(logoLines, infoLines)
	fmt.Println()
}

// Detailed renders the full section-based breakdown view.
func Detailed(info *sysinfo.DetailedInfo, cfg *config.Config) {
	accent := cfg.Display.AccentColor
	if accent == "" {
		accent = "cyan"
	}

	fmt.Println()
	printDetailedSections(info, cfg, accent)
	fmt.Println()
}

// ── summary helpers ──────────────────────────────────────────────────────────

func buildSummaryLines(info *sysinfo.SummaryInfo, cfg *config.Config, accent string) []string {
	var lines []string
	addLine := func(label, value string) {
		if value == "" {
			value = "N/A"
		}
		// Fixed-width label column (9 visible chars) regardless of ANSI codes.
		lbl := BoldWrap(accent, label+":")
		lines = append(lines, fmt.Sprintf("%-*s %s", 9+len(lbl)-VisibleLen(lbl), lbl, value))
	}
	sc := cfg.Summary

	if sc.ShowOS {
		addLine("OS", info.OS)
	}
	if sc.ShowKernel {
		addLine("Kernel", info.Kernel)
	}
	if sc.ShowHostname {
		addLine("Host", info.Hostname)
	}
	if sc.ShowUptime {
		addLine("Uptime", info.Uptime)
	}
	if sc.ShowShell {
		addLine("Shell", info.Shell)
	}
	if sc.ShowCPU {
		cpu := info.CPU
		if info.CPUCores > 0 {
			cpu += fmt.Sprintf(" (%d)", info.CPUCores)
		}
		addLine("CPU", cpu)
	}
	if sc.ShowRAM {
		addLine("RAM", fmt.Sprintf("%s / %s",
			sysinfo.FormatBytes(info.RAMUsed), sysinfo.FormatBytes(info.RAMTotal)))
	}
	if sc.ShowDisks {
		addLine("Disk", fmt.Sprintf("%s / %s",
			sysinfo.FormatBytes(info.DiskUsed), sysinfo.FormatBytes(info.DiskTotal)))
	}
	if sc.ShowNetwork {
		net := info.NetIface
		if info.NetIP != "" {
			net += " @ " + info.NetIP
		}
		addLine("Network", net)
	}
	return lines
}

func printSideBySide(left, right []string) {
	// Find the maximum visible width of the left column.
	maxLeft := 0
	for _, l := range left {
		if w := VisibleLen(l); w > maxLeft {
			maxLeft = w
		}
	}

	height := len(left)
	if len(right) > height {
		height = len(right)
	}

	for i := 0; i < height; i++ {
		var l, r string
		if i < len(left) {
			l = left[i]
		}
		if i < len(right) {
			r = right[i]
		}
		fmt.Fprintf(os.Stdout, "%s%s%s%s\n",
			l,
			strings.Repeat(" ", maxLeft-VisibleLen(l)),
			logoInfoGap,
			r,
		)
	}
}

// ── detailed helpers ─────────────────────────────────────────────────────────

func sectionHeader(title, accent string) string {
	dashes := strings.Repeat("─", max(0, 44-len(title)))
	header := fmt.Sprintf("── %s %s", title, dashes)
	if NoColorMode {
		return header
	}
	return C(accent) + Bold + header + R()
}

func field(label, value string) string {
	if value == "" {
		value = "N/A"
	}
	return fmt.Sprintf("  %-14s %s", label+":", value)
}

func printDetailedSections(info *sysinfo.DetailedInfo, cfg *config.Config, accent string) {
	dc := cfg.Detailed

	// ── CPU ──────────────────────────────────────────────────────────────────
	if dc.ShowCPU {
		fmt.Println(sectionHeader("CPU", accent))
		if len(info.CPUModels) > 0 {
			for _, m := range info.CPUModels {
				fmt.Println(field("Model", m))
			}
		} else {
			fmt.Println(field("Model", "N/A"))
		}
		fmt.Println(field("Physical", fmt.Sprintf("%d cores", info.PhysCores)))
		fmt.Println(field("Logical", fmt.Sprintf("%d threads", info.LogiCores)))
		if len(info.CoreUsage) > 0 {
			for i, pct := range info.CoreUsage {
				bar := progressBar(pct, 20)
				fmt.Printf("  Core %-3d      %s %s%.1f%%%s\n",
					i, Wrap(accent, bar), C("white"), pct, R())
			}
		}
		if dc.ShowLoad && info.LoadAvg[0] >= 0 {
			fmt.Println(field("Load avg",
				fmt.Sprintf("%.2f  %.2f  %.2f  (1m 5m 15m)",
					info.LoadAvg[0], info.LoadAvg[1], info.LoadAvg[2])))
		}
		fmt.Println()
	}

	// ── Memory ───────────────────────────────────────────────────────────────
	if dc.ShowRAM {
		fmt.Println(sectionHeader("Memory", accent))
		ramPct := 0.0
		if info.RAMTotal > 0 {
			ramPct = float64(info.RAMUsed) / float64(info.RAMTotal) * 100
		}
		bar := progressBar(ramPct, 20)
		fmt.Printf("  %-14s %s / %s  %s  %.1f%%\n",
			"RAM:",
			sysinfo.FormatBytes(info.RAMUsed),
			sysinfo.FormatBytes(info.RAMTotal),
			Wrap(accent, bar),
			ramPct,
		)
		fmt.Println(field("Free", sysinfo.FormatBytes(info.RAMFree)))
		if info.RAMCached > 0 {
			fmt.Println(field("Cached", sysinfo.FormatBytes(info.RAMCached)))
		}
		if dc.ShowSwap {
			swapPct := 0.0
			if info.SwapTotal > 0 {
				swapPct = float64(info.SwapUsed) / float64(info.SwapTotal) * 100
				swapBar := progressBar(swapPct, 20)
				fmt.Printf("  %-14s %s / %s  %s  %.1f%%\n",
					"Swap:",
					sysinfo.FormatBytes(info.SwapUsed),
					sysinfo.FormatBytes(info.SwapTotal),
					Wrap(accent, swapBar),
					swapPct,
				)
			} else {
				fmt.Println(field("Swap", "none"))
			}
		}
		fmt.Println()
	}

	// ── Disks ─────────────────────────────────────────────────────────────────
	if dc.ShowDisks && len(info.Disks) > 0 {
		fmt.Println(sectionHeader("Disks", accent))
		for _, d := range info.Disks {
			bar := progressBar(d.Percent, 20)
			fmt.Printf("  %-20s %s  %s  %-6s  %s / %s  (%.1f%%)\n",
				d.Mountpoint,
				Wrap(accent, bar),
				Wrap("white", d.FSType),
				"",
				sysinfo.FormatBytes(d.Used),
				sysinfo.FormatBytes(d.Total),
				d.Percent,
			)
			fmt.Printf("    %-12s %s\n", "Device:", d.Device)
			if dc.ShowDiskIO {
				fmt.Printf("    %-12s %s   %-12s %s\n",
					"Read:",
					formatRate(d.ReadBps),
					"Write:",
					formatRate(d.WriteBps),
				)
			}
		}
		fmt.Println()
	}

	// ── Network ───────────────────────────────────────────────────────────────
	if dc.ShowNetwork && len(info.NetIfaces) > 0 {
		fmt.Println(sectionHeader("Network", accent))
		for _, n := range info.NetIfaces {
			status := Wrap("red", "DOWN")
			if n.IsUp {
				status = Wrap("green", "UP  ")
			}
			fmt.Printf("  %s %-12s  %s\n",
				status, n.Name, formatIface(n))
			if n.MAC != "" {
				fmt.Printf("    %-12s %s\n", "MAC:", n.MAC)
			}
			fmt.Printf("    %-12s %s   %-12s %s\n",
				"Total sent:", sysinfo.FormatBytes(n.BytesSent),
				"Total recv:", sysinfo.FormatBytes(n.BytesRecv),
			)
			if dc.ShowNetThroughput {
				fmt.Printf("    %-12s %s   %-12s %s\n",
					"↑ Upload:", formatRate(n.SendBps),
					"↓ Download:", formatRate(n.RecvBps),
				)
			}
		}
		fmt.Println()
	}

	// ── Host ─────────────────────────────────────────────────────────────────
	fmt.Println(sectionHeader("Host", accent))
	if dc.ShowHostname {
		fmt.Println(field("Hostname", info.Hostname))
	}
	if dc.ShowOS {
		fmt.Println(field("OS", info.OS))
	}
	if dc.ShowKernel {
		kinfo := info.Kernel
		if info.Arch != "" {
			kinfo += "  " + info.Arch
		}
		fmt.Println(field("Kernel", kinfo))
	}
	if dc.ShowUptime {
		fmt.Println(field("Uptime", info.Uptime))
	}
	if info.BootTime != "" {
		fmt.Println(field("Boot time", info.BootTime))
	}
	if dc.ShowProcesses && info.Processes > 0 {
		fmt.Println(field("Processes", fmt.Sprintf("%d running", info.Processes)))
	}
	fmt.Println()
}

// ── utilities ─────────────────────────────────────────────────────────────────

func progressBar(pct float64, width int) string {
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	filled := int(math.Round(pct / 100.0 * float64(width)))
	return "[" + strings.Repeat("█", filled) + strings.Repeat("░", width-filled) + "]"
}

func formatRate(bps float64) string {
	if bps < 0 {
		return "N/A"
	}
	return sysinfo.FormatBps(bps)
}

func formatIface(n sysinfo.NetInfo) string {
	parts := []string{}
	if n.IPv4 != "" {
		parts = append(parts, n.IPv4)
	}
	if n.IPv6 != "" {
		parts = append(parts, n.IPv6)
	}
	if len(parts) == 0 {
		return "no address"
	}
	return strings.Join(parts, "  ")
}

func loadCustomLogo(path string) logos.Logo {
	data, err := os.ReadFile(path)
	if err != nil {
		return logos.Fallback()
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	return logos.Logo{Lines: lines, Color: "white"}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
