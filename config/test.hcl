
pipeline "DFIR-ORC" {

  // find "ese" {
  //       patterns = [".*"]
  //       mime_types = [ "Extensible storage engine" ]
  //       process "srum" {}
  // }


  // find "ese" {
  //       patterns = [".*"]
  //       mime_types = [ "Extensible storage engine" ]
  //       process "ese" {}
  // }

    //   find "System registry hives" {
    //     patterns = [".*"]
    //     mime_types = ["MS Windows registry file"]

    //     process "registry" {}
    //   }
    // }

  process "sigma" {
    config = {
      evtx_extension = ".evtx_data"
    }
  }

}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}