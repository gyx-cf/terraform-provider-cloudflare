package tiered_cache_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func testTieredCacheConfig(rnd, zoneID, cacheType string) string {
	return fmt.Sprintf(`
resource "cloudflare_tiered_cache" "%[1]s" {
	zone_id = "%[2]s"
	cache_type = "%[3]s"
}
`, rnd, zoneID, cacheType)
}

func TestAccCloudflareTieredCache_Smart(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTieredCacheConfig(rnd, zoneID, "smart"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "cache_type", "smart"),
				),
			},
		},
	})
}

func TestAccCloudflareTieredCache_Generic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_tiered_cache." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTieredCacheConfig(rnd, zoneID, "generic"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "cache_type", "generic"),
				),
			},
		},
	})
}