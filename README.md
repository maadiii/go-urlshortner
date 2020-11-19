# Url shotner

## create urlshortner user in postgres
`create user urlshortner with password 'urlshortner';`

## create urlshortner database
`create database urlshortner;`

## grant database to user
`grant all privileges on database urlshortner to urlshortner ;`

## create table urls
login with urlshortner user and create table urls
```create table url(
shorted varchar(10) primary key,
original varchar(500) not null,
created_at timestamp not null,
expired_at timestamp not null);
```

## run app
`go run main.go serve`
