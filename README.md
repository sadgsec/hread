This is a playabout repo, don't use this code

Probably trying to make an image board, or a c2 server if i get bored

lets seee

# GETTING STARTED
1. create postgres database - 
    `docker run --name dev-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres` #TODO put the actual command here for dev env
2. Apply db schema
3. export DBURL="postgresql://postgres:mysecretpassword@localhost/hread"
4. go run .

