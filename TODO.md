# Todo

1. begin ocpp implementation
    1. add request (done)
    2. add response (done)
    3. handle call request
    4. handle call response

2. add ocpp start up process
    1. send boot notification once on connect
    2. update heartbeat interval
    3. send connector states on ocpp ready event

3. connector
    1. implement state management to process transactions

4. add transaction handler
    1. link request and response
    2. store events in file
    3. delete events from file if they have been fulfilled

5. simulate hardware such as:
    1. RIFD
