package sql

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var (
	testDir = "test_steam_sql_"
	clean   bool
)

func setup() (string, error) {
	// Create Temporary Directory
	dir, err := ioutil.TempDir("", testDir)
	return dir, errors.Wrap(err, "creating temp directory")
}

func Err(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("%s: %+v", message, err)
		t.Fail()
	}
}

func FatalErr(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("%s: %+v", message, err)
		t.FailNow()
	}
}

func Check(t *testing.T, expected, actual interface{}, message string) {
	if actual != expected {
		panic(fmt.Sprintf("%s: expected %v, actual %v", message, expected, actual))
		t.Fail()
	}
}

func FatalCheck(t *testing.T, expected, actual interface{}, message string) {
	if actual != expected {
		t.Errorf("%s: expected %v, actual %v", message, expected, actual)
		t.FailNow()
	}
}

func setupSqlite() (string, string, error) {
	dir, err := setup()
	if err != nil {
		return "", "", err
	}

	dbPath := filepath.Join(dir, "sl.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return "", "", errors.Wrap(err, "opening database")
	}

	tx, err := db.Begin()
	if err != nil {
		return "", "", errors.Wrap(err, "initializing database")
	}
	defer tx.Rollback()

	for tbl, query := range schema {
		_, err := tx.Exec(query)
		if err != nil {
			return "", "", errors.Wrapf(err, "intitializing %s", tbl)
		}
	}

	return dir, dbPath, errors.Wrap(tx.Commit(), "initializing database")
}

func TestMain(m *testing.M) {
	flag.BoolVar(&clean, "clean", true, "Set false to leave temp dir")
	flag.Parse()

	os.Exit(m.Run())
}

func TestSqliteDB(t *testing.T) {

	// -- Setup --

	dir, dbPath, err := setupSqlite()
	if err != nil {
		t.Fatalf("Failed setup: %+v", err)
	}

	// -- Setup --

	ds, err := NewDatastore("sqlite3", dbPath)
	if err != nil {
		t.Errorf("Failed creating database: %+v", err)
		t.FailNow()
	}

	t.Run("Sqlite_Cluster", func(t *testing.T) { testClusterCRUD(t, ds) })
	t.Run("Sqlite_Identity", func(t *testing.T) { testIdentityCRUD(t, ds) })

	if clean {
		os.RemoveAll(dir)
	} else {
		log.Println("Database for TestSqliteDB at", dbPath)
	}
}

func testClusterCRUD(t *testing.T, ds *Datastore) {
	clustersTable := [...]Cluster{
		Cluster{Name: "TestCloud1", ClusterTypeId: ds.ClusterType.External},
		Cluster{Name: "TestCloud2", ClusterTypeId: ds.ClusterType.External},
	}

	// Create
	clusterIds := make([]int64, 0)
	for _, cluster := range clustersTable {
		clusterId, err := ds.CreateCluster(cluster.Name, cluster.ClusterTypeId, WithState(ds.State.Started))
		clusterIds = append(clusterIds, clusterId)
		FatalErr(t, err, "creating cluster")
	}
	Check(t, len(clustersTable), len(clusterIds), "cluster ids returned")

	// Read
	clusters, err := ds.ReadClusters()
	Err(t, err, "reading clusters")
	Check(t, len(clustersTable), len(clusters), "clusters returned")

	cluster, ok, err := ds.ReadCluster(ById(clusterIds[0]))
	Err(t, err, "reading cluster")
	Check(t, true, ok, "cluster exists")
	Check(t, clustersTable[0].Name, cluster.Name, "cluster returned by name")

	// Update
	Err(t, ds.UpdateCluster(clusterIds[0], WithAddress("1.1.1.1"), WithState(ds.State.Started)), "updating cluster")
	cluster, ok, err = ds.ReadCluster(ById(clusterIds[0]))
	Err(t, err, "reading cluster")

	// Delete
	for _, clusterId := range clusterIds {
		Err(t, ds.DeleteCluster(clusterId), "deleting cluster")
	}
}
