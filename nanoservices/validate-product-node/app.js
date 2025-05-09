
const express = require('express');
const app = express();
app.use(express.json());

app.post('/validate-product', (req, res) => {
  const { productId } = req.body;
  if (!productId) return res.status(400).json({ error: 'Missing productId' });

  const isValid = productId.startsWith('PROD-');
  return res.json({ valid: isValid });
});

app.listen(3000, () => console.log("Node nanoservice running on port 3000"));
