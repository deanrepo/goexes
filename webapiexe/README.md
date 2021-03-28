### Build image 
```bash
    $ docker build -t gowebapi .
```

### Run a container from the built image
```bash
    $ docker run -d -p 3333:3000 -e APPID=3333 --name mywebapi gowebapi
```

### Test
open a browser and input the url below

* http://localhost:3333/ 
* http://localhost:3333/health_check