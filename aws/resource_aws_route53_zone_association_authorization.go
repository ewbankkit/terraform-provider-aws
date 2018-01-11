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
	conn := meta.(*AWSClient).r53conn

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

	log.Printf("[DEBUG] Creating Route53 zone association authorization: %#v", req)
	_, err := conn.CreateVPCAssociationAuthorization(req)
	if err != nil {
		return fmt.Errorf("Error creating Route53 zone association authorization: %s", err.Error())
	}

	d.SetId(fmt.Sprintf("%s:%s", *req.HostedZoneId, *req.VPC.VPCId))
	d.Set("vpc_region", req.VPC.VPCRegion)

	return resourceAwsRoute53ZoneAssociationAuthorizationRead(d, meta)
}

func resourceAwsRoute53ZoneAssociationAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).r53conn

	zoneId, vpcId := resourceAwsRoute53ZoneAssociationAuthorizationParseId(d.Id())
	req := route53.ListVPCAssociationAuthorizationsInput{
		HostedZoneId: aws.String(zoneId),
	}
	for {
		res, err := conn.ListVPCAssociationAuthorizations(&req)
		if err != nil {
			return fmt.Errorf("Error reading Route53 zone association authorizations: %s", err.Error())
		}

		for _, vpc := range res.VPCs {
			if vpcId == *vpc.VPCId {
				return nil
			}
		}

		// Loop till we find our authorization or we reach the end.
		if res.NextToken == nil {
			break
		}
		req.NextToken = res.NextToken
	}

	// No authorization found.
	d.SetId("")
	return nil
}

func resourceAwsRoute53ZoneAssociationAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).r53conn

	log.Printf("[DEBUG] Deleting Route53 zone association authorization: %s", d.Id())
	zoneId, vpcId := resourceAwsRoute53ZoneAssociationAuthorizationParseId(d.Id())
	_, err := conn.DeleteVPCAssociationAuthorization(&route53.DeleteVPCAssociationAuthorizationInput{
		HostedZoneId: aws.String(zoneId),
		VPC: &route53.VPC{
			VPCId:     aws.String(vpcId),
			VPCRegion: aws.String(d.Get("vpc_region").(string)),
		},
	})
	if err != nil {
		if isAWSErr(err, "VPCAssociationAuthorizationNotFound", "") {
			log.Printf("[DEBUG] Route53 zone association authorization %s is already gone", d.Id())
		} else {
			return fmt.Errorf("Error deleting Route53 zone association authorization: %s", err.Error())
		}
	}

	return nil
}

func resourceAwsRoute53ZoneAssociationAuthorizationParseId(id string) (string, string) {
	parts := strings.SplitN(id, ":", 2)
	return parts[0], parts[1]
}
