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
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	psnet "github.com/shirou/gopsutil/v3/net"
)

const sampleInterval = 500 * time.Millisecond

// CollectDetailed gathers all data for the full breakdown, taking one ~500ms
// sample to compute live CPU, disk I/O, and network throughput rates.
func CollectDetailed() *DetailedInfo {
	info := &DetailedInfo{}
	info.LoadAvg = [3]float64{-1, -1, -1}

	// ── pre-sample readings ──────────────────────────────────────────────────
	netBefore, netErr := psnet.IOCounters(true)
	diskBefore, diskErr := disk.IOCounters()
	t0 := time.Now()

	// cpu.Percent with interval > 0 sleeps internally — this is our one wait.
	coreUsage, cpuErr := cpu.Percent(sampleInterval, true)

	elapsed := time.Since(t0).Seconds()
	netAfter, _ := psnet.IOCounters(true)
	diskAfter, _ := disk.IOCounters()

	// ── host ─────────────────────────────────────────────────────────────────
	if h, err := host.Info(); err == nil {
		info.Hostname = h.Hostname
		info.OS = FormatOSName(h)
		info.Kernel = h.KernelVersion
		info.Arch = h.KernelArch
		info.Uptime = FormatUptime(h.Uptime)
		info.Processes = int(h.Procs)
		if h.BootTime > 0 {
			info.BootTime = time.Unix(int64(h.BootTime), 0).Format("2006-01-02 15:04:05")
		}
	}

	// ── CPU ──────────────────────────────────────────────────────────────────
	if cpuErr == nil {
		info.CoreUsage = coreUsage
	}
	if cpuInfos, err := cpu.Info(); err == nil {
		seen := map[string]bool{}
		for _, c := range cpuInfos {
			model := CleanCPUModel(c.ModelName)
			if !seen[model] {
				info.CPUModels = append(info.CPUModels, model)
				seen[model] = true
			}
		}
	}
	info.PhysCores, _ = cpu.Counts(false)
	info.LogiCores, _ = cpu.Counts(true)

	if la, err := load.Avg(); err == nil {
		info.LoadAvg = [3]float64{la.Load1, la.Load5, la.Load15}
	}

	// ── memory ───────────────────────────────────────────────────────────────
	if vm, err := mem.VirtualMemory(); err == nil {
		info.RAMUsed = vm.Used
		info.RAMFree = vm.Free
		info.RAMTotal = vm.Total
		info.RAMCached = vm.Cached
	}
	if sm, err := mem.SwapMemory(); err == nil {
		info.SwapUsed = sm.Used
		info.SwapTotal = sm.Total
	}

	// ── disks ─────────────────────────────────────────────────────────────────
	if parts, err := disk.Partitions(false); err == nil {
		for _, p := range parts {
			if PseudoFS[p.Fstype] {
				continue
			}
			di := DiskInfo{
				Device:     p.Device,
				Mountpoint: p.Mountpoint,
				FSType:     p.Fstype,
				ReadBps:    -1,
				WriteBps:   -1,
			}
			if du, err := disk.Usage(p.Mountpoint); err == nil {
				di.Used = du.Used
				di.Total = du.Total
				di.Percent = du.UsedPercent
			}
			// Compute I/O rate from the sampled delta.
			if diskErr == nil && elapsed > 0 {
				devName := diskDevName(p.Device)
				if b, ok := diskBefore[devName]; ok {
					if a, ok := diskAfter[devName]; ok {
						di.ReadBps = float64(a.ReadBytes-b.ReadBytes) / elapsed
						di.WriteBps = float64(a.WriteBytes-b.WriteBytes) / elapsed
					}
				}
			}
			info.Disks = append(info.Disks, di)
		}
	}

	// ── network ───────────────────────────────────────────────────────────────
	if ifaces, err := psnet.Interfaces(); err == nil {
		// Build lookup maps for IO counters.
		beforeMap := make(map[string]psnet.IOCountersStat)
		afterMap := make(map[string]psnet.IOCountersStat)
		if netErr == nil {
			for _, c := range netBefore {
				beforeMap[c.Name] = c
			}
		}
		for _, c := range netAfter {
			afterMap[c.Name] = c
		}

		for _, iface := range ifaces {
			ni := NetInfo{
				Name:    iface.Name,
				MAC:     iface.HardwareAddr,
				SendBps: -1,
				RecvBps: -1,
			}
			for _, f := range iface.Flags {
				if f == "up" {
					ni.IsUp = true
				}
			}
			for _, addr := range iface.Addrs {
				ip := SplitAddr(addr.Addr)
				if IsIPv4(ip) {
					ni.IPv4 = ip
				} else if ni.IPv6 == "" && ip != "" {
					ni.IPv6 = ip
				}
			}
			if a, ok := afterMap[iface.Name]; ok {
				ni.BytesSent = a.BytesSent
				ni.BytesRecv = a.BytesRecv
				if b, ok := beforeMap[iface.Name]; ok && elapsed > 0 {
					ni.SendBps = float64(a.BytesSent-b.BytesSent) / elapsed
					ni.RecvBps = float64(a.BytesRecv-b.BytesRecv) / elapsed
				}
			}
			info.NetIfaces = append(info.NetIfaces, ni)
		}
	}

	return info
}

// diskDevName extracts the base device name from a path like /dev/sda1 → sda1.
func diskDevName(dev string) string {
	for i := len(dev) - 1; i >= 0; i-- {
		if dev[i] == '/' || dev[i] == '\\' {
			return dev[i+1:]
		}
	}
	return dev
}

var _ = fmt.Sprintf // silence unused import warning
