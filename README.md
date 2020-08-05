# fake provider for pokemon-api

## commands

```bash
go mod init terraform-provider-pokemons
go mod vendor
go build -o terraform-provider-pokemons
mkdir -p ~/.terraform.d/plugins/jcanon.com/collec/pokemons/1.0/linux_amd64/
mv terraform-provider-pokemons ~/.terraform.d/plugins/jcanon.com/collec/pokemons/1.0/linux_amd64/terraform-provider-pokemons
```

## inspiration

https://learn.hashicorp.com/terraform/providers/setup-implement-read

https://github.com/hashicorp/terraform-provider-hashicups

