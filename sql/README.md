# SQL

- [SQL](#sql)
  - [Entities](#entities)
  - [Relationships](#relationships)
  - [Usecases](#usecases)
    - [Author](#author)
      - [Queries](#queries)
      - [Commands](#commands)


## Entities

1. Author
2. Book
3. Publisher

## Relationships

- Author <- m2m -> Book
- Book -> Publisher

## Usecases

### Author

#### Queries

- List with `books_count`
- Get with `books`, `publisher`

#### Commands

- Create
- Update
- Upsert
- Delete
  

