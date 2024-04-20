# I'm to lazy to  make my own structure type project
Btw im this person^


 This is a general guideline, which you can adjust to match the specific needs of your project. 
 Every file structure is kept as it's separate file which makes it easier to modify

 ## USAGE ##

first install this thing (https://go.dev/doc/tutorial/compile-install) < great thing
0. BE INSIDE THIS DIRECTORY
1. go build
2. go list -f '{{.Target}}'
3. copy output from command above
4. $ export PATH=$PATH:/path/to/directory/you/copied
5. go install
6. then type headstartgo projectname

thing that you type after headstartgo will be the name for root of your template
NO SPACES ALLOWED IN PROJECT NAME

then pick your template

have fun


## Folder Structures built-in

1. **Flat Structure**:
  Feels very nice in small projects, harder to maintaint in more complex projects.

   ```go
   project-root/
       ├── main.go
       ├── handler.go
       ├── config.go
       ├── database.go
       ├── ...
       ├── static/
       ├── templates/
       ├── scripts/
       ├── configs/
       ├── tests/
       └── docs/
   ```

2. **Layered Structure**:
   Main benefit is that every part of project is kept in it's own place.

   ```go
   project-root/
       ├── main.go
       ├── web/
       │   ├── handler.go
       │   ├── static/
       │   ├── templates/
       ├── api/
       │   ├── routes.go
       │   ├── middleware/
       ├── data/
       │   ├── database.go
       │   ├── repository.go
       ├── configs/
       ├── tests/
       ├── docs/
   ```

3. **Domain-Driven Design (DDD)**:
  Each domain has its own directory. It can be nice in large projects

   ```go
   project-root/
       ├── cmd/
       │   ├── app1/
       │   ├── app2/
       ├── internal/
       │   ├── auth/
       │   │   ├── handler.go
       │   │   ├── service.go
       │   ├── orders/
       │   │   ├── handler.go
       │   │   ├── service.go
       │   ├── ...
       ├── pkg/
       │   ├── utility/
       │   │   ├── ...
       │   ├── ...
       ├── api/
       │   ├── app1/
       │   │   ├── ...
       │   ├── app2/
       │   │   ├── ...
       ├── web/
       │   ├── app1/
       │   │   ├── ...
       │   ├── app2/
       │   │   ├── ...
       ├── scripts/
       ├── configs/
       ├── tests/
       └── docs/
   ```

4. **Clean Architecture**:
   It emphasizes a separation of concerns between different layers in your application.

   ```go
   project-root/
       ├── cmd/
       │   ├── your-app/
       │   │   ├── main.go
       ├── internal/
       │   ├── app/
       │   │   ├── handler.go
       │   │   ├── service.go
       │   ├── domain/
       │   │   ├── model.go
       │   │   ├── repository.go
       ├── pkg/
       │   ├── utility/
       │   │   ├── ...
       ├── api/
       │   ├── ...
       ├── web/
       │   ├── ...
       ├── scripts/
       ├── configs/
       ├── tests/
       └── docs/
   ```

5. **Modular Structure**:
   This approach of keeping every part of the program in individual module can be useful when developing multiple independent components within single project.

   ```go
   project-root/
       ├── module1/
       │   ├── cmd/
       │   ├── internal/
       │   ├── pkg/
       │   ├── api/
       │   ├── web/
       │   ├── scripts/
       │   ├── configs/
       │   ├── tests/
       │   └── docs/
       ├── module2/
       │   ├── ...
   ```
