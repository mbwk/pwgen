package pwgen

import (
	"testing"
)

func TestSelectCaseLower(t *testing.T) {
	expected := Lowercase
	actual, _ := SelectCase(false, true)
	if expected != actual {
		t.Fatalf("select case should have returned lowercase pool of letters")
	}
}
func TestSelectCaseUpper(t *testing.T) {
	expected := Uppercase
	actual, _ := SelectCase(true, false)
	if expected != actual {
		t.Fatalf("select case should have returned uppercase pool of letters")
	}
}
func TestSelectCaseMixed(t *testing.T) {
	expected := Mixedcase
	actual, _ := SelectCase(true, true)
	if expected != actual {
		t.Fatalf("select case should have returned mixed case pool of letters")
	}
}

func TestSelectNoneCaseError(t *testing.T) {
	_, err := SelectCase(false, false)
	if err == nil {
		t.Fatalf("error should be raised if neither case is selected")
	}
}
