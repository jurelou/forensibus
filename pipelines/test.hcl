
pipeline "testpipe" {

    extract "zip archives" {
      patterns = [".*.zip$"]

      process "evtxdump" {}

    }

  // extract "7z archives" {
  //     patterns = [".*.7z$"]
  
      find co files" {
        patterns = [".*"]

        process "prefetch" {}

      }

    // }


  // find "first" {
  //   patterns = [".zip$"]
  //   // mime_types = ["application/evtx"]

  //   process "evtxdump" {}
  //   extract "lol" {
  //     patterns = [".*"]
  //   }

  //   find "first_second" {
  //     patterns = ["main.go$"]

  //     process "evtxdump" {}

  //   }

  // }

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