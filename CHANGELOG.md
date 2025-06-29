# Changelog

## 1.4.4 (2025-06-28)

Full Changelog: [v1.4.3...v1.4.4](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.4.3...v1.4.4)

### Bug Fixes

* **ci:** release-doctor — report correct token name ([a2fc9e3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a2fc9e3cf36c645f7ad3013a1f42c0b1b9300646))
* null nested attribute decoding ([51df8c7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/51df8c753c1906bfac73349e8061274387d8b5b1))


### Chores

* **ci:** only run for pushes and fork pull requests ([5a4afa7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5a4afa7b1d77385dd5f095f64296d2becac1096f))

## 1.4.3 (2025-06-18)

Full Changelog: [v1.4.2...v1.4.3](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.4.2...v1.4.3)

### Chores

* **internal:** codegen related update ([f1abd5b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f1abd5b63703d435fa594236df2140274d3b69c6))

## 1.4.2 (2025-06-17)

Full Changelog: [v1.4.1...v1.4.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.4.1...v1.4.2)

### Chores

* **ci:** enable for pull requests ([3ca9943](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/3ca9943dae47474444d10886e7208548df5a165f))

## 1.4.1 (2025-06-04)

Full Changelog: [v1.4.0...v1.4.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.4.0...v1.4.1)

### Chores

* **internal:** version bump ([b6d18cd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b6d18cd5186fa5afd44ec485bfa935df84331fa0))

## 1.4.0 (2025-06-04)

Full Changelog: [v1.3.2...v1.4.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.3.2...v1.4.0)

### Features

* **api:** api update ([66ec9df](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/66ec9df1d9b9afe7643c43059469dec7a739edf4))
* **api:** api update ([4da6373](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4da6373c22fbd38532dd0948f0eaccdff6b77c7f))


### Bug Fixes

* **resource/vm:** remove check on public_ip_enabled since it's not returned by API ([#187](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/187)) ([0c337a0](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/0c337a07f99565e5d98834f323b0194af6efdebf))


### Chores

* **internal:** codegen related update ([33225f3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/33225f31df972579f2761398763b47620de1ba79))
* **internal:** codegen related update ([83c7a1f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/83c7a1f12f08b847f8e0a526b92241a2785a350d))
* **internal:** codegen related update ([7a46906](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7a46906b6c43de86dc75d5ef559c63cbfb00f194))

## 1.3.2 (2025-05-30)

Full Changelog: [v1.3.1...v1.3.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.3.1...v1.3.2)

### Chores

* **internal:** version bump ([a05a49d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a05a49da72e1a4c019c10325c2879c6ae14b19b4))

## 1.3.1 (2025-05-30)

Full Changelog: [v1.3.0...v1.3.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.3.0...v1.3.1)

### Chores

* bump deps to avoid GetResourceIdentitySchemas errors for Terraform CLI v1.12+ ([849dc18](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/849dc18bdb0ab2e14f8905504ab3ce0e18af3f39))

## 1.3.0 (2025-05-29)

Full Changelog: [v1.2.1...v1.3.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.2.1...v1.3.0)

### Features

* **client:** support environments property from Stainless config ([9509b09](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/9509b09ac5d5120672c02258181713a759e52c01))

## 1.2.1 (2025-05-23)

Full Changelog: [v1.2.1...v1.2.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.2.1...v1.2.1)

### Features

* **docs:** update base path ([c1312d3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c1312d38dcdeb0e6e827919f7892bab7ad6fe184))

## 1.2.1 (2025-05-23)

Full Changelog: [v1.2.0...v1.2.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.2.0...v1.2.1)

### Features

* **api:** add availability resources for VM and Volumes ([9285abc](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/9285abca1792182f940478f1aeb9f77498c8dc1b))
* **api:** api update ([25f3eb3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/25f3eb31c02719b36617e21293103a0b3f717443))


### Bug Fixes

* **build:** enable building against private Go production repos ([b0440fd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b0440fd06918b959a2844c0b4eacef5f3219ec2f))


### Chores

* **docs:** grammar improvements ([f7528ae](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f7528ae7bd895c51fe073cc76de94c9dd4b376fc))

## 1.2.0 (2025-05-16)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.1.0...v1.2.0)

### Features

* **api:** api update ([387f507](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/387f507fe9272507dbebb47ee7d501ac13dcde81))


### Bug Fixes

* **release:** update README and version correctly in release PRs ([07676b7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/07676b7c43c50d836f04545aa88c47f209036cc6))

## 1.1.0 (2025-05-13)

Full Changelog: [v1.0.7...v1.1.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.7...v1.1.0)

### Features

* **api:** api update ([c481bf6](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c481bf66d94006993f1c7ea36ce79d21de8b472c))
* **api:** manual updates ([2ba3d0f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2ba3d0f6698827b13e1e1befb20edd7f8feacba0))
* **api:** manual updates ([4701ea6](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4701ea6279454dbb5026c770136bcf3d2711b1b2))


### Bug Fixes

* only unmarshal attributes that exist on the read response schema during refresh ([f420da5](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f420da566b4bab48174dcb9b94b3232d773ec60d))


### Chores

* **build:** update go.mod indirect dependencies ([9e31d01](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/9e31d013a74a5252db5e6321717f0eda198e78d8))
* **internal:** codegen related update ([64d2ac6](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/64d2ac66434398b485fd8b0fdb93cd11f0b9250a))

## 1.0.7 (2025-05-06)

Full Changelog: [v1.0.6...v1.0.7](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.6...v1.0.7)

### Bug Fixes

* fix caching issue between Unmarshal and UnmarshalComputed ([fea1c6f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/fea1c6f541df3fef170fa7a9db861db5a6a3e6e2))

## 1.0.6 (2025-05-01)

Full Changelog: [v1.0.5...v1.0.6](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.5...v1.0.6)

### Chores

* **internal:** codegen related update ([2a807a2](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2a807a2e26bcd823a174de0f489aa02ba86553c8))

## 1.0.5 (2025-04-26)

Full Changelog: [v1.0.4...v1.0.5](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.4...v1.0.5)

### Bug Fixes

* **build:** do not fail if go mod tidy fails during bootstrapping ([30f4051](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/30f4051ebe54bf3933d1bd9e1f1fcfcc6161f1e2))


### Chores

* **internal:** codegen related update ([de4991c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/de4991cfc6f6e9b4d9c07f08a96f5665c5d1eb9b))

## 1.0.4 (2025-04-25)

Full Changelog: [v1.0.3...v1.0.4](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.3...v1.0.4)

### Chores

* **ci:** run on more branches ([c62c9c0](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c62c9c0489245bcaa090821ee294e074835d9694))

## 1.0.3 (2025-04-17)

Full Changelog: [v1.0.2...v1.0.3](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.2...v1.0.3)

### Chores

* **internal:** codegen related update ([97bccfa](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/97bccfa832306a77e5a7ef70d5d3467fc8c7ba8b))

## 1.0.2 (2025-04-15)

Full Changelog: [v1.0.1...v1.0.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.1...v1.0.2)

### Features

* **api:** api update ([dc0d3d3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/dc0d3d3676e1bfb78697f1c2720538c2dcf3733a))


### Documentation

* update documentation links to be more uniform ([bbf40d2](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/bbf40d2af4e13fd4b5d6d258f8a47dc80b668851))

## 1.0.1 (2025-04-13)

Full Changelog: [v1.0.0...v1.0.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v1.0.0...v1.0.1)

### Chores

* **internal:** version bump ([152a49d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/152a49d46bb0e2bb385d6ff537addceedbf8c476))

## 1.0.0 (2025-04-13)

Full Changelog: [v0.7.0...v1.0.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.7.0...v1.0.0)

### Features

* **api:** api update ([befd2f8](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/befd2f87ec8a6a96ec12231914e99fb00e447534))

## 0.7.0 (2025-04-12)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.6.1...v0.7.0)

### Features

* **api:** api update ([ddba6d2](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/ddba6d23f2efd916e86e3f7e7d34facb68a537ad))
* **api:** api update ([1c3d99c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1c3d99c4b0a16374219db9897e971f1466e77d35))
* **api:** api update ([8560e43](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8560e4359bb496dc8b5bc742c0c760f9dc7de773))
* **api:** api update ([937525b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/937525b3f624a100d5f383320f5e1467ae20022b))

## 0.6.1 (2025-04-11)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.6.0...v0.6.1)

### Chores

* **internal:** codegen related update ([71838eb](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/71838ebe3bc1868ff5cf049e4a9be269681b6b22))

## 0.6.0 (2025-04-10)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([200b2c4](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/200b2c4d9559c7eac1c4f1171c5053d09dc89034))
* **api:** api update ([b65702c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b65702caabe77362e02b9c93de597959bef5e58f))


### Chores

* configure new SDK language ([002bf8c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/002bf8c49d2e9b8e96beb0dfa237d286be04a311))

## 0.5.0 (2025-04-09)

Full Changelog: [v0.4.6...v0.5.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.6...v0.5.0)

### Chores

* **internal:** version bump ([74b9cb1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/74b9cb1a14a84d15b72e1229e33c8b8c3058cbe7))

## 0.4.6 (2025-04-09)

Full Changelog: [v0.4.5...v0.4.6](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.5...v0.4.6)

### Features

* **api:** api update ([4fa4a27](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4fa4a27e456e7b4c12a28ee0d6c1b9a7b1fbebb6))

## 0.4.5 (2025-04-09)

Full Changelog: [v0.4.4...v0.4.5](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.4...v0.4.5)

### Chores

* **tests:** improve enum examples ([#160](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/160)) ([1d93015](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1d930158a897c32a53bf54db6597e0168fd628cc))

## 0.4.4 (2025-04-08)

Full Changelog: [v0.4.3...v0.4.4](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.3...v0.4.4)

### Chores

* **build:** scripts/format should not fail if generate-docs fails ([#158](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/158)) ([7bcdac2](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7bcdac20e73f9a5dc4dd8dfc0c8efef346e8c1f8))

## 0.4.3 (2025-04-07)

Full Changelog: [v0.4.2...v0.4.3](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.2...v0.4.3)

### Chores

* **internal:** version bump ([#156](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/156)) ([5ddb287](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5ddb28704df5ab59cc40e6f5131e52831dd39911))

## 0.4.2 (2025-04-07)

Full Changelog: [v0.4.1...v0.4.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.1...v0.4.2)

### Features

* **api:** api update ([#154](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/154)) ([2290869](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2290869a3b5215a916d91d45dfbccfc09ca0e7c9))

## 0.4.1 (2025-04-07)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.4.0...v0.4.1)

### Chores

* **internal:** version bump ([#152](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/152)) ([97fabfd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/97fabfdb74a14633aaff06a12ad4332008dc7df2))

## 0.4.0 (2025-04-07)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([#150](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/150)) ([f632e2a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f632e2a393d69545412b168b02e6e827728cccbf))

## 0.3.0 (2025-04-04)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.2.1...v0.3.0)

### Features

* **api:** api update ([#146](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/146)) ([6a33051](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6a3305118db366a6f45b088be0c6dc142f012eeb))
* **api:** api update ([#148](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/148)) ([4e6a326](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4e6a3264a8ddb0b1b9e0a72afa53136f571b1e7c))
* **api:** api update ([#149](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/149)) ([fb18909](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/fb18909532cf9ca05747aed3ed95f9852ae16b03))

## 0.2.1 (2025-04-04)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.2.0...v0.2.1)

### Chores

* **internal:** version bump ([#144](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/144)) ([bacfd4a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/bacfd4ad31b29fee3f75e086709980fc614797c3))

## 0.2.0 (2025-04-03)

Full Changelog: [v0.1.4...v0.2.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.4...v0.2.0)

### Features

* **api:** api update ([#143](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/143)) ([d610d25](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d610d250bdc4965ec3a406ef550cd5ab44c5a876))


### Chores

* **internal:** codegen related update ([#141](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/141)) ([7dc2e9e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7dc2e9e3c1d7ad536abd48b9a9b41454f7c36996))

## 0.1.4 (2025-03-27)

Full Changelog: [v0.1.3...v0.1.4](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.3...v0.1.4)

### Chores

* **internal:** codegen related update ([#138](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/138)) ([625eb0e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/625eb0ec2715f2669d3d022d7fc538344f290b8c))
* **internal:** codegen related update ([#140](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/140)) ([687ed88](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/687ed8898d3b7a80b2c26225ee4064865f2ffbb5))

## 0.1.3 (2025-03-25)

Full Changelog: [v0.1.2...v0.1.3](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.2...v0.1.3)

### Chores

* **internal:** codegen related update ([#135](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/135)) ([815a783](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/815a783655d41534335e9eac1253d1c4ace61b44))
* update dependencies ([#137](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/137)) ([6db17d9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6db17d9a75c8a976fc7114d13897149e261f8abb))

## 0.1.2 (2025-03-22)

Full Changelog: [v0.1.1...v0.1.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.1...v0.1.2)

### Bug Fixes

* **api:** better support for environment variables as provider properties ([#133](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/133)) ([109aad9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/109aad94377f0a9bf89724a514bf782277207332))

## 0.1.1 (2025-03-22)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0...v0.1.1)

### Chores

* comment - trigger new release ([#131](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/131)) ([5bfc23e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5bfc23efa10bc57260ad605317d09f5d05c3c466))

## 0.1.0 (2025-03-21)

Full Changelog: [v0.1.0-alpha.15...v0.1.0](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.15...v0.1.0)

### Bug Fixes

* **resource/compute_vm:** fetch boot volume size and preserve values which are not returned by the API response for os_image_name and ssh_key ([#128](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/128)) ([3d47908](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/3d479084d3eabd6305ddbc59b133b71bd0b9c680))
* simplify error checking ([#130](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/130)) ([d7681ac](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d7681acffb0bc629461874fceac167ebb40f0c64))

## 0.1.0-alpha.15 (2025-03-21)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **api:** api update ([#126](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/126)) ([702e430](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/702e43075b5b09ef061e6c977ae4dd8dd4de62f9))

## 0.1.0-alpha.14 (2025-03-21)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* **api:** api update ([#124](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/124)) ([8453cc8](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8453cc8fb7c45702b2cfd6ad11d8524a0af25968))

## 0.1.0-alpha.13 (2025-03-21)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Chores

* **internal:** version bump ([#122](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/122)) ([268b6cd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/268b6cd5d9a58d85e96e828a65adb3ddf78ab51f))

## 0.1.0-alpha.12 (2025-03-21)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** api update ([#120](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/120)) ([030ffbc](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/030ffbc6181491d1be1bae450c018d01fef18bd6))
* **api:** api update ([#121](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/121)) ([bf9907b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/bf9907b642388b291c4d72ecd43368ccdbc6c17e))


### Chores

* **internal:** codegen related update ([#118](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/118)) ([2f71e4f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2f71e4f07020fe7ef88d242d5553ad377eacbabd))

## 0.1.0-alpha.11 (2025-03-19)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* **api:** api update ([#115](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/115)) ([393aa9f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/393aa9f552c7c036426f77199d9a84f8cfa499c8))
* **api:** api update ([#117](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/117)) ([fd7b63c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/fd7b63c0604de24119e08c2f8094f8a11e57560b))
* handle responses that differ from request shape ([#114](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/114)) ([4868e39](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4868e399d2bcdc26a8d924f093eb5e2449fd8adc))


### Bug Fixes

* **docs:** generate docs by default ([#116](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/116)) ([2b65a60](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2b65a60aef7e8b83329641f3a91694508857401d))


### Chores

* **internal:** version bump ([#112](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/112)) ([dd2a80e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/dd2a80edf3566b06be72cc20c4ac4539916ce52d))

## 0.1.0-alpha.10 (2025-03-18)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** api update ([#105](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/105)) ([77ddc2a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/77ddc2ad10ce44e8e3b27214163877310aba80e4))
* **api:** api update ([#110](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/110)) ([2098820](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/209882066ee1f9e012ece4f5915fefdb2d3399c4))
* **api:** api update ([#111](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/111)) ([f720bf3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f720bf3a4610b2d20a17d0f393987a5f1b57d6bd))

## 0.1.0-alpha.9 (2025-03-17)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Chores

* trigger rebuild of terraform ([#106](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/106)) ([5dc6454](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5dc645433201c4db62f4e40d23c6ea25438cbb2f))

## 0.1.0-alpha.8 (2025-03-17)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** api update ([#104](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/104)) ([8f7d601](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8f7d60146cb20c2fdf159e08e5758129b085a40a))


### Chores

* **internal:** version bump ([#102](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/102)) ([05dd5e9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/05dd5e9bc49d375697cbc7bd183e61ead158cd43))

## 0.1.0-alpha.7 (2025-03-17)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** api update ([#101](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/101)) ([5bd326d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5bd326d3be40862944d7e86b98d9e31e4039bee3))


### Chores

* **internal:** codegen related update ([#99](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/99)) ([03fd206](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/03fd2068d91c125c5d896a33b401eb74de4e2c3c))

## 0.1.0-alpha.6 (2025-03-14)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** api update ([#98](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/98)) ([44932a4](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/44932a41d7e2bbeeef94724b21cdc9a036516ef4))


### Chores

* **internal:** version bump ([#96](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/96)) ([bfefda1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/bfefda1cb5c406ee4c4da051c988277fee815b47))

## 0.1.0-alpha.5 (2025-03-13)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** api update ([#94](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/94)) ([8afd114](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8afd114d6dca106e55df3fb111fcb1017c781919))
* **api:** api update ([#95](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/95)) ([65491fb](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/65491fb388ec7116aa0de940bff12af56ccd64de))
* **api:** manual updates ([#92](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/92)) ([a84db6d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a84db6d9f293a240faf94b2a15eb66a39a6cd06d))

## 0.1.0-alpha.4 (2025-03-13)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Chores

* **internal:** version bump ([#90](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/90)) ([38e5e3d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/38e5e3d850ab3e0bac055d5dcea6ec4195558dda))

## 0.1.0-alpha.3 (2025-03-13)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([#87](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/87)) ([f31f7a8](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f31f7a8839624d2bd81e3130a29e37bd462eab0d))
* **resource/networking_vpc:** update Update method to use operation ([#89](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/89)) ([44f16bd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/44f16bd2c77796bcda7efe2acc60b9b61f974fde))

## 0.1.0-alpha.2 (2025-03-13)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* add docs generation to format script ([#80](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/80)) ([859a18e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/859a18ea52c6da609b0b61378208bffdcc69630e))
* add SKIP_BREW env var to ./scripts/bootstrap ([#79](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/79)) ([d518a4c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d518a4c00b03a361417c4a30e5d0a726ac83ef4c))
* add support for wait for an operation after create/update/delete ([#71](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/71)) ([380c6cc](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/380c6cc1708853681347b1f84e63eaa87c86416c))
* **api:** api update ([#83](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/83)) ([c5e1dcb](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c5e1dcb3b0fef13de3038ff33c33725e2c8303f8))
* **api:** enable doc generation ([#84](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/84)) ([8274111](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8274111318681edd625050456ef0279181e4cde5))
* **api:** manual updates ([#72](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/72)) ([8eb7edf](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8eb7edfccf70bbb944fd5aa0a57677657886f558))
* **api:** manual updates ([#74](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/74)) ([032eff1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/032eff19b70fdca60d3f8d9577f7de1561066b76))
* **api:** manual updates ([#75](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/75)) ([5db684f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5db684fd1a6bc9c8663e2e0685db1b794a6bea78))
* **api:** manual updates ([#76](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/76)) ([d6da4c9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d6da4c981d5b9828f259f9b16439cc029ca2f915))
* **api:** manual updates ([#78](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/78)) ([1893abb](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1893abbc52170e22c447764d22e144bf816cbd86))


### Chores

* **internal:** codegen related update ([#82](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/82)) ([de5b1e7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/de5b1e77aaceec0eb6b442721df4e991304f6799))
* **internal:** codegen related update ([#86](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/86)) ([1d1e036](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1d1e036ecbf9ea6fc06ad3e8181c5052884e9647))

## 0.1.0-alpha.1 (2025-03-10)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/nirvana-labs/terraform-provider-nirvana/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* add doc string to specify what legal terraform values are for enums ([#56](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/56)) ([e0c2d79](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/e0c2d79e36c32a9e9c9ee1a2a04348825e8aa997))
* **api:** api update ([85c74cd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/85c74cd6143e91762f53db4468232aefaa827d32))
* **api:** api update ([d1bac0c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d1bac0cff85a825a65dac1dab82cc61e638f174a))
* **api:** api update ([326a6f5](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/326a6f51a7e2de05021acdea5338aa43a6809e82))
* **api:** api update ([b98cd84](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b98cd84e2aab7de6ed954d2d7c0e71643567a8fe))
* **api:** api update ([1b922d2](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1b922d26dfeb77220e49cd4b22a1a2123a087129))
* **api:** api update ([6c7e85b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6c7e85b4373bf1115d76fa72a0360b7e8db770a2))
* **api:** api update ([c8f7c2e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c8f7c2e6f89d6397a02d746bc4dc47982517f3a8))
* **api:** api update ([57085f1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/57085f13b59aa30827f95909d2b0cad6a545b4f6))
* **api:** api update ([e065f8a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/e065f8aa938f78d8df8c9cd209994e85248beacc))
* **api:** api update ([8c913cc](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8c913cc87765b4e818e0829c95ede77d791a87f3))
* **api:** api update ([#1](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/1)) ([5a2aadf](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5a2aadf386d7242aa1b20d6b83a59241b2454f2a))
* **api:** api update ([#10](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/10)) ([00a6f90](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/00a6f90a059dd79cf2c39a403015cfa552a05c54))
* **api:** api update ([#12](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/12)) ([cd87f0b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/cd87f0be395b5d7b78792c203cffe44463603308))
* **api:** api update ([#13](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/13)) ([4f02052](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4f02052f5c77af4a848a9c1bddba2d49b4e9f503))
* **api:** api update ([#14](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/14)) ([6c589b9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6c589b975ba71d9ddc9765cfb35bcf1fd2c6d1be))
* **api:** api update ([#17](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/17)) ([1a8c5e5](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1a8c5e5888c0124cac07874e1eb88c60d1adc183))
* **api:** api update ([#2](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/2)) ([cfb4f2a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/cfb4f2a101167dc4c273476bbdbe7ad32454092b))
* **api:** api update ([#20](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/20)) ([ab4225a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/ab4225a6a8314f40be8b9eee5c74c14a93bbdc28))
* **api:** api update ([#21](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/21)) ([01f562f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/01f562f9f1653dd253e71771740d8fd0abfd916e))
* **api:** api update ([#3](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/3)) ([34103f1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/34103f1f3a9353dee7566dfcb54a18df628500f5))
* **api:** api update ([#32](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/32)) ([e27f158](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/e27f15833a742a966c893ba7381f5973efb24a3b))
* **api:** api update ([#34](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/34)) ([9c59f47](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/9c59f4738b8cfa9d2d37e9dffb936c4d434d5bac))
* **api:** api update ([#37](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/37)) ([d481deb](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d481deb3598409f28b4a7a3d29346d45b0a753fc))
* **api:** api update ([#38](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/38)) ([28f73c1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/28f73c1f3162d17a541426685177bc5d0ca5b88d))
* **api:** api update ([#39](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/39)) ([4386cb7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4386cb7bbb79b482611720b721f0fda27af14442))
* **api:** api update ([#4](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/4)) ([5ed2e8c](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5ed2e8c03d75dd239fe145797d5372ad5ccc3b6c))
* **api:** api update ([#43](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/43)) ([00dfa77](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/00dfa77468e2735bdb290e222e03e008e6c158b7))
* **api:** api update ([#44](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/44)) ([664139f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/664139f822d13286daf6d4e35bcc76bf122a5afb))
* **api:** api update ([#46](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/46)) ([8b13b5b](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8b13b5b84a2beb329eb80ff2ddc9023265db6f48))
* **api:** api update ([#49](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/49)) ([4723c5a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4723c5a75db11786cea8cb95e12f39a6815f32cf))
* **api:** api update ([#5](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/5)) ([b657551](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b657551b38b587c4459f6a88da70754fa0f11352))
* **api:** api update ([#6](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/6)) ([c9fcef9](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c9fcef95830a27d5acb7c0a88f6dfea5077ba151))
* **api:** api update ([#61](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/61)) ([8dcc651](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8dcc651f37d86fd346d409b48cb79061554f3efd))
* **api:** api update ([#62](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/62)) ([90781b6](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/90781b6a1a35b65fb2e46c9d2eb21812a28e70b9))
* **api:** api update ([#64](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/64)) ([3017d9e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/3017d9e2633b4109673d50f895b41292890400e5))
* **api:** api update ([#65](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/65)) ([351a323](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/351a323139f39e4e264880678403fdfd98926502))
* **api:** api update ([#7](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/7)) ([ba06c9d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/ba06c9d45561c2a901952283231cae381fb5fe80))
* **api:** api update ([#8](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/8)) ([739dd01](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/739dd012d9129679565eb3c6ebd707cc6db70bf5))
* **api:** api update ([#9](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/9)) ([9b29291](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/9b2929142c489528fd91291703d3e58ca770d077))
* **api:** manual updates ([#22](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/22)) ([c6d953f](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c6d953f5882b133a93878157d0615d784ea66087))
* **api:** manual updates ([#31](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/31)) ([d09a9f3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/d09a9f35bf90aa69952c1572628f6e463519fe61))
* **api:** manual updates ([#35](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/35)) ([5312d98](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5312d98a22e91290dd9a3d0b0ea6c3c6f42586ca))
* **api:** manual updates ([#50](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/50)) ([5eb18fc](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/5eb18fcc2a5d68cd54a613837f4323004f4d2c0b))
* **api:** manual updates ([#52](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/52)) ([14e8bf3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/14e8bf30e90e8dba7c00b78d428176ee59a9812a))
* **api:** manual updates ([#67](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/67)) ([2506d57](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/2506d572120278d2b93d7c159d2c9e360c249943))
* **api:** manual updates ([#68](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/68)) ([6b54a7a](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6b54a7af928dd891e0296d96488fbf94dfc15608))
* **api:** update via SDK Studio ([1750182](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/1750182636ca15bad92c452eac69c65b98ed9e45))
* **build:** allow for building against private go repos ([#29](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/29)) ([4ee8a62](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4ee8a62c794cff4668e755d4c6a589779e2a10fd))
* **docs:** improve readme ([#19](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/19)) ([a513344](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a513344b40c23c352888f60b6578b0f56a0d136d))
* support using environment variables as provider attributes ([f5cd506](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f5cd506918de8cc5a47e0687d4b54b565e4e9d98))


### Bug Fixes

* **api:** improve drift detection for resources ([#26](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/26)) ([a8a6797](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a8a6797a502f80a7d6c5a512d58ea44cfea2569d))
* **api:** improve drift detection for resources ([#48](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/48)) ([a3eb4af](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a3eb4af641f18294861109195a530ba44561adf7))
* **build:** ensure scripts/generate-docs works regardless of PATH ([#16](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/16)) ([584a779](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/584a779082afcf7a7bde8cb6c3e50a31ab3a053d))
* **build:** improve release process ([#15](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/15)) ([82d79cf](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/82d79cf6dd550f2c9e3e1a0e0969e8f4510ff887))
* **datasource:** improve configurability of path parameters on data sources ([#45](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/45)) ([96d5ce7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/96d5ce732a6f28e1373652401a59b0ee9af316c6))
* do not call path.Base on ContentType ([#41](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/41)) ([605027e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/605027ea14774913937542ccfeb42cbb6fa10002))
* **docs:** skip tfplugindocs generation if `generate_docs` is false ([#25](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/25)) ([00ff134](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/00ff134ee8cecb45f61e02d810350bae5cda0a97))


### Chores

* casing change in doc string ([#57](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/57)) ([7f1f51d](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7f1f51d5551e3f6265b06fb3646aa2ed5b55aa92))
* **internal:** bumps go dependencies ([0b382bd](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/0b382bd0d619f2c9b56711eebacd0f7a7196ad99))
* **internal:** codegen related update ([52c7920](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/52c79205ddb57620c66b405e3947d617a01691b7))
* **internal:** codegen related update ([#23](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/23)) ([a3de893](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/a3de8933276859d5cd230d89d5baf960ede8b5d6))
* **internal:** codegen related update ([#24](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/24)) ([54a2fa6](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/54a2fa69a9a96c0846ea007b767669036083c7a9))
* **internal:** codegen related update ([#27](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/27)) ([bd81db4](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/bd81db43653ed9531c4a78856e38fffe61cc15b1))
* **internal:** codegen related update ([#28](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/28)) ([b045b92](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/b045b926300f5552baf068cf921f1f312fd45bd6))
* **internal:** codegen related update ([#30](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/30)) ([fc77254](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/fc77254e1520006ce40758e59a7ce9b5e4e14d8b))
* **internal:** codegen related update ([#33](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/33)) ([23762d1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/23762d1aa4b8121c95e232c6c03f7abfa3709329))
* **internal:** codegen related update ([#36](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/36)) ([415c1d7](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/415c1d7bd9ff77c1127c7fe643d6ae100d155bd7))
* **internal:** codegen related update ([#40](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/40)) ([c2170c1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/c2170c1e8fe6dd399650dba861e3f8dd9c00ec95))
* **internal:** codegen related update ([#42](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/42)) ([eabc5be](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/eabc5be00e9af79427e86ff71b570cdde62a7e85))
* **internal:** codegen related update ([#47](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/47)) ([4b6651e](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/4b6651e249749a7dce745189b7d3c82a6bf34d6f))
* **internal:** codegen related update ([#51](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/51)) ([afddd65](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/afddd65a2b79e23752d9f665386392d9b9e99d41))
* **internal:** codegen related update ([#53](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/53)) ([f939587](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/f939587130fedd0026ef2c3d78c826b5e88e6807))
* **internal:** codegen related update ([#54](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/54)) ([17d3239](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/17d32397660d02d553ac25f2c2f5dc93ea52a45b))
* **internal:** codegen related update ([#55](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/55)) ([7f14f52](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7f14f52acd65acd5be25d6de289af7d8e2191f7b))
* **internal:** codegen related update ([#60](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/60)) ([8f07eda](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/8f07eda7aca3874bc1eceada160733f8a53b87fa))
* **internal:** codegen related update ([#63](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/63)) ([85a60db](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/85a60db611841e1b576517c02b56f0d9b380f010))
* simplify string literals ([#59](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/59)) ([336fbb1](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/336fbb1f35a7d6502e8899f0a39e7371a6e1d27c))
* update SDK settings ([db42895](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/db42895e99d0b3ef621bb16b9041d23c2a56beb9))
* update SDK settings ([#11](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/11)) ([ad145b3](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/ad145b3ab3aded3a2df2c70f60c35dcbdb21e37e))
* update SDK settings ([#18](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/18)) ([7b0bbaf](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/7b0bbaf49c0e9c25ed0ef91f9ca9c440fd0e91c9))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#58](https://github.com/nirvana-labs/terraform-provider-nirvana/issues/58)) ([6cf5270](https://github.com/nirvana-labs/terraform-provider-nirvana/commit/6cf5270c714f0f941555371c4116ac900633b334))
