#!/bin/bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data ' {"title":"Test title", "subTitle":"Test subtitle", "text":"Test text"}' \
  http://localhost:3000/api/note