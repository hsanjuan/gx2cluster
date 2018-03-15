# gxc

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> Pin your project&#39;s Gx dependency tree in ipfs-cluster.

Run your `gxc` command in your [Gx'ed](https://github.com/whyrusleeping/gx) project and all your dependencies will be submitted to ipfs-cluster for pinning, correctly named.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

Safest way is to build manually rewriting gx'ed deps:

```
go get -u -d https://github.com/hsanjuan/gxc
cd $GOPATH/src/github.com/hsanjuan/gxc
gx install --global
gx-go rw
go install
```

## Usage

Submit to the local, default ipfs-cluster API endpoint (`/ip4/localhost/tcp/9094`)

```
$ gxc
```

Submit to your remote ipfs-cluster peer:

```
$ gxc --peer <multiaddress>
```

Other options (`-h`):

```
Usage of gxc:
  -peer string
        multiaddress of the IPFS Cluster API (default "/ip4/127.0.0.1/tcp/9094")
  -pnet string
        pnet key
  -pw string
        basic auth pw
  -ssl
        enable ssl
  -user string
        basic auth username
```

## Contribute

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Hector Sanjuan
