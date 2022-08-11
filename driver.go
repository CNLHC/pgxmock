package pgxmock

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgxmockConn struct {
	pgxmock
}

type MockOptFn func(*pgxmock) error

// NewConn creates PgxConnIface database connection and a mock to manage expectations.
// Accepts options, like ValueConverterOption, to use a ValueConverter from
// a specific driver.
// Pings db so that all expectations could be
// asserted.
func NewConn(options ...MockOptFn) (PgxConnIface, error) {
	smock := &pgxmockConn{}
	smock.ordered = true
	return smock, smock.open(options)
}

func (c *pgxmockConn) Close(ctx context.Context) error {
	return c.close(ctx)
}

func WithTestingReporter(t *testing.T) MockOptFn {
	return func(p *pgxmock) error {
		p.t = t
		return nil
	}
}

type pgxmockPool struct {
	pgxmock
}

// NewPool creates PgxPoolIface pool of database connections and a mock to manage expectations.
// Accepts options, like ValueConverterOption, to use a ValueConverter from
// a specific driver.
// Pings db so that all expectations could be
// asserted.
func NewPool(options ...MockOptFn) (PgxPoolIface, error) {
	smock := &pgxmockPool{}
	smock.ordered = true
	return smock, smock.open(options)
}

func (p *pgxmockPool) Close() {
	_ = p.close(context.Background())
}

func (p *pgxmockPool) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return nil, errors.New("pgpool.Acquire() method is not implemented")
}
