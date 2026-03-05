# Changelog

## [0.3.0](https://github.com/suzutan/sdcd-cli/compare/v0.2.0...v0.3.0) (2026-03-05)


### New Features

* move main package to cmd/sdcd for clean go install path ([6e3a466](https://github.com/suzutan/sdcd-cli/commit/6e3a466c3495191b16f3e28a389af94212ce8cc3))
* move main package to cmd/sdcd for clean go install path ([c6cdcb4](https://github.com/suzutan/sdcd-cli/commit/c6cdcb4fe773db12d7c95bcbed6c82e8cb135529))

## [0.2.0](https://github.com/suzutan/sdcd-cli/compare/v0.1.0...v0.2.0) (2026-03-05)


### New Features

* add build artifact command to download individual artifact files ([#7](https://github.com/suzutan/sdcd-cli/issues/7)) ([b07f214](https://github.com/suzutan/sdcd-cli/commit/b07f214cd117ce130c9c9e3cb0501ac889dd49fc))
* flatten auth context to top-level context command ([ef73344](https://github.com/suzutan/sdcd-cli/commit/ef733448920c7d364e518100108b6edada33bf28))
* initial release of sdcd CLI ([c0dcf4b](https://github.com/suzutan/sdcd-cli/commit/c0dcf4bc94d89acd3752b851cde1f13d2fcae529))
* rename get subcommand to view across all resources ([81ba329](https://github.com/suzutan/sdcd-cli/commit/81ba3290c720877adfc90dc6552768446817dbb3))
* rename job latest to job latest-build ([dc452b6](https://github.com/suzutan/sdcd-cli/commit/dc452b621b19490cecdd56af6b14c485d40a884c))


### Bug Fixes

* parse ZIP response from build artifacts endpoint ([#6](https://github.com/suzutan/sdcd-cli/issues/6)) ([8b98bf1](https://github.com/suzutan/sdcd-cli/commit/8b98bf1b708ea22c2b0a81deb967fa48042d6b88))
* remove unused secretUpdateHasAllow variable ([f87cc0e](https://github.com/suzutan/sdcd-cli/commit/f87cc0ed40b14ca33ff5ee90e4c8231902d55a7b))
* start build log pagination from page 1 instead of page 0 ([#4](https://github.com/suzutan/sdcd-cli/issues/4)) ([11ac610](https://github.com/suzutan/sdcd-cli/commit/11ac610b139dc6bdba39a71dd333fbb9b987370f))
* use from (line offset) instead of page for build log pagination ([#5](https://github.com/suzutan/sdcd-cli/issues/5)) ([5b3d8c1](https://github.com/suzutan/sdcd-cli/commit/5b3d8c1b858765fbd2ee16e44a3236b5e8328f41))
