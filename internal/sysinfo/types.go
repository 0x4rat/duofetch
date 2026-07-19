// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package sysinfo

// SummaryInfo holds instantaneous data for the compact view.
type SummaryInfo struct {
	OS       string
	Kernel   string
	Hostname string
	Uptime   string
	Shell    string
	CPU      string
	CPUCores int
	RAMUsed  uint64
	RAMTotal uint64
	DiskUsed uint64
	DiskTotal uint64
	NetIface string
	NetIP    string
	DistroID string
}

// DetailedInfo holds all data for the full breakdown view.
type DetailedInfo struct {
	// Host
	Hostname  string
	OS        string
	Kernel    string
	Arch      string
	BootTime  string
	Uptime    string
	Processes int

	// CPU
	CPUModels []string
	PhysCores int
	LogiCores int
	CoreUsage []float64  // per-core %, sampled
	LoadAvg   [3]float64 // 1/5/15 min; negative if unavailable

	// Memory
	RAMUsed   uint64
	RAMFree   uint64
	RAMTotal  uint64
	RAMCached uint64
	SwapUsed  uint64
	SwapTotal uint64

	// Disks
	Disks []DiskInfo

	// Network
	NetIfaces []NetInfo
}

// DiskInfo holds per-partition data.
type DiskInfo struct {
	Device     string
	Mountpoint string
	FSType     string
	Used       uint64
	Total      uint64
	Percent    float64
	ReadBps    float64 // bytes/sec; negative if unavailable
	WriteBps   float64
}

// NetInfo holds per-interface data.
type NetInfo struct {
	Name      string
	IPv4      string
	IPv6      string
	MAC       string
	IsUp      bool
	BytesSent uint64
	BytesRecv uint64
	SendBps   float64 // bytes/sec current rate; negative if unavailable
	RecvBps   float64
}
