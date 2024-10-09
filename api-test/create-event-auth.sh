curl -X POST \
-H 'content-type: application/json' \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJleHAiOjE3MjU3ODU3OTMsInVzZXJJZCI6MX0.0Pi5qvBCLo-lyNTeuQocF0doB1TxyM_l9RAd1D8OhSE" \
-d '{"title": "Music","description": "Harry Styles Concert","location": "All Around","datetime": "2025-01-01T15:40:00.000Z"}' \
http://localhost:8080/events