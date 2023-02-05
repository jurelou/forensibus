---
title: "Supported artifacts"
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 130
toc: true
---

## Common file formats

{{< details "csv" >}}
Parse comma separated values.
{{< /details >}}

{{< details "ini" >}}
Parse [ini](https://en.wikipedia.org/wiki/INI_file) files.
For example:

```ini
; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = "payroll.dat"
```
{{< /details >}}

{{< details "Extended Storage Engine" >}}
Extensible Storage Engine (ESE), also known as JET Blue is a data storage technology from Microsoft. ESE databases are dumped using [go-ese](https://github.com/Velocidex/go-ese/tree/master/bin)
{{< /details >}}

## Malware analysis

{{< details "Yara" >}}
The yara processor matches files to the rules contained in the `external/yara_rules/` folder. You can configure the default folder by editing `yaraRulesFolder` in the config file.
__Rules with external variables are not supported yet__
{{< /details >}}


## Windows

{{< details "Master File Table" >}}
The MFT is analysed using [go-ntfs](https://github.com/Velocidex/go-ntfs)
{{< /details >}}

{{< details "Sigma" >}}
Forensibus uses [hayabusa](https://github.com/Yamato-Security/hayabusa) to match sigma rules. Hayabusa was chosen instead of [zircolite](https://github.com/wagga40/Zircolite) or [chainsaw](https://github.com/WithSecureLabs/chainsaw) because it is performant and it respects the [sigma specification](https://github.com/SigmaHQ/sigma-specification). Also, hayabusa allows us to create timelines easily.
{{< /details >}}

{{< details "System Resource Usage Monitor (SRUM)" >}}
The SRUM tracks process and network statistics over time in a database. Internally, the SRUM processor uses the ESE processor.
{{< /details >}}

{{< details "Registry hives" >}}
Windows registry hives are analysed using [regparser](https://github.com/Velocidex/regparser)
{{< /details >}}

{{< details "Prefetch" >}}
Prefetch files are analysed using [go-prefetch](https://github.com/Velocidex/go-prefetch)
{{< /details >}}

{{< details "Windows XML EventLog (EVTX)" >}}
At the moment, evt files are not supported.
{{< /details >}}

## Linux 
