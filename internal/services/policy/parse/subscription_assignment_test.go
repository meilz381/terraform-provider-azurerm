package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = SubscriptionAssignmentId{}

func TestSubscriptionAssignmentIDFormatter(t *testing.T) {
	actual := NewSubscriptionAssignmentID("12345678-1234-9876-4563-123456789012", "assignment1").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/assignment1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestSubscriptionAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionAssignmentId
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
			// missing PolicyAssignmentName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/",
			Error: true,
		},

		{
			// missing value for PolicyAssignmentName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Authorization/policyAssignments/assignment1",
			Expected: &SubscriptionAssignmentId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				PolicyAssignmentName: "assignment1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/PROVIDERS/MICROSOFT.AUTHORIZATION/POLICYASSIGNMENTS/ASSIGNMENT1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := SubscriptionAssignmentID(v.Input)
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
		if actual.PolicyAssignmentName != v.Expected.PolicyAssignmentName {
			t.Fatalf("Expected %q but got %q for PolicyAssignmentName", v.Expected.PolicyAssignmentName, actual.PolicyAssignmentName)
		}
	}
}
