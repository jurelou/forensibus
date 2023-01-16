
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
  }


}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}