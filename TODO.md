## TODO: GENERAL 
- solve spacer in hearder
- lazy loading, loading spinners
- change docker-compose db credentials for better security
- change JWT_key before deployment
- remove env variables for deployment
- Railway/Vercel deployment
- README.md with overview and demo
- change mongo tmpfs (stored in memory then permanently deleted) volume type to some lasting volume type

### TODO: FRONTEND
- Fix width of carousels
- Fix repertoire
- Login
- Sign In
- Logout
- Protected resources
	- Navbar (icons, greetings, etc)
	Admin:
		- admin panel (crud operations)
	User:
		- review
		- reservations
		- add to watch list
		- user panel
- Components
	- ChooseSeat
	- ReviewEditor
	- ReviewsList
	- User Panel - reservations, reviews, watchlist
	- Admin Panel - CRUD operations for resources
- Footer Pages:
	- Contact Page
	- Newsletter Page
	- LocationInfo Page
	- Informations
    

### Minor fixes and features:
- [ ] BE: HTTP Caching (c.Header("Cache-Control", "public, max-age=3600") / c.Header("ETag", computeETag(data))) / for very dynamic data use Redis
- [ ] BE: validate showingId, userId and movieRef in reservations method
- [ ] BE: extract generic method like: GetAll, GetByKeyID,
- [ ] BE: gin-swagger
- [ ] FE: Split types into Model and Props
- [ ] FE: Night mode
- [ ] FE: Add component att to every MUI component
- [ ] FE: Make scrollbar not taking space
- [ ] FE: Constraints on calendar and disable past dates




