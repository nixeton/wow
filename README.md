# Word of Wisdom

Design and implement “Word of Wisdom” tcp server.
* TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
* The choice of the POW algorithm should be explained.
* After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
* Docker file should be provided both for the server and for the client that solves the POW challenge

#

* Service uses clean architectural design
* Implemented hash algorithm for POW. Possible change complexity of work. Can be dynamically changeable if, for example, analytics detect DDOS.
* Implemented challenge-response protocol 


 Run server locally:
```
 make start-server
```

Run client locally
```
make start-client
```

 Run docker
 ```
 make compose-up
```
 
 Stop Docker
 ```
 make compose-down
```
