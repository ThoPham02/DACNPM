gen-service:
	goctl api go -api ./account.api -dir ./api
gen-model:
	goctl model pg datasource -url="postgres://thopham:hdp0pylJq8EoLvAy8gI15XWU4vytfabF@dpg-ck0hgej6fquc73e0nkl0-a.singapore-postgres.render.com/dscnpm_db" -table="user_tbl" -dir="./api/model"
	
runs:
	go run api/account.go
