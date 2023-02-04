
pipeline "test_invalid_processor" {

  extract "invalid ep" {
    patterns = ["*test"]

  }

}

output "splunk" {
  address = "localhost"
  username = "admin"
  password = "admin"
}