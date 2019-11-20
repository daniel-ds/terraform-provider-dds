package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccFile_basic(t *testing.T) {
	filename := "test-file"
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"dds": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDdsFileConfig(filename, "content"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("dds_file.file", "filename", filename),
					resource.TestCheckResourceAttr("dds_file.file", "content", "content"),
				),
			},
			{
				Config: testAccDdsFileConfig(filename, "updated content"),
				Check:  resource.TestCheckResourceAttr("dds_file.file", "content", "updated content"),
			},
		},
	})
}

func testAccDdsFileConfig(filename string, content string) string {
	return fmt.Sprintf(`
resource "dds_file" "file" {
  filename = "%s"
  content = "%s"
}
`, filename, content)
}
