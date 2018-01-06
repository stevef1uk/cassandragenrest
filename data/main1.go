/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	var id int
	var name string
	var events [] int
	var sex bool

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	if err := session.Query(`SELECT id, name, events, sex FROM accounts2 WHERE id = ? LIMIT 1`,
		"1").Consistency(gocql.One).Scan(&id, &name, &events, &sex); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Accounts:", id, name, events, sex)

  // Now try a Setvar events [] int

	var events1[] int64

	if err := session.Query(`SELECT id, name, events, sex FROM accounts3 WHERE id = ? LIMIT 1`,
		"1").Consistency(gocql.One).Scan(&id, &name, &events1, &sex); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Accounts:", id, name, events1, sex)


	session.Close()
}
