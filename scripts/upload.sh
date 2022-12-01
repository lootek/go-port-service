#!/bin/bash

curl -X POST http://localhost:8000/upload \
  -F "file=@../testdata/ports.json" \
  -H "Content-Type: multipart/form-data"
