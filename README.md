# Microexchange

Microexchange is a monolith exchange implementation.  It is composed of 4 major components:
`orderserver`, `orderload`, `orderbook` and the `dispatcher`.

## Components

1. orderserver: takes in orders and records clients in a subcomponent `clientManager`
2. orderload: receives orders from the `orderserver` sorts them by time and the submits them to the `orderbook`-
              changes the state on the orderbook.
3. orderbook: provides api to manage the orderbook state.
4. dispatcher: dispatches state changes on the orderbook to the given clients.

## Future improvements

1. complete clientManager
1. tie dispatcher to clientManager
1. provide main function - tie the components together
1. provide fault tolerance through etcd, or 2 phase commits for completed trades. i.e . snapshot the database
