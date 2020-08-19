package pokemons
 
import (
  "context"

  "github.com/jcanongfi/pokemon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
 
// Provider -
func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "url": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("POKE_URL", nil),
      },
      "username": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("POKE_USERNAME", nil),
      },
      "password": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        Sensitive:   true,
        DefaultFunc: schema.EnvDefaultFunc("POKE_PASSWORD", nil),
      },
    },
    ResourcesMap: map[string]*schema.Resource{
        "pocketmonster": resourcePokemon(),
    },
    DataSourcesMap: map[string]*schema.Resource{
        "pocketmonster": dataSourcePokemons(),
    },
    ConfigureContextFunc: providerConfigure,
  }
}
 
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
  url := d.Get("url").(string)
  username := d.Get("username").(string)
  password := d.Get("password").(string)

  var diags diag.Diagnostics

  if (url != "") && (username != "") && (password != "") {
    c, err := pokemon.NewClient(&url, &username, &password)
    if err != nil {
      return nil, diag.FromErr(err)
    }

    return c, diags
  }

  if (username != "") && (password != "") {
    c, err := pokemon.NewClient(nil, &username, &password)
    if err != nil {
      return nil, diag.FromErr(err)
    }

    return c, diags
  }

  if (url != "") {
    c, err := pokemon.NewClient(&url, nil, nil)
    if err != nil {
      return nil, diag.FromErr(err)
    }

    return c, diags
  }

  c, err := pokemon.NewClient(nil, nil, nil)
  if err != nil {
    return nil, diag.FromErr(err)
  }

  return c, diags
}

