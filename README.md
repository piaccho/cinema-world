<h1 align="center" id="title">Cinema World 🎬</h1>

## 🔎 Overview  
<p id="description">"Cinema World" is a comprehensive full-stack web application tailored for cinema management. It serves as a centralized platform for users to explore movie offerings, check repertoires, create accounts, book tickets, and provide reviews for movies. With a user-friendly interface and a suite of features, it aims to enhance the movie-going experience for both casual viewers and cinema enthusiasts.</p>

## ✨ Features 

Here're some of the project's best features:

-   **Movie Browsing:** Explore a diverse collection of movies with intuitive navigation like a carousel, and genre selection, and a search bar.
-   **Movie Details:** Access comprehensive movie information and details including plot summaries, duration, languages, and trailers.
-   **Repertoires Viewing:** You can view showing schedules for each movie, which include details like time, price, available seats, and the type of showing.
-   **User Registration:** Authentication and registration as a new user allow you to unlock more features.
-   **Ticket Booking:** Book tickets for desired movies and decide which places you prefer.
-   **Movie Reviews:** You can express your opinion about the film by leaving a rating and review.
-   **Admin Panel:** Manage movie listings, user accounts, and other resources through a dedicated admin panel.


<div id="used-technologies" />

## ⚙️ Used technologies


- **Frontend**: React with TypeScript (Vite.js)
- **Backend**: Golang with Gin
- **Database**: MongoDB

The application is containerized using <i>Docker Compose</i>, streamlining the deployment process and facilitating deployment across different environments.


<div id="setup" />

## 🛠️ Setup 

This application is containerized using Docker Compose, which makes the setup process straightforward. Here are the steps to get it up and running:

1. **Prerequisites**: 

Make sure you have Docker and Docker Compose installed on your machine. If not, you can download them from [the official Docker website](https://docs.docker.com/get-docker/).

2. **Clone the Repository**

Clone this repository to your local machine using this command:
```bash
git clone https://github.com/piaccho/cinema-world.git
```
Navigate to the cloned project directory:
```bash
cd cinema-world
```

3. **Build and Run the Application**: 

Now you should be in project root directory where the `docker-compose.yml` file is located. Run the following command to build and start the application:

```bash
docker compose up --build
```

This command builds the Docker images (if they don’t exist) and starts the containers. Your application should now be running at  http://localhost:5000/. That’s it! You have successfully set up and run the application.

4.  **Stop the Application**: 

To stop the running services, use the following command:

```bash
docker compose down
```
