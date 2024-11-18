# Todo

1. load configuration
    1. inject payload to charger
    2. forward relevant configuration to each hardware

2. connector
    1. implement state management to process transactions
    2. integrate with event management

3. add transaction handler
    1. link request and response
    2. store events in file
    3. delete events from file if they have been fulfilled

4. begin ocpp implementation
    1. send boot notification once on connect
    2. update heartbeat interval
    3. send connector states on ocpp ready event

5. simulate hardware such as:
    1. RIFD
    2. Main board (DONE)
    3. Network controller. (DONE)
