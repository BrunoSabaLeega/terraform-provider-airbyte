---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_s3 Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceS3 DataSource
---

# airbyte_source_s3 (Data Source)

SourceS3 DataSource

## Example Usage

```terraform
data "airbyte_source_s3" "my_source_s3" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) NOTE: When this Spec is changed, legacy_config_transformer.py must also be modified to uptake the changes
because it is responsible for converting legacy S3 v3 configs into v4 configs using the File-Based CDK. (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `aws_access_key_id` (String) In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `aws_secret_access_key` (String) In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `bucket` (String) Name of the S3 bucket where the file(s) exist.
- `dataset` (String) Deprecated and will be removed soon. Please do not use this field anymore and use streams.name instead. The name of the stream you would like this source to output. Can contain letters, numbers, or underscores.
- `endpoint` (String) Endpoint to an S3 compatible service. Leave empty to use AWS.
- `format` (Attributes) Deprecated and will be removed soon. Please do not use this field anymore and use streams.format instead. The format of the files you'd like to replicate (see [below for nested schema](#nestedatt--configuration--format))
- `path_pattern` (String) Deprecated and will be removed soon. Please do not use this field anymore and use streams.globs instead. A regular expression which tells the connector which files to replicate. All files which match this pattern will be replicated. Use | to separate multiple patterns. See <a href="https://facelessuser.github.io/wcmatch/glob/" target="_blank">this page</a> to understand pattern syntax (GLOBSTAR and SPLIT flags are enabled). Use pattern <strong>**</strong> to pick up all files.
- `provider` (Attributes) Deprecated and will be removed soon. Please do not use this field anymore and use bucket, aws_access_key_id, aws_secret_access_key and endpoint instead. Use this to load files from S3 or S3-compatible services (see [below for nested schema](#nestedatt--configuration--provider))
- `schema` (String) Deprecated and will be removed soon. Please do not use this field anymore and use streams.input_schema instead. Optionally provide a schema to enforce, as a valid JSON string. Ensure this is a mapping of <strong>{ "column" : "type" }</strong>, where types are valid <a href="https://json-schema.org/understanding-json-schema/reference/type.html" target="_blank">JSON Schema datatypes</a>. Leave as {} to auto-infer the schema.
- `source_type` (String) must be one of ["s3"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.
- `streams` (Attributes List) Each instance of this configuration defines a <a href="https://docs.airbyte.com/cloud/core-concepts#stream">stream</a>. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table. (see [below for nested schema](#nestedatt--configuration--streams))

<a id="nestedatt--configuration--format"></a>
### Nested Schema for `configuration.format`

Read-Only:

- `source_s3_file_format_avro` (Attributes) This connector utilises <a href="https://fastavro.readthedocs.io/en/latest/" target="_blank">fastavro</a> for Avro parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_file_format_avro))
- `source_s3_file_format_csv` (Attributes) This connector utilises <a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank">PyArrow (Apache Arrow)</a> for CSV parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_file_format_csv))
- `source_s3_file_format_jsonl` (Attributes) This connector uses <a href="https://arrow.apache.org/docs/python/json.html" target="_blank">PyArrow</a> for JSON Lines (jsonl) file parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_file_format_jsonl))
- `source_s3_file_format_parquet` (Attributes) This connector utilises <a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank">PyArrow (Apache Arrow)</a> for Parquet parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_file_format_parquet))
- `source_s3_update_file_format_avro` (Attributes) This connector utilises <a href="https://fastavro.readthedocs.io/en/latest/" target="_blank">fastavro</a> for Avro parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_update_file_format_avro))
- `source_s3_update_file_format_csv` (Attributes) This connector utilises <a href="https: // arrow.apache.org/docs/python/generated/pyarrow.csv.open_csv.html" target="_blank">PyArrow (Apache Arrow)</a> for CSV parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_update_file_format_csv))
- `source_s3_update_file_format_jsonl` (Attributes) This connector uses <a href="https://arrow.apache.org/docs/python/json.html" target="_blank">PyArrow</a> for JSON Lines (jsonl) file parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_update_file_format_jsonl))
- `source_s3_update_file_format_parquet` (Attributes) This connector utilises <a href="https://arrow.apache.org/docs/python/generated/pyarrow.parquet.ParquetFile.html" target="_blank">PyArrow (Apache Arrow)</a> for Parquet parsing. (see [below for nested schema](#nestedatt--configuration--format--source_s3_update_file_format_parquet))

<a id="nestedatt--configuration--format--source_s3_file_format_avro"></a>
### Nested Schema for `configuration.format.source_s3_file_format_avro`

Read-Only:

- `filetype` (String) must be one of ["avro"]


<a id="nestedatt--configuration--format--source_s3_file_format_csv"></a>
### Nested Schema for `configuration.format.source_s3_file_format_csv`

Read-Only:

- `additional_reader_options` (String) Optionally add a valid JSON string here to provide additional options to the csv reader. Mappings must correspond to options <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ConvertOptions.html#pyarrow.csv.ConvertOptions" target="_blank">detailed here</a>. 'column_types' is used internally to handle schema so overriding that would likely cause problems.
- `advanced_options` (String) Optionally add a valid JSON string here to provide additional <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ReadOptions.html#pyarrow.csv.ReadOptions" target="_blank">Pyarrow ReadOptions</a>. Specify 'column_names' here if your CSV doesn't have header, or if you want to use custom column names. 'block_size' and 'encoding' are already used above, specify them again here will override the values above.
- `block_size` (Number) The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.
- `delimiter` (String) The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
- `double_quote` (Boolean) Whether two quotes in a quoted CSV value denote a single quote in the data.
- `encoding` (String) The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
- `escape_char` (String) The character used for escaping special characters. To disallow escaping, leave this field blank.
- `filetype` (String) must be one of ["csv"]
- `infer_datatypes` (Boolean) Configures whether a schema for the source should be inferred from the current data or not. If set to false and a custom schema is set, then the manually enforced schema is used. If a schema is not manually set, and this is set to false, then all fields will be read as strings
- `newlines_in_values` (Boolean) Whether newline characters are allowed in CSV values. Turning this on may affect performance. Leave blank to default to False.
- `quote_char` (String) The character used for quoting CSV values. To disallow quoting, make this field blank.


<a id="nestedatt--configuration--format--source_s3_file_format_jsonl"></a>
### Nested Schema for `configuration.format.source_s3_file_format_jsonl`

Read-Only:

- `block_size` (Number) The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.
- `filetype` (String) must be one of ["jsonl"]
- `newlines_in_values` (Boolean) Whether newline characters are allowed in JSON values. Turning this on may affect performance. Leave blank to default to False.
- `unexpected_field_behavior` (String) must be one of ["ignore", "infer", "error"]
How JSON fields outside of explicit_schema (if given) are treated. Check <a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank">PyArrow documentation</a> for details


<a id="nestedatt--configuration--format--source_s3_file_format_parquet"></a>
### Nested Schema for `configuration.format.source_s3_file_format_parquet`

Read-Only:

- `batch_size` (Number) Maximum number of records per batch read from the input files. Batches may be smaller if there aren’t enough rows in the file. This option can help avoid out-of-memory errors if your data is particularly wide.
- `buffer_size` (Number) Perform read buffering when deserializing individual column chunks. By default every group column will be loaded fully to memory. This option can help avoid out-of-memory errors if your data is particularly wide.
- `columns` (List of String) If you only want to sync a subset of the columns from the file(s), add the columns you want here as a comma-delimited list. Leave it empty to sync all columns.
- `filetype` (String) must be one of ["parquet"]


<a id="nestedatt--configuration--format--source_s3_update_file_format_avro"></a>
### Nested Schema for `configuration.format.source_s3_update_file_format_avro`

Read-Only:

- `filetype` (String) must be one of ["avro"]


<a id="nestedatt--configuration--format--source_s3_update_file_format_csv"></a>
### Nested Schema for `configuration.format.source_s3_update_file_format_csv`

Read-Only:

- `additional_reader_options` (String) Optionally add a valid JSON string here to provide additional options to the csv reader. Mappings must correspond to options <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ConvertOptions.html#pyarrow.csv.ConvertOptions" target="_blank">detailed here</a>. 'column_types' is used internally to handle schema so overriding that would likely cause problems.
- `advanced_options` (String) Optionally add a valid JSON string here to provide additional <a href="https://arrow.apache.org/docs/python/generated/pyarrow.csv.ReadOptions.html#pyarrow.csv.ReadOptions" target="_blank">Pyarrow ReadOptions</a>. Specify 'column_names' here if your CSV doesn't have header, or if you want to use custom column names. 'block_size' and 'encoding' are already used above, specify them again here will override the values above.
- `block_size` (Number) The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.
- `delimiter` (String) The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
- `double_quote` (Boolean) Whether two quotes in a quoted CSV value denote a single quote in the data.
- `encoding` (String) The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
- `escape_char` (String) The character used for escaping special characters. To disallow escaping, leave this field blank.
- `filetype` (String) must be one of ["csv"]
- `infer_datatypes` (Boolean) Configures whether a schema for the source should be inferred from the current data or not. If set to false and a custom schema is set, then the manually enforced schema is used. If a schema is not manually set, and this is set to false, then all fields will be read as strings
- `newlines_in_values` (Boolean) Whether newline characters are allowed in CSV values. Turning this on may affect performance. Leave blank to default to False.
- `quote_char` (String) The character used for quoting CSV values. To disallow quoting, make this field blank.


<a id="nestedatt--configuration--format--source_s3_update_file_format_jsonl"></a>
### Nested Schema for `configuration.format.source_s3_update_file_format_jsonl`

Read-Only:

- `block_size` (Number) The chunk size in bytes to process at a time in memory from each file. If your data is particularly wide and failing during schema detection, increasing this should solve it. Beware of raising this too high as you could hit OOM errors.
- `filetype` (String) must be one of ["jsonl"]
- `newlines_in_values` (Boolean) Whether newline characters are allowed in JSON values. Turning this on may affect performance. Leave blank to default to False.
- `unexpected_field_behavior` (String) must be one of ["ignore", "infer", "error"]
How JSON fields outside of explicit_schema (if given) are treated. Check <a href="https://arrow.apache.org/docs/python/generated/pyarrow.json.ParseOptions.html" target="_blank">PyArrow documentation</a> for details


<a id="nestedatt--configuration--format--source_s3_update_file_format_parquet"></a>
### Nested Schema for `configuration.format.source_s3_update_file_format_parquet`

Read-Only:

- `batch_size` (Number) Maximum number of records per batch read from the input files. Batches may be smaller if there aren’t enough rows in the file. This option can help avoid out-of-memory errors if your data is particularly wide.
- `buffer_size` (Number) Perform read buffering when deserializing individual column chunks. By default every group column will be loaded fully to memory. This option can help avoid out-of-memory errors if your data is particularly wide.
- `columns` (List of String) If you only want to sync a subset of the columns from the file(s), add the columns you want here as a comma-delimited list. Leave it empty to sync all columns.
- `filetype` (String) must be one of ["parquet"]



<a id="nestedatt--configuration--provider"></a>
### Nested Schema for `configuration.provider`

Read-Only:

- `aws_access_key_id` (String) In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `aws_secret_access_key` (String) In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
- `bucket` (String) Name of the S3 bucket where the file(s) exist.
- `endpoint` (String) Endpoint to an S3 compatible service. Leave empty to use AWS.
- `path_prefix` (String) By providing a path-like prefix (e.g. myFolder/thisTable/) under which all the relevant files sit, we can optimize finding these in S3. This is optional but recommended if your bucket contains many folders/files which you don't need to replicate.
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any file modified before this date will not be replicated.


<a id="nestedatt--configuration--streams"></a>
### Nested Schema for `configuration.streams`

Read-Only:

- `days_to_sync_if_history_is_full` (Number) When the state history of the file store is full, syncs will only read files that were last modified in the provided day range.
- `file_type` (String) The data file type that is being extracted for a stream.
- `format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format))
- `globs` (List of String) The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look <a href="https://en.wikipedia.org/wiki/Glob_(programming)">here</a>.
- `input_schema` (String) The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.
- `legacy_prefix` (String) The path prefix configured in v3 versions of the S3 connector. This option is deprecated in favor of a single glob.
- `name` (String) The name of the stream.
- `primary_key` (String) The column or columns (for a composite key) that serves as the unique identifier of a record.
- `schemaless` (Boolean) When enabled, syncs will not validate or structure records against the stream's schema.
- `validation_policy` (String) must be one of ["Emit Record", "Skip Record", "Wait for Discover"]
The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.

<a id="nestedatt--configuration--streams--format"></a>
### Nested Schema for `configuration.streams.format`

Read-Only:

- `source_s3_file_based_stream_config_format_avro_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_avro_format))
- `source_s3_file_based_stream_config_format_csv_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_csv_format))
- `source_s3_file_based_stream_config_format_jsonl_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_jsonl_format))
- `source_s3_file_based_stream_config_format_parquet_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_parquet_format))
- `source_s3_update_file_based_stream_config_format_avro_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_avro_format))
- `source_s3_update_file_based_stream_config_format_csv_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_csv_format))
- `source_s3_update_file_based_stream_config_format_jsonl_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_jsonl_format))
- `source_s3_update_file_based_stream_config_format_parquet_format` (Attributes) The configuration options that are used to alter how to read incoming files that deviate from the standard formatting. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format))

<a id="nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_avro_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `double_as_string` (Boolean) Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers.
- `filetype` (String) must be one of ["avro"]


<a id="nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_csv_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `delimiter` (String) The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
- `double_quote` (Boolean) Whether two quotes in a quoted CSV value denote a single quote in the data.
- `encoding` (String) The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
- `escape_char` (String) The character used for escaping special characters. To disallow escaping, leave this field blank.
- `false_values` (List of String) A set of case-sensitive strings that should be interpreted as false values.
- `filetype` (String) must be one of ["csv"]
- `header_definition` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition))
- `inference_type` (String) must be one of ["None", "Primitive Types Only"]
How to infer the types of the columns. If none, inference default to strings.
- `null_values` (List of String) A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.
- `quote_char` (String) The character used for quoting CSV values. To disallow quoting, make this field blank.
- `skip_rows_after_header` (Number) The number of rows to skip after the header row.
- `skip_rows_before_header` (Number) The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field.
- `strings_can_be_null` (Boolean) Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself.
- `true_values` (List of String) A set of case-sensitive strings that should be interpreted as true values.

<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition`

Read-Only:

- `source_s3_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated))
- `source_s3_file_based_stream_config_format_csv_format_csv_header_definition_from_csv` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_from_csv))
- `source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided))

<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `header_definition_type` (String) must be one of ["Autogenerated"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_from_csv"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `header_definition_type` (String) must be one of ["From CSV"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `column_names` (List of String) The column names that will be used while emitting the CSV records
- `header_definition_type` (String) must be one of ["User Provided"]




<a id="nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_jsonl_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `filetype` (String) must be one of ["jsonl"]


<a id="nestedatt--configuration--streams--format--source_s3_file_based_stream_config_format_parquet_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `decimal_as_float` (Boolean) Whether to convert decimal fields to floats. There is a loss of precision when converting decimals to floats, so this is not recommended.
- `filetype` (String) must be one of ["parquet"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_avro_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `double_as_string` (Boolean) Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers.
- `filetype` (String) must be one of ["avro"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_csv_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `delimiter` (String) The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
- `double_quote` (Boolean) Whether two quotes in a quoted CSV value denote a single quote in the data.
- `encoding` (String) The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
- `escape_char` (String) The character used for escaping special characters. To disallow escaping, leave this field blank.
- `false_values` (List of String) A set of case-sensitive strings that should be interpreted as false values.
- `filetype` (String) must be one of ["csv"]
- `header_definition` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition))
- `inference_type` (String) must be one of ["None", "Primitive Types Only"]
How to infer the types of the columns. If none, inference default to strings.
- `null_values` (List of String) A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.
- `quote_char` (String) The character used for quoting CSV values. To disallow quoting, make this field blank.
- `skip_rows_after_header` (Number) The number of rows to skip after the header row.
- `skip_rows_before_header` (Number) The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field.
- `strings_can_be_null` (Boolean) Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself.
- `true_values` (List of String) A set of case-sensitive strings that should be interpreted as true values.

<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition`

Read-Only:

- `source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated))
- `source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_from_csv` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_from_csv))
- `source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided` (Attributes) How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows. (see [below for nested schema](#nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided))

<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_autogenerated"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `header_definition_type` (String) must be one of ["Autogenerated"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_from_csv"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `header_definition_type` (String) must be one of ["From CSV"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format--header_definition--source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format.header_definition.source_s3_update_file_based_stream_config_format_csv_format_csv_header_definition_user_provided`

Read-Only:

- `column_names` (List of String) The column names that will be used while emitting the CSV records
- `header_definition_type` (String) must be one of ["User Provided"]




<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_jsonl_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `filetype` (String) must be one of ["jsonl"]


<a id="nestedatt--configuration--streams--format--source_s3_update_file_based_stream_config_format_parquet_format"></a>
### Nested Schema for `configuration.streams.format.source_s3_update_file_based_stream_config_format_parquet_format`

Read-Only:

- `decimal_as_float` (Boolean) Whether to convert decimal fields to floats. There is a loss of precision when converting decimals to floats, so this is not recommended.
- `filetype` (String) must be one of ["parquet"]


