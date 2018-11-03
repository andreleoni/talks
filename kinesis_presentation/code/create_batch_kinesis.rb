require 'aws-sdk'
require 'securerandom'

items = [
  {
    user_id: 20,
    created_at: "#{Time.now.to_i}-#{SecureRandom.hex}",
    text: 'Em call First item'
  }, {
    user_id: 20,
    text: 'Em call Second item'
  }, {
    user_id: 20,
    created_at: "#{Time.now.to_i}-#{SecureRandom.hex}",
    text: 'Em call Third item'
  }
]

kinesis_client = Aws::Kinesis::Client.new(
  region: 'us-east-1',
  credentials: Aws::Credentials.new(
    'AKIAJNVFGVGRPGAZMWQQ',
    'pSV0Db3j0ytBU8C5IcItC5vc1+WbhMQybjuwmiLl'
  )
)

items.each do |item|
  kinesis_client.put_record(
    stream_name: 'test_enjoei_kinesis',
    data: item.to_json,
    partition_key: SecureRandom.hex(32)
  )
end
