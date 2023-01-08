
pipeline "test_simple_pipeline" {

  find "firstFind" {
    patterns = ["aa", "bb"]

    extract "firstExtract" {
        patterns = []
        mime_types = ["some/file"]
        process "firstProcess" {

        }
    }
    process "secondProcess" {}

    find "secondFind" {
        mime_types = ["application/something"]
        patterns = ["*.a"]

        extract "nested" {
            patterns = []

            process "thirdProcess" {
                config = {"rules_path": "/tmp/myrules"}
            }

            process "fifthProcess" {
                config = {"rules_path": "/tmp/myrules"}
            }


        }
        extract "nestedBis" {
          patterns = []
          mime_types = []
          process "thirdBis" {}
          process "fifthBis" {}
        }
    }

  }
  process "fourthProcess" {

  }
}

output "splunk" {
  address = "localhost"
  username = "admin"
  password = "admin"
}