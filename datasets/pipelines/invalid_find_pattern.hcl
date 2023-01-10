
pipeline "test_invalid_processor" {

  find "invalid pattern" {
    patterns = [".(this_is_invalid"]

  }

}

output "splunk" {
  address = "localhost"
  username = "admin"
  password = "admin"
}