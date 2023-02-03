
pipeline "DFIR-ORC" {

  // find "ese" {
  //       patterns = [".*"]
  //       mime_types = [ "Extensible storage engine" ]
  //       process "srum" {}
  // }

  find "ese" {
        patterns = [".*"]
        mime_types = [ "Extensible storage engine" ]
        process "ese" {}
  }


  extract "DFIR-ORC archives" {
    patterns = ["Collect_[^/]*.7z$", "DFIR-[^/]*.7z$", "ORC_[^/]*.7z$"]
    mime_types = ["7-zip archive data"]


    // extract "System hives archives" {
    //   patterns = ["SystemHives.7z$"]
    //   mime_types = ["7-zip archive data"]

    //   find "System registry hives" {
    //     patterns = [".*"]
    //     mime_types = ["MS Windows registry file"]

    //     process "registry" {}
    //   }
    // }



  }


}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}