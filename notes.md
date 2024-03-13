* No authorizations for admins. Only admins can add products. Can add later on if needed
* Connecting with postgresql DB: One simple way of connecting is to manually add tables and insert records come Go server
  * Check out DbConnection()
  * The host, user name, password, dbname and port must be named correctly
  * db.Exec() will execute the given SQL command on the database
  *  