# Changelog

## [0.2.0](https://github.com/dominikhaska/codegen/compare/v0.1.8...v0.2.0) (2024-11-05)


### ⚠ BREAKING CHANGES

* lower json schema version, rename number to float ([#12](https://github.com/dominikhaska/codegen/issues/12))

### 🐛 Bug Fixes

* container copy command ([#40](https://github.com/dominikhaska/codegen/issues/40)) ([8448543](https://github.com/dominikhaska/codegen/commit/8448543fda56a3d68851cf44a4735c1902bf5b98))
* docker publishing ([c663816](https://github.com/dominikhaska/codegen/commit/c663816e33d0a020c1bd4db110ac0e4f451ff7b1))
* docker publishing ([2d24d51](https://github.com/dominikhaska/codegen/commit/2d24d5141c0822edb7254f38efdabaa6e9b5b351))
* release permissions ([#25](https://github.com/dominikhaska/codegen/issues/25)) ([dc07cdf](https://github.com/dominikhaska/codegen/commit/dc07cdfe5487c0a22209c54d0ee195bbdcf1b5ed))
* set github token for release process ([a2fe4aa](https://github.com/dominikhaska/codegen/commit/a2fe4aa33e380e86925480e7233eeed4bfb9ed90))


### ✨ New Features

* add basic react support ([#31](https://github.com/dominikhaska/codegen/issues/31)) ([757ab66](https://github.com/dominikhaska/codegen/commit/757ab66b7fde7103ca6f5cb7f10c0632073b58d8))
* add version command ([#38](https://github.com/dominikhaska/codegen/issues/38)) ([c13a448](https://github.com/dominikhaska/codegen/commit/c13a4486b9b42f3e4a6f34abd43a87aecf91355e))
* Fixing problems with generated code for golang and adding sample manifest for testing. ([#22](https://github.com/dominikhaska/codegen/issues/22)) ([558e964](https://github.com/dominikhaska/codegen/commit/558e9640b8756e9cacccfdb23f136d95bd81629b))
* initial CLI for codegen with support for golang strongly typed accessors ([#13](https://github.com/dominikhaska/codegen/issues/13)) ([e8f3d3e](https://github.com/dominikhaska/codegen/commit/e8f3d3ea2815b7d5473746e71f1bedc856e723c8))
* lower json schema version, rename number to float ([#12](https://github.com/dominikhaska/codegen/issues/12)) ([ed844b4](https://github.com/dominikhaska/codegen/commit/ed844b43a3d05113b49b39a1e368d0ee3c308dc9))


### 🧹 Chore

* **main:** release 0.1.0 ([#24](https://github.com/dominikhaska/codegen/issues/24)) ([1c53caa](https://github.com/dominikhaska/codegen/commit/1c53caa25101821a90b0984ba4dbb9b7dc14f34f))
* **main:** release 0.1.1 ([#26](https://github.com/dominikhaska/codegen/issues/26)) ([33efe94](https://github.com/dominikhaska/codegen/commit/33efe94d6126047211c67d2ea3e62268a8acbab5))
* **main:** release 0.1.2 ([#29](https://github.com/dominikhaska/codegen/issues/29)) ([3499c45](https://github.com/dominikhaska/codegen/commit/3499c45ed5619a40d8a229d1ceac54acd3e67152))
* **main:** release 0.1.3 ([#35](https://github.com/dominikhaska/codegen/issues/35)) ([7d619ad](https://github.com/dominikhaska/codegen/commit/7d619ad23c017086a88f0794eb36c73243ac9132))
* **main:** release 0.1.4 ([#39](https://github.com/dominikhaska/codegen/issues/39)) ([630f730](https://github.com/dominikhaska/codegen/commit/630f73027b7c7c56bd59b19c1a0785b3b5993522))
* **main:** release 0.1.5 ([#41](https://github.com/dominikhaska/codegen/issues/41)) ([8fa83b0](https://github.com/dominikhaska/codegen/commit/8fa83b0bdb493cb6e0f73f67767c47bce56f3e5c))
* **main:** release 0.1.6 ([#48](https://github.com/dominikhaska/codegen/issues/48)) ([bbc18a9](https://github.com/dominikhaska/codegen/commit/bbc18a91bf85404f1d00766ea52e7b01b40a5bf3))
* **main:** release 0.1.7 ([#50](https://github.com/dominikhaska/codegen/issues/50)) ([784b8ab](https://github.com/dominikhaska/codegen/commit/784b8ab63dbada0517cbe47e1fb35233f015efba))
* **main:** release 0.1.8 ([#51](https://github.com/dominikhaska/codegen/issues/51)) ([165c6f3](https://github.com/dominikhaska/codegen/commit/165c6f39a6d726a81246d8617d49b596d83ad31c))
* remove empty testutils package ([#55](https://github.com/dominikhaska/codegen/issues/55)) ([9dc1d9f](https://github.com/dominikhaska/codegen/commit/9dc1d9fbc3751b53956e4c61cd43df63edca9f19))


### 📚 Documentation

* Add initial flag manifest schema ([#9](https://github.com/dominikhaska/codegen/issues/9)) ([fac35ca](https://github.com/dominikhaska/codegen/commit/fac35caff88e1ef9a9c5ff1e8624040d91db9307))
* switch from code gen to cli ([#47](https://github.com/dominikhaska/codegen/issues/47)) ([7a1f9f3](https://github.com/dominikhaska/codegen/commit/7a1f9f304cc9c512b407b19986fbd82e3b80fe53))


### 🔄 Refactoring

* change folder, package structure; integrate with cobra ([#27](https://github.com/dominikhaska/codegen/issues/27)) ([850c694](https://github.com/dominikhaska/codegen/commit/850c694c84fad1a71722a1b1e620f1473bc2d2ab))
* change name of go module ([#46](https://github.com/dominikhaska/codegen/issues/46)) ([e3058db](https://github.com/dominikhaska/codegen/commit/e3058db6d7f4feef4780df6a5f1772e05b82571a))
* change the case of the flag manifest to camel case. ([#19](https://github.com/dominikhaska/codegen/issues/19)) ([fbac8ce](https://github.com/dominikhaska/codegen/commit/fbac8ce70dda766aff437b59286beb0579aa8472))
* embed flag manifest schema into code and moves files around ([#18](https://github.com/dominikhaska/codegen/issues/18)) ([aa9d3b0](https://github.com/dominikhaska/codegen/commit/aa9d3b03f0ece5295f6ce7be1f9093ed8ee9200f))

## [0.1.8](https://github.com/open-feature/cli/compare/v0.1.7...v0.1.8) (2024-10-31)


### 🐛 Bug Fixes

* docker publishing ([c663816](https://github.com/open-feature/cli/commit/c663816e33d0a020c1bd4db110ac0e4f451ff7b1))

## [0.1.7](https://github.com/open-feature/cli/compare/v0.1.6...v0.1.7) (2024-10-31)


### 🐛 Bug Fixes

* docker publishing ([2d24d51](https://github.com/open-feature/cli/commit/2d24d5141c0822edb7254f38efdabaa6e9b5b351))

## [0.1.6](https://github.com/open-feature/cli/compare/v0.1.5...v0.1.6) (2024-10-31)


### 📚 Documentation

* switch from code gen to cli ([#47](https://github.com/open-feature/cli/issues/47)) ([7a1f9f3](https://github.com/open-feature/cli/commit/7a1f9f304cc9c512b407b19986fbd82e3b80fe53))


### 🔄 Refactoring

* change name of go module ([#46](https://github.com/open-feature/cli/issues/46)) ([e3058db](https://github.com/open-feature/cli/commit/e3058db6d7f4feef4780df6a5f1772e05b82571a))

## [0.1.5](https://github.com/open-feature/codegen/compare/v0.1.4...v0.1.5) (2024-10-22)


### 🐛 Bug Fixes

* container copy command ([#40](https://github.com/open-feature/codegen/issues/40)) ([8448543](https://github.com/open-feature/codegen/commit/8448543fda56a3d68851cf44a4735c1902bf5b98))

## [0.1.4](https://github.com/open-feature/codegen/compare/v0.1.3...v0.1.4) (2024-10-22)


### ✨ New Features

* add version command ([#38](https://github.com/open-feature/codegen/issues/38)) ([c13a448](https://github.com/open-feature/codegen/commit/c13a4486b9b42f3e4a6f34abd43a87aecf91355e))

## [0.1.3](https://github.com/open-feature/codegen/compare/v0.1.2...v0.1.3) (2024-10-22)


### 🐛 Bug Fixes

* set github token for release process ([a2fe4aa](https://github.com/open-feature/codegen/commit/a2fe4aa33e380e86925480e7233eeed4bfb9ed90))

## [0.1.2](https://github.com/open-feature/codegen/compare/v0.1.1...v0.1.2) (2024-10-22)


### ✨ New Features

* add basic react support ([#31](https://github.com/open-feature/codegen/issues/31)) ([757ab66](https://github.com/open-feature/codegen/commit/757ab66b7fde7103ca6f5cb7f10c0632073b58d8))


### 🔄 Refactoring

* change folder, package structure; integrate with cobra ([#27](https://github.com/open-feature/codegen/issues/27)) ([850c694](https://github.com/open-feature/codegen/commit/850c694c84fad1a71722a1b1e620f1473bc2d2ab))

## [0.1.1](https://github.com/open-feature/codegen/compare/v0.1.0...v0.1.1) (2024-10-04)


### 🐛 Bug Fixes

* release permissions ([#25](https://github.com/open-feature/codegen/issues/25)) ([dc07cdf](https://github.com/open-feature/codegen/commit/dc07cdfe5487c0a22209c54d0ee195bbdcf1b5ed))

## [0.1.0](https://github.com/open-feature/codegen/compare/v0.0.1...v0.1.0) (2024-10-04)


### ⚠ BREAKING CHANGES

* lower json schema version, rename number to float ([#12](https://github.com/open-feature/codegen/issues/12))

### ✨ New Features

* Fixing problems with generated code for golang and adding sample manifest for testing. ([#22](https://github.com/open-feature/codegen/issues/22)) ([558e964](https://github.com/open-feature/codegen/commit/558e9640b8756e9cacccfdb23f136d95bd81629b))
* initial CLI for codegen with support for golang strongly typed accessors ([#13](https://github.com/open-feature/codegen/issues/13)) ([e8f3d3e](https://github.com/open-feature/codegen/commit/e8f3d3ea2815b7d5473746e71f1bedc856e723c8))
* lower json schema version, rename number to float ([#12](https://github.com/open-feature/codegen/issues/12)) ([ed844b4](https://github.com/open-feature/codegen/commit/ed844b43a3d05113b49b39a1e368d0ee3c308dc9))


### 📚 Documentation

* Add initial flag manifest schema ([#9](https://github.com/open-feature/codegen/issues/9)) ([fac35ca](https://github.com/open-feature/codegen/commit/fac35caff88e1ef9a9c5ff1e8624040d91db9307))


### 🔄 Refactoring

* change the case of the flag manifest to camel case. ([#19](https://github.com/open-feature/codegen/issues/19)) ([fbac8ce](https://github.com/open-feature/codegen/commit/fbac8ce70dda766aff437b59286beb0579aa8472))
* embed flag manifest schema into code and moves files around ([#18](https://github.com/open-feature/codegen/issues/18)) ([aa9d3b0](https://github.com/open-feature/codegen/commit/aa9d3b03f0ece5295f6ce7be1f9093ed8ee9200f))
