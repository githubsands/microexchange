# Microexchange

Microexchange is a monolith exchange implementation.  It is composed of 4 major components:
`orderserver`, `orderload`, `orderbook` and the `dispatcher`.

## Components

1. orderserver: takes in orders and records clients in a subcomponent `clientManager`
2. orderload: receives orders from the `orderserver` sorts and then submits them to the `orderbook`-
              changes the state on the orderbook.
3. orderbook: provides api to manage the orderbook state.
4. dispatcher: dispatches state changes on the orderbook to the given clients.

## Future improvements

1. complete clientManager. cache clients in a LIFO queue
1. tie dispatcher to clientManager
1. provide main function - tie the components together
1. completed trades should be posted in sqlLite with some fault tolerance (two phase commits)
1. cleanup linked-list implmentation (needs unsafe freepointers to fully replicate voyager's C implementation)

## Inspired by

1. Voyager's approach on leveraging preallocated memory over standard red black tree approach: https://gist.github.com/druska/d6ce3f2bac74db08ee9007cdf98106ef
2. Limit order books a queueing systems perspective: https://www0.gsb.columbia.edu/faculty/cmaglaras/papers/IC-Lectures-2015.pdf
3. Market Microstructure in Pratice https://www.amazon.com/Market-Microstructure-Practice-Charles-Albert-Lehalle/dp/9813231122/ref=pd_lpo_3?pd_rd_i=9813231122&psc=1
4. The two-phase-commit https://bravenewgeek.com/tag/two-phase-commit/
