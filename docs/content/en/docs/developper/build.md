---
title: "Build from sources"
description: ""
lead: ""
draft: false
images: []
menu:
  docs:
    parent: "Developer guide"
weight: 310
toc: true
---

These steps have been tested on linux only, this allows you to compile the application for linux and windows.
Actually, forensibus has not been tested on OSX.


## Install dependencies

- golang

See the docs for installing the go programming language here: [](https://go.dev/doc/install)

- build tools

```bash
sudo apt-get install automake libtool make gcc pkg-config libssl-dev libmagic-dev
```

- musl (linux compiler)

```bash
wget http://musl.libc.org/releases/musl-1.2.3.tar.gz
tar -xvf musl-1.2.3.tar.gz
cd musl-1.2.3
./configure
make
sudo make install
```
This will install musl in `/usr/local/musl` – it will not conflict with your glibc. Alternatively, you can pass --prefix to ./configure to install it elsewhere; for example in your home directory.

- mingw (windows compiler)

```bash
apt-get install mingw-w64-x86-64-dev gcc-mingw-w64-x86-64 gcc-mingw-w64 gcc-multilib -y
```

- yara

```bash
wget https://github.com/VirusTotal/yara/archive/refs/tags/v4.2.3.tar.gz
tar -xzvf v4.2.3.tar.gz
cd yara-4.2.3/
./bootstrap.sh
./configure --enable-magic --enable-dotnet
make
sudo make install
make check
```

## Build

```bash
make
```

This will generate statically linked binaries in the `./bin` folder.
eg:

```bash
$ ldd bin/forensibus_linux_amd64 
        linux-vdso.so.1 (0x00007ffd481f7000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f398a339000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f398a147000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f398a375000)
```