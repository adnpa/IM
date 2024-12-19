package graph

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

//type Node struct {
//}
//
//type Relation struct {
//}

type Conn struct {
	driver neo4j.DriverWithContext
}

func (c *Conn) InsertNode(ctx context.Context, node neo4j.Node) {
	//neo4j.ExecuteQuery()
}

func (c *Conn) InsertPath(path neo4j.Path) {

}

func (c *Conn) Close() {
	_ = c.driver.Close(context.Background())
}

// insert node
func init() {
	ctx := context.Background()
	// URI examples: "graph://localhost", "graph+s://xxx.databases.graph.io"
	dbUri := "<URI for Neo4j database>"
	dbUser := "<Username>"
	dbPassword := "<Password>"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established.")
}
