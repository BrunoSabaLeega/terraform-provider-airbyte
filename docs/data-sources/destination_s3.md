---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_s3 Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationS3 DataSource
---

# airbyte_destination_s3 (Data Source)

DestinationS3 DataSource

## Example Usage

```terraform
data "airbyte_destination_s3" "my_destination_s3" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `access_key_id` (String) The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.
- `destination_type` (String) must be one of ["s3"]
- `file_name_pattern` (String) The pattern allows you to set the file-name format for the S3 staging file(s)
- `format` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format))
- `s3_bucket_name` (String) The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
- `s3_bucket_path` (String) Directory under the S3 bucket where data will be written. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=to%20format%20the-,bucket%20path,-%3A">here</a>
- `s3_bucket_region` (String) must be one of ["", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]
The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
- `s3_endpoint` (String) Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>
- `s3_path_format` (String) Format string on how data will be organized inside the S3 bucket directory. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=The%20full%20path%20of%20the%20output%20data%20with%20the%20default%20S3%20path%20format">here</a>
- `secret_access_key` (String) The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>

<a id="nestedatt--configuration--format"></a>
### Nested Schema for `configuration.format`

Read-Only:

- `destination_s3_output_format_avro_apache_avro` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro))
- `destination_s3_output_format_csv_comma_separated_values` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values))
- `destination_s3_output_format_json_lines_newline_delimited_json` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json))
- `destination_s3_output_format_parquet_columnar_storage` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_parquet_columnar_storage))
- `destination_s3_update_output_format_avro_apache_avro` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro))
- `destination_s3_update_output_format_csv_comma_separated_values` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values))
- `destination_s3_update_output_format_json_lines_newline_delimited_json` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json))
- `destination_s3_update_output_format_parquet_columnar_storage` (Attributes) Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_parquet_columnar_storage))

<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro`

Read-Only:

- `compression_codec` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--compression_codec))
- `format_type` (String) must be one of ["Avro"]

<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--compression_codec"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type`

Read-Only:

- `destination_s3_output_format_avro_apache_avro_compression_codec_bzip2` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_bzip2))
- `destination_s3_output_format_avro_apache_avro_compression_codec_deflate` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_deflate))
- `destination_s3_output_format_avro_apache_avro_compression_codec_no_compression` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_no_compression))
- `destination_s3_output_format_avro_apache_avro_compression_codec_snappy` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_snappy))
- `destination_s3_output_format_avro_apache_avro_compression_codec_xz` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_xz))
- `destination_s3_output_format_avro_apache_avro_compression_codec_zstandard` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_zstandard))

<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_bzip2"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_bzip2`

Read-Only:

- `codec` (String) must be one of ["bzip2"]


<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_deflate"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_deflate`

Read-Only:

- `codec` (String) must be one of ["Deflate"]
- `compression_level` (Number) 0: no compression & fastest, 9: best compression & slowest.


<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_no_compression`

Read-Only:

- `codec` (String) must be one of ["no compression"]


<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_snappy"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_snappy`

Read-Only:

- `codec` (String) must be one of ["snappy"]


<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_xz"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_xz`

Read-Only:

- `codec` (String) must be one of ["xz"]
- `compression_level` (Number) See <a href="https://commons.apache.org/proper/commons-compress/apidocs/org/apache/commons/compress/compressors/xz/XZCompressorOutputStream.html#XZCompressorOutputStream-java.io.OutputStream-int-">here</a> for details.


<a id="nestedatt--configuration--format--destination_s3_output_format_avro_apache_avro--format_type--destination_s3_output_format_avro_apache_avro_compression_codec_zstandard"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_avro_apache_avro.format_type.destination_s3_output_format_avro_apache_avro_compression_codec_zstandard`

Read-Only:

- `codec` (String) must be one of ["zstandard"]
- `compression_level` (Number) Negative levels are 'fast' modes akin to lz4 or snappy, levels above 9 are generally for archival purposes, and levels above 18 use a lot of memory.
- `include_checksum` (Boolean) If true, include a checksum with each data block.




<a id="nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_csv_comma_separated_values`

Read-Only:

- `compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--compression))
- `flattening` (String) must be one of ["No flattening", "Root level flattening"]
Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
- `format_type` (String) must be one of ["CSV"]

<a id="nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--compression"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_csv_comma_separated_values.format_type`

Read-Only:

- `destination_s3_output_format_csv_comma_separated_values_compression_gzip` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--format_type--destination_s3_output_format_csv_comma_separated_values_compression_gzip))
- `destination_s3_output_format_csv_comma_separated_values_compression_no_compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--format_type--destination_s3_output_format_csv_comma_separated_values_compression_no_compression))

<a id="nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--format_type--destination_s3_output_format_csv_comma_separated_values_compression_gzip"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_csv_comma_separated_values.format_type.destination_s3_output_format_csv_comma_separated_values_compression_gzip`

Read-Only:

- `compression_type` (String) must be one of ["GZIP"]


<a id="nestedatt--configuration--format--destination_s3_output_format_csv_comma_separated_values--format_type--destination_s3_output_format_csv_comma_separated_values_compression_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_csv_comma_separated_values.format_type.destination_s3_output_format_csv_comma_separated_values_compression_no_compression`

Read-Only:

- `compression_type` (String) must be one of ["No Compression"]




<a id="nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_json_lines_newline_delimited_json`

Read-Only:

- `compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--compression))
- `flattening` (String) must be one of ["No flattening", "Root level flattening"]
Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
- `format_type` (String) must be one of ["JSONL"]

<a id="nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--compression"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_json_lines_newline_delimited_json.format_type`

Read-Only:

- `destination_s3_output_format_json_lines_newline_delimited_json_compression_gzip` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--format_type--destination_s3_output_format_json_lines_newline_delimited_json_compression_gzip))
- `destination_s3_output_format_json_lines_newline_delimited_json_compression_no_compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--format_type--destination_s3_output_format_json_lines_newline_delimited_json_compression_no_compression))

<a id="nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--format_type--destination_s3_output_format_json_lines_newline_delimited_json_compression_gzip"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_json_lines_newline_delimited_json.format_type.destination_s3_output_format_json_lines_newline_delimited_json_compression_gzip`

Read-Only:

- `compression_type` (String) must be one of ["GZIP"]


<a id="nestedatt--configuration--format--destination_s3_output_format_json_lines_newline_delimited_json--format_type--destination_s3_output_format_json_lines_newline_delimited_json_compression_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_json_lines_newline_delimited_json.format_type.destination_s3_output_format_json_lines_newline_delimited_json_compression_no_compression`

Read-Only:

- `compression_type` (String) must be one of ["No Compression"]




<a id="nestedatt--configuration--format--destination_s3_output_format_parquet_columnar_storage"></a>
### Nested Schema for `configuration.format.destination_s3_output_format_parquet_columnar_storage`

Read-Only:

- `block_size_mb` (Number) This is the size of a row group being buffered in memory. It limits the memory usage when writing. Larger values will improve the IO when reading, but consume more memory when writing. Default: 128 MB.
- `compression_codec` (String) must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "LZO", "BROTLI", "LZ4", "ZSTD"]
The compression algorithm used to compress data pages.
- `dictionary_encoding` (Boolean) Default: true.
- `dictionary_page_size_kb` (Number) There is one dictionary page per column per row group when dictionary encoding is used. The dictionary page size works like the page size but for dictionary. Default: 1024 KB.
- `format_type` (String) must be one of ["Parquet"]
- `max_padding_size_mb` (Number) Maximum size allowed as padding to align row groups. This is also the minimum size of a row group. Default: 8 MB.
- `page_size_kb` (Number) The page size is for compression. A block is composed of pages. A page is the smallest unit that must be read fully to access a single record. If this value is too small, the compression will deteriorate. Default: 1024 KB.


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro`

Read-Only:

- `compression_codec` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--compression_codec))
- `format_type` (String) must be one of ["Avro"]

<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--compression_codec"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type`

Read-Only:

- `destination_s3_update_output_format_avro_apache_avro_compression_codec_bzip2` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_bzip2))
- `destination_s3_update_output_format_avro_apache_avro_compression_codec_deflate` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_deflate))
- `destination_s3_update_output_format_avro_apache_avro_compression_codec_no_compression` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_no_compression))
- `destination_s3_update_output_format_avro_apache_avro_compression_codec_snappy` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_snappy))
- `destination_s3_update_output_format_avro_apache_avro_compression_codec_xz` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_xz))
- `destination_s3_update_output_format_avro_apache_avro_compression_codec_zstandard` (Attributes) The compression algorithm used to compress data. Default to no compression. (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_zstandard))

<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_bzip2"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_bzip2`

Read-Only:

- `codec` (String) must be one of ["bzip2"]


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_deflate"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_deflate`

Read-Only:

- `codec` (String) must be one of ["Deflate"]
- `compression_level` (Number) 0: no compression & fastest, 9: best compression & slowest.


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_no_compression`

Read-Only:

- `codec` (String) must be one of ["no compression"]


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_snappy"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_snappy`

Read-Only:

- `codec` (String) must be one of ["snappy"]


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_xz"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_xz`

Read-Only:

- `codec` (String) must be one of ["xz"]
- `compression_level` (Number) See <a href="https://commons.apache.org/proper/commons-compress/apidocs/org/apache/commons/compress/compressors/xz/XZCompressorOutputStream.html#XZCompressorOutputStream-java.io.OutputStream-int-">here</a> for details.


