
# 1. start server
  go run server.go

# 2. basic auth
  username: progmatic
  password: reviews

# 3. test data
 1. POST /api/videos
    {
        "title": "Cool title",
        "description": "desc 1",
        "url": "https://www.youtube.com/embed/qR0WnWL2o1Q",
        "author": {
            "firstname": "test",
            "lastname": "haha",
            "age": 1,
            "email": "eee@email.com"
        }
    }
