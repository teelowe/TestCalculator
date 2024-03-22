package calculator

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Global test setup here...")
	exitVal := m.Run()
	log.Println("Global test teardown here...")
	os.Exit(exitVal)
}

func SubTestSetup(t *testing.T) {
	t.Logf("[[[ SUBTEST SETUP HERE... ]]]")
}

func SubTestTearDown(t *testing.T) {
	t.Logf("[[[ SUBTEST TEARDOWN HERE... ]]]")
}
