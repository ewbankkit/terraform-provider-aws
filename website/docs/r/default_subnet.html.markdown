---
layout: "aws"
page_title: "AWS: aws_default_subnet"
sidebar_current: "docs-aws-resource-default-subnet"
description: |-
  Manage a default VPC subnet resource.
---

# aws_default_subnet

Provides a resource to manage a [default AWS VPC subnet](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html#default-vpc-basics)
in the current region.

The `aws_default_subnet` behaves differently from normal resources, in that
Terraform does not _create_ this resource, but instead "adopts" it
into management.

## Example Usage

Basic usage with tags:

```hcl
resource "aws_default_subnet" "default_az1" {
  availability_zone = "us-west-2a"

	tags {
		Name = "Default subnet for us-west-2a"
	}
}
```

## Argument Reference

The arguments of an `aws_default_subnet` differ from `aws_subnet` resources.
Namely, the `availability_zone` argument is required and the `vpc_id` and `cidr_block` arguments are computed.
The following arguments are still supported:

* `ipv6_cidr_block` - (Optional) The IPv6 network range for the subnet,
    in CIDR notation. The subnet size must use a /64 prefix length.
* `map_public_ip_on_launch` -  (Optional) Specify true to indicate
    that instances launched into the subnet should be assigned
    a public IP address.
* `assign_ipv6_address_on_creation` - (Optional) Specify true to indicate
    that network interfaces created in the specified subnet should be
    assigned an IPv6 address.
* `tags` - (Optional) A mapping of tags to assign to the resource.

### Removing `aws_default_subnet` from your configuration

The `aws_default_subnet` resource allows you to manage a region's default VPC subnet,
but Terraform cannot destroy it. Removing this resource from your configuration
will remove it from your statefile and management, but will not destroy the subnet.
You can resume managing the subnet via the AWS Console.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the subnet
* `availability_zone`- The AZ for the subnet.
* `cidr_block` - The CIDR block for the subnet.
* `vpc_id` - The VPC ID.
* `ipv6_association_id` - The association ID for the IPv6 CIDR block.
* `ipv6_cidr_block` - The IPv6 CIDR block.

## Timeouts

`aws_default_subnet` provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `update` - (Default `3 minutes`) Used for default subnet updates
