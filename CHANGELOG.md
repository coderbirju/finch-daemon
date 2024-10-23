# Changelog

## [0.9.1](https://github.com/coderbirju/finch-daemon/compare/v0.9.0...v0.9.1) (2024-10-23)


### Build System or External Dependencies

* **deps:** bump github.com/containerd/go-cni from 1.1.9 to 1.1.10 ([#53](https://github.com/coderbirju/finch-daemon/issues/53)) ([31583b0](https://github.com/coderbirju/finch-daemon/commit/31583b0bd25dfdcf5c53ae78882b9df3ac36cc11))
* **deps:** bump github.com/onsi/ginkgo/v2 from 2.17.1 to 2.20.2 ([#19](https://github.com/coderbirju/finch-daemon/issues/19)) ([e282c25](https://github.com/coderbirju/finch-daemon/commit/e282c253bfdd2bad7e97866e75598291892fb7fa))
* **deps:** bump github.com/onsi/gomega from 1.32.0 to 1.34.2 ([#18](https://github.com/coderbirju/finch-daemon/issues/18)) ([ea72df3](https://github.com/coderbirju/finch-daemon/commit/ea72df3f479e10ef0de0357a31a1686d626f5041))
* **deps:** bump github.com/spf13/cobra from 1.8.0 to 1.8.1 ([#49](https://github.com/coderbirju/finch-daemon/issues/49)) ([3eff666](https://github.com/coderbirju/finch-daemon/commit/3eff666f81e4ea655b9d70e5fa7e8043283ec959))
* **deps:** bump github.com/vishvananda/netlink from 1.2.1-beta.2 to 1.3.0 ([#50](https://github.com/coderbirju/finch-daemon/issues/50)) ([e3cffc7](https://github.com/coderbirju/finch-daemon/commit/e3cffc77ac28451c15d5c6a04ab63fd89c34fe4b))


### Features

* add container create options ([#27](https://github.com/coderbirju/finch-daemon/issues/27)) ([504dcaf](https://github.com/coderbirju/finch-daemon/commit/504dcaf9eff1316c9dd40db82a4ecce9b3e1796d))
* allow custom socket path ([#7](https://github.com/coderbirju/finch-daemon/issues/7)) ([4c17545](https://github.com/coderbirju/finch-daemon/commit/4c1754576d5beb3bd6b12e36893a588b2bb95825))
* implement container restart API ([#23](https://github.com/coderbirju/finch-daemon/issues/23)) ([5d9b1e0](https://github.com/coderbirju/finch-daemon/commit/5d9b1e0f4e1565fd374b0f0941f373a094dc749c))
* Port 'implement container restart API' patch ([5d9b1e0](https://github.com/coderbirju/finch-daemon/commit/5d9b1e0f4e1565fd374b0f0941f373a094dc749c))


### Bug Fixes

* doc nits and parameter casing ([#57](https://github.com/coderbirju/finch-daemon/issues/57)) ([e22c156](https://github.com/coderbirju/finch-daemon/commit/e22c156cc8bcb97f25c6f41a14e833203e8798ce))
* filter unsupported enable_icc option ([#36](https://github.com/coderbirju/finch-daemon/issues/36)) ([6c5e72d](https://github.com/coderbirju/finch-daemon/commit/6c5e72d4e8c9f6a5be12bf38078798423d11064f))
* image load should close stream after copy ([#34](https://github.com/coderbirju/finch-daemon/issues/34)) ([5ee657b](https://github.com/coderbirju/finch-daemon/commit/5ee657b17de96c1d2302e9ee7490ccfdc64cd907))
* Make release binaries static ([090da7c](https://github.com/coderbirju/finch-daemon/commit/090da7c219f790922aff699df53fe32b3295932b))
* README changes re: systemd setup ([#59](https://github.com/coderbirju/finch-daemon/issues/59)) ([2096ded](https://github.com/coderbirju/finch-daemon/commit/2096ded2283a8582186be01eeee42a8c0ab6161d))
* Set release version to 0.9.0 ([#56](https://github.com/coderbirju/finch-daemon/issues/56)) ([024768a](https://github.com/coderbirju/finch-daemon/commit/024768a6937ab2917870f9a3348dc0be114d3523))
* truncate image id on publish tag event ([#35](https://github.com/coderbirju/finch-daemon/issues/35)) ([6aa5b7c](https://github.com/coderbirju/finch-daemon/commit/6aa5b7ce76979682ad1cf2b49ac0237a74cac809))
* Update makefile for build target ([5887d3b](https://github.com/coderbirju/finch-daemon/commit/5887d3b53e2154943a374f41d82f658a57b5312f))

## 0.9.0  Finch-Daemon Init

This is the first release of the Finch Daemon.
The Finch Daemon project is an open source container runtime engine that enables users to integrate software that uses Docker's RESTful APIs as a programmatic dependency.
