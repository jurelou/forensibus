
pipeline "testpipe" {

    find "csv files" {
      patterns = [".*.[Cc][Ss][Vv]$"]

      process "csv" {
        config = {
          aa = 1
        }
      }
    }
}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}