---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_langchain Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationLangchain DataSource
---

# airbyte_destination_langchain (Data Source)

DestinationLangchain DataSource

## Example Usage

```terraform
data "airbyte_destination_langchain" "my_destination_langchain" {
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

- `destination_type` (String) must be one of ["langchain"]
- `embedding` (Attributes) Embedding configuration (see [below for nested schema](#nestedatt--configuration--embedding))
- `indexing` (Attributes) Indexing configuration (see [below for nested schema](#nestedatt--configuration--indexing))
- `processing` (Attributes) (see [below for nested schema](#nestedatt--configuration--processing))

<a id="nestedatt--configuration--embedding"></a>
### Nested Schema for `configuration.embedding`

Read-Only:

- `destination_langchain_embedding_fake` (Attributes) Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs. (see [below for nested schema](#nestedatt--configuration--embedding--destination_langchain_embedding_fake))
- `destination_langchain_embedding_open_ai` (Attributes) Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions. (see [below for nested schema](#nestedatt--configuration--embedding--destination_langchain_embedding_open_ai))
- `destination_langchain_update_embedding_fake` (Attributes) Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs. (see [below for nested schema](#nestedatt--configuration--embedding--destination_langchain_update_embedding_fake))
- `destination_langchain_update_embedding_open_ai` (Attributes) Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions. (see [below for nested schema](#nestedatt--configuration--embedding--destination_langchain_update_embedding_open_ai))

<a id="nestedatt--configuration--embedding--destination_langchain_embedding_fake"></a>
### Nested Schema for `configuration.embedding.destination_langchain_embedding_fake`

Read-Only:

- `mode` (String) must be one of ["fake"]


<a id="nestedatt--configuration--embedding--destination_langchain_embedding_open_ai"></a>
### Nested Schema for `configuration.embedding.destination_langchain_embedding_open_ai`

Read-Only:

- `mode` (String) must be one of ["openai"]
- `openai_key` (String)


<a id="nestedatt--configuration--embedding--destination_langchain_update_embedding_fake"></a>
### Nested Schema for `configuration.embedding.destination_langchain_update_embedding_fake`

Read-Only:

- `mode` (String) must be one of ["fake"]


<a id="nestedatt--configuration--embedding--destination_langchain_update_embedding_open_ai"></a>
### Nested Schema for `configuration.embedding.destination_langchain_update_embedding_open_ai`

Read-Only:

- `mode` (String) must be one of ["openai"]
- `openai_key` (String)



<a id="nestedatt--configuration--indexing"></a>
### Nested Schema for `configuration.indexing`

Read-Only:

- `destination_langchain_indexing_chroma_local_persistance` (Attributes) Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_indexing_chroma_local_persistance))
- `destination_langchain_indexing_doc_array_hnsw_search` (Attributes) DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_indexing_doc_array_hnsw_search))
- `destination_langchain_indexing_pinecone` (Attributes) Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_indexing_pinecone))
- `destination_langchain_update_indexing_chroma_local_persistance` (Attributes) Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_update_indexing_chroma_local_persistance))
- `destination_langchain_update_indexing_doc_array_hnsw_search` (Attributes) DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_update_indexing_doc_array_hnsw_search))
- `destination_langchain_update_indexing_pinecone` (Attributes) Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain. (see [below for nested schema](#nestedatt--configuration--indexing--destination_langchain_update_indexing_pinecone))

<a id="nestedatt--configuration--indexing--destination_langchain_indexing_chroma_local_persistance"></a>
### Nested Schema for `configuration.indexing.destination_langchain_indexing_chroma_local_persistance`

Read-Only:

- `collection_name` (String) Name of the collection to use.
- `destination_path` (String) Path to the directory where chroma files will be written. The files will be placed inside that local mount.
- `mode` (String) must be one of ["chroma_local"]


<a id="nestedatt--configuration--indexing--destination_langchain_indexing_doc_array_hnsw_search"></a>
### Nested Schema for `configuration.indexing.destination_langchain_indexing_doc_array_hnsw_search`

Read-Only:

- `destination_path` (String) Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.
- `mode` (String) must be one of ["DocArrayHnswSearch"]


<a id="nestedatt--configuration--indexing--destination_langchain_indexing_pinecone"></a>
### Nested Schema for `configuration.indexing.destination_langchain_indexing_pinecone`

Read-Only:

- `index` (String) Pinecone index to use
- `mode` (String) must be one of ["pinecone"]
- `pinecone_environment` (String) Pinecone environment to use
- `pinecone_key` (String)


<a id="nestedatt--configuration--indexing--destination_langchain_update_indexing_chroma_local_persistance"></a>
### Nested Schema for `configuration.indexing.destination_langchain_update_indexing_chroma_local_persistance`

Read-Only:

- `collection_name` (String) Name of the collection to use.
- `destination_path` (String) Path to the directory where chroma files will be written. The files will be placed inside that local mount.
- `mode` (String) must be one of ["chroma_local"]


<a id="nestedatt--configuration--indexing--destination_langchain_update_indexing_doc_array_hnsw_search"></a>
### Nested Schema for `configuration.indexing.destination_langchain_update_indexing_doc_array_hnsw_search`

Read-Only:

- `destination_path` (String) Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.
- `mode` (String) must be one of ["DocArrayHnswSearch"]


<a id="nestedatt--configuration--indexing--destination_langchain_update_indexing_pinecone"></a>
### Nested Schema for `configuration.indexing.destination_langchain_update_indexing_pinecone`

Read-Only:

- `index` (String) Pinecone index to use
- `mode` (String) must be one of ["pinecone"]
- `pinecone_environment` (String) Pinecone environment to use
- `pinecone_key` (String)



<a id="nestedatt--configuration--processing"></a>
### Nested Schema for `configuration.processing`

Read-Only:

- `chunk_overlap` (Number) Size of overlap between chunks in tokens to store in vector store to better capture relevant context
- `chunk_size` (Number) Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
- `text_fields` (List of String) List of fields in the record that should be used to calculate the embedding. All other fields are passed along as meta fields. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.


