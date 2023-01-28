# Pagination Example

This code demonstrates how to implement pagination in a Go web server. It sets up an HTTP server that handles a /users endpoint and returns a slice of users for the current page.

## Getting Started

1.Make sure you have Go installed on your machine. You can download it from the official website: [Golang](https://golang.org/dl/)

2.Clone the repository to your local machine using git clone 
```bash 
https://github.com/olgunozer/paginationExampleWithGo.git 
```

3.Open a terminal in the project's root directory and run ```bash go run main.go ```

4.Open your browser and navigate to ```bash http://localhost:8080/users ``` to see the pagination in action

## Endpoint
'/users' endpoint accepts two query parameters:

- 'page': The page number you want to retrieve, with a default value of 1
- 'page_size': The number of items per page, with a default value of 10

For example, ```bash http://localhost:8080/users?page=2&page_size=5 ``` will retrieve the second page with 5 items per page.

## Response
The server will return a JSON object with the following properties:

- 'page': The current page number
- 'page_size': The number of items per page
- 'data': An array of User objects for the current page

Additionally, the response will include some headers:

- 'X-Total-Pages': The total number of pages
- 'X-Total-Items': The total number of items

### Example
```go

{
    "page": 1,
    "page_size": 10,
    "data": [
        {
            "ID": 1,
            "Name": "User 1"
        },
        {
            "ID": 2,
            "Name": "User 2"
        },
        // ...
    ]
}
```

## Note
This is just an example, you may want to use some database to get pagination data in real-life projects.
