@run_us:
  cd ./services/user_service/ && GO_ENV=dev go run ./...

@apply:
  ls
  cd ./scripts/k8s && bash apply-all.sh
