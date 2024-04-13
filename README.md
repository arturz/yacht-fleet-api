# Yacht Charter Management API

A RESTful API for managing yacht charters, marinas, and yacht migrations between ports.

## Resource Hierarchy

| URI                   | GET                   | POST                    | PUT                      | DELETE        |
| --------------------- | --------------------- | ----------------------- | ------------------------ | ------------- |
| /charters             | List charters         | Charter a yacht         | -                        | -             |
| /charters/{cid}       | Charter details       | -                       | Update charter           | End charter   |
| /yachts               | List yachts           | Add new yacht           | -                        | -             |
| /yachts/{yid}         | Yacht details         | -                       | Update yacht information | Remove yacht  |
| /marinas              | List marinas          | Add marina              | -                        | -             |
| /marinas/{mid}        | Marina details        | -                       | Update marina details    | Remove marina |
| /marinas/{mid}/yachts | List yachts in marina | Assign yacht to marina  | -                        | -             |
| /migrations           | Migration history     | Migrate yacht to marina | -                        | -             |
| /migrations/{mig_id}  | Migration details     | -                       | -                        | -             |

## Features

### Collection Resources

Resources representing collections of other resources with pagination support:

- `GET /yachts` - See `TestListYachtsPagination` for implementation details

### Controller Resources

Resources enabling atomic updates of multiple related resources:

- `POST /migrations`
  - Updates `yacht.MarinaID`
  - Creates new migration record

### POST-Once-Exactly Pattern

Resources implementing exactly-once semantics to prevent duplicate submissions:

- `GET /tokens` - Generate single-use token
- `POST /migrations?token=...` - See `TestCreateMigration` for implementation details

### Optimistic Concurrency Control

Implementation of lost update prevention using verification mode:

- `PUT /yachts/<yid>` - See `TestUpdateYacht` for implementation details

## Testing

The API includes comprehensive test coverage. Key test cases are referenced in the features section above.
