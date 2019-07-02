// Copyright (c) 2019 Aiven, Helsinki, Finland. https://aiven.io/
package main

import (
	"github.com/aiven/aiven-go-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceProjectVPC() *schema.Resource {
	return &schema.Resource{
		Read: datasourceProjectVPCRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The project the VPC belongs to",
				Required:    true,
				Type:        schema.TypeString,
			},
			"vpc_id": {
				Description: "The ID of the VPC",
				Required:    true,
				Type:        schema.TypeString,
			},
			"cloud_name": {
				Description: "Cloud the VPC is in",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"network_cidr": {
				Description: "Network address range used by the VPC like 192.168.0.0/24",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"state": {
				Computed:    true,
				Description: "State of the VPC (APPROVED, ACTIVE, DELETING, DELETED)",
				Type:        schema.TypeString,
			},
		},
	}
}

func datasourceProjectVPCRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)

	projectName := d.Get("project").(string)
	vpcID := d.Get("vpc_id").(string)

	vpc, err := client.VPCs.Get(projectName, vpcID)
	if err != nil {
		return err
	}

	d.SetId(buildResourceID(projectName, vpc.ProjectVPCID))
	return copyVPCPropertiesFromAPIResponseToTerraform(d, vpc, projectName)
}
