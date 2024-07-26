# Embedded FerretDB

This repository contains an example of embedding FerretDB into the Go program.

To see it in action:

1. Start PostgreSQL for storing data with:
```
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres
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
Current Mongosh Log ID: 66a36fe425b4669aac9f9769
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.12
Using MongoDB:          7.0.42
Using Mongosh:          2.2.12

For mongosh info see: https://docs.mongodb.com/mongodb-shell/

------
   The server generated these startup warnings when booting
   2024-07-26T09:44:04.726Z: Powered by FerretDB unknown and PostgreSQL 16.3.
   2024-07-26T09:44:04.726Z: Please star us on GitHub: https://github.com/FerretDB/FerretDB.
------

test> db.test.find()
[ { _id: ObjectId('66a36fdbf693bdd9d33739e8'), hello: 'world' } ]
test>
```
