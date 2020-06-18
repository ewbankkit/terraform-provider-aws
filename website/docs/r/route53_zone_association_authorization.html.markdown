---
subcategory: "Route53"
layout: "aws"
page_title: "AWS: aws_route53_zone_association_authorization"
description: |-
  Provides a Route53 private Hosted Zone to VPC association authorization resource.
---

# aws_route53_zone_association_authorization

Provides a Route53 private Hosted Zone to VPC association authorization resource.

This resource is used to authorize the AWS account that created a specified VPC to [associate](route53_zone_association.html)
that VPC with a specified Route53 hosted zone created by a different account.

This resource should be used in the AWS account that created the [Route53 hosted zone](route53_zone.html).
The [association](route53_zone_association.html) should be created in the AWS account that created the [VPC](vpc.html).

## Example Usage

```hcl
provider "aws" {
  // Zone creator's credentials.
}

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
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The ID of the private hosted zone that you want to authorize associating a VPC with.
* `vpc_id` - (Required) The ID of the VPC to associate with the private hosted zone.
* `vpc_region` - (Optional) The VPC's region. Defaults to the region of the AWS provider.

## Attributes Reference

The following attributes are exported:

* `id` - The authorization ID.
