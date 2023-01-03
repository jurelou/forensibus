
archives_passwords = ["virus", "avproof"]
temporary_folder = "/tmp/forensibus"

pipeline "mypipeline4orc" {

  extract "DFIR-ORC archives" {
    patterns = ["Collect_.*.7z$", "DFIR-.*.7z$"]

    extract "evtx archives" {
      patterns = ["Evtx.7z$", "Events.7z$"]

      find "evtx files" {
        patterns = []
        mime_types = ["application/evtx"]

        process "evtxdump" {}

      }

      process "hayabusa" {
        config = {"rules_folder": "/tmp/myrules"}
        // Process le dossier evtx complet
      }

    }

    extract "artifacts archive" {
      patterns = ["Artifacts.7z"]

      find "files matching lnk" {
        patterns = [".*.lnk$"]

        process "lnk" {}

      }

      process "yara" {
        config = {"rules_path": "/tmp/myrules"}
        // Process sur le dossier artifacts
      }

    }
  }

  process "something" {
    config = {"abc": "def"}
  }

}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}