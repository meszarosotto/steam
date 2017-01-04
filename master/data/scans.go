// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package data

import "database/sql"

func ScanBinomialModel(r *sql.Row) (binomialModel, error) {
	var s binomialModel
	if err := r.Scan(
		&s.ModelId,
		&s.Mse,
		&s.RSquared,
		&s.Logloss,
		&s.Auc,
		&s.Gini,
	); err != nil {
		return binomialModel{}, err
	}
	return s, nil
}

func ScanBinomialModels(rs *sql.Rows) ([]binomialModel, error) {
	structs := make([]binomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s binomialModel
		if err = rs.Scan(
			&s.ModelId,
			&s.Mse,
			&s.RSquared,
			&s.Logloss,
			&s.Auc,
			&s.Gini,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanCluster(r *sql.Row) (Cluster, error) {
	var s Cluster
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.ContextPath,
		&s.ClusterTypeId,
		&s.DetailId,
		&s.Address,
		&s.State,
		&s.Token,
		&s.Created,
	); err != nil {
		return Cluster{}, err
	}
	return s, nil
}

func ScanClusters(rs *sql.Rows) ([]Cluster, error) {
	structs := make([]Cluster, 0, 16)
	var err error
	for rs.Next() {
		var s Cluster
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.ContextPath,
			&s.ClusterTypeId,
			&s.DetailId,
			&s.Address,
			&s.State,
			&s.Token,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanClusterType(r *sql.Row) (clusterType, error) {
	var s clusterType
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return clusterType{}, err
	}
	return s, nil
}

func ScanClusterTypes(rs *sql.Rows) ([]clusterType, error) {
	structs := make([]clusterType, 0, 16)
	var err error
	for rs.Next() {
		var s clusterType
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanClusterYarnDetail(r *sql.Row) (ClusterYarnDetail, error) {
	var s ClusterYarnDetail
	if err := r.Scan(
		&s.Id,
		&s.EngineId,
		&s.Size,
		&s.ApplicationId,
		&s.Memory,
		&s.OutputDir,
	); err != nil {
		return ClusterYarnDetail{}, err
	}
	return s, nil
}

func ScanClusterYarnDetails(rs *sql.Rows) ([]ClusterYarnDetail, error) {
	structs := make([]ClusterYarnDetail, 0, 16)
	var err error
	for rs.Next() {
		var s ClusterYarnDetail
		if err = rs.Scan(
			&s.Id,
			&s.EngineId,
			&s.Size,
			&s.ApplicationId,
			&s.Memory,
			&s.OutputDir,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEngine(r *sql.Row) (Engine, error) {
	var s Engine
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Location,
		&s.Created,
	); err != nil {
		return Engine{}, err
	}
	return s, nil
}

func ScanEngines(rs *sql.Rows) ([]Engine, error) {
	structs := make([]Engine, 0, 16)
	var err error
	for rs.Next() {
		var s Engine
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Location,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEntityType(r *sql.Row) (entityType, error) {
	var s entityType
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return entityType{}, err
	}
	return s, nil
}

func ScanEntityTypes(rs *sql.Rows) ([]entityType, error) {
	structs := make([]entityType, 0, 16)
	var err error
	for rs.Next() {
		var s entityType
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanHistory(r *sql.Row) (History, error) {
	var s History
	if err := r.Scan(
		&s.Id,
		&s.Action,
		&s.IdentityId,
		&s.EntityTypeId,
		&s.EntityId,
		&s.Description,
		&s.Created,
	); err != nil {
		return History{}, err
	}
	return s, nil
}

func ScanHistorys(rs *sql.Rows) ([]History, error) {
	structs := make([]History, 0, 16)
	var err error
	for rs.Next() {
		var s History
		if err = rs.Scan(
			&s.Id,
			&s.Action,
			&s.IdentityId,
			&s.EntityTypeId,
			&s.EntityId,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentity(r *sql.Row) (Identity, error) {
	var s Identity
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Password,
		&s.WorkgroupId,
		&s.IsActive,
		&s.LastLogin,
		&s.Created,
	); err != nil {
		return Identity{}, err
	}
	return s, nil
}

func ScanIdentitys(rs *sql.Rows) ([]Identity, error) {
	structs := make([]Identity, 0, 16)
	var err error
	for rs.Next() {
		var s Identity
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Password,
			&s.WorkgroupId,
			&s.IsActive,
			&s.LastLogin,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentityRole(r *sql.Row) (identityRole, error) {
	var s identityRole
	if err := r.Scan(
		&s.IdentityId,
		&s.RoleId,
	); err != nil {
		return identityRole{}, err
	}
	return s, nil
}

func ScanIdentityRoles(rs *sql.Rows) ([]identityRole, error) {
	structs := make([]identityRole, 0, 16)
	var err error
	for rs.Next() {
		var s identityRole
		if err = rs.Scan(
			&s.IdentityId,
			&s.RoleId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentityWorkgroup(r *sql.Row) (identityWorkgroup, error) {
	var s identityWorkgroup
	if err := r.Scan(
		&s.IdentityId,
		&s.WorkgroupId,
	); err != nil {
		return identityWorkgroup{}, err
	}
	return s, nil
}

func ScanIdentityWorkgroups(rs *sql.Rows) ([]identityWorkgroup, error) {
	structs := make([]identityWorkgroup, 0, 16)
	var err error
	for rs.Next() {
		var s identityWorkgroup
		if err = rs.Scan(
			&s.IdentityId,
			&s.WorkgroupId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanLabel(r *sql.Row) (Label, error) {
	var s Label
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.ModelId,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Label{}, err
	}
	return s, nil
}

func ScanLabels(rs *sql.Rows) ([]Label, error) {
	structs := make([]Label, 0, 16)
	var err error
	for rs.Next() {
		var s Label
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.ModelId,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanMeta(r *sql.Row) (meta, error) {
	var s meta
	if err := r.Scan(
		&s.Id,
		&s.Key,
		&s.Value,
	); err != nil {
		return meta{}, err
	}
	return s, nil
}

func ScanMetas(rs *sql.Rows) ([]meta, error) {
	structs := make([]meta, 0, 16)
	var err error
	for rs.Next() {
		var s meta
		if err = rs.Scan(
			&s.Id,
			&s.Key,
			&s.Value,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanModelCategory(r *sql.Row) (modelCategory, error) {
	var s modelCategory
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return modelCategory{}, err
	}
	return s, nil
}

func ScanModelCategorys(rs *sql.Rows) ([]modelCategory, error) {
	structs := make([]modelCategory, 0, 16)
	var err error
	for rs.Next() {
		var s modelCategory
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanModel(r *sql.Row) (Model, error) {
	var s Model
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.ClusterId,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumn,
		&s.LogicalName,
		&s.Location,
		&s.ModelObjectType,
		&s.MaxRunTime,
		&s.Schema,
		&s.SchemaVersion,
		&s.Created,
	); err != nil {
		return Model{}, err
	}
	return s, nil
}

func ScanModels(rs *sql.Rows) ([]Model, error) {
	structs := make([]Model, 0, 16)
	var err error
	for rs.Next() {
		var s Model
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.ClusterId,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumn,
			&s.LogicalName,
			&s.Location,
			&s.ModelObjectType,
			&s.MaxRunTime,
			&s.Schema,
			&s.SchemaVersion,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanMultinomialModel(r *sql.Row) (multinomialModel, error) {
	var s multinomialModel
	if err := r.Scan(
		&s.ModelId,
		&s.Mse,
		&s.RSquared,
		&s.Logloss,
	); err != nil {
		return multinomialModel{}, err
	}
	return s, nil
}

func ScanMultinomialModels(rs *sql.Rows) ([]multinomialModel, error) {
	structs := make([]multinomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s multinomialModel
		if err = rs.Scan(
			&s.ModelId,
			&s.Mse,
			&s.RSquared,
			&s.Logloss,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPermission(r *sql.Row) (Permission, error) {
	var s Permission
	if err := r.Scan(
		&s.Id,
		&s.Code,
		&s.Description,
	); err != nil {
		return Permission{}, err
	}
	return s, nil
}

func ScanPermissions(rs *sql.Rows) ([]Permission, error) {
	structs := make([]Permission, 0, 16)
	var err error
	for rs.Next() {
		var s Permission
		if err = rs.Scan(
			&s.Id,
			&s.Code,
			&s.Description,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPrivilege(r *sql.Row) (Privilege, error) {
	var s Privilege
	if err := r.Scan(
		&s.Type,
		&s.WorkgroupId,
		&s.EntityType,
		&s.EntityId,
	); err != nil {
		return Privilege{}, err
	}
	return s, nil
}

func ScanPrivileges(rs *sql.Rows) ([]Privilege, error) {
	structs := make([]Privilege, 0, 16)
	var err error
	for rs.Next() {
		var s Privilege
		if err = rs.Scan(
			&s.Type,
			&s.WorkgroupId,
			&s.EntityType,
			&s.EntityId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanProject(r *sql.Row) (Project, error) {
	var s Project
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.ModelCategory,
		&s.Created,
	); err != nil {
		return Project{}, err
	}
	return s, nil
}

func ScanProjects(rs *sql.Rows) ([]Project, error) {
	structs := make([]Project, 0, 16)
	var err error
	for rs.Next() {
		var s Project
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.ModelCategory,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRegressionModel(r *sql.Row) (regressionModel, error) {
	var s regressionModel
	if err := r.Scan(
		&s.ModelId,
		&s.Mse,
		&s.RSquared,
		&s.MeanResidualDeviance,
	); err != nil {
		return regressionModel{}, err
	}
	return s, nil
}

func ScanRegressionModels(rs *sql.Rows) ([]regressionModel, error) {
	structs := make([]regressionModel, 0, 16)
	var err error
	for rs.Next() {
		var s regressionModel
		if err = rs.Scan(
			&s.ModelId,
			&s.Mse,
			&s.RSquared,
			&s.MeanResidualDeviance,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRole(r *sql.Row) (Role, error) {
	var s Role
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Role{}, err
	}
	return s, nil
}

func ScanRoles(rs *sql.Rows) ([]Role, error) {
	structs := make([]Role, 0, 16)
	var err error
	for rs.Next() {
		var s Role
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRolePermission(r *sql.Row) (rolePermission, error) {
	var s rolePermission
	if err := r.Scan(
		&s.RoleId,
		&s.PermissionId,
	); err != nil {
		return rolePermission{}, err
	}
	return s, nil
}

func ScanRolePermissions(rs *sql.Rows) ([]rolePermission, error) {
	structs := make([]rolePermission, 0, 16)
	var err error
	for rs.Next() {
		var s rolePermission
		if err = rs.Scan(
			&s.RoleId,
			&s.PermissionId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanState(r *sql.Row) (state, error) {
	var s state
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return state{}, err
	}
	return s, nil
}

func ScanStates(rs *sql.Rows) ([]state, error) {
	structs := make([]state, 0, 16)
	var err error
	for rs.Next() {
		var s state
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanService(r *sql.Row) (Service, error) {
	var s Service
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.ModelId,
		&s.Name,
		&s.Host,
		&s.Port,
		&s.ProcessId,
		&s.State,
		&s.Created,
	); err != nil {
		return Service{}, err
	}
	return s, nil
}

func ScanServices(rs *sql.Rows) ([]Service, error) {
	structs := make([]Service, 0, 16)
	var err error
	for rs.Next() {
		var s Service
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.ModelId,
			&s.Name,
			&s.Host,
			&s.Port,
			&s.ProcessId,
			&s.State,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanWorkgroup(r *sql.Row) (Workgroup, error) {
	var s Workgroup
	if err := r.Scan(
		&s.Id,
		&s.Type,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Workgroup{}, err
	}
	return s, nil
}

func ScanWorkgroups(rs *sql.Rows) ([]Workgroup, error) {
	structs := make([]Workgroup, 0, 16)
	var err error
	for rs.Next() {
		var s Workgroup
		if err = rs.Scan(
			&s.Id,
			&s.Type,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
