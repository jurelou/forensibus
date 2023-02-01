<div align="center">
<p><img width="210" src="assets/logo.png"/></p>
<!-- <h2><strong>Forensibus</strong></h2> -->
<p>Forensibus is a forensics <a href="https://en.wikipedia.org/wiki/Extract,_transform,_load">ETL</a> that focuses on processing digital forensics artefacts</p>
<p>
    <a href="https://goreportcard.com/report/github.com/jurelou/forensibus">
    <img src="https://goreportcard.com/badge/github.com/jurelou/forensibus"/>
    </a>
    <a href="https://deepsource.io/gh/jurelou/forensibus">
    <img src="https://deepsource.io/gh/jurelou/forensibus.svg/?label=active+issues&show_trend=true&token=grbhROhG-pVFvqhtQJUvozTU)](https://deepsource.io/gh/jurelou/forensibus/?ref=repository-badge"/>
    </a>

</p>
<p>
  <a href="#Features">Features</a> •
  <a href="#getting-started">Getting started</a> •
  <a href="#license">License</a>
</p>
</div>

--

# Features

- ⚡ Blazingly fast - Horizontal scaling and high performance parallelism
- 🖥️ Works with splunk right off the bat
- ⚙️ Modular - Add your own artefacts processors easily

# Getting started

# TODO
- A single processor processes multiple files at once
- Add support for yara arguments (filepath, filename, extension, ...)
- wrap errors properly
- processors as grpc plugins
- utils/filesystem: add some file types to the filetype lib. (evtx, registry, text ...)
- Add a --force flag to the run command, force exec even if some errors appear (some workers are unavailable, ....)
- Find (hcl) Add a `recurse` boolean
- Find (hcl) Add a `type` option file/dir (file by default)
- set windows releases PE ressources : https://github.com/tc-hib/go-winres- cli: Add worker ping, plugin list, plugin info, ....
- Add more archives extraction types
- splunk: configure event_renderers.conf to set event colors
- splunk HEC configure indexer acknowledgement
- https://parsiya.net/blog/2018-11-01-windows-filetime-timestamps-and-byte-wrangling-with-go/
- extract: add parralelism
- check goroutines /file descriptors leaks

# License

Source code in `forensibus` is available under the [GNU General Public License v3.0](/LICENSE).


# Install development dependencies

sudo apt-get install automake libtool make gcc pkg-config libssl-dev

## Yara 4.2.x


```
wget https://github.com/VirusTotal/yara/archive/refs/tags/v4.2.3.tar.gz
tar -xzvf v4.2.3.tar.gz
./bootstrap.sh
./configure
make
sudo make install
make check
```

sudo apt-get install mingw-w64-x86-64-dev gcc-mingw-w64-x86-64 gcc-mingw-w64 gcc-multilib

# Update the protocol buffers

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/worker.proto
```
