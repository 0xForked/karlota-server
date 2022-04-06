FILE=.env
if test -f "$FILE"; then
  echo "Everything is OK"
  echo "Validate dependencies . . ."
  go mod tidy -compat=1.17
  echo "Re-generate Swagger File (api-spec docs) . . ."
  swag init --parseDependency --parseInternal --parseDepth 1
  echo "Trying to run the tests . . ."
  go test ./... -v
  echo "Trying to run the server . . ."
  go run main.go
else
  echo "==========================================================="
  echo "|  $FILE (environment) file does not exist.                |"
  echo "|  Please Crete new .env file from .env.example.          |"
  echo "|  by running this script: //:~$ cp .env.example .env     |"
  echo "==========================================================="
  exit 0
fi

