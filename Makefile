test:
	@docker run -d --name "kubikvest_db" -p 3306:3306 imega/mysql
	@docker run --rm \
		-v $(CURDIR)/sql:/sql \
		--link kubikvest_db:kubikvest_db \
		imega/mysql-client \
		mysql --host=kubikvest_db -e "source /sql/kubikvest.sql"