<a id="nestedatt--configuration--format--destination_s3_update_output_format_avro_apache_avro--format_type--destination_s3_update_output_format_avro_apache_avro_compression_codec_zstandard"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_avro_apache_avro.format_type.destination_s3_update_output_format_avro_apache_avro_compression_codec_zstandard`

Read-Only:

- `codec` (String) must be one of ["zstandard"]
- `compression_level` (Number) Negative levels are 'fast' modes akin to lz4 or snappy, levels above 9 are generally for archival purposes, and levels above 18 use a lot of memory.
- `include_checksum` (Boolean) If true, include a checksum with each data block.




<a id="nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_csv_comma_separated_values`

Read-Only:

- `compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--compression))
- `flattening` (String) must be one of ["No flattening", "Root level flattening"]
Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
- `format_type` (String) must be one of ["CSV"]

<a id="nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--compression"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_csv_comma_separated_values.format_type`

Read-Only:

- `destination_s3_update_output_format_csv_comma_separated_values_compression_gzip` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--format_type--destination_s3_update_output_format_csv_comma_separated_values_compression_gzip))
- `destination_s3_update_output_format_csv_comma_separated_values_compression_no_compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--format_type--destination_s3_update_output_format_csv_comma_separated_values_compression_no_compression))

<a id="nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--format_type--destination_s3_update_output_format_csv_comma_separated_values_compression_gzip"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_csv_comma_separated_values.format_type.destination_s3_update_output_format_csv_comma_separated_values_compression_gzip`

Read-Only:

- `compression_type` (String) must be one of ["GZIP"]


<a id="nestedatt--configuration--format--destination_s3_update_output_format_csv_comma_separated_values--format_type--destination_s3_update_output_format_csv_comma_separated_values_compression_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_csv_comma_separated_values.format_type.destination_s3_update_output_format_csv_comma_separated_values_compression_no_compression`

Read-Only:

- `compression_type` (String) must be one of ["No Compression"]




<a id="nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_json_lines_newline_delimited_json`

Read-Only:

- `compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--compression))
- `flattening` (String) must be one of ["No flattening", "Root level flattening"]
Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
- `format_type` (String) must be one of ["JSONL"]

<a id="nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--compression"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_json_lines_newline_delimited_json.format_type`

Read-Only:

- `destination_s3_update_output_format_json_lines_newline_delimited_json_compression_gzip` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--format_type--destination_s3_update_output_format_json_lines_newline_delimited_json_compression_gzip))
- `destination_s3_update_output_format_json_lines_newline_delimited_json_compression_no_compression` (Attributes) Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz"). (see [below for nested schema](#nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--format_type--destination_s3_update_output_format_json_lines_newline_delimited_json_compression_no_compression))

<a id="nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--format_type--destination_s3_update_output_format_json_lines_newline_delimited_json_compression_gzip"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_json_lines_newline_delimited_json.format_type.destination_s3_update_output_format_json_lines_newline_delimited_json_compression_gzip`

Read-Only:

- `compression_type` (String) must be one of ["GZIP"]


<a id="nestedatt--configuration--format--destination_s3_update_output_format_json_lines_newline_delimited_json--format_type--destination_s3_update_output_format_json_lines_newline_delimited_json_compression_no_compression"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_json_lines_newline_delimited_json.format_type.destination_s3_update_output_format_json_lines_newline_delimited_json_compression_no_compression`

Read-Only:

- `compression_type` (String) must be one of ["No Compression"]




<a id="nestedatt--configuration--format--destination_s3_update_output_format_parquet_columnar_storage"></a>
### Nested Schema for `configuration.format.destination_s3_update_output_format_parquet_columnar_storage`

Read-Only:

- `block_size_mb` (Number) This is the size of a row group being buffered in memory. It limits the memory usage when writing. Larger values will improve the IO when reading, but consume more memory when writing. Default: 128 MB.
- `compression_codec` (String) must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "LZO", "BROTLI", "LZ4", "ZSTD"]
The compression algorithm used to compress data pages.
- `dictionary_encoding` (Boolean) Default: true.
- `dictionary_page_size_kb` (Number) There is one dictionary page per column per row group when dictionary encoding is used. The dictionary page size works like the page size but for dictionary. Default: 1024 KB.
- `format_type` (String) must be one of ["Parquet"]
- `max_padding_size_mb` (Number) Maximum size allowed as padding to align row groups. This is also the minimum size of a row group. Default: 8 MB.
- `page_size_kb` (Number) The page size is for compression. A block is composed of pages. A page is the smallest unit that must be read fully to access a single record. If this value is too small, the compression will deteriorate. Default: 1024 KB.

