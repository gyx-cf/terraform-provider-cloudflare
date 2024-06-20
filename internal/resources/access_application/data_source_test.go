package access_application_test

import (
	"fmt"
	"testing"

	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudflareAccessApplicationDataSource_AccountName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountName(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountName(accountID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_AccountDomain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationAccountDomain(accountID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.AccountIDSchemaKey, accountID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationAccountDomain(accountID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		account_id = "%[2]s"
		domain = "%[1]s.%[3]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, accountID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneName(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneName(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneName(zoneID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, zoneID, domain)
}

func TestAccCloudflareAccessApplicationDataSource_ZoneDomain(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "data.cloudflare_access_application." + rnd
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "domain", rnd+"."+domain),
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttrSet(name, "aud"),
				),
			},
		},
	})
}

func testAccCheckCloudflareAccessApplicationZoneDomain(zoneID, name, domain string) string {
	return fmt.Sprintf(`
	resource "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_access_application" "%[1]s" {
		zone_id = "%[2]s"
		domain = "%[1]s.%[3]s"
		depends_on = [cloudflare_access_application.%[1]s]
	}
	`, name, zoneID, domain)
}