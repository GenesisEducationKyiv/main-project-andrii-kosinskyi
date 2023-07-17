package tests

import (
	"strings"
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestPackage_ShouldNotDependOn(t *testing.T) {
	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/validator").
			ShouldNotDependOn("bitcoin_checker_api/internal/renderer")

		assertNoError(t, mockT)
	})

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/validator").
			ShouldNotDependOn("bitcoin_checker_api/internal/usecase")

		assertNoError(t, mockT)
	})

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/renderer").
			ShouldNotDependOn("bitcoin_checker_api/internal/usecase")

		assertNoError(t, mockT)
	})

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/pkg/email").
			ShouldNotDependOn("bitcoin_checker_api/internal/pkg/exchange-rate")

		assertNoError(t, mockT)
	})

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/pkg/email").
			ShouldNotDependOn("bitcoin_checker_api/internal/repository")

		assertNoError(t, mockT)
	})

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/pkg/exchange-rate").
			ShouldNotDependOn("bitcoin_checker_api/internal/repository")

		assertNoError(t, mockT)
	})

	t.Run("Fails on dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/pkg/email").
			ShouldNotDependOn("bitcoin_checker_api/internal/model")

		assertError(t, mockT,
			"bitcoin_checker_api/internal/pkg/email",
			"bitcoin_checker_api/internal/model")
	})

	t.Run("Fails on dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/repository").
			ShouldNotDependOn("bitcoin_checker_api/internal/model")

		assertError(t, mockT,
			"bitcoin_checker_api/internal/repository",
			"bitcoin_checker_api/internal/model")
	})

	t.Run("Fails on dependencies", func(t *testing.T) {
		mockT := new(testingT)
		archtest.Package(mockT, "bitcoin_checker_api/internal/pkg/exchange-rate").
			ShouldNotDependOn("bitcoin_checker_api/internal/model")

		assertError(t, mockT,
			"bitcoin_checker_api/internal/pkg/exchange-rate",
			"bitcoin_checker_api/internal/model")
	})
}

func assertNoError(t *testing.T, mockT *testingT) {
	t.Helper()
	if mockT.errored() {
		t.Fatalf("archtest should not have failed but, %s", mockT.message())
	}
}

func assertError(t *testing.T, mockT *testingT, dependencyTrace ...string) {
	t.Helper()
	if !mockT.errored() {
		t.Fatal("archtest did not fail on dependency")
	}

	if dependencyTrace == nil {
		return
	}

	s := strings.Builder{}
	s.WriteString("Error:\n")
	for i, v := range dependencyTrace {
		s.WriteString(strings.Repeat("\t", i))
		s.WriteString(v + "\n")
	}

	if mockT.message() != s.String() {
		t.Errorf("expected %s got error message: %s", s.String(), mockT.message())
	}
}

type testingT struct {
	errors [][]interface{}
}

func (t *testingT) Error(args ...interface{}) {
	t.errors = append(t.errors, args)
}

func (t testingT) errored() bool {
	if len(t.errors) != 0 {
		return true
	}

	return false
}

func (t *testingT) message() interface{} {
	return t.errors[0][0]
}
