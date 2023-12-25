

```markdown
# Golang CRUD Operations with Hardcoded Data

This repository demonstrates basic CRUD (Create, Read, Update, Delete) operations in Golang using hardcoded data.
It provides a simple example to understand how to implement these operations in a Golang application.

## Project Structure

- `main.go`: The main Golang file containing the CRUD operations.
- `README.md`: Documentation explaining the purpose and usage of the project.

## Prerequisites

Make sure you have Golang installed on your machine. If not, follow the instructions on the [official Golang website](https://golang.org/doc/install).
```
## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/EuclidStellar/CRUD-API-Go/
   ```

2. Navigate to the project folder:

   ```bash
   cd CRUD-API-Go
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. Open your web browser and access the following endpoints:

   - List all movies: [http://localhost:8000/movies](http://localhost:8000/movies)
   - Get a specific movie by ID: [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id})
   - Create a new movie: Send a POST request to [http://localhost:8000/movies](http://localhost:8000/movies) with a JSON payload.
   - Update a movie by ID: Send a PUT request to [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id}) with a JSON payload.
   - Delete a movie by ID: Send a DELETE request to [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id}).


### Create a new movie

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": "6", "title": "New Movie", "year": "2023", "director": {"firstname": "John", "lastname": "Doe", "age": 40}}' http://localhost:8000/movies
```

### Update a movie

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": "6", "title": "Updated Movie", "year": "2023", "director": {"firstname": "John", "lastname": "Doe", "age": 40}}' http://localhost:8000/movies/6
```

### Delete a movie

```bash
curl -X DELETE http://localhost:8000/movies/6
```

## Running Locally

To run the project locally, make sure you have Golang installed. Follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/EuclidStellar/CRUD-API-Go/
   ```

2. Navigate to the project folder:

   ```bash
   cd CRUD-API-Go
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. Access the endpoints as described in the "Getting Started" section.

## Contributions

Contributions are welcome! If you have improvements, suggestions, or bug fixes, feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

- [Gaurav](https://github.com/euclidstellar)
```

Customize the content further based on your repository's specifics.
```
