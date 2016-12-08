/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package data

import (
	"database/sql"

	"github.com/pkg/errors"
)

func createSQLiteDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "beginning transaction")
	}
	defer tx.Rollback()

	for tbl, query := range schema {
		_, err := tx.Exec(query)
		if err != nil {
			return errors.Wrapf(err, "initializing %s", tbl)
		}
	}

	return errors.Wrap(tx.Commit(), "committing transaction")
}

var schema = map[string]string{
	"cluster":             createTableCluster,
	"cluster_type":        createTableClusterType,
	"cluster_yarn_detail": createTableClusterYarnDetail,
	"engine":              createTableEngine,
	"entity_type":         createTableEntityType,
	"identity":            createTableIdentity,
	"identity_workgroup":  createTableIdentityWorkgroup,
	"identity_role":       createTableIdentityRole,
	"history":             createTableHistory,
	"meta":                createTableMeta,
	"permission":          createTablePermission,
	"privilege":           createTablePrivilege,
	"role":                createTableRole,
	"role_permission":     createTableRolePermission,
	"state":               createTableState,
	"workgroup":           createTableWorkgroup,
}

var createTableCluster = `
CREATE TABLE cluster (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    type_id integer NOT NULL,
    detail_id integer,
    address text UNIQUE,
    state string NOT NULL,
    created datetime NOT NULL,

    CONSTRAINT type_id FOREIGN KEY (type_id) REFERENCES cluster_type(id),
    CONSTRAINT detail_id FOREIGN KEY (detail_id) REFERENCES cluster_yarn_detail(id),
    CONSTRAINT state FOREIGN KEY (state) REFERENCES state(name)
)
`

var createTableClusterType = `
CREATE TABLE cluster_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
)
`

var createTableClusterYarnDetail = `
CREATE TABLE cluster_yarn_detail (
    id integer PRIMARY KEY AUTOINCREMENT,
    engine_id integer NOT NULL,
    size integer NOT NULL,
    application_id text NOT NULL,
    memory text NOT NULL,
    output_dir text NOT NULL,

    FOREIGN KEY (engine_id) REFERENCES engine(id)
)
`

var createTableEngine = `
CREATE TABLE engine (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    location text NOT NULL,
    created datetime NOT NULL
)
`

var createTableEntityType = `
CREATE TABLE entity_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
)
`

var createTableHistory = `
CREATE TABLE history (
    id integer PRIMARY KEY AUTOINCREMENT,
    action text NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,
    description text,
    created datetime NOT NULL,

    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (identity_id) REFERENCES identity(id)
)
`

var createTableIdentity = `
CREATE TABLE identity (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    password text,
    workgroup_id integer,
    is_active boolean NOT NULL,
    last_login integer with time zone,
    created datetime NOT NULL
)
`

var createTableIdentityRole = `
CREATE TABLE identity_role (
    identity_id integer NOT NULL,
    role_id integer NOT NULL,

    PRIMARY KEY (identity_id, role_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
)
`

var createTableIdentityWorkgroup = `
CREATE TABLE identity_workgroup (
    identity_id integer NOT NULL,
    workgroup_id integer NOT NULL,

    PRIMARY KEY (identity_id, workgroup_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
)
`

var createTableMeta = `
CREATE TABLE meta (
    id integer NOT NULL,
    key text NOT NULL UNIQUE,
    value text NOT NULL,

    PRIMARY KEY (id)
)
`

var createTablePermission = `
CREATE TABLE permission (
    id integer PRIMARY KEY AUTOINCREMENT,
    code text NOT NULL UNIQUE,
    description text NOT NULL
)
`

var createTablePrivilege = `
CREATE TABLE privilege (
    privilege_type text NOT NULL,
    workgroup_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,

    PRIMARY KEY (privilege_type, workgroup_id, entity_type_id, entity_id),
    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
)
`

var createTableProject = `
CREATE TABLE project (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    model_category text NOT NULL,
    created datetime NOT NULL  
)
`

var createTableRole = `
CREATE TABLE role (
    id integer NOT NULL,
    name text NOT NULL UNIQUE,
    description text,
    created datetime NOT NULL,

    PRIMARY KEY (id)
)
`

var createTableRolePermission = `
CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,

    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE
)
`

var createTableService = `
CREATE TABLE service (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer NOT NULL,
    name text NOT NULL,
    host text,
    port integer,
    process_id integer,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id)
    FOREIGN KEY (model_id) REFERENCES model(id)
)
`

var createTableState = `
CREATE TABLE state (
    id integer NOT NULL,
    name text NOT NULL UNIQUE,

    PRIMARY KEY (id)
)
`

var createTableWorkgroup = `
CREATE TABLE workgroup (
    id integer PRIMARY KEY AUTOINCREMENT,
    type workgroup_type NOT NULL,
    name text NOT NULL UNIQUE,
    description text,
    created datetime NOT NULL 
)
`