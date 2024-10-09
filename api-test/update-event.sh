curl -X PUT \
-H 'content-type: application/json' \
-d '{"title": "Updated Event_1","description": "Updated Description 1","location": "London","datetime": "2026-01-01T15:40:00.000Z"}' \
http://localhost:8080/events/1