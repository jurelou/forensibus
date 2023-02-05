
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


  find "all" {
        patterns = [".*"]

        process "ini" {}
  }

}

output "splunk" {
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}