# Trip entity

--------



## Entity Resource Name (TRN)

The entity `TRN` (Torpedo Resource Name) is the unique resource name into the torpedo platform.

```text
trn::entity::trip::1234567890QWERTYUIOP
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
    "name":"trip"
}
```

## Schema

The schema of the `trip` entity is given by the next fields:

| Name      | Type    | Description                                           |
|-----------|---------|-------------------------------------------------------|
| id        | `ulid`    | The entity unique identifier                      |
| created   | `date`    | The entity creation timestamp in milliseconds UTC     |
| updated   | `date`    | The entity modification timestamp in milliseconds UTC |
| departure | `string`  | The trip departure airport |
| arrival | `string`  | The trip arrival airport |
| miles | `int64`  | The trip miles |
| from | `int64`  | The trip from date |
| to | `int64`  | The trip to date |
| userId | `string`  | The userId relationship |


## Adapters

### Inputs

#### HTTP
The HTTP input exposes a CRUD API:

```
- Create: [POST]   /trips
- Read:   [GET]    /trips/:id
- Update: [PUT]    /trips/:id
- Delete: [DELETE] /trips/:id
```

And a query API based on the [Torpedo Query Language:](https://darksubmarine.com/docs/torpedo/tql.html)

```
- TQL: [POST] /trips/query
```

### Outputs


 - Memory
 

 - MongoDB 
   - Collection: trips


