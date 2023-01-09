
pipeline "simple_pipeline" {

    extract "extract all 7z" {
        patterns = [".*"]
        mime_types = ["application/x-7z-compressed"]

        find "compressed prefetch files" {
            patterns = [".*.pf$"]

            process "prefetch" {}

        }

    }
  
    find "prefetch files" {
        patterns = [".*.pf$"]

        process "prefetch" {}

    }
}

output "splunk" {
  address = "test"
  username = "test"
  password = "test"
}