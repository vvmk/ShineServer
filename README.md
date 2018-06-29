# bounce

I'm developing a client app against a multi-service backend (spring) and i'm noticing that locally 
running 6 java services plus the db and the client app was not a sustainable development practice. 
Here entereth bounce: a tiny-in-comparison web service mocking the API of its big brother (and a chance 
to exercise my Go muscles). 

**6/29/18:** Go is fun. I'm just going to refactor this into the prod server.
