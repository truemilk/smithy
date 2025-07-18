package tar

import (
	"context"
	"os"

	"github.com/go-errors/errors"

	"github.com/smithy-security/smithy/components/targets/source-code-artifact/internal/artifact/extractor/common"
	"github.com/smithy-security/smithy/components/targets/source-code-artifact/internal/reader"
)

type extractor struct{}

// NewExtractor returns a new extractor.
func NewExtractor() extractor {
	return extractor{}
}

// ExtractArtifact plainly uses the Untar helper.
func (e extractor) ExtractArtifact(ctx context.Context, sourcePath, destPath string) error {
	tmpArchive, err := os.OpenFile(sourcePath, os.O_RDONLY, 0600)
	if err != nil {
		return errors.Errorf("could not open temporary archive file for extracting: %w", err)
	}
	defer reader.CloseReader(ctx, tmpArchive)

	return common.Untar(ctx, tmpArchive, destPath)
}
