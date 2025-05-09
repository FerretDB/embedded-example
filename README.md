# Embedded FerretDB

This repository contains an example of embedding FerretDB v2 into the Go program.

To see it in action:

1. Start PostgreSQL with [the DocumentDB extension](https://github.com/microsoft/documentdb) with:
```
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password ghcr.io/ferretdb/postgres-documentdb:17
```

2. Start example with:
```
git clone https://github.com/FerretDB/embedded-example.git
cd embedded-example
go run main.go
```

3. Connect with `mongosh`:

```
$ mongosh
Current Mongosh Log ID: 681e49a5494194df75a2277c
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.1
Using MongoDB:          7.0.77
Using Mongosh:          2.5.1

For mongosh info see: https://www.mongodb.com/docs/mongodb-shell/

------
   The server generated these startup warnings when booting
   2025-05-09T18:29:58.382Z: Powered by FerretDB v2.2.0 and DocumentDB 0.103.0 (PostgreSQL 17.5).
   2025-05-09T18:29:58.382Z: Please star 🌟 us on GitHub: https://github.com/FerretDB/FerretDB.
   2025-05-09T18:29:58.382Z: The telemetry state is undecided. Read more about FerretDB telemetry and how to opt out at https://beacon.ferretdb.com.
------
test> db.test.find()
[ { _id: ObjectId('67ec0f1dd93b7a55588017e3'), hello: 'world' } ]

test>
```
