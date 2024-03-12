package test

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type BasicSuite struct {
	suite.Suite
}

func (s *BasicSuite) TestFloppaTest(t provider.T) {
	t.NewStep("Floppa Step")
}

func (s *BasicSuite) Test69Test(t provider.T) {
	t.WithNewStep("My 69 Step", func(sCtx provider.StepCtx) {
		sCtx.NewStep("SubStep 69")
	})
}

func TestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(BasicSuite))
}
