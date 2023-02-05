---
title: "Quick Start"
description: "One page summary of how to start a new Doks project."
lead: "One page summary of how to start a new Doks project."
date: 2020-11-16T13:59:39+01:00
lastmod: 2020-11-16T13:59:39+01:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 210
toc: true
---

While Internet access is required to build forensibus, it is not required at runtime.

dfTimewolf is typically run by specifying a recipe name and any arguments the recipe defines. For example:

>> dftimewolf plaso_ts /tmp/path1,/tmp/path2 --incident_id 12345

This will launch the plaso_ts recipe against path1 and path2 in /tmp. In this recipe --incident_id is used by Timesketch as a sketch description.

Details on a recipe can be obtained using the standard python help flags:

$ dftimewolf -h
[2020-10-06 14:29:42,111] [dftimewolf          ] INFO     Logging to stdout and /tmp/dftimewolf.log
[2020-10-06 14:29:42,111] [dftimewolf          ] DEBUG    Recipe data path: /Users/tomchop/code/dftimewolf/data
[2020-10-06 14:29:42,112] [dftimewolf          ] DEBUG    Configuration loaded from: /Users/tomchop/code/dftimewolf/data/config.json
usage: dftimewolf [-h]
                             {aws_forensics,gce_disk_export,gcp_forensics,gcp_logging_cloudaudit_ts,gcp_logging_cloudsql_ts,...}

Available recipes:

 aws_forensics                      Copies a volume from an AWS account to an analysis VM.
 gce_disk_export                    Export disk image from a GCP project to Google Cloud Storage.
 gcp_forensics                      Copies disk from a GCP project to an analysis VM.
 gcp_logging_cloudaudit_ts          Collects GCP logs from a project and exports them to Timesketch.
 [...]

positional arguments:
  {aws_forensics,gce_disk_export,gcp_forensics,gcp_logging_cloudaudit_ts,...}

optional arguments:
  -h, --help            show this help message and exit




To get details on an individual recipe, call the recipe with the -h flag.









Running a recipe¶
One typically invokes dftimewolf with a recipe name and a few arguments. For example:

$ dftimewolf <RECIPE_NAME> arg1 arg2 --optarg1 optvalue1
Given the help output above, you can then use the recipe like this:

$ dftimewolf grr_artifacts_ts tomchop.greendale.xyz collection_reason
If you only want to collect browser activity:

$ dftimewolf grr_artifacts_ts tomchop.greendale.xyz collection_reason --artifact_list=BrowserHistory
In the same way, if you want to specify one (or more) approver(s):

$ dftimewolf grr_artifacts_ts tomchop.greendale.xyz collection_reason --artifact_list=BrowserHistory --approvers=admin
$ dftimewolf grr_artifacts_ts tomchop.greendale.xyz collection_reason --artifact_list=BrowserHistory --approve







## Requirements

- [Git](https://git-scm.com/) — latest source release
- [Node.js](https://nodejs.org/) — latest LTS version or newer

{{< details "Why Node.js?" >}}
Doks uses npm (included with Node.js) to centralize dependency management, making it [easy to update]({{< relref "how-to-update" >}}) resources, build tooling, plugins, and build scripts.
{{< /details >}}

## Start a new Doks project

Create a new site, change directories, install dependencies, and start development server.

### Create a new site

Doks is available as a child theme and a starter theme.

#### Child theme

- Intended for novice to intermediate users
- Intended for minor customizations
- [Easily update npm packages]({{< relref "how-to-update" >}}) — __including__ [Doks](https://www.npmjs.com/package/@hyas/doks)

```bash
git clone https://github.com/h-enk/doks-child-theme.git my-doks-site
```

#### Starter theme

- Intended for intermediate to advanced users
- Intended for major customizations
- [Easily update npm packages]({{< relref "how-to-update" >}})

```bash
git clone https://github.com/h-enk/doks.git my-doks-site
```

{{< details "Help me choose" >}}
Not sure which one is for you? Pick the child theme.
{{< /details >}}

### Change directories

```bash
cd my-doks-site
```

### Install dependencies

```bash
npm install
```

### Start development server

```bash
npm run start
```

Doks will start the Hugo development webserver accessible by default at `http://localhost:1313`. Saved changes will live reload in the browser.

## Other commands

Doks comes with commands for common tasks. [Commands →]({{< relref "commands" >}})
