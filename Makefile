test:
	go test -v ./testing

testacc:
	JERAKIA_ACC=1 go test -v ./acceptance -run="$(TEST)"
