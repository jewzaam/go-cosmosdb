package example

import (
	"net/http"
	"os"
	"testing"

	"github.com/jim-minter/go-cosmosdb/example/cosmosdb"
	"github.com/jim-minter/go-cosmosdb/example/types"
)

const (
	dbid      = "testdb"
	collid    = "people"
	triggerid = "trigger"
	personid  = "jim"
)

func TestE2E(t *testing.T) {
	account, found := os.LookupEnv("COSMOSDB_ACCOUNT")
	if !found {
		t.Fatal("must set COSMOSDB_ACCOUNT")
	}

	key, found := os.LookupEnv("COSMOSDB_KEY")
	if !found {
		t.Fatal("must set COSMOSDB_KEY")
	}

	dbc, err := cosmosdb.NewDatabaseClient(http.DefaultClient, account, key)
	if err != nil {
		t.Fatal(err)
	}

	db, err := dbc.Create(&cosmosdb.Database{ID: dbid})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", db)

	dbs, err := dbc.ListAll()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", dbs)

	db, err = dbc.Get(dbid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", db)

	collc := cosmosdb.NewCollectionClient(dbc, dbid)

	coll, err := collc.Create(&cosmosdb.Collection{
		ID: collid,
		PartitionKey: &cosmosdb.PartitionKey{
			Paths: []string{
				"/id",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", coll)

	colls, err := collc.ListAll()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", colls)

	coll, err = collc.Get(collid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", coll)

	pkrs, err := collc.PartitionKeyRanges(collid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", pkrs)

	triggerc := cosmosdb.NewTriggerClient(collc, collid)

	trigger, err := triggerc.Create(&cosmosdb.Trigger{
		ID:               triggerid,
		TriggerOperation: cosmosdb.TriggerOperationAll,
		TriggerType:      cosmosdb.TriggerTypePre,
		Body:             "function trigger() {}",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", trigger)

	triggers, err := triggerc.ListAll()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", triggers)

	trigger, err = triggerc.Get(triggerid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", trigger)

	dc := cosmosdb.NewPersonClient(collc, collid)

	doc, err := dc.Create(personid, &types.Person{
		ID:      personid,
		Surname: "Minter",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	docs, err := dc.ListAll()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)

	doc, err = dc.Get(personid, personid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	docs, err = dc.QueryAll(personid, &cosmosdb.Query{
		Query: "SELECT * FROM people WHERE people.surname = @surname",
		Parameters: []cosmosdb.Parameter{
			{
				Name:  "@surname",
				Value: "Minter",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)

	oldETag := doc.ETag
	doc, err = dc.Replace(personid, &types.Person{
		ID:      personid,
		ETag:    doc.ETag,
		Surname: "Morrison",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	_, err = dc.Replace(personid, &types.Person{
		ID:      personid,
		ETag:    oldETag,
		Surname: "Henson",
	})
	if !cosmosdb.IsErrorStatusCode(err, http.StatusPreconditionFailed) {
		t.Error(err)
	}

	err = dc.Delete(personid, doc)
	if err != nil {
		t.Error(err)
	}

	err = triggerc.Delete(trigger)
	if err != nil {
		t.Error(err)
	}

	err = collc.Delete(coll)
	if err != nil {
		t.Error(err)
	}

	err = dbc.Delete(db)
	if err != nil {
		t.Error(err)
	}
}
