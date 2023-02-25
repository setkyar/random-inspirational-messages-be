# random-inspirational-messages

Printing out random inspirational messages.

### Build
```
docker build -t setkyar/random-inspirational-messages-be:1.0.1 . --platform linux/amd64
```

### Run and Test

```
docker run -p 8080:8080 setkyar/random-inspirational-messages-be:1.0.1
curl -XGET http://localhost:8080/message
```

### Push

```
docker push setkyar/random-inspirational-messages-be:1.0.1
```