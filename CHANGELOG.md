# Changelog

## [3.1.0](https://github.com/lgdd/lfr-cli/compare/v3.0.0...v3.1.0) (2024-05-02)


### Features

* add config command ([45a932c](https://github.com/lgdd/lfr-cli/commit/45a932cdd579b3f0b1a34e8068a2e65524d4f263))
* use github urls by default ([66281ea](https://github.com/lgdd/lfr-cli/commit/66281eaf43e1a806fe0ec9fd7c2fea2b34a26a98))


### Bug Fixes

* output printed twice when running 'lfr c cx' ([9857217](https://github.com/lgdd/lfr-cli/commit/9857217f20866b69148aaa88525674250f823916))
* wrong github urls for portal workspaces ([#17](https://github.com/lgdd/lfr-cli/issues/17)) ([4ca0436](https://github.com/lgdd/lfr-cli/commit/4ca0436d45030f914be27be38b0b1b7524dad85b))

## [3.0.0](https://github.com/lgdd/lfr-cli/compare/v2.0.0...v3.0.0) (2024-04-28)


### ⚠ BREAKING CHANGES

* rework client extension command & prompts

### Features

* add alternative github bundle urls for dxp ([3468c35](https://github.com/lgdd/lfr-cli/commit/3468c35f81ce378d304260d11fcca642819946c8))
* add config file for default flags ([2fc7ef9](https://github.com/lgdd/lfr-cli/commit/2fc7ef991baa0d8d47bc60bd9e480a893caffd1c))
* add github workflows into workspace ([ae2d650](https://github.com/lgdd/lfr-cli/commit/ae2d65024425557c459927fd8f2a339de345db06))
* add quarterly releases support ([3f30ef0](https://github.com/lgdd/lfr-cli/commit/3f30ef0b2a1455c1c1854733e71e0c1719d8f40f))
* fetch latest workspace plugin version ([244105e](https://github.com/lgdd/lfr-cli/commit/244105e4ed43ff076e39884d3cdc464b93584524))
* initiate git for workspace ([032208d](https://github.com/lgdd/lfr-cli/commit/032208d130a3af5ff1deacf5fdfe386caa94bfd0))
* make prompts accessible via config ([96f5261](https://github.com/lgdd/lfr-cli/commit/96f5261e1dffd1ab5d4c8c71c3a1fba701705f4c))


### Bug Fixes

* **ci:** wrong client extension samples path ([0d7e770](https://github.com/lgdd/lfr-cli/commit/0d7e770d8614c0779a2e9203e1879e2e77ffd3e9))
* maven bom for dxp ([963167a](https://github.com/lgdd/lfr-cli/commit/963167ad67d73dae8eba4cf42997557c6c7db5a2))
* offline workspace creation failure ([e0a61e4](https://github.com/lgdd/lfr-cli/commit/e0a61e49efc1629fd7e16d5d70055aa92698f955))
* wrong bom for dxp workspaces ([1b0b672](https://github.com/lgdd/lfr-cli/commit/1b0b6723b6936de6910da853286c9937a7003025))
* wrong docker base image in gradle workspaces ([8ee3e0b](https://github.com/lgdd/lfr-cli/commit/8ee3e0b01a867cffa72c14cb3efc1449d1e4fdac))
* wrong path for docker build dirs ([9f6f5b1](https://github.com/lgdd/lfr-cli/commit/9f6f5b1e6da04c5d1d346f6c9b8310d81e8d2677))


### Code Refactoring

* rework client extension command & prompts ([697e877](https://github.com/lgdd/lfr-cli/commit/697e877082ab8881df5a38f74a938ba9797348ae))

## [2.0.0](https://github.com/lgdd/liferay-cli/compare/v1.4.0...v2.0.0) (2024-04-20)


### ⚠ BREAKING CHANGES

* rename project to lfr-cli

### Bug Fixes

* **ci:** auto update assets merge issue ([62cfab5](https://github.com/lgdd/liferay-cli/commit/62cfab5c9d717af7965ac0573f5f2a64ae0f8b48))


### Code Refactoring

* rename project to lfr-cli ([75bb827](https://github.com/lgdd/liferay-cli/commit/75bb827379700e50a8ea25b302ce517868820924)), closes [#14](https://github.com/lgdd/liferay-cli/issues/14)
