# Go GraphQL PostgreSQL

```bash
docker-compose up -d
docker exec -i postgres psql -U user go_graphql_postgresql < data/db.sql
```

[http://localhost:8088/graphql](http://localhost:8088/graphql)

```graphql
{
 "query": "{users(name:\"User%\"){id, name, age, active}}"
}
```
