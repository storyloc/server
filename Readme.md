# Usage

## Run server

`go run cmd/server/main.go server --graphiql=true`

## Open Browser

on `localhost:3000`

Example Queries:

```
query getNameById($id: String) {
  profile(id: $id) {
    name
  }
}

query getProfile($id: String) {
  profile(id: $id) {
    name
    createdAt
  }
}

mutation createProfile {
  createProfile(name: "Andrej") {
    id
  }
}
```

Query Variables: (needs correct ID)
```
{
  "id": "9UFFp7veQ2AH-lbsgu7dg"
}
```

## Check Data 

- file system: `ls -la ~/.storylock/`

# Usage with IPFS
todo


# Download schema

To download the schema install Apollo CLI.
On macOS use `brew install apollo-cli`, there are some fixed issued included. 

Donwload schema.json by running

`apollo schema:download --endpoint=http://localhost:3000/graphql schema.json`

Note: the server has to be running.
