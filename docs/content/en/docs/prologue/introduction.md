---
title: "Introduction"
lead: ""
date: 2020-10-06T08:48:57+00:00
lastmod: 2020-10-06T08:48:57+00:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 100
toc: true
---


Forensibus is a modern data processing tool dedicated to forensic analysts and incident responders. It was built to solve the following problems:

- Forensic data comes in a wide variety of formats. In case of incident response, the format may not be known in advance (DFIR collection tools, raw disks, event logs, EDR telemetry, ...).

- These challenges are exacerbated by the increase in the size of the networks and the volume of data collected. Improved efficiency of the digital investigation process is needed, in terms of increasing the speed and reducing the human effort.

- The Digital Forensic Investigation process is largely manual in nature, or at best semi-automated

Forensibus aims to improve efficiency by emphasizing automation within the digital forensic triage process.


## Architecture Overview

The objective is to carry out the ETL (extract - transform - load) process by reading the raw artifacts and then formatting them so that they can be loaded into a database for further analysis.


![Green Sea Turtle Hatchling by Hannah Le Leu](/flow.png)

{{< alert icon="👉" text="Currently, forensibus does not supports writing to elasticsearch." />}}

### Tutorial

{{< alert icon="👉" text="The Tutorial is intended for novice to intermediate users." />}}

Step-by-step instructions on how to start a new Doks project. [Tutorial →](https://getdoks.org/tutorial/introduction/)

### Quick Start

{{< alert icon="👉" text="The Quick Start is intended for intermediate to advanced users." />}}

One page summary of how to start a new Doks project. [Quick Start →]({{< relref "quick-start" >}})

## Go further

Recipes, Reference Guides, Extensions, and Showcase.

### Recipes

Get instructions on how to accomplish common tasks with Doks. [Recipes →](https://getdoks.org/docs/recipes/project-configuration/)

### Reference Guides

Learn how to customize Doks to fully make it your own. [Reference Guides →](https://getdoks.org/docs/reference-guides/security/)

### Extensions

Get instructions on how to add even more to Doks. [Extensions →](https://getdoks.org/docs/extensions/breadcrumb-navigation/)

### Showcase

See what others have build with Doks. [Showcase →](https://getdoks.org/showcase/electric-blocks/)

## Contributing

Find out how to contribute to Doks. [Contributing →](https://getdoks.org/docs/contributing/how-to-contribute/)

## Help

Get help on Doks. [Help →]({{< relref "how-to-update" >}})
