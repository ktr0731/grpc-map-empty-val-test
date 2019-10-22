Start the server that returns a response has a value is not nil.
``` bash
$ make run/server/nonnil
```

Start the server that returns a response has a value is nil.
``` bash
$ make run/server/nil
```

Run the generated client (it will not panic on `run/server/nil`).
``` bash
$ make run/client
```

Run gRPCurl (it will panic on `run/server/nil`).
``` bash
$ make run/grpcurl
```

Run Evans (it will panic on `run/server/nil`).
``` bash
$ make run/evans
```
