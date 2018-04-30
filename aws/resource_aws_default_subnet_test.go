package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSDefaultSubnet_basic(t *testing.T) {
	var v ec2.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSDefaultSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSDefaultSubnetConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet"),
				),
			},
		},
	})
}

func TestAccAWSDefaultSubnet_publicIp(t *testing.T) {
	var v ec2.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSDefaultSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSDefaultSubnetConfigPublicIp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "map_public_ip_on_launch", "true"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "assign_ipv6_address_on_creation", "false"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet-public-ip"),
				),
			},
			{
				Config: testAccAWSDefaultSubnetConfigNoPublicIp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "map_public_ip_on_launch", "false"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "assign_ipv6_address_on_creation", "false"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet-no-public-ip"),
				),
			},
		},
	})
}

func TestAccAWSDefaultSubnet_ipv6(t *testing.T) {
	var v ec2.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSDefaultSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSDefaultSubnetConfigIpv6,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					testAccCheckAwsSubnetIpv6(&v, true, true),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet-ipv6"),
				),
			},
			{
				Config: testAccAWSDefaultSubnetConfigIpv6UpdateAssignIpv6OnCreation,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					testAccCheckAwsSubnetIpv6(&v, true, false),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet-assign-ipv6-on-creation"),
				),
			},
			{
				Config: testAccAWSDefaultSubnetConfigIpv6UpdateIpv6Cidr,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					testAccCheckAwsSubnetIpv6(&v, true, false),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet-ipv6-update-cidr"),
				),
			},
			{
				Config: testAccAWSDefaultSubnetConfigNoIpv6,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists("aws_default_subnet.foo", &v),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "availability_zone", "us-west-2a"),
					testAccCheckAwsSubnetIpv6(&v, false, false),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.%", "1"),
					resource.TestCheckResourceAttr(
						"aws_default_subnet.foo", "tags.Name", "terraform-testacc-default-subnet"),
				),
			},
		},
	})
}

func testAccCheckAWSDefaultSubnetDestroy(s *terraform.State) error {
	// We expect subnet to still exist
	return nil
}

const testAccAWSDefaultSubnetConfigBasic = `
resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  tags {
    Name = "terraform-testacc-default-subnet"
  }
}
`

const testAccAWSDefaultSubnetConfigPublicIp = `
resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  map_public_ip_on_launch = true
  tags {
    Name = "terraform-testacc-default-subnet-public-ip"
  }
}
`

const testAccAWSDefaultSubnetConfigNoPublicIp = `
resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  map_public_ip_on_launch = false
  tags {
    Name = "terraform-testacc-default-subnet-no-public-ip"
  }
}
`

const testAccAWSDefaultSubnetConfigIpv6 = `
data "aws_vpc" "default" {
  default = true
}

resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  ipv6_cidr_block = "${cidrsubnet(data.aws_vpc.default.ipv6_cidr_block, 8, 1)}"
  assign_ipv6_address_on_creation = true
  tags {
    Name = "terraform-testacc-default-subnet-ipv6"
  }
}
`

const testAccAWSDefaultSubnetConfigIpv6UpdateAssignIpv6OnCreation = `
data "aws_vpc" "default" {
  default = true
}

resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  ipv6_cidr_block = "${cidrsubnet(data.aws_vpc.default.ipv6_cidr_block, 8, 1)}"
  assign_ipv6_address_on_creation = false
  tags {
    Name = "terraform-testacc-default-subnet-assign-ipv6-on-creation"
  }
}
`

const testAccAWSDefaultSubnetConfigIpv6UpdateIpv6Cidr = `
data "aws_vpc" "default" {
  default = true
}

resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  ipv6_cidr_block = "${cidrsubnet(data.aws_vpc.default.ipv6_cidr_block, 8, 3)}"
  assign_ipv6_address_on_creation = false
  tags {
    Name = "terraform-testacc-default-subnet-ipv6-update-cidr"
  }
}
`

const testAccAWSDefaultSubnetConfigNoIpv6 = `
resource "aws_default_subnet" "foo" {
  availability_zone = "us-west-2a"
  ipv6_cidr_block = ""
  tags {
    Name = "terraform-testacc-default-subnet"
  }
}
`
