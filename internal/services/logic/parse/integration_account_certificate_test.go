package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = IntegrationAccountCertificateId{}

func TestIntegrationAccountCertificateIDFormatter(t *testing.T) {
	actual := NewIntegrationAccountCertificateID("12345678-1234-9876-4563-123456789012", "group1", "account1", "certificate1").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/certificate1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestIntegrationAccountCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IntegrationAccountCertificateId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing IntegrationAccountName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/",
			Error: true,
		},

		{
			// missing value for IntegrationAccountName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/",
			Error: true,
		},

		{
			// missing CertificateName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/",
			Error: true,
		},

		{
			// missing value for CertificateName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/certificate1",
			Expected: &IntegrationAccountCertificateId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroup:          "group1",
				IntegrationAccountName: "account1",
				CertificateName:        "certificate1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/GROUP1/PROVIDERS/MICROSOFT.LOGIC/INTEGRATIONACCOUNTS/ACCOUNT1/CERTIFICATES/CERTIFICATE1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := IntegrationAccountCertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.IntegrationAccountName != v.Expected.IntegrationAccountName {
			t.Fatalf("Expected %q but got %q for IntegrationAccountName", v.Expected.IntegrationAccountName, actual.IntegrationAccountName)
		}
		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}
	}
}
