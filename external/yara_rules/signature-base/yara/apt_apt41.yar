/*
   Yara Rule Set
   Author: Florian Roth
   Date: 2019-08-07
   Identifier: APT41
   Reference: https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html
   License: https://creativecommons.org/licenses/by-nc/4.0/
*/

/* Rule Set ----------------------------------------------------------------- */

import "pe"

rule APT_APT41_POISONPLUG_3 {
   meta:
      description = "Detects APT41 malware POISONPLUG"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 80
      hash1 = "70c03ce5c80aca2d35a5555b0532eedede24d4cc6bdb32a2c8f7e630bba5f26e"
   strings:
      $s1 = "Rundll32.exe \"%s\", DisPlay 64" fullword ascii
      $s2 = "tcpview.exe" fullword ascii
      $s3 = "nuR\\noisreVtnerruC\\swodniW\\tfosorciM\\erawtfoS" fullword ascii /* reversed goodware string 'Software\\Microsoft\\Windows\\CurrentVersion\\Run' */
      $s4 = "AxEeulaVteSgeR" fullword ascii /* reversed goodware string 'RegSetValueExA' */
      $s5 = "%04d-%02d-%02d_%02d-%02d-%02d.dmp" fullword ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 900KB and 3 of them
}

rule APT_APT41_POISONPLUG_SHADOW {
   meta:
      description = "Detects APT41 malware POISONPLUG SHADOW"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 85
      hash1 = "462a02a8094e833fd456baf0a6d4e18bb7dab1a9f74d5f163a8334921a4ffde8"
   condition:
      uint16(0) == 0x5a4d and filesize < 500KB and pe.imphash() == "c67de089f2009b21715744762fc484e8"
}

rule APT_APT41_CRACKSHOT {
   meta:
      description = "Detects APT41 malware CRACKSHOT"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 85
      hash1 = "993d14d00b1463519fea78ca65d8529663f487cd76b67b3fd35440bcdf7a8e31"
   strings:
      $x1 = ";procmon64.exe;netmon.exe;tcpview.exe;MiniSniffer.exe;smsniff.exe" ascii

      $s1 = "RunUrlBinInMem" fullword ascii
      $s2 = "DownRunUrlFile" fullword ascii
      $s3 = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36" fullword ascii
      $s4 = "%s|%s|%s|%s|%s|%s|%s|%dx%d|%04x|%08X|%s|%s" fullword ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 250KB and ( 1 of ($x*) or 2 of them )
}

rule APT_APT41_POISONPLUG_2 {
   meta:
      description = "Detects APT41 malware POISONPLUG"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 70
      hash1 = "0055dfaccc952c99b1171ce431a02abfce5c6f8fb5dc39e4019b624a7d03bfcb"
   strings:
      $s1 = "ma_lockdown_service.dll" fullword wide
      $s2 = "acbde.dll" fullword ascii
      $s3 = "MA lockdown Service" fullword wide
      $s4 = "McAfee Agent" fullword wide
   condition:
      uint16(0) == 0x5a4d and filesize < 11000KB and all of them
}

rule APT_APT41_POISONPLUG {
   meta:
      description = "Detects APT41 malware POISONPLUG"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 80
      hash1 = "2eea29d83f485897e2bac9501ef000cc266ffe10019d8c529555a3435ac4aabd"
      hash2 = "5d971ed3947597fbb7e51d806647b37d64d9fe915b35c7c9eaf79a37b82dab90"
      hash3 = "f4d57acde4bc546a10cd199c70cdad09f576fdfe66a36b08a00c19ff6ae19661"
      hash4 = "3e6c4e97cc09d0432fbbbf3f3e424d4aa967d3073b6002305cd6573c47f0341f"
   strings:
      $s1 = "TSMSISrv.DLL" fullword wide
      $s2 = "[-]write failed[%d]" fullword ascii
      $s3 = "[-]load failed" fullword ascii
      $s4 = "Remote Desktop Services" fullword wide
   condition:
      uint16(0) == 0x5a4d and filesize < 10000KB and (
         pe.imphash() == "1b074ef7a1c0888ef31337c8ad2f2e0a" or
         2 of them
      )
}

rule APT_APT41_HIGHNOON {
   meta:
      description = "Detects APT41 malware HIGHNOON"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 85
      hash1 = "63e8ed9692810d562adb80f27bb1aeaf48849e468bf5fd157bc83ca83139b6d7"
      hash2 = "4aa6970cac04ace4a930de67d4c18106cf4004ba66670cfcdaa77a4c4821a213"
   strings:
      $x1 = "workdll64.dll" fullword ascii

      $s1 = "\\Fonts\\Error.log" ascii
      $s2 = "[%d/%d/%d/%d:%d:%d]" fullword ascii
      $s3 = "work_end" fullword ascii
      $s4 = "work_start" fullword ascii
      $s5 = "\\svchost.exe" ascii
      $s6 = "LoadAppInit_DLLs" fullword ascii
      $s7 = "netsvcs" fullword ascii
      $s8 = "HookAPIs ...PID %d " fullword ascii
      $s9 = "SOFTWARE\\Microsoft\\HTMLHelp" fullword ascii
      $s0 = "DllMain_mem" fullword ascii
      $s10 = "%s\\NtKlRes.dat" fullword ascii
      $s11 = "Global\\%s-%d" fullword ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 2000KB and ( 1 of ($x*) or 4 of them )
}

rule APT_APT41_HIGHNOON_2 {
   meta:
      description = "Detects APT41 malware HIGHNOON"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      hash1 = "79190925bd1c3fae65b0d11db40ac8e61fb9326ccfed9b7e09084b891089602d"
   strings:
      $x1 = "H:\\RBDoor\\" ascii

      $s1 = "PlusDll.dll" fullword ascii
      $s2 = "ShutDownEvent.dll" fullword ascii
      $s3 = "\\svchost.exe" ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 600KB and (
         pe.imphash() == "b70358b00dd0138566ac940d0da26a03" or
         pe.exports("DllMain_mem") or
         $x1 or 3 of them
      )
}

rule APT_APT41_HIGHNOON_BIN {
   meta:
      description = "Detects APT41 malware HIGHNOON.BIN"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 90
      hash1 = "490c3e4af829e85751a44d21b25de1781cfe4961afdef6bb5759d9451f530994"
      hash2 = "79190925bd1c3fae65b0d11db40ac8e61fb9326ccfed9b7e09084b891089602d"
   strings:
      $s1 = "PlusDll.dll" fullword ascii
      $s2 = "\\Device\\PORTLESS_DeviceName" wide
      $s3 = "%s%s\\Security" fullword ascii
      $s4 = "%s\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Internet Settings" fullword ascii
      $s5 = "%s%s\\Enum" fullword ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 600KB and (
         pe.imphash() == "b70358b00dd0138566ac940d0da26a03" or
         3 of them
      )
}

rule APT_APT41_HIGHNOON_BIN_2 {
   meta:
      description = "Detects APT41 malware HIGHNOON.BIN"
      author = "Florian Roth"
      reference = "https://www.fireeye.com/blog/threat-research/2019/08/apt41-dual-espionage-and-cyber-crime-operation.html"
      date = "2019-08-07"
      score = 85
      hash1 = "63e8ed9692810d562adb80f27bb1aeaf48849e468bf5fd157bc83ca83139b6d7"
      hash2 = "c51c5bbc6f59407286276ce07f0f7ea994e76216e0abe34cbf20f1b1cbd9446d"
   strings:
      $x1 = "\\Double\\Door_wh\\" ascii
      $x2 = "[Stone] Config --> 2k3 TCP Positive Logout." fullword ascii
      $x3 = "\\RbDoorX64.pdb" ascii
      $x4 = "RbDoor, Version 1.0" fullword wide
      $x5 = "About RbDoor" fullword wide
   condition:
      uint16(0) == 0x5a4d and filesize < 2000KB and 1 of them
}

rule APT_APT41_CN_ELF_Speculoos_Backdoor {
   meta:
      description = "Detects Speculoos Backdoor used by APT41"
      author = "Florian Roth"
      reference = "https://unit42.paloaltonetworks.com/apt41-using-new-speculoos-backdoor-to-target-organizations-globally/"
      date = "2020-04-14"
      score = 90
      hash1 = "6943fbb194317d344ca9911b7abb11b684d3dca4c29adcbcff39291822902167"
      hash2 = "99c5dbeb545af3ef1f0f9643449015988c4e02bf8a7164b5d6c86f67e6dc2d28"
   strings:
      $xc1 = { 2F 70 72 69 76 61 74 65 2F 76 61 72 00 68 77 2E
               70 68 79 73 6D 65 6D 00 68 77 2E 75 73 65 72 6D
               65 6D 00 4E 41 2D 4E 41 2D 4E 41 2D 4E 41 2D 4E
               41 2D 4E 41 00 6C 6F 30 00 00 00 00 25 30 32 78
               2D 25 30 32 78 2D 25 30 32 78 2D 25 30 32 78 2D
               25 30 32 78 2D 25 30 32 78 0A 00 72 00 4E 41 00
               75 6E 61 6D 65 20 2D 76 }
      
      $s1 = "badshell" ascii fullword
      $s2 = "hw.physmem" ascii fullword
      $s3 = "uname -v" ascii fullword
      $s4 = "uname -s" ascii fullword
      $s5 = "machdep.tsc_freq" ascii fullword
      $s6 = "/usr/sbin/config.bak" ascii fullword
      $s7 = "enter MessageLoop..." ascii fullword
      $s8 = "exit StartCBProcess..." ascii fullword

      $sc1 = { 72 6D 20 2D 72 66 20 22 25 73 22 00 2F 70 72 6F
               63 2F }
   condition:
      uint16(0) == 0x457f and
      filesize < 600KB and
      1 of ($x*) or 4 of them
}
