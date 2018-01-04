package operations

import (
    middleware "github.com/go-openapi/runtime/middleware"
    //"cassandragenrest/models"
    "github.com/stevef1uk/cassandragenrest/models"
    "github.com/gocql/gocql"
    "fmt"
    "os"
    //"log"
)

/** func Search(params GetAccountsParams) middleware.Responder {
    payload1 := new(models.GetAccountsOKBodyItems)

    num := int64(params.ID)
    payload1.ID = &num
    //name := "me"
    payload1.Name = &params.Name

    var ret  models.GetAccountsOKBody
    ret  = make(models.GetAccountsOKBody,2)
    ret[0] = payload1

    return NewGetAccountsOK().WithPayload(ret)
  }*/

  /*func retRow( id int, name string ) (local_id int, local_name string)  {
      local_id = id;
      local_name = name
      return
  }*/


  var session *gocql.Session

  func SetUp() {
      fmt.Println("Connecting to Cassandra")
      cluster := gocql.NewCluster(os.Getenv("REST1_SERVICE_HOST"))
      cluster.Keyspace = "demo"
      cluster.Consistency = gocql.One
      session, _ = cluster.CreateSession()
  }

  func Stop() {
    fmt.Println("Stopping Cassandra")
    session.Close()
  }
  func Search(params GetAccountsParams) middleware.Responder {
      var id int;
      var name string

      fmt.Println("Id = ", params.ID)
      fmt.Println("Name = ", params.Name)
      if err := session.Query(`SELECT id, name FROM accounts WHERE id = ? and name = ?`,
                              params.ID, params.Name).Consistency(gocql.One).Scan(&id, &name); err != nil {
          fmt.Println("No data? ", err)
      }
      var ret  models.GetAccountsOKBody
      fmt.Println("Row  = ", id, name )
      num := int64(id)
      payload1 := new(models.GetAccountsOKBodyItems)
      payload1.ID = &num
      payload1.Name = &name
      ret  = make(models.GetAccountsOKBody,1)
      ret[0] = payload1


      /*Code to retrive more than one row
      iter := session.Query(`SELECT id, name FROM accounts WHERE id = ?`,
                              params.ID).Consistency(gocql.One).Iter();

      var ret  models.GetAccountsOKBody
      ret  = make(models.GetAccountsOKBody,iter.NumRows())

      fmt.Println("Retried rows count = ", iter.NumRows() )

      for i := 0; iter.Scan(&id, &name); i++  {
        fmt.Println("Rows  = ", id, name )
        payload1 := new(models.GetAccountsOKBodyItems)
        //var id, name = retRow( id, name )
        var id =  id
        var name = name
        num := int64(id)
        payload1.ID = &num
        payload1.Name = &name
        ret[i] = payload1
      }*/

      return NewGetAccountsOK().WithPayload(ret)
    }
