package cyral

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCyralRepository() *schema.Resource {
        return &schema.Resource{
                Create: resourceCyralRepositoryCreate,
                Read:   resourceCyralRepositoryRead,
                Update: resourceCyralRepositoryUpdate,
                Delete: resourceCyralRepositoryDelete,

                Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },

                Importer: &schema.ResourceImporter{
                        State: schema.ImportStatePassthrough,
                },
        }
}

func resourceCyralRepositoryCreate(d *schema.ResourceData, m interface{}) error {
        d.SetId(d.Get("name").(string))
        return resourceCyralRepositoryRead(d, m)
}

func resourceCyralRepositoryRead(d *schema.ResourceData, m interface{}) error {
        d.Set("name", d.Id)
        return nil
}

func resourceCyralRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceCyralRepositoryRead(d, m)
}

func resourceCyralRepositoryDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
