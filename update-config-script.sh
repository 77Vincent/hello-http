curl -X PATCH http://localhost:8001/plugins/1e3df258-04aa-52d5-92f1-f8f2027fdc73 \
  -H "Content-Type: application/json" \
  -d '{
    "config": {
      "second": 10000,
      "minute": 20000
    }
  }'
