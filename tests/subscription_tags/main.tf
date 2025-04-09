# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "azurex" {
  subscription_id = "1b75aa77-e080-4f83-b893-564dc9df7de5"
}

resource "azurex_subscription_tags" "testing" {
  tags = {
    "tag1" = "value1"
    "tag2" = "value2"
  }
}

terraform {
  required_providers {
    azurex = {
      source  = "ekristen/azurex"
    }
  }
}
