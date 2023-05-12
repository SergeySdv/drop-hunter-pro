#### How to deploy migrations:

1) Raise up postgres database
2) install migration tool [goose](https://github.com/pressly/goose)
3) Check status 
>goose postgres "user=postgres password=postgres dbname=bback sslmode=disable port=6667" status
>
4) Deploy

 >goose postgres "user=postgres password=postgres dbname=bback sslmode=disable port=6666" up