// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

// Package logos provides ASCII art logos indexed by distro ID.
// To add a new logo: add an entry to All using the ID= value from /etc/os-release.
package logos

// Logo holds the ASCII art lines and the primary accent color name.
type Logo struct {
	Lines []string
	Color string // matches render color names: red/green/yellow/blue/magenta/cyan/white/orange
}

// All maps distro IDs (from /etc/os-release ID=) to their logos.
var All = map[string]Logo{
	"ubuntu": {
		Color: "orange",
		Lines: []string{
			"            .-/+oossssoo+/-.            ",
			"        `:+ssssssssssssssssss+:`        ",
			"      -+ssssssssssssssssssyyssss+-      ",
			"    .ossssssssssssssssss dMMMNysssso.   ",
			"   /ssssssssssshdmmNNmmyNMMMMhssssss/   ",
			"  +ssssssssshmydMMMMMMMNddddyssssssss+  ",
			" /sssssssshNMMMyhhyyyyhmNMMMNhssssssss\\ ",
			".ssssssssdMMMNhsssssssssshNMMMdssssssss.",
			"+sssshhhyNMMNyssssssssssssyNMMMysssssss+",
			"ossyNMMMNyMMhsssssssssssssshmmmhssssssso",
			"ossyNMMMNyMMhsssssssssssssshmmmhssssssso",
			"+sssshhhyNMMNyssssssssssssyNMMMysssssss+",
			".ssssssssdMMMNhsssssssssshNMMMdssssssss.",
			" \\sssssssshNMMMyhhyyyyhmNMMMNhssssssss/ ",
			"  +ssssssssshmydMMMMMMMNddddyssssssss+  ",
			"   \\ssssssssssshdmmNNmmyNMMMMhssssss/   ",
			"    .osssssssssssssssssssdMMMNysssso.   ",
			"      -+sssssssssssssssssyyyssss+-      ",
			"        `:+ssssssssssssssssss+:`        ",
			"            .-/+oossssoo+/-.            ",
		},
	},
	"arch": {
		Color: "cyan",
		Lines: []string{
			"                  -`          ",
			"                 .o+`         ",
			"                `ooo/         ",
			"               `+oooo:        ",
			"              `+oooooo:       ",
			"              -+oooooo+:      ",
			"            `/:-:++oooo+:     ",
			"           `/++++/+++++++:    ",
			"          `/++++++++++++++:   ",
			"         `/+++ooooooooooooo/` ",
			"        ./ooosssso++osssssso+`",
			"       .oossssso-    /ossssss+`",
			"      -osssssso.      :ssssssso.",
			"     :osssssss/        osssso+++.",
			"    /ossssssss/        +ssssooo/-",
			"  `/ossssso+/:-        -:/+osssso+-",
			" `+sso+:-`                 `.-/+oso:",
			"`++:.                           `-/+/",
			".`                                 `/",
		},
	},
	"archlinux": {
		Color: "cyan",
		Lines: []string{
			"                  -`          ",
			"                 .o+`         ",
			"                `ooo/         ",
			"               `+oooo:        ",
			"              `+oooooo:       ",
			"              -+oooooo+:      ",
			"            `/:-:++oooo+:     ",
			"           `/++++/+++++++:    ",
			"          `/++++++++++++++:   ",
			"         `/+++ooooooooooooo/` ",
			"        ./ooosssso++osssssso+`",
			"       .oossssso-    /ossssss+`",
			"      -osssssso.      :ssssssso.",
			"     :osssssss/        osssso+++.",
			"    /ossssssss/        +ssssooo/-",
			"  `/ossssso+/:-        -:/+osssso+-",
			" `+sso+:-`                 `.-/+oso:",
			"`++:.                           `-/+/",
			".`                                 `/",
		},
	},
	"debian": {
		Color: "red",
		Lines: []string{
			"       _,met$$$$$gg.       ",
			"    ,g$$$$$$$$$$$$$$$P.    ",
			"  ,g$$P\"     \"\"\"Y$$.\"`.    ",
			" ,$$P'              `$$$.  ",
			"',$$P       ,ggs.     `$$b:",
			"`d$$'     ,$P\"'   .    $$$  ",
			" $$P      d$'     ,    $$P  ",
			" $$:      $$.   -    ,d$$'  ",
			" $$;      Y$b._   _,d$P'   ",
			" Y$$.    `.`\"Y$$$$P\"'       ",
			" `$$b      \"-.__           ",
			"  `Y$$b        \"-._        ",
			"   `Y$$.           `.      ",
			"     `$$b.           `-.   ",
			"       `Y$$b.         `\\  ",
			"         `\"Y$b._____...`   ",
			"              `\"\"\"\"\"\"\"``  ",
		},
	},
	"fedora": {
		Color: "blue",
		Lines: []string{
			"          /:-------------:\\         ",
			"       :-------------------::       ",
			"     :-----------/shhOHbmp---:\\     ",
			"   /-----------omMMMNNNMMD  ---:    ",
			"  :-----------sMMMMNMNMP.    ---:   ",
			" :-----------:MMMdP-------    ---\\  ",
			",------------:MMMd--------    ---:  ",
			":------------:MMMd-------    .---:  ",
			":----    oNMMMMMMMMMNho     .----:  ",
			":--  .+shhhMMMMMMMMMNNMonNM   ---:  ",
			": . /     /MMMMMMMNNE .-   ---:     ",
			"\\          -MMMMNMNDI         .\\   ",
			" \\.    .----+MMMMMMhia          ..\\ ",
			"  +/     ./+sMMMNNNMNH            ..+",
			"   +:        :MMMMMMMMN.           .-:",
			"    :         `+MMMMMMM`.         .:-",
			"     \\          `+hhhhhh`          \\:",
		},
	},
	"centos": {
		Color: "yellow",
		Lines: []string{
			"                 ..          ",
			"               .PLTJ.        ",
			"             <><><><>        ",
			"    KKSSV' 4KKK LJ KKKL.'VSSKK",
			"    KKV' 4KKKKK LJ KKKKAL 'VKK",
			"    V' ' 'VKKKK LJ KKKKV' ' 'V",
			"    .4MA.' 'VKK LJ KKV' '.4MAF.",
			"   KKKKKAL.'' LJ ''-.LKKKKKK  ",
			"   KKKKK.4MA LJ MAF.4KKKKK   ",
			"   KKKKK' V' LJ V' 'KKKKK    ",
			"   'VKKKK. LJ .KKKKV'        ",
			"     'VKK. LJ .KKV'          ",
			"        V' LJ 'V             ",
			"           LJ               ",
			"           LJ               ",
			"           ' '              ",
		},
	},
	"rhel": {
		Color: "red",
		Lines: []string{
			"           .MMM..:MMMMMMM       ",
			"          MMMMMMMMMMMMMMMMMM    ",
			"          MMMMMMMMMMMMMMMMMMMM. ",
			" -MMMM-.  MMMMMMMMMMMMMMMMMMMMMM",
			"MMMMMMMM- MMMMMMMMMMMMMMMMMMMMM ",
			"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM ",
			"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM ",
			"MMMMMMMMMMMMMMMMMMMMMMMMMMMM MM ",
			"MMMMMMMMMMMMMMMMMMMMMMMMMMMM MM ",
			"MMMMMMMMM. MMMMMMMMMMMMMMMM  MM ",
			" MMMMMM.   MMMMMMMMMMMMMMM.  MM ",
			"           MMMMMMMM: MMMM       ",
			"            .MMMM.    MM        ",
		},
	},
	"linux": {
		Color: "white",
		Lines: []string{
			"        _nnnn_        ",
			"       dGGGGMMb       ",
			"      @p~qp~~qMb      ",
			"      M|@||@) M|      ",
			"      @,----.JM|      ",
			"     JS^\\__/  qKL    ",
			"    dZP        qKRb   ",
			"   dZP          qKKb  ",
			"  fZP            SMMb ",
			"  HZM            MMMM ",
			"  FqM            MMMM ",
			"__| \".        |\\dS\"qML",
			"|    `.       | `' \\Zq",
			"_)      \\.___..|     .'",
			"\\____   )MMMMMM|   .'  ",
			"     `-'       `--'    ",
		},
	},
	"windows": {
		Color: "cyan",
		Lines: []string{
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
			"                                  ",
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
			"################  ################",
		},
	},
	"darwin": {
		Color: "white",
		Lines: []string{
			"                 ,MMMM.        ",
			"               .MMMMMM         ",
			"               MMMMM,          ",
			"     .;MMMMM:' MMMMMMMMMM;.    ",
			"   MMMMMMMMMMMMNWMMMMMMMMMMM:  ",
			" .MMMMMMMMMMMMMMMMMMMMMMMMWM.  ",
			" MMMMMMMMMMMMMMMMMMMMMMMMM.    ",
			";MMMMMMMMMMMMMMMMMMMMMMMM:     ",
			":MMMMMMMMMMMMMMMMMMMMMMMM:     ",
			".MMMMMMMMMMMMMMMMMMMMMMMM.     ",
			" MMMMMMMMMMMMMMMMMMMMMMMMM.    ",
			"  MMMMMMMMMMMMMMMMMMMMMMMMMM.  ",
			"   \\MMMMMMMMMMMMMMMMMMMMMMMM:  ",
			"     ;MMMMMMMMMMMMMMMMMMMM:.   ",
			"       `\"\"``\"\"``\"\"``\"\"``       ",
		},
	},
	"alpine": {
		Color: "blue",
		Lines: []string{
			"       /\\ /\\      ",
			"      /  V  \\     ",
			"     / /   \\ \\   ",
			"    / /     \\ \\  ",
			"   / /       \\ \\ ",
			"  / /  Alpine  \\ \\",
			" /_/____________\\_\\",
		},
	},
	"opensuse": {
		Color: "green",
		Lines: []string{
			"          .;ldkO0000Okdl;.         ",
			"      .;d00xl:^''''''^:ok00d;.     ",
			"    .d00l'                'o00d.   ",
			"  .d0Kd'  Okxol:;,.          :O0d.",
			" .OKKKK0kOKKKKKKKKKKOxo:,      lKO",
			",0KKKKKKKKKKKKKKKK0P^,,,^dx:    ;00",
			".OKKKKKKKKKKKKKKKk.          ;xOO' ",
			":KKKKKKKKKKKKKKKKK: .oOPPb.   'ODk.",
			"'KKKKKKKKKKKKKKKKK' .OKKKK0.  OKKKb",
			" d0KKKKKKKKKKKKKK. lKKKKKd  .OKKKK.",
			"  l0KKK000KKKKKx  d0KKKKx   xKKKK0 ",
			"   `^^^^^okKK0' .oKKKKx   dKKKKKK` ",
			"       ````      xKKd.  .lKKKd^     ",
			"                 `^     `^^         ",
		},
	},
	"opensuse-leap": {
		Color: "green",
		Lines: []string{
			"          .;ldkO0000Okdl;.         ",
			"      .;d00xl:^''''''^:ok00d;.     ",
			"    .d00l'                'o00d.   ",
			"  .d0Kd'  Okxol:;,.          :O0d.",
			" .OKKKK0kOKKKKKKKKKKOxo:,      lKO",
			",0KKKKKKKKKKKKKKKK0P^,,,^dx:    ;00",
			".OKKKKKKKKKKKKKKKk.          ;xOO' ",
			":KKKKKKKKKKKKKKKKK: .oOPPb.   'ODk.",
			"'KKKKKKKKKKKKKKKKK' .OKKKK0.  OKKKb",
			" d0KKKKKKKKKKKKKK. lKKKKKd  .OKKKK.",
			"  l0KKK000KKKKKx  d0KKKKx   xKKKK0 ",
			"   `^^^^^okKK0' .oKKKKx   dKKKKKK` ",
			"       ````      xKKd.  .lKKKd^     ",
			"                 `^     `^^         ",
		},
	},
}

// Fallback returns the generic Linux (Tux) logo.
func Fallback() Logo {
	return All["linux"]
}

// Get returns the logo for the given distro ID, falling back to the Linux logo.
func Get(distroID string) Logo {
	if logo, ok := All[distroID]; ok {
		return logo
	}
	return All["linux"]
}
