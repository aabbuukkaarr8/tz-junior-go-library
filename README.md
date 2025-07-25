# tz-junior-go-library
1. cd tz-junior-go-library
2. cd docker
3. in the bash docker-compose up --build
4. run docker
5. go build main
6. POST /books/create for create book 
7. body for create
8. Title  string `json:"title"`
   Author string `json:"author"`
   Pages  int    `json:"pages"`
9. GET /books Get All your books list
10. GET /books/:id get book by id
11. PUT /books/:id update your book by id
12. DELETE /books/:id delete book by id
