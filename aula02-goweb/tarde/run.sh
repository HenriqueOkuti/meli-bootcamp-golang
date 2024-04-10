# Verify if nodemon package is installed:

# If Nodemon is not installed, exit with error message

if ! [ -x "$(command -v nodemon)" ]; then
  echo 'Error: nodemon is not installed.' >&2
  exit 1
fi

# If Nodemon is installed, run the application

nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run ./cmd/server/main.go