
# Send Email Service (Go)

This nanoservice mocks sending an email notification.

## Tech Stack
- Go (net/http)

## Run Locally
```bash
go run send-email.go
```

## Docker
```bash
docker build -t send-email .
docker run -p 8080:8080 send-email
```

## Endpoint
POST `/send-email`
```json
{
  "to": "user@example.com",
  "message": "Welcome!"
}
```
