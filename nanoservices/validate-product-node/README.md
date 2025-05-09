
# Validate Product Service (Node.js)

This nanoservice validates a product ID using a simple rule (must start with `PROD-`).

## Tech Stack
- Node.js
- Express

## Run Locally
```bash
npm install
npm start
```

## Docker
```bash
docker build -t validate-product .
docker run -p 3000:3000 validate-product
```

## Endpoint
POST `/validate-product`
```json
{
  "productId": "PROD-123"
}
```
