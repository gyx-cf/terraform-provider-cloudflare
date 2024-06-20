package tunnel_test

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccCloudflareTunnelCreate_Basic(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelBasic(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_tunnel" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	}`, accID, name)
}

func TestAccCloudflareTunnelCreate_Managed(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelManaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
					resource.TestCheckResourceAttr(name, "config_src", "cloudflare"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelManaged(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_tunnel" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
		config_src = "cloudflare"
	}`, accID, name)
}

func TestAccCloudflareTunnelCreate_Unmanaged(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Argo Tunnel
	// endpoint does not yet support the API tokens.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("cloudflare_tunnel.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareTunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareTunnelUnmanaged(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rnd),
					resource.TestCheckResourceAttr(name, "secret", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="),
					resource.TestMatchResourceAttr(name, "cname", regexp.MustCompile(".*\\.cfargotunnel\\.com")),
					resource.TestCheckResourceAttr(name, "config_src", "local"),
				),
			},
		},
	})
}

func testAccCheckCloudflareTunnelUnmanaged(accID, name string) string {
	return fmt.Sprintf(`
	resource "cloudflare_tunnel" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
		config_src = "local"
	}`, accID, name)
}

func testAccCheckCloudflareTunnelDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_tunnel" {
			continue
		}

		accountID := rs.Primary.Attributes[consts.AccountIDSchemaKey]
		tunnelID := rs.Primary.ID
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		tunnel, err := client.GetTunnel(context.Background(), cloudflare.AccountIdentifier(accountID), tunnelID)

		if err != nil {
			return err
		}

		if tunnel.DeletedAt == nil {
			return fmt.Errorf("argo tunnel with ID %s still exists", tunnel.ID)
		}
	}

	return nil
}