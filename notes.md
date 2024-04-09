* No authorizations for admins. Only admins can add products. Can add later on if needed
* Connecting with postgresql DB: One simple way of connecting is to manually add tables and insert records come Go server
  * Check out DbConnection()
  * The host, user name, password, dbname and port must be named correctly
  * db.Exec() will execute the given SQL command on the database
*  Dependency injection: Objects depend on other objects. An object would be integrated with othe objects. Analogy of Macbook using different companies hardrive. Concept of abstraction and interfaces are used so that in the future it is simpler to change. `new` is tight coupling. Inject hard drive object inside can be done by an external dependency.
   *  Dependency injection containers are responsible for creating object. It will then be injected in our class
   *  Helpful for testing, no need to test dependency since loosely copuled
*  Importing functions - If importing lot of functions from same package but present in different directories, give alias names for the directories in the beginning
* Workflow -> Data(Requests and Responses structures) -> Model(Represents table in system) -> Repository(Storage of entity/model bean in system) -> Service(Business logic implementation) -> Controller(Manages REST interface of business logic) -> Router(Responsible for routing requests for APIs) -> Initialize all files later in main file
* If any data is getting repeated, move it to the child rather than having it in the parent. All must be connected through foreign keys. Have single source of truth
* E-commerce vertical domains - SKU, stock keeping, shelf