package mongo

import (
	"context"
	"fmt"
	m "github.com/mongodb/mongo-go-driver/mongo"
)

func ConnectionString(host string, port int) string {
	return fmt.Sprintf("mongodb://%v:%v", host, port)
}

func MustConnect(connection string) (client *m.Client, ctx context.Context) {
	var err error

	if client, err = m.NewClient(connection); err != nil {
		fmt.Println(fmt.Sprintf(`Error establishing Mongo Client (With Connection String %v): %v`, connection, err.Error()))
		panic(err)
	}

	ctx = context.Background()

	if err = client.Connect(ctx); err != nil {
		panic(err)
	}

	return
}

func Connect(connectionString string) (client *m.Client, ctx context.Context, err error) {

	if client, err = m.NewClient(connectionString); err != nil {
		fmt.Println(fmt.Sprintf(`Error establishing Mongo Client (With Connection String %v): %v`, connectionString, err.Error()))
		return
	}

	ctx = context.Background()

	if err = client.Connect(ctx); err != nil {
		fmt.Println(fmt.Sprintf(`Error establishing Mongo Client Connection: %v`, err.Error()))
	}

	return
}
