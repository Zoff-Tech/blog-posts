
# Apply Discount Service (Python)

This nanoservice calculates a discount based on a coupon code.

## Tech Stack
- FastAPI
- Uvicorn

## Run Locally
```bash
pip install -r requirements.txt
uvicorn apply_discount:app --reload --port 8000
```

## Docker
```bash
docker build -t apply-discount .
docker run -p 8000:8000 apply-discount
```

## Endpoint
POST `/apply-discount`
```json
{
  "code": "SAVE10",
  "subtotal": 100.0
}
```
