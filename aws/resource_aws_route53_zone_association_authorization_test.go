package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func TestAccAWSRoute53ZoneAssociationAuthorization_basic(t *testing.T) {
	var zone route53.HostedZone

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoute53ZoneAssociationAuthorizationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRoute53ZoneAssociationAuthorizationConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute53ZoneAssociationAuthorizationExists("aws_route53_zone_association_authorization.foo", &zone),
				),
			},
		},
	})
}

func testAccCheckRoute53ZoneAssociationAuthorizationDestroy(s *terraform.State) error {
	return testAccCheckRoute53ZoneAssociationAuthorizationDestroyWithProvider(s, testAccProvider)
}

func testAccCheckRoute53ZoneAssociationAuthorizationDestroyWithProvider(s *terraform.State, provider *schema.Provider) error {
	conn := provider.Meta().(*AWSClient).r53conn
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_route53_zone_association_authorization" {
			continue
		}

		zone_id, vpc_id := resourceAwsRoute53ZoneAssociationAuthorizationParseId(rs.Primary.ID)

		req := route53.ListVPCAssociationAuthorizationsInput{HostedZoneId: aws.String(zone_id)}
		res, err := conn.ListVPCAssociationAuthorizations(&req)
		if err != nil {
			return err
		}

		exists := false
		for _, vpc := range res.VPCs {
			if vpc_id == *vpc.VPCId {
				exists = true
			}
		}

		if exists {
			return fmt.Errorf("Zone association authorization for zone %v with VPC %v still exists", zone_id, vpc_id)
		}
	}
	return nil
}

func testAccCheckRoute53ZoneAssociationAuthorizationExists(n string, zone *route53.HostedZone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return testAccCheckRoute53ZoneAssociationAuthorizationExistsWithProvider(s, n, zone, testAccProvider)
	}
}

func testAccCheckRoute53ZoneAssociationAuthorizationExistsWithProvider(s *terraform.State, n string, zone *route53.HostedZone, provider *schema.Provider) error {
	rs, ok := s.RootModule().Resources[n]
	if !ok {
		return fmt.Errorf("Not found: %s", n)
	}

	if rs.Primary.ID == "" {
		return fmt.Errorf("No zone association authorization ID is set")
	}

	zone_id, vpc_id := resourceAwsRoute53ZoneAssociationAuthorizationParseId(rs.Primary.ID)
	conn := provider.Meta().(*AWSClient).r53conn

	req := route53.ListVPCAssociationAuthorizationsInput{HostedZoneId: aws.String(zone_id)}
	res, err := conn.ListVPCAssociationAuthorizations(&req)
	if err != nil {
		return err
	}

	exists := false
	for _, vpc := range res.VPCs {
		if vpc_id == *vpc.VPCId {
			exists = true
		}
	}

	if !exists {
		return fmt.Errorf("Zone association authorization not found")
	}

	return nil
}

const testAccRoute53ZoneAssociationAuthorizationConfig = `
provider "aws" {
  alias = "bar"
  // VPC creator's credentials.
}

resource "aws_vpc" "foo" {
  cidr_block           = "10.6.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "aws_route53_zone" "foo" {
  name   = "example.com"
  vpc_id = "${aws_vpc.foo.id}"
}

resource "aws_vpc" "bar" {
  provider = "aws.bar"

  cidr_block           = "10.7.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "aws_route53_zone_association_authorization" "foo" {
  zone_id = "${aws_route53_zone.foo.id}"
  vpc_id  = "${aws_vpc.bar.id}"
}

resource "aws_route53_zone_association" "bar" {
  provider = "aws.bar"

  zone_id = "${aws_route53_zone_association_authorization.foo.zone_id}"
  vpc_id  = "${aws_vpc.bar.id}"
}
`
