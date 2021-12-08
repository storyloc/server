# Usage

## Run server

`go run cmd/server/main.go server --server-graphiql`

## Open Browser

on `localhost:3000`

Example Queries:

```
query getStories {
  stories {
    id
    name
    owner {
      id
    }
    createdAt
  }
}

query getNameById($userId: String!) {
  profile(id: $userId) {
    name
  }
}

query getProfile($userId: String!) {
  profile(id: $userId) {
    id
    name
    createdAt
    updatedAt
  }
}

mutation createProfile($userName:String!) {
  createProfile(input: {name: $userName}) {
    id
  }
}

mutation createStory($storyName:String!, $userId:String!) {
  createStory(
    input:{
      name: $storyName,
      ownerId: $userId
    }
  ) {
    id
  }
}

query schema {
  __schema {
    types {
      name
      description
    }
  }
}
```

Query Variables: (needs correct ID)
```
{
  "userName": "Felix",
  "userId": "N9klwfiU_3DySdUFI8DZP",
  "storyName": "Waterfall"
}
```

1. Run "createProfile"
2. Copy `id`to Query Variables `userId`
3. Run "createStory"
4. Run "getStories"

Repeat Step 3 and change `storyName` and / or `userId` to create more stories.


## Check Data

- file system: `ls -la ~/.storyloc/`

Remove this folder to clean data.

# Download schema for Clients

To download the schema install Apollo CLI.
On macOS use `brew install apollo-cli`, there are some fixed issued included. 

Donwload schema.json by running

`apollo schema:download --endpoint=http://localhost:3000/graphql schema.json`

Note: the server has to be running.

# Usage with IPFS
todo
