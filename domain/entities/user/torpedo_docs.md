# User entity

--------



## Entity Resource Name (TRN)

The entity `TRN` (Torpedo Resource Name) is the unique resource name into the torpedo platform.

```text
trn::entity::user::1234567890QWERTYUIOP
```

### TRN JSON representation
```json
{
    "metadata":{
        "objectType":"TRN",
        "exportedDateMillis":1682083031839
    },
    "id":"1234567890QWERTYUIOP",
    "kind":"entity",
    "name":"user"
}
```

## Schema

The schema of the `user` entity is given by the next fields:

| Name      | Type    | Description                                           |
|-----------|---------|-------------------------------------------------------|
| id        | `ulid`    | The entity unique identifier                      |
| created   | `date`    | The entity creation timestamp in milliseconds UTC     |
| updated   | `date`    | The entity modification timestamp in milliseconds UTC |
| name | `string`  | The user full name |
| email | `string`  | The user contact email |
| password | `string`  | The user system password |
| plan | `string`  | The user membership plan |
| miles | `int64`  | The accumulated flyer miles |


## Adapters

### Inputs

#### HTTP
The HTTP input exposes a CRUD API:

```
- Create: [POST]   /users
- Read:   [GET]    /users/:id
- Update: [PUT]    /users/:id
- Delete: [DELETE] /users/:id
```

And a query API based on the [Torpedo Query Language:](https://darksubmarine.com/docs/torpedo/tql.html)

```
- TQL: [POST] /users/query
```

### Outputs


 - Memory
 

 - MongoDB 
   - Collection: users


