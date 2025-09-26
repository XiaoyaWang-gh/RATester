if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <model> <project>"
    exit 1
fi

model=$1
project=$2

echo "Processing model: $model"

python run_test.py --model IntelliTester_$model --project $project

cd projects/$project || { echo "Project directory not found"; exit 1; }

go test ./... -timeout 180s -coverpkg=./... -coverprofile=IntelliTester_$model\_coverage.out -v -gcflags=all=-l > IntelliTester_$model\_test.out

gocov convert IntelliTester_$model\_coverage.out | gocov report > IntelliTester_$model\_coverage.txt

echo "All models processed successfully."