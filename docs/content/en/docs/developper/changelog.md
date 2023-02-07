---
title: "Changelog"
description: ""
lead: ""
draft: false
images: []
menu:
  docs:
    parent: "Developer guide"
weight: 330
toc: true
---

Forensibus changelog is generated using [git-chglog](https://github.com/git-chglog/git-chglog)

# Install dependencies

```bash
go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
```

# Generate the changelog

```bash
make changelog
```

This will generate a `./CHANELOG.md` file
