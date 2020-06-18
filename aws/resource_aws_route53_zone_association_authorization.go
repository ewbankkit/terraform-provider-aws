package aws

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func resourceAwsRoute53ZoneAssociationAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsRoute53ZoneAssociationAuthorizationCreate,
		Read:   resourceAwsRoute53ZoneAssociationAuthorizationRead,
		Delete: resourceAwsRoute53ZoneAssociationAuthorizationDelete,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"vpc_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsRoute53ZoneAssociationAuthorizationCreate(d *schema.ResourceData, meta interface{}) error {
	r53 := meta.(*AWSClient).r53conn

	req := &route53.CreateVPCAssociationAuthorizationInput{
		HostedZoneId: aws.String(d.Get("zone_id").(string)),
		VPC: &route53.VPC{
			VPCId: aws.String(d.Get("vpc_id").(string)),
		},
	}
	if v, ok := d.GetOk("vpc_region"); ok {
		req.VPC.VPCRegion = aws.String(v.(string))
	} else {
		req.VPC.VPCRegion = aws.String(meta.(*AWSClient).region)
	}

	log.Printf("[DEBUG] Creating Route53 VPC Association Authorization for hosted zone %s with VPC %s and region %s", *req.HostedZoneId, *req.VPC.VPCId, *req.VPC.VPCRegion)
	var err error
	_, err = r53.CreateVPCAssociationAuthorization(req)
	if err != nil {
		return err
	}

	// Store association id
	d.SetId(fmt.Sprintf("%s:%s", *req.HostedZoneId, *req.VPC.VPCId))
	d.Set("vpc_region", req.VPC.VPCRegion)

	return resourceAwsRoute53ZoneAssociationAuthorizationRead(d, meta)
}

func resourceAwsRoute53ZoneAssociationAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	r53 := meta.(*AWSClient).r53conn
	zone_id, vpc_id := resourceAwsRoute53ZoneAssociationAuthorizationParseId(d.Id())
	req := route53.ListVPCAssociationAuthorizationsInput{HostedZoneId: aws.String(zone_id)}
	for {
		res, err := r53.ListVPCAssociationAuthorizations(&req)
		if err != nil {
			return err
		}

		for _, vpc := range res.VPCs {
			if vpc_id == *vpc.VPCId {
				return nil
			}
		}

		// Loop till we find our authorization or we reach the end
		if res.NextToken != nil {
			req.NextToken = res.NextToken
		} else {
			break
		}
	}

	// no association found
	d.SetId("")
	return nil
}

func resourceAwsRoute53ZoneAssociationAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	r53 := meta.(*AWSClient).r53conn
	zone_id, vpc_id := resourceAwsRoute53ZoneAssociationAuthorizationParseId(d.Id())
	log.Printf("[DEBUG] Deleting Route53 Assocatiation Authorization for (%s) from vpc %s)",
		zone_id, vpc_id)

	req := route53.DeleteVPCAssociationAuthorizationInput{
		HostedZoneId: aws.String(zone_id),
		VPC: &route53.VPC{
			VPCId:     aws.String(vpc_id),
			VPCRegion: aws.String(d.Get("vpc_region").(string)),
		},
	}

	_, err := r53.DeleteVPCAssociationAuthorization(&req)
	if err != nil {
		return err
	}

	return nil
}

func resourceAwsRoute53ZoneAssociationAuthorizationParseId(id string) (string, string) {
	parts := strings.SplitN(id, ":", 2)
	return parts[0], parts[1]
}
