require 'aws-sdk'
require 'securerandom'

dynamodb = Aws::DynamoDB::Client.new(
  region: 'us-east-1',
  credentials: Aws::Credentials.new(
    'AKIAJNVFGVGRPGAZMWQQ',
    'pSV0Db3j0ytBU8C5IcItC5vc1+WbhMQybjuwmiLl'
  ),
  http_read_timeout: 3
)

request_params = {
  table_name: 'test_enjoei_dynamo',
  select: 'ALL_ATTRIBUTES'
}

puts dynamodb.scan(request_params)
