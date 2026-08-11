package config

import (
	"context"
	"testing"
)

const (
	testProjectID = testProjectID
	testVPCID     = testVPCID
	testDbaasID   = testDbaasID
	testSGID      = testSGID
)

// TestLeafIDFromSlash verifies the helper that extracts a specific segment.
func TestLeafIDFromSlash(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		index int
		want  string
	}{
		{"subnet leaf at index 2", "proj-1/vpc-2/sub-3", 2, "sub-3"},
		{"sg leaf at index 2", "proj-1/vpc-2/sg-3", 2, "sg-3"},
		{"rule leaf at index 3", "proj-1/vpc-2/sg-3/rule-4", 3, "rule-4"},
		{"fallback when short", "only-one-part", 2, "only-one-part"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fn := leafIDFromSlash(tc.index)
			got, err := fn(map[string]any{"id": tc.id})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// TestSubnetGetIDFn verifies the subnet composite ID reconstruction.
func TestSubnetGetIDFn(t *testing.T) {
	e := subnetExternalName()
	id, err := e.GetIDFn(context.Background(), "sub-xyz", map[string]any{
		"project_id": testProjectID,
		"vpc_id":     testVPCID,
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-abc/vpc-123/sub-xyz"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

func TestSubnetGetIDFn_MissingVpcID(t *testing.T) {
	e := subnetExternalName()
	_, err := e.GetIDFn(context.Background(), "sub-xyz", map[string]any{
		"project_id": testProjectID,
	}, nil)
	if err == nil {
		t.Error("expected error when vpc_id is missing, got nil")
	}
}

// TestSecurityGroupGetIDFn verifies the 3-part sg ID.
func TestSecurityGroupGetIDFn(t *testing.T) {
	e := securityGroupExternalName()
	id, err := e.GetIDFn(context.Background(), "sg-xyz", map[string]any{
		"project_id": testProjectID,
		"vpc_id":     testVPCID,
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-abc/vpc-123/sg-xyz"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

// TestSecurityRuleGetIDFn verifies the 4-part rule ID.
func TestSecurityRuleGetIDFn(t *testing.T) {
	e := securityRuleExternalName()
	id, err := e.GetIDFn(context.Background(), "rule-xyz", map[string]any{
		"project_id":        testProjectID,
		"vpc_id":            testVPCID,
		"security_group_id": testSGID,
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-abc/vpc-123/sg-456/rule-xyz"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

// TestSnapshotExternalName verifies the unusual project/snap/billing_period format.
func TestSnapshotGetExternalName(t *testing.T) {
	e := snapshotExternalName()
	got, err := e.GetExternalNameFn(map[string]any{"id": "proj-abc/snap-xyz/Hour"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "snap-xyz" {
		t.Errorf("got %q, want %q", got, "snap-xyz")
	}
}

func TestSnapshotGetIDFn(t *testing.T) {
	e := snapshotExternalName()
	id, err := e.GetIDFn(context.Background(), "snap-xyz", map[string]any{
		"project_id":     testProjectID,
		"billing_period": "Hour",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-abc/snap-xyz/Hour"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

// TestDatabaseGetIDFn verifies the project/dbaas/db composite ID.
func TestDatabaseGetIDFn(t *testing.T) {
	e := databaseExternalName()
	id, err := e.GetIDFn(context.Background(), "db-xyz", map[string]any{
		"project_id": testProjectID,
		"dbaas_id":   testDbaasID,
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-abc/dbaas-123/db-xyz"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

// TestIdentifierFromProviderWithProjectID verifies 2-part project/resource ID.
func TestIdentifierFromProviderWithProjectID(t *testing.T) {
	e := identifierFromProviderWithProjectID()
	id, err := e.GetIDFn(context.Background(), "vpc-abc", map[string]any{
		"project_id": "proj-xyz",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "proj-xyz/vpc-abc"
	if id != want {
		t.Errorf("got %q, want %q", id, want)
	}
}

func TestIdentifierFromProviderWithProjectID_Missing(t *testing.T) {
	e := identifierFromProviderWithProjectID()
	_, err := e.GetIDFn(context.Background(), "vpc-abc", map[string]any{}, nil)
	if err == nil {
		t.Error("expected error when project_id is missing, got nil")
	}
}
