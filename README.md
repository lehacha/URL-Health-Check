# URL-Health-Check

URL-Health-Check is a simple Go application to check the health of URLs. It can ping multiple URLs for their active status.

## Running Locally

To run the URL-Health-Check locally, use the following command:

```bash
go run cmd\main.go
```

## Testing the URL Health Check API
The API can be tested with various sets of URLs to check their health status. Below are some examples of how to test the API using curl.

### Test with Standard URLs
This test checks the health of a few standard URLs:

```bash
curl --location 'http://localhost:8080/pingURLs' \
--header 'Content-Type: application/json' \
--data '{
    "links": [
        "http://example.com",
        "http://example.com/resource",
        "http://youtube.com",
        ""
    ]
}'
```

### Test with Timeout
This test includes a URL that will timeout, simulating a scenario where a URL takes too long to respond:

```bash
curl --location 'http://localhost:8080/pingURLs' \
--header 'Content-Type: application/json' \
--data '{
    "links": [
        "http://example.com",
        "http://example.com/resource",
        "http://youtube.com",
        "http://localhost:8080/delay/31"
    ]
}'
```

Note: The last URL in this set (http://localhost:8080/delay/31) is configured to respond after a 30 sec delay.