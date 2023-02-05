
pipeline "DFIR-ORC" {

  extract "DFIR-ORC archives" {
    patterns = ["/Collect_[^/]*.7z$", "/DFIR-[^/]*.7z$", "/ORC_[^/]*.7z$"]
    mime_types = ["7-zip archive data"]

    extract "GetSamples archives" {
      patterns = ["/GetSamples.7z$"]
      mime_types = ["7-zip archive data"]

      find "All files" {
        patterns = [".*"]

        process "yara" {}
      }
    }

    extract "EXE_TMP archives" {
      patterns = ["/EXE_TMP.7z$"]
      mime_types = ["7-zip archive data"]

      find "All files" {
        patterns = ["/EXE_TMP/"]

        process "yara" {}
      }
    }

    extract "Scripts archives" {
      patterns = ["/[Ss]cripts.7z$"]
      mime_types = ["7-zip archive data"]

      find "All files" {
        patterns = [".*"]

        process "yara" {}
      }
    }

    extract "Evtx archives" {
      patterns = ["/[Ee]vents?.7z$"]
      mime_types = ["7-zip archive data"]

      process "sigma" {
        config = {
          evtx_extension = ".evtx_data"
        }
      }

      find "Evtx files" {
        patterns = [".evtx$", ".evtx_data$"]
        mime_types = ["MS Windows Vista Event Log"]

        process "evtxdump" {}
      }
    }

    extract "Alternate data streams archives" {
      patterns = ["/[Aa][Dd][Ss].7z$"]
      mime_types = ["7-zip archive data"]

      find "Alternate data streams" {
        patterns = ["/ads/"]

        process "ini" {
          sourcetype = "forensibus:ads"
        }
      }

      find "SRUM database" {
        patterns = ["/[Ss][Rr][Uu][Mm]/"]
        mime_types = ["Extensible storage engine"]

        process "srum" {}
      }
    }

    extract "Artifacts archives" {
      patterns = ["/[Aa]rtefacts?.7z$"]
      mime_types = ["7-zip archive data"]

      find "Prefetch files" {
        patterns = [".pf$", ".pf_data$"]

        process "prefetch" {}
      }

      find "SRUM database" {
        patterns = ["/[Ss][Rr][Uu][Mm]/"]
        mime_types = ["Extensible storage engine"]

        process "srum" {}
      }
    }

    extract "NTFSInfo archives" {
      patterns = ["/NTFSInfo.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS info files" {
        patterns = ["NTFSInfo_[^/]*.csv$"]

        process "csv" {
          sourcetype = "forensibus:ntfs:ntfsinfo"
        }
      }
    }

    extract "NTFS I30 archives" {
      patterns = ["/NTFSInfo_[Ii]30Info.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS I30 files" {
        patterns = ["I30Info_[^/]*.csv$"]

        process "csv" {
          sourcetype = "forensibus:ntfs:i30info"
        }
      }
    }

    extract "NTFS security descriptors archives" {
      patterns = ["/NTFSInfo_SecDescr?.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS security descriptors files" {
        patterns = ["SecDescr?_[^/]*.csv$"]

        process "csv" {
          sourcetype = "forensibus:ntfs:security_descriptors"
        }
      }
    }

    extract "USN info archives" {
      patterns = ["/USNInfo.7z$"]
      mime_types = ["7-zip archive data"]

      find "USN info files" {
        patterns = ["/USNInfo_[^/]*.csv$"]

        process "csv" {
          sourcetype = "forensibus:ntfs:usninfo"
        }
      }
    }

    extract "UEFI archives" {
      patterns = ["/UefiFull.7z$"]
      mime_types = ["7-zip archive data"]

      find "UEFI files" {
        patterns = ["/GetSectors.csv$"]

        process "csv" {
          sourcetype = "forensibus:boot_sectors"
        }
      }
    }

    extract "System hives archives" {
      patterns = ["/SystemHives.7z$"]
      mime_types = ["7-zip archive data"]

      find "System registry hives" {
        patterns = [".*"]
        mime_types = ["MS Windows registry file"]

        process "registry" {}
      }
    }

    extract "User hives archives" {
      patterns = ["/UserHives.7z$"]
      mime_types = ["7-zip archive data"]

      find "User registry hives" {
        patterns = [".*"]
        mime_types = ["MS Windows registry file"]

        process "registry" {}
      }
    }

    extract "SAM hives archives" {
      patterns = ["/[Ss][Aa][Mm].7z$"]
      mime_types = ["7-zip archive data"]

      find "SAM registry hives" {
        patterns = [".*"]
        mime_types = ["MS Windows registry file"]

        process "registry" {}
      }
    }

    extract "Browser history archives" {
      patterns = ["/Browsers?.7z$"]
      mime_types = ["7-zip archive data"]

      find "Browser ESE databases" {
        patterns = [".*"]
        mime_types = ["Extensible storage engine"]

        process "ese" {
          sourcetype = "forensibus:browsers_history"
        }
      }
    }
  }
}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}