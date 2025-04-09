resource "nirvana_networking_vpc" "example_networking_vpc" {
  name = "my-vpc"
  region = "us-wdc-1"
  subnet_name = "my-subnet"
}
