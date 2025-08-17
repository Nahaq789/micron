@run_us:
  cd ./services/user_service/ && GO_ENV=dev go run ./...

@wire_us:
  cd ./services/user_service/cmd/di/ && wire

@apply:
  ls
  cd ./scripts/k8s && bash apply-all.sh
