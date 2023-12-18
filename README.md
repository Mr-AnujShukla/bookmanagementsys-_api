# Book Management System

### About the Project:-
**Description:** This project is a simple Book Management System built using the GoFr framework with a MySQL database for storing book information.
The system allows users to perform **CRUD**(_Create, Read, Update, Delete_) operations on a collection of books.

Before starting this project make sure your go is up-to-date and intalled in your system.
if not then you can download it from [here](https://go.dev/).
Once you have done that then you can proceed with the steps below to start the project.

### Setup to get started :
1. In an empty folder, initialize the go module by using:`go mod init <your package_name>`.it is recommended to name your module like this: `go mod init github.com/{USERNAME}/{REPO}`<br>
Note:You have to use terminal to run your command, and you can use any ide like vscode or goland.<br>
2. Now we have to add dependencies:`go get gofr.dev`
3. Create `main.go ` file and add the code in it.
4. Start the server:`go run main.go`, by default server starts at [http://localhost:8000/books](http://localhost:8000/books), but you can override default port by mentioning an environmental variable
like this:`HTTP_PORT=9000 go run main.go`
 or,<br>
by adding an **_configs_** folder and making `.env` file and adding code
    ```
    APP_NAME=test-service
    HTTP_PORT=9000
    ```
   
**_These were the steps to start the project now lets look at how to run the project_**

## Installation
1. Clone the Repository:`git clone https://github.com/Mr-AnujShukla/bookmanagementsys-_api.git`<br>then change directory to:`cd bookmanagementsys-_api`
2. Setup MySql docker container:
```azure
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE books (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL);"
```
3. Update the GoFr configuration in .env:
```azure
-- APP_NAME=your-app-name
-- HTTP_PORT=8000 //its default in my case.

REDIS_HOST=localhost
REDIS_PORT=6379

DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root123
DB_NAME=test_db
DB_PORT=3306
DB_DIALECT=mysql

```
4. Run the GoFr application:`go run main.go`

## Usage commands for CRUD operations:
1. Adding a Book:
```azure
curl --location --request POST 'http://localhost:8000/book/YourBookTitle/YourAuthor'
```
2. Retrieving a Book by ID:
```azure
curl --location --request GET 'http://localhost:8000/book/1'
```
3. Retrieving All Books:
```azure
curl --location --request GET 'http://localhost:8000/books'
```
4. Updating a Book by ID:
```azure
curl --location --request PUT 'http://localhost:8000/book/1/NewBookTitle/NewAuthor'
``` 
5. Deleting a Book by ID:
```azure
curl --location --request DELETE 'http://localhost:8000/book/1'
```
## To Run Using Postman
1. Add a Book (Create - POST):
```azure
Method: POST
URL: http://localhost:8000/book/BookTitle/AuthorName
```
2. Retrieve a Book (Read - GET) by Id:
```azure
Method: GET
URL: http://localhost:8000/book/1
```
3.  Retrieve All Books (Read - GET):
```azure
Method: GET
URL: http://localhost:8000/books
```
4. Update a Book (Update - PUT):
```azure
Method: PUT
URL: http://localhost:8000/book/1/NewTitle/NewAuthor
```
5. Delete a Book (Delete - DELETE):
```azure
Method: DELETE
URL: http://localhost:8000/book/1
```
- make sure to replace Title and Author name with the respective names as per actual books.
- By using Postman you can easily perform these CRUD Operations and observe the responses returned.

## Sequence Diagram

![Sequence Diagram.png](assests/Sequence%20Diagram.png)

## UML Diagram
![UML Class diagram.png](assests/UML%20Class%20diagram.png)