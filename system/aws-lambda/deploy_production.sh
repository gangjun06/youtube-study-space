# TODO: ################## これは本番環境用です!!!!!!!!!!! ###################


# set_desired_max_seats, rooms_state, youtube_organize_database, daily_organize_database, check_live_stream_status
# lambda_sandbox, transfer_collection_history_bigquery, process_user_rp_parallel


# Windows (PowerShell)
cd system; cd aws-lambda;  # ディレクトリを移動
$env:CGO_ENABLED = "0"; $env:GOOS = "linux"; $env:GOARCH = "amd64"; aws configure set region ap-northeast-1
go build -o main     reset_daily_total_study_time.go
C:\Users\momom\go\bin\build-lambda-zip.exe -output main.zip main
aws lambda create-function --function-name     rooms_state     --runtime go1.x --zip-file fileb://main.zip --handler
main --role arn:aws:iam::652333062396:role/service-role/my-first-golang-lambda-function-role-cb8uw4th --timeout 120 --profile soraride
aws lambda update-function-code --function-name     reset_daily_total_study_time     --zip-file fileb://main.zip --profile soraride


# TODO: ################## これは本番環境用です!!!!!!!!!!! ###################

# Mac OS
cd system; cd aws-lambda;  # ディレクトリを移動
GOARCH=amd64 GOOS=linux && aws configure set region ap-northeast-1
go build -o main    youtube_organize_database.go
zip main.zip main

aws lambda create-function --function-name change_user_info --runtime go1.x --zip-file fileb://main.zip --handler
main --role arn:aws:iam::652333062396:role/service-role/my-first-golang-lambda-function-role-cb8uw4th --timeout 120 --profile soraride

aws lambda update-function-code --function-name   youtube_organize_database   --zip-file fileb://main.zip --profile soraride
