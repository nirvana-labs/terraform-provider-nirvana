resource "nirvana_networking_vpc" "example_networking_vpc" {
  name = "my-vpc"
  region = "amsterdam"
  subnet_name = "my-subnet"
}
