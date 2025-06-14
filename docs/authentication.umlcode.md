@startuml
actor User
participant "Client App" as Client
participant "Backend Server" as Server
database "Database" as DB

User -> Client : Login
Client -> Server : Send credentials (username/password)
Server -> DB : Validate credentials
DB --> Server : User data (if valid)
Server -> Client : Access Token + Refresh Token

Client -> Server : Request protected API with Access Token
Server -> Server : Verify Access Token (valid)
Server --> Client : 200 OK

== After 15 minutes (token expires) ==

Client -> Server : Request API with expired Access Token
Server -> Server : Check Access Token (expired)
Server --> Client : 401 Unauthorized

Client -> Server : Send Refresh Token
Server -> DB : Validate Refresh Token
DB --> Server : Valid
Server -> Client : New Access Token (+ Refresh Token)

Client -> Server : Retry API with new Access Token
Server -> Server : Verify Access Token (valid)
Server --> Client : 200 OK

@enduml
