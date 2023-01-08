
pipeline "prefetch" {

    extract "all archives" {
        patterns = [".*"]
        mime_types = ["application/x-7z-compressed", "application/x-zip", "custom/zip"]

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
  address = "localhost:9997"
  username = "xx"
  password = "yy"
}