toto = "tutu"

pipeline "test_simple_pipeline" {
  toto = "hello"

  find "1Find" {
    patterns = ["aa", "bb"]
  }
  process "1Process" {
  }
  find "2Find" {
    patterns = ["aa", "bb"]
  }
  process "2Process" {
  }
  process "3Process" {
  }
  process "4Process" {
  }

}

output "splunk" {
  address = "localhost"
  username = "admin"
  password = "admin"
}