# Nirvana Labs Terraform Provider

The [Nirvana Labs Terraform provider](https://registry.terraform.io/providers/nirvana-labs/nirvana/latest/docs) provides convenient access to
the [Nirvana Labs REST API](https://docs.nirvanalabs.io) from Terraform.

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    nirvana = {
      source  = "nirvana-labs/nirvana"
      version = "~> 1.10.2"
    }
  }
}

# Initialize the provider
provider "nirvana" {
  api_key = "My API Key" # or set NIRVANA_LABS_API_KEY env variable
}

# Configure a resource
resource "nirvana_compute_vm" "example_compute_vm" {
  boot_volume = {
    size = 100
  }
  cpu_config = {
    vcpu = 2
  }
  memory_config = {
    size = 2
  }
  name = "my-vm"
  os_image_name = "ubuntu-noble-2025-04-03"
  public_ip_enabled = true
  region = "us-wdc-1"
  ssh_key = {
    public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2"
  }
  subnet_id = "123e4567-e89b-12d3-a456-426614174000"
  data_volumes = [{
    name = "my-data-volume"
    size = 100
  }]
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/nirvana-labs/nirvana/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property | Environment variable   | Required | Default value |
| -------- | ---------------------- | -------- | ------------- |
| api_key  | `NIRVANA_LABS_API_KEY` | true     | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/nirvana-labs/terraform-provider-nirvana/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
