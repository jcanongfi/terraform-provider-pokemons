terraform {
  required_providers {
    pokemons = {
      versions = ["1.0"]
      source = "jcanon.com/collec/pokemons"
    }
  }
}

provider "pokemons" {}

