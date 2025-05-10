# Database Error Handling Guidelines

We have two utility functions for handling database errors:

1. `errorutils.AnalyzeDBErr(err)` - Basic error analysis
2. `errorutils.AnalyzeDBResults(err, result)` - Extended analysis including "no rows affected"

## When to use which function:

### Use `AnalyzeDBResults(err, result)` for:

- UPDATE operations where the record should exist
- DELETE operations where the record should exist
- Any operation where "no rows affected" should be treated as an error

Example:

```go
func (r *repository) UpdateUser(ctx context.Context, id uuid.UUID, req UpdateUserReq) error {
    query := `UPDATE users SET ... WHERE id = :id`
    result, err := r.db.NamedExecContext(ctx, query, params)
    return errorutils.AnalyzeDBResults(err, result)
}
```
