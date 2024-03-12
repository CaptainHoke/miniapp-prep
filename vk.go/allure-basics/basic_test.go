package test

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"testing"
)

func TestRunner(t *testing.T) {
	runner.Run(t, "First Test", func (t provider.T) {
		t.NewStep("First Step Yuuu")
	})

	runner.Run(t, "Second Test", func(t provider.T) {
		t.WithNewStep("Second Step I can't believe it", func(sCtx provider.StepCtx) {
			sCtx.NewStep("Stepsister, I'm StepContext")
		})
	})
}
