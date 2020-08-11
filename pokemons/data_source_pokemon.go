package pokemons
 
import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "time"
  
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
 
func dataSourcePokemons() *schema.Resource {
  return &schema.Resource{
    ReadContext: dataSourcePokemonsRead,
    Schema: map[string]*schema.Schema{
      "pokemons": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
              Type:     schema.TypeInt,
              Computed: true,
            },
            "nom": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "type": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
          },
        },
      },
    },
  }
}


func dataSourcePokemonsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  client := &http.Client{Timeout: 10 * time.Second}
  
  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics
  
  req, err := http.NewRequest("GET", fmt.Sprintf("%s/pokemon", "http://localhost:8888"), nil)
  if err != nil {
    return diag.FromErr(err)
  }
  
  r, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer r.Body.Close()
  
  pokemons := make([]map[string]interface{}, 0)
  err = json.NewDecoder(r.Body).Decode(&pokemons)
  if err != nil {
    return diag.FromErr(err)
  }
  
  if err := d.Set("pokemons", pokemons); err != nil {
    return diag.FromErr(err)
  }
  
  // always run
  d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
  
  return diags
}

