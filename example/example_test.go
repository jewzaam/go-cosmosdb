package example

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"

	"github.com/jim-minter/go-cosmosdb/example/cosmosdb"
	"github.com/jim-minter/go-cosmosdb/example/types"
)

const (
	dbid      = "testdb"
	collid    = "people"
	triggerid = "trigger"
	personid  = "jim"

	triggerbody = `function trigger() {
	var request = getContext().getRequest();
	var body = request.getBody();
	var ts = new Date();
	body["updateTime"] = ts.getTime();
	request.setBody(body);
}`
)

func TestE2E(t *testing.T) {
	ctx := context.Background()
	log := logrus.NewEntry(logrus.StandardLogger())

	account, found := os.LookupEnv("COSMOSDB_ACCOUNT")
	if !found {
		t.Fatal("must set COSMOSDB_ACCOUNT")
	}

	key, found := os.LookupEnv("COSMOSDB_KEY")
	if !found {
		t.Fatal("must set COSMOSDB_KEY")
	}

	jsonHandle := &codec.JsonHandle{
		BasicHandle: codec.BasicHandle{
			DecodeOptions: codec.DecodeOptions{
				ErrorIfNoField: true,
			},
		},
	}

	dbc, err := cosmosdb.NewDatabaseClient(log, http.DefaultClient, jsonHandle, account, key)
	if err != nil {
		t.Fatal(err)
	}

	db, err := dbc.Create(ctx, &cosmosdb.Database{ID: dbid})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", db)

	dbs, err := dbc.ListAll(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", dbs)

	db, err = dbc.Get(ctx, dbid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", db)

	collc := cosmosdb.NewCollectionClient(dbc, dbid)

	coll, err := collc.Create(ctx, &cosmosdb.Collection{
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

	colls, err := collc.ListAll(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", colls)

	coll, err = collc.Get(ctx, collid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", coll)

	pkrs, err := collc.PartitionKeyRanges(ctx, collid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", pkrs)

	triggerc := cosmosdb.NewTriggerClient(collc, collid)

	trigger, err := triggerc.Create(ctx, &cosmosdb.Trigger{
		ID:               triggerid,
		TriggerOperation: cosmosdb.TriggerOperationAll,
		TriggerType:      cosmosdb.TriggerTypePre,
		Body:             triggerbody,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", trigger)

	triggers, err := triggerc.ListAll(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", triggers)

	trigger, err = triggerc.Get(ctx, triggerid)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", trigger)

	dc := cosmosdb.NewPersonClient(collc, collid)

	doc, err := dc.Create(ctx, personid, &types.Person{
		ID:      personid,
		Surname: "Minter",
	}, &cosmosdb.Options{PreTriggers: []string{triggerid}})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	docs, err := dc.ListAll(ctx, nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)

	doc, err = dc.Get(ctx, personid, personid, nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	docs, err = dc.QueryAll(ctx, personid, &cosmosdb.Query{
		Query: "SELECT * FROM people WHERE people.surname = @surname",
		Parameters: []cosmosdb.Parameter{
			{
				Name:  "@surname",
				Value: "Minter",
			},
		},
	}, nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)

	i := dc.ChangeFeed(nil)
	docs, err = i.Next(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)
	if len(docs.People) != 1 || docs.People[0].Surname != "Minter" {
		t.Error(len(docs.People))
	}

	docs, err = i.Next(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)
	if docs != nil {
		t.Error(docs)
	}

	oldETag := doc.ETag
	doc, err = dc.Replace(ctx, personid, &types.Person{
		ID:      personid,
		ETag:    doc.ETag,
		Surname: "Morrison",
	}, &cosmosdb.Options{PreTriggers: []string{triggerid}})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", doc)

	_, err = dc.Replace(ctx, personid, &types.Person{
		ID:      personid,
		ETag:    oldETag,
		Surname: "Henson",
	}, &cosmosdb.Options{PreTriggers: []string{triggerid}})
	if !cosmosdb.IsErrorStatusCode(err, http.StatusPreconditionFailed) {
		t.Error(err)
	}

	docs, err = i.Next(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)
	if len(docs.People) != 1 || docs.People[0].Surname != "Morrison" {
		t.Error(len(docs.People))
	}

	docs, err = i.Next(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", docs)
	if docs != nil {
		t.Error(docs)
	}

	err = dc.Delete(ctx, personid, doc, nil)
	if err != nil {
		t.Error(err)
	}

	err = triggerc.Delete(ctx, trigger)
	if err != nil {
		t.Error(err)
	}

	err = collc.Delete(ctx, coll)
	if err != nil {
		t.Error(err)
	}

	err = dbc.Delete(ctx, db)
	if err != nil {
		t.Error(err)
	}
}
