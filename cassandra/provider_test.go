package cassandra

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cassandra": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func TestProvider_configure1(t *testing.T) {
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{
		"username": "cassanrda",
		"password": "cassanrda",
		"port":     9042,
		"host":     "asdf",
	})
	p := Provider()
	err := p.Configure(rc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProvider_configure2(t *testing.T) {
	rc := terraform.NewResourceConfigRaw(map[string]interface{}{
		"username": "cassanrda",
		"password": "cassanrda",
		"port":     9042,
		"hosts":    []interface{}{"asd"},
	})
	p := Provider()
	err := p.Configure(rc)
	if err != nil {
		t.Fatal(err)
	}
}

func testAccPreCheck(t *testing.T) {
	url := os.Getenv("CASSANDRA_HOST")
	if url == "" {
		t.Fatal("CASSANDRA_HOST must be set for acceptance tests")
	}

	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}