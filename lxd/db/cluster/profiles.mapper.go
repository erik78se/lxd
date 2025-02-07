//go:build linux && cgo && !agent

package cluster

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var profileID = RegisterStmt(`
SELECT profiles.id FROM profiles JOIN projects ON profiles.project_id = projects.id
  WHERE projects.name = ? AND profiles.name = ?
`)

// GetProfileID return the ID of the profile with the given key.
// generator: profile ID
func GetProfileID(ctx context.Context, tx *sql.Tx, project string, name string) (int64, error) {
	stmt := stmt(tx, profileID)
	rows, err := stmt.Query(project, name)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"profiles\" ID: %w", err)
	}

	defer rows.Close()

	// Ensure we read one and only one row.
	if !rows.Next() {
		return -1, api.StatusErrorf(http.StatusNotFound, "Profile not found")
	}
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Failed to scan ID: %w", err)
	}

	if rows.Next() {
		return -1, fmt.Errorf("More than one row returned")
	}
	err = rows.Err()
	if err != nil {
		return -1, fmt.Errorf("Result set failure: %w", err)
	}

	return id, nil
}
