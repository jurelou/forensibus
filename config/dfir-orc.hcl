
pipeline "DFIR-ORC" {

  extract "DFIR-ORC archives" {
    patterns = ["Collect_.*.7z$", "DFIR-.*.7z$", "ORC_.*.7z$"]
    mime_types = ["7-zip archive data"]

    extract "Evtx archives" {
      patterns = ["[Ee]vents?.7z$"]
      mime_types = ["7-zip archive data"]

      find "Evtx files" {
        patterns = [".evtx$", ".evtx_data$"]
        mime_types = ["MS Windows Vista Event Log"]

        process "evtxdump" {}
      }

    }

    extract "Artifacts" {
      patterns = ["[Aa]rtefacts?.7z$"]
      mime_types = ["7-zip archive data"]

      find "Prefetch files" {
        patterns = [".pf$", ".pf_data$"]

        process "prefetch" {}
      }
    }

    extract "NTFSInfo archives" {
      patterns = ["NTFSInfo.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS info files" {
        patterns = ["NTFSInfo_.*.csv$"]

        process "csv" {}
      }
    }

    extract "NTFS I30 archives" {
      patterns = ["NTFSInfo_[Ii]30Info.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS I30 files" {
        patterns = ["I30Info_.*.csv$"]

        process "csv" {}
      }
    }


    extract "NTFS security descriptors archives" {
      patterns = ["NTFSInfo_SecDesc.7z$"]
      mime_types = ["7-zip archive data"]

      find "NTFS security descriptors files" {
        patterns = ["SecDescr_.*.csv$"]

        process "csv" {}
      }
    }

    extract "System hives archives" {
      patterns = ["SystemHives.7z$"]
      mime_types = ["7-zip archive data"]

      find "System registry hives" {
        patterns = [".*"]
        mime_types = ["MS Windows registry file"]

        process "registry" {}
      }
    }



  }


}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}