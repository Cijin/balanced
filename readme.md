## Run load balancer using:

`go run *.go`
or 
`go build -o lb && ./lb`

## Parallel requests

There are about 10,000 urls in the config file, you can use this command to run and get response 
from the backend serers: `curl --parallel --parallel-immediate --parallel-max 3 --config urls.txt` 
