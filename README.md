Development:

```
docker build -t bbb-api-recordings .
docker run -ti --rm -p 8080:8080 bbb-api-recordings
```

Test requests with:

```
curl -X POST -H "Content-Type: application/json" -d '{ "attributes": { "name": "-- POSTED NAME --", "roomId": "my-amazing-room" } }' -k "http://localhost:8080/recordings"

curl -X GET -H "Content-Type: application/json" -k "http://localhost:8080/recordings"
```
