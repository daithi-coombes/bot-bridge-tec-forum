# TECommons Discourse & Conviction Voting Bot

This bot solves a fringe case that will automatically update a forum post with a link to the relevant proposal.
More information on the required solution: 


### Usage
```
go run main.go \
    --endpoint wss://rinkeby.infura.io/ws/v3/1234 \
    --discourse-key 1234 \
    --discourse-endpoint http://localhost:9292 \
    --dao 0x0
```

### Tests
```
make test
```