@run_us:
  cd ./services/user_service/ && GO_ENV=dev go run ./...

@run_us_test:
  cd ./services/user_service/ && go test -v ./...

@wire_us:
  cd ./services/user_service/cmd/di/ && wire

@apply:
  cd ./scripts/k8s && bash apply-all.sh
