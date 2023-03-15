# Embedded FerretDB

This repository contains an example of embedding FerretDB into the Go program.

To see it in action:

1. Start PostgreSQL for storing data with `docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres`.
2. Start example with:
```
git clone https://github.com/FerretDB/embedded-example.git
cd embedded-example
go run main.go
```

3. Connect with `mongosh`:

```
$ mongosh
Current Mongosh Log ID:	64116d42ca2d5e0cab0c58d8
Connecting to:		mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.0
Using MongoDB:		6.0.42
Using Mongosh:		1.8.0

For mongosh info see: https://docs.mongodb.com/mongodb-shell/

------
   The server generated these startup warnings when booting
   2023-03-15T07:01:22.38Z: Powered by FerretDB unknown and PostgreSQL 15.2.
   2023-03-15T07:01:22.38Z: Please star us on GitHub: https://github.com/FerretDB/FerretDB.
------

test> db.test.find()
[ { _id: ObjectId("62c85f2a3bd8164619555cc0"), hello: 'world' } ]
```
