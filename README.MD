## Project Structure

```
├── internal
|   ├── config                  # configuration variables and functions
|   ├── controller              # api controller or usecase
|   ├── db                      # database config
|   ├── httpsvc                 # handle http req and res
|   ├── model                   # database model
|   ├── repository              # handle data and database func
|   ├── view                    # User interface
├── main.go
```

## How to run:

```
- Run docker compose to create container for postgres
- Create database 'weight' inside the container
- Create .env file and copy the variables from .env.example file to .env
- Fill db username, host and password
- type go run main.go in the terminal to start server
- In view folder, open index.html and preview the page
- The web app is ready
```
