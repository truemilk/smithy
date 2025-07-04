package component

import (
	"context"

	"github.com/smithy-security/smithy/sdk/component/store"
	"github.com/smithy-security/smithy/sdk/component/uuid"
	finding "github.com/smithy-security/smithy/sdk/component/vulnerability-finding"
	ocsf "github.com/smithy-security/smithy/sdk/gen/ocsf_schema/v1"
)

// Helpers interfaces for common functionalities.
type (
	// Validator allows validating vulnerability findings by a specified criteria.
	Validator interface {
		// Validate validates the supplied vulnerability finding and returns an error if invalid.
		Validate(finding *ocsf.VulnerabilityFinding) error
	}

	// Reader allows reading vulnerability findings from a storage.
	Reader interface {
		// Read reads vulnerability findings from a storage.
		Read(ctx context.Context, instanceID uuid.UUID, queryOpts *store.QueryOpts) ([]*finding.VulnerabilityFinding, error)
	}

	// Updater allows updating vulnerability findings in an underlying storage.
	Updater interface {
		// Update updates existing vulnerability findings.
		Update(ctx context.Context, instanceID uuid.UUID, findings []*finding.VulnerabilityFinding) error
	}

	// Writer allows writing non-existent vulnerability findings in an underlying storage.
	Writer interface {
		// Write writes non-existing vulnerability findings.
		Write(ctx context.Context, instanceID uuid.UUID, findings []*ocsf.VulnerabilityFinding) error
	}

	// Closer allows to define behaviours to close component dependencies gracefully.
	Closer interface {
		// Close can be implemented to gracefully close component dependencies.
		Close(context.Context) error
	}

	// Storer allows storing vulnerability findings in an underlying storage.
	Storer interface {
		Closer
		Validator
		Reader
		Updater
		Writer
	}
)

// Components interfaces.
type (
	// Target prepares the workflow environment.
	Target interface {
		// Prepare prepares the target to be scanned.
		Prepare(ctx context.Context) error
	}

	// Scanner reads a scan's result and produces vulnerability findings.
	Scanner interface {
		// Transform transforms the raw scan data into vulnerability finding format.
		Transform(ctx context.Context) ([]*ocsf.VulnerabilityFinding, error)
	}

	// Filter allows filtering out vulnerability findings by some criteria.
	Filter interface {
		// Filter returns filtered findings from the supplied ones applying some criteria.
		// It returns false if no findings have been filtered out.
		Filter(ctx context.Context, findings []*finding.VulnerabilityFinding) ([]*finding.VulnerabilityFinding, bool, error)
	}

	// Enricher allows enriching vulnerability findings by some criteria.
	Enricher interface {
		// Annotate enriches vulnerability findings by some criteria.
		Annotate(ctx context.Context, findings []*finding.VulnerabilityFinding) ([]*finding.VulnerabilityFinding, error)
	}

	// Reporter advertises behaviours for reporting vulnerability findings.
	Reporter interface {
		// Report reports vulnerability findings on a specified destination.
		// i.e. raises them as tickets on your favourite ticketing system.
		Report(ctx context.Context, findings []*finding.VulnerabilityFinding) error
	}
)
