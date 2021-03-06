package pokemons

import (
  "context"
  "strconv"
  
  hc "github.com/jcanongfi/pokemon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePokemon() *schema.Resource {
  return &schema.Resource{
    CreateContext: resourcePokemonCreate,
    ReadContext:   resourcePokemonRead,
    UpdateContext: resourcePokemonUpdate,
    DeleteContext: resourcePokemonDelete,
    Schema: map[string]*schema.Schema{
        "nom": &schema.Schema{
          Type:     schema.TypeString,
          Required: true,
    //      ForceNew: true,
        },
        "type": &schema.Schema{
          Type:     schema.TypeString,
          Required: true,
        },
    },

  }
}

func resourcePokemonCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*hc.Client)
  
    var diags diag.Diagnostics

    v_nom := d.Get("nom").(string)
    v_type := d.Get("type").(string)

    v_poke := hc.Pokemon{
      Nom: v_nom,
      Type: v_type,
    }

    o, err := c.CreatePokemon(v_poke)
    if err != nil {
        return diag.FromErr(err)
    }

    d.SetId(strconv.Itoa(o.ID))

    resourcePokemonRead(ctx, d, m)

    return diags
}




func resourcePokemonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  c := m.(*hc.Client)
  
  var diags diag.Diagnostics
  
  pokemonID := d.Id()
  
  pokemon, err := c.GetPokemon(pokemonID)
  if err != nil {
    return diag.FromErr(err)
  }

  if err := d.Set("nom", pokemon.Nom); err != nil {
    return diag.FromErr(err)
  }
  
  if err := d.Set("type", pokemon.Type); err != nil {
    return diag.FromErr(err)
  }
  
  return diags
}


func resourcePokemonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  c := m.(*hc.Client)

  pokemonID := d.Id()

// Todo : tester si seulement changement de "type"
//  if d.HasChange("nom") {
    v_nom := d.Get("nom").(string)
    v_type := d.Get("type").(string)

    v_poke := hc.Pokemon{
      Nom: v_nom,
      Type: v_type,
    }

    _, err := c.UpdatePokemon(pokemonID, v_poke)
    if err != nil {
      return diag.FromErr(err)
    }

    resourcePokemonRead(ctx, d, m)
//    d.Set("last_updated", time.Now().Format(time.RFC850))
//  }
  return resourcePokemonRead(ctx, d, m)
}




func resourcePokemonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  c := m.(*hc.Client)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  pokemonID := d.Id()

  err := c.DeletePokemon(pokemonID)
  if err != nil {
    return diag.FromErr(err)
  }

  // d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
  d.SetId("")

  return diags
}

