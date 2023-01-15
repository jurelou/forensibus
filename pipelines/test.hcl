
pipeline "testpipe" {

    find "evtx files" {
      patterns = [".*"]
      mime_types = ["MS Windows Vista Event Log"]

      process "evtxdump" {}

    }


}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}