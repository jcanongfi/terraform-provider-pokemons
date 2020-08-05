terraform {
  required_providers {
    pokemons = {
      versions = ["1.0"]
      source = "jcanon.com/collec/pokemons"
    }
  }
}

provider "pokemons" {}

#module "test1" {
#  source = "./mod_pokemon"
#
#  poke_name = "Giranta"
#}

#output "output-test1" {
#  value = module.test1.coffee
#}
