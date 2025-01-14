resource "airbyte_destination_firebolt" "my_destination_firebolt" {
  configuration = {
    account          = "...my_account..."
    database         = "...my_database..."
    destination_type = "firebolt"
    engine           = "...my_engine..."
    host             = "api.app.firebolt.io"
    loading_method = {
      destination_firebolt_loading_method_external_table_via_s3 = {
        aws_key_id     = "...my_aws_key_id..."
        aws_key_secret = "...my_aws_key_secret..."
        method         = "S3"
        s3_bucket      = "...my_s3_bucket..."
        s3_region      = "us-east-1"
      }
    }
    password = "...my_password..."
    username = "username@email.com"
  }
  name         = "Roman Kulas"
  workspace_id = "c7abd74d-d39c-40f5-92cf-f7c70a45626d"
}