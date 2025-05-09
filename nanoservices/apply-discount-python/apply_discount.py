
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class DiscountRequest(BaseModel):
    code: str
    subtotal: float

@app.post("/apply-discount")
def apply_discount(req: DiscountRequest):
    if req.code == "SAVE10":
        discount = req.subtotal * 0.10
    else:
        discount = 0.0
    return {"discount": round(discount, 2)}
