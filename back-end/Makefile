gen-service:
	goctl api go -api ./account.api -dir ./api
gen-model:
	goctl model mysql ddl -src="./account.sql" -dir ./api/model
	
runs:
	go run api/account.go
