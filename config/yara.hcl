pipeline "yara" {

  process "yara" {}

  extract "all archives" {
    patterns = [".*"]
    mime_types = ["7-zip archive data", "Zip archive", "gzip compressed data", "POSIX tar archive"]

    find "all files" {
        patterns = [".*"]
        process "yara" {}

    }
  }
}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}