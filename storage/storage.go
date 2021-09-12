package storage

import (
	"context"

	gcp "cloud.google.com/go/storage"
)

func Save(ctx context.Context, c *gcp.Client, bucket string, key string, b []byte) error {
	w := c.Bucket(bucket).Object(key).NewWriter(ctx)

	if _, err := w.Write(b); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
