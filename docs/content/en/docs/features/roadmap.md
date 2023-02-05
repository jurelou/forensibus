---
title: "Roadmap"
lead: ""
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 130
toc: true
---


# CLI

- Add a --force flag to the run command, force exec even if some errors appear (some workers are unavailable, ....)
- cli: Add worker ping, plugin list, plugin info, ....
- make config file optional (options can be set via CLI params, --config file to set a config file)

## Code quality 

- processors as grpc plugins
- check goroutines /file descriptors leaks

## dsl

- utils/filesystem: add some file types to the filetype lib. (evtx, registry, text ...)
- Find/extract (hcl) patterns is optional and defaults to ".*"
- Find (hcl) Add a `recurse` boolean
- Find (hcl) Add a `type` option file/dir (file by default)
- find / extract (hcl): Add exclusion patterns / exclusion mime types (eg: yara pipeline should not scan archives)

## Processors

- A single processor processes multiple files at once
- Add support for yara arguments (filepath, filename, extension, ...)
- extract: add parralelism (make it a processor)
- Add capa https://github.com/mandiant/capa

## Splunk

- splunk: configure event_renderers.conf to set event colors
- splunk HEC configure indexer acknowledgement

## Others

- set windows releases PE ressources : https://github.com/tc-hib/go-winres
- Add more archives extraction types
- Bufferize progress bar (it becomes slow when there is tons of fast tasks like yara)
