package opentelekomcloud

/*
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/vips"
)

// SKIP deprecated
func TestAccLBV1VIP_basic(t *testing.T) {
	var vip vips.VirtualIP

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDeprecated(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV1VIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLBV1VIP_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBV1VIPExists("opentelekomcloud_lb_vip_v1.vip_1", &vip),
				),
			},
			resource.TestStep{
				Config: testAccLBV1VIP_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("opentelekomcloud_lb_vip_v1.vip_1", "name", "vip_1_updated"),
				),
			},
		},
	})
}

// SKIP deprecated
func TestAccLBV1VIP_timeout(t *testing.T) {
	var vip vips.VirtualIP

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDeprecated(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV1VIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLBV1VIP_timeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBV1VIPExists("opentelekomcloud_lb_vip_v1.vip_1", &vip),
				),
			},
		},
	})
}

func testAccCheckLBV1VIPDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkingClient, err := config.networkingV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud networking client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_lb_vip_v1" {
			continue
		}

		_, err := vips.Get(networkingClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("LB VIP still exists")
		}
	}

	return nil
}

func testAccCheckLBV1VIPExists(n string, vip *vips.VirtualIP) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkingClient, err := config.networkingV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating OpenTelekomCloud networking client: %s", err)
		}

		found, err := vips.Get(networkingClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("VIP not found")
		}

		*vip = *found

		return nil
	}
}

const testAccLBV1VIP_basic = `
resource "opentelekomcloud_networking_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "opentelekomcloud_networking_subnet_v2" "subnet_1" {
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${opentelekomcloud_networking_network_v2.network_1.id}"
}

resource "opentelekomcloud_lb_pool_v1" "pool_1" {
  name = "pool_1"
  protocol = "HTTP"
  lb_method = "ROUND_ROBIN"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"
}

resource "opentelekomcloud_lb_vip_v1" "vip_1" {
  name = "vip_1"
  protocol = "HTTP"
  port = 80
  admin_state_up = true
  pool_id = "${opentelekomcloud_lb_pool_v1.pool_1.id}"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"

  persistence {
    type = "SOURCE_IP"
  }
}
`

const testAccLBV1VIP_update = `
resource "opentelekomcloud_networking_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "opentelekomcloud_networking_subnet_v2" "subnet_1" {
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${opentelekomcloud_networking_network_v2.network_1.id}"
}

resource "opentelekomcloud_lb_pool_v1" "pool_1" {
  name = "pool_1"
  protocol = "HTTP"
  lb_method = "ROUND_ROBIN"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"
}

resource "opentelekomcloud_lb_vip_v1" "vip_1" {
  name = "vip_1_updated"
  protocol = "HTTP"
  port = 80
  admin_state_up = true
  pool_id = "${opentelekomcloud_lb_pool_v1.pool_1.id}"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"

  persistence {
    type = "SOURCE_IP"
  }
}
`

const testAccLBV1VIP_timeout = `
resource "opentelekomcloud_networking_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "opentelekomcloud_networking_subnet_v2" "subnet_1" {
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${opentelekomcloud_networking_network_v2.network_1.id}"
}

resource "opentelekomcloud_lb_pool_v1" "pool_1" {
  name = "pool_1"
  protocol = "HTTP"
  lb_method = "ROUND_ROBIN"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"
}

resource "opentelekomcloud_lb_vip_v1" "vip_1" {
  name = "vip_1"
  protocol = "HTTP"
  port = 80
  admin_state_up = true
  pool_id = "${opentelekomcloud_lb_pool_v1.pool_1.id}"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet_1.id}"

  persistence {
    type = "SOURCE_IP"
  }

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
*/