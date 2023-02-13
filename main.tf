terraform {
    required_version = ">= 0.12.0"
    required_providers {
        azurerm = {
            source  = "hashicorp/azurerm"
            version = "~>2.0.0"
        }
    }
}

provider "azurerm" {
    features {}
}
# provider "azapi" {}

// zenithはChatGPTがつけてくれた名前。頂点という意味らしい
resource "azurerm_resource_group" "zenith" {
    name     = "zenith-rg"
    location = "japaneast"
}

resource "azurerm_container_registry" "zenith" {
  name                     = "zenithcontaineracr"
  location                 = azurerm_resource_group.zenith.location
  resource_group_name      = azurerm_resource_group.zenith.name
  admin_enabled            = true
  sku                      = "Standard"
}

# resource "azapi_resource" "build_and_push" {
#   depends_on = [azurerm_container_registry.zenith]

#   script = <<EOT
#     #!/bin/bash
#     az acr build --image zenithacr.azurecr.io/zenithqueue:latest .
#     az acr build --image zenithacr.azurecr.io/zenithdequeue:latest .
#     az acr login --name zenithacr
#     docker push zenithacr.azurecr.io/zenithqueue:latest
#     docker push zenithacr.azurecr.io/zenithdequeue:latest
#   EOT

#   env = {
#     ACR_ID = azurerm_container_registry.zenith.id
#   }
# }