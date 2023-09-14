// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/pkg/pusher"
	"github.com/ethersphere/bee/pkg/storage"
	"github.com/ethersphere/bee/pkg/swarm"
	"github.com/ethersphere/bee/pkg/topology"
	"golang.org/x/sync/errgroup"
)

// DirectUpload is the implementation of the NetStore.DirectUpload method.
func (db *DB) DirectUpload() PutterSession {
	// egCtx will allow early exit of Put operations if we have
	// already encountered error.
	eg, egCtx := errgroup.WithContext(context.Background())

	return &putterSession{
		Putter: putterWithMetrics{
			storage.PutterFunc(func(ctx context.Context, ch swarm.Chunk) error {
				db.directUploadLimiter <- struct{}{}
				eg.Go(func() error {
					defer func() { <-db.directUploadLimiter }()

					for {
						op := &pusher.Op{Chunk: ch, Err: make(chan error, 1), Direct: true}
						select {
						case <-ctx.Done():
							return ctx.Err()
						case <-egCtx.Done():
							return egCtx.Err()
						case <-db.quit:
							return ErrDBQuit
						case db.pusherFeed <- op:
							select {
							case <-ctx.Done():
								return ctx.Err()
							case <-egCtx.Done():
								return egCtx.Err()
							case <-db.quit:
								return ErrDBQuit
							case err := <-op.Err:
								if errors.Is(err, pusher.ErrShallowReceipt) {
									db.logger.Debug("direct upload: shallow receipt received, retrying", "chunk", ch.Address())
								} else if errors.Is(err, topology.ErrNotFound) {
									db.logger.Debug("direct upload: no peers available, retrying", "chunk", ch.Address())
								} else {
									return err
								}
							}
						}
					}
				})
				return nil
			}),
			db.metrics,
			"netstore",
		},
		done:    func(_ swarm.Address) error { return eg.Wait() },
		cleanup: func() error { _ = eg.Wait(); return nil },
	}
}

// Download is the implementation of the NetStore.Download method.
func (db *DB) Download(cache bool) storage.Getter {
	return getterWithMetrics{
		storage.GetterFunc(func(ctx context.Context, address swarm.Address) (swarm.Chunk, error) {
			ch, err := db.Lookup().Get(ctx, address)
			switch {
			case err == nil:
				return ch, nil
			case errors.Is(err, storage.ErrNotFound):
				if db.retrieval != nil {
					// if chunk is not found locally, retrieve it from the network
					ch, err = db.retrieval.RetrieveChunk(ctx, address, swarm.ZeroAddress)
					if err == nil && cache {
						select {
						case <-ctx.Done():
						case <-db.quit:
						case db.cacheLimiter.sem <- struct{}{}:
							db.cacheLimiter.wg.Add(1)
							go func() {
								defer func() {
									<-db.cacheLimiter.sem
									db.cacheLimiter.wg.Done()
								}()

								err := db.Cache().Put(db.cacheLimiter.ctx, ch)
								if err != nil {
									db.logger.Debug("putting chunk to cache failed", "error", err, "chunk_address", ch.Address())
								}
							}()
						}
					}
				}
			}
			if err != nil {
				return nil, err
			}
			return ch, nil
		}),
		db.metrics,
		"netstore",
	}
}

// PusherFeed is the implementation of the NetStore.PusherFeed method.
func (db *DB) PusherFeed() <-chan *pusher.Op {
	return db.pusherFeed
}
