package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/thejaredhooper/util/common/lib"
	m "github.com/thejaredhooper/util/common/lib/mongo"
	"github.com/thejaredhooper/util/common/lib/mongo/models"
)

type Application struct {
	Host string
	Port int

	Client  *mongo.Client
	Context context.Context
}

func (app *Application) GetCollections(database string) (collections []models.Collection) {
	cursor, err := app.
		Client.
		Database(database).
		ListCollections(app.Context, bson.M{})

	if err != nil {
		lib.Fatal("GetCollections - Failed to get Collections", err)
	}

	defer cursor.Close(app.Context)

	for cursor.Next(app.Context) {
		var collection models.Collection
		err := cursor.Decode(&collection)

		if err != nil {
			lib.Fatal("Failed to decode Collection into its model", err)
		} else {
			collections = append(collections, collection)
		}
	}

	return
}

func (app *Application) commandLine() {
	hPtr, pPtr := m.FlagsNoParse()

	flag.Parse()

	app.Host = *hPtr
	app.Port = *pPtr
}

func (app *Application) mongo() {
	connection := m.ConnectionString(app.Host, app.Port)
	app.Client, app.Context = m.MustConnect(connection)
}

func Startup() (app Application) {
	fmt.Print(lib.LineBreak + "\n Starting..." + lib.LineBreak)

	app.commandLine()
	app.mongo()

	return
}
