package runner

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"testing"
)

func TestRunner(t *testing.T) {
	r := runner.NewRunner(t, "Flabbergasting Suite")

	r.NewTest("Flamboyant Test", func(t provider.T) {
		t.NewStep("Flaming Step")
	})

	r.NewTest("Fascinating Test", func(t provider.T) {
		t.WithNewStep("Fabulous Step", func(sCtx provider.StepCtx) {
			sCtx.NewStep("Fancy SubStep")
		})
	})

	r.RunTests()
}

