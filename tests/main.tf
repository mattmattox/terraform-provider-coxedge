terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "GM3COPLOU6nOI12/NZ7HNg=="
}


data "coxedge_vpcs" "vpcs" {
  environment = "test-backend"
}

output "out_vpc" {
  value = data.coxedge_vpcs.vpcs
}
