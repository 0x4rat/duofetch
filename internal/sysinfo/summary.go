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
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	psnet "github.com/shirou/gopsutil/v3/net"
)

// CollectSummary gathers all data for the compact view with no sampling delay.
func CollectSummary() *SummaryInfo {
	info := &SummaryInfo{}

	info.DistroID = GetDistroID()
	info.Shell = GetShell()

	if h, err := host.Info(); err == nil {
		info.Hostname = h.Hostname
		info.OS = FormatOSName(h)
		info.Kernel = h.KernelVersion
		info.Uptime = FormatUptime(h.Uptime)
	}

	if cpuInfos, err := cpu.Info(); err == nil && len(cpuInfos) > 0 {
		info.CPU = CleanCPUModel(cpuInfos[0].ModelName)
	}
	if n, err := cpu.Counts(true); err == nil {
		info.CPUCores = n
	}

	if vm, err := mem.VirtualMemory(); err == nil {
		info.RAMUsed = vm.Used
		info.RAMTotal = vm.Total
	}

	if du, err := disk.Usage(RootPath()); err == nil {
		info.DiskUsed = du.Used
		info.DiskTotal = du.Total
	}

	if ifaces, err := psnet.Interfaces(); err == nil {
		for _, iface := range ifaces {
			if !IsActiveIface(iface) {
				continue
			}
			info.NetIface = iface.Name
			for _, addr := range iface.Addrs {
				ip := SplitAddr(addr.Addr)
				if IsIPv4(ip) {
					info.NetIP = ip
					break
				}
			}
			break
		}
	}

	return info
}
