# Todo

1. add ocpp start up process
    1. send boot notification once on connect
    2. update heartbeat interval
    3. send connector states on ocpp ready event

2. connector
    1. implement state management to process transactions

3. add transaction handler
    1. link request and response
    2. store events in file
    3. delete events from file if they have been fulfilled

4. simulate hardware such as:
    1. RIFD
