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
