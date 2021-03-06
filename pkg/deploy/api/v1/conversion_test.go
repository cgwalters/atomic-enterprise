package v1

import (
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"

	newer "github.com/projectatomic/atomic-enterprise/pkg/deploy/api"
)

func TestTriggerRoundTrip(t *testing.T) {
	p := DeploymentTriggerImageChangeParams{
		From: kapi.ObjectReference{
			Kind: "DockerImage",
			Name: "",
		},
	}
	out := &newer.DeploymentTriggerImageChangeParams{}
	if err := api.Scheme.Convert(&p, out); err == nil {
		t.Errorf("unexpected error: %v", err)
	}
	p.From.Name = "a/b:test"
	out = &newer.DeploymentTriggerImageChangeParams{}
	if err := api.Scheme.Convert(&p, out); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if out.RepositoryName != "a/b" && out.Tag != "test" {
		t.Errorf("unexpected output: %#v", out)
	}
}
