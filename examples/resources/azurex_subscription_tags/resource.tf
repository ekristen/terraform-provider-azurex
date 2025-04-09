resource "azurex_subscription_tags" {
  tags = {
    "Environment" = "Production"
    "Owner"       = "DevOps Team"
  }
}