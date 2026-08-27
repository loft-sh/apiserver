package apiserver

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenericConfigWithoutKubeContext(t *testing.T) {
	// GenericConfig writes self-signed certs relative to the working directory.
	t.Chdir(t.TempDir())

	kubeconfig := filepath.Join(t.TempDir(), "kubeconfig")
	if err := os.WriteFile(kubeconfig, nil, 0o600); err != nil {
		t.Fatalf("write empty kubeconfig: %v", err)
	}
	t.Setenv("KUBECONFIG", kubeconfig)
	t.Setenv("KUBERNETES_SERVICE_HOST", "")
	t.Setenv("KUBERNETES_SERVICE_PORT", "")

	_, err := newAPIServerOptions(nil).GenericConfig(nil)
	if err == nil {
		t.Fatal("expected an error, got none")
	}
	if !strings.Contains(err.Error(), "failed to build loopback client") {
		t.Fatalf("expected a loopback client error, got: %v", err)
	}
}
