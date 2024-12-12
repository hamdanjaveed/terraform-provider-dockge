// Copyright (c) HashiCorp, Inc.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccStackResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStackResourceConfig("test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("stack.test", "name", "tests"),
				),
			},
		},
	})
}

func testAccStackResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "stack" "%s" {
}
`, name)
}
