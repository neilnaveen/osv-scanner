package reporter_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/osv-scanner/pkg/reporter"
	"github.com/google/osv-scanner/pkg/reporter/sbom"
)

func TestCycloneDXReporter_Errorf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		version sbom.CycloneDXVersion
	}{
		{version: sbom.CycloneDXVersion14},
		{version: sbom.CycloneDXVersion15},
	}

	text := "hello world!"
	for _, test := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewCycloneDXReporter(io.Discard, writer, test.version, reporter.ErrorLevel)

		r.Errorf(text)

		if writer.String() != text {
			t.Error("Error level message should have been printed")
		}
		if !r.HasErrored() {
			t.Error("HasErrored() should have returned true")
		}
	}
}

func TestCycloneDXReporter_Warnf(t *testing.T) {
	t.Parallel()

	text := "hello world!"
	tests := []struct {
		lvl              reporter.VerbosityLevel
		expectedPrintout string
		version          sbom.CycloneDXVersion
	}{
		{lvl: reporter.WarnLevel, expectedPrintout: text, version: sbom.CycloneDXVersion14},
		{lvl: reporter.WarnLevel, expectedPrintout: text, version: sbom.CycloneDXVersion15},
		{lvl: reporter.ErrorLevel, expectedPrintout: "", version: sbom.CycloneDXVersion14},
		{lvl: reporter.ErrorLevel, expectedPrintout: "", version: sbom.CycloneDXVersion15},
	}

	for _, test := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewCycloneDXReporter(io.Discard, writer, test.version, test.lvl)

		r.Warnf(text)

		if writer.String() != test.expectedPrintout {
			t.Errorf("expected \"%s\", got \"%s\"", test.expectedPrintout, writer.String())
		}
	}
}

func TestCycloneDXReporter_Infof(t *testing.T) {
	t.Parallel()

	text := "hello world!"
	tests := []struct {
		lvl              reporter.VerbosityLevel
		expectedPrintout string
		version          sbom.CycloneDXVersion
	}{
		{lvl: reporter.InfoLevel, expectedPrintout: text, version: sbom.CycloneDXVersion14},
		{lvl: reporter.InfoLevel, expectedPrintout: text, version: sbom.CycloneDXVersion15},
		{lvl: reporter.WarnLevel, expectedPrintout: "", version: sbom.CycloneDXVersion14},
		{lvl: reporter.WarnLevel, expectedPrintout: "", version: sbom.CycloneDXVersion15},
	}

	for _, test := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewCycloneDXReporter(io.Discard, writer, test.version, test.lvl)

		r.Infof(text)

		if writer.String() != test.expectedPrintout {
			t.Errorf("expected \"%s\", got \"%s\"", test.expectedPrintout, writer.String())
		}
	}
}

func TestCycloneDXReporter_Verbosef(t *testing.T) {
	t.Parallel()
	text := "hello world!"
	tests := []struct {
		version          sbom.CycloneDXVersion
		lvl              reporter.VerbosityLevel
		expectedPrintout string
	}{
		{
			version:          sbom.CycloneDXVersion14,
			lvl:              reporter.VerboseLevel,
			expectedPrintout: text,
		},
		{
			version:          sbom.CycloneDXVersion15,
			lvl:              reporter.VerboseLevel,
			expectedPrintout: text,
		},
		{
			version:          sbom.CycloneDXVersion14,
			lvl:              reporter.InfoLevel,
			expectedPrintout: "",
		},
		{
			version:          sbom.CycloneDXVersion15,
			lvl:              reporter.InfoLevel,
			expectedPrintout: "",
		},
	}

	for _, test := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewCycloneDXReporter(io.Discard, writer, test.version, test.lvl)

		r.Verbosef(text)

		if writer.String() != test.expectedPrintout {
			t.Errorf("expected \"%s\", got \"%s\"", test.expectedPrintout, writer.String())
		}
	}
}
