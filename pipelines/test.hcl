
archives_passwords = ["virus", "avproof"]
temporary_folder = "/tmp/forensibus"

pipeline "testpipe" {

  find "first" {
    patterns = [".go$"]
    // mime_types = ["application/evtx"]

    process "evtxdump" {}

    find "first_second" {
      patterns = ["main.go$"]

      process "evtxdump" {}

    }

  }

  find "bis" {
    patterns = []
    mime_types = ["application/evtx"]

    process "evtxdump" {}

  }

}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}