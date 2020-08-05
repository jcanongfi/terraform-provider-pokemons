terraform {
  required_providers {
    pokemon = {
      versions = ["1.0"]
      source = "jcanon.com/collec/pokemons"
    }
  }
}

variable "poke_name" {
  type    = string
  default = "Miniature"
}

data "mes_pokemons" "all" {}

output "all_pokemons" {
  value = data.mes_pokemons.all.pokemons
}

output "best" {
  value = {
    for toto in data.mes_pokemons.all.pokemons :
    toto.id => toto
    if toto.name == var.poke_name
  }
}
