
pipeline "test_invalid_processor" {

  find "" {
    patterns = ["aa"]
    process "ThisProcessorWillNeverExists" {}

  }

}

output "splunk" {
  address = "localhost"
  username = "admin"
  password = "admin"
}