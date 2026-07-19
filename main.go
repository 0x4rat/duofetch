// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0x4rat/duofetch/internal/config"
	"github.com/0x4rat/duofetch/internal/render"
	"github.com/0x4rat/duofetch/internal/sysinfo"
)

const version = "1.0.0"

func main() {
	detailed := flag.Bool("d", false, "show full detailed breakdown")
	flag.BoolVar(detailed, "detailed", false, "show full detailed breakdown")
	noColor := flag.Bool("no-color", false, "disable ANSI color output")
	noLogo := flag.Bool("no-logo", false, "hide ASCII logo")
	genConfig := flag.Bool("gen-config", false, "write default config file and exit")
	force := flag.Bool("force", false, "overwrite existing config (use with --gen-config)")
	versionFlag := flag.Bool("version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "duofetch %s — fast cross-platform system info\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  duofetch              compact summary (default)\n")
		fmt.Fprintf(os.Stderr, "  duofetch -d           full detailed breakdown (~1s)\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("duofetch %s\n", version)
		return
	}

	if *genConfig {
		config.Generate(*force)
		return
	}

	// Enable VT100 colors on Windows terminals; no-op on Linux.
	render.EnableVTProcessing()

	cfg := config.Load()

	// CLI flags override config file.
	if *noColor || cfg.Display.NoColor || !render.IsTerminal() {
		render.NoColorMode = true
		cfg.NoColor = true
	}
	if *noLogo {
		cfg.NoLogo = true
	}

	if *detailed {
		info := sysinfo.CollectDetailed()
		render.Detailed(info, &cfg)
	} else {
		info := sysinfo.CollectSummary()
		render.Summary(info, &cfg)
	}
}
