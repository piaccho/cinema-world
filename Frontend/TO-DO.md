### To do:
- [ ] JWT Auth
    - [ ] /refresh endpoint
    - [ ] /login endpoint
    - [ ] /register endpoint
    - [ ] /logout endpoint
- [ ] Routing
- [ ] Carousel -> Movie OverView Page
- [ ] Header Admin/Login icon
- [ ] Pages:
    - [ ] ChooseSeat seats visualization
    - [x] Repretoires Page
        - [x] Calendar
        - [ ] Display Repertoire (default from today)
    - [x] Login Page
    - [x] Register Page
    - [x] Movie Overview Page
        - [x] Movie Description
        - [ ] Want to watch button
        - [x] Buy Ticket - get repertoire with the first next showing of specific movie 
        - [x] List of showings
    - [ ] Admin panel
        - [ ] Header Icon
        - [ ] Delete movies - list
    - [ ] User Panel
        - [ ] Header Icon
        - [ ] Want to watch - list
        - [ ] Reservations - list
    

Minor fixes and features:
- [ ] fix searchbar from same route
- [ ] Split types into Model and Props
- [x] Fix footer margin copyrights
- [ ] Fixed Header taking space
- [ ] Fix Footer to be sticky on bottom of parent element
- [ ] Night mode
- [ ] Extract FooterColumn component
- [ ] Extract CarouselSection component
- [ ] Add component att to every MUI component
- [ ] Make scrollbar not taking space
- [ ] Constraints on calendar and disable past dates
- [ ] Contact Page
- [ ] Newsletter Page
- [ ] LocationInfo Page


Endpoints:
- GET all movies by category
- GET top 20 popular movies
- GET top 20 upcoming movies
- GET all showings by id
- GET all showings by dacie
- POST /register - {login, password, firstname, lastname } -> {tokenJWT} - powinien zwracać token JWT
- POST /login - {login, password} -> {tokenJWT} - powinien zwracać token JWT
- DELETE /logout - { tokenJWT } - powinien usuwać token JWT w bazie
- POST /reserve - { userId, showingId } 
- POST /addToWatch - { userId, movieId } 