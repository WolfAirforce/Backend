for d in cmd/*; do
  echo "Building $d ..."
  go build -o ./bin airforce/$d
done