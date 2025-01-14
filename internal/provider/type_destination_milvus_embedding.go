// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationMilvusEmbedding struct {
	DestinationMilvusEmbeddingCohere          *DestinationMilvusEmbeddingCohere    `tfsdk:"destination_milvus_embedding_cohere"`
	DestinationMilvusEmbeddingFake            *DestinationLangchainEmbeddingFake   `tfsdk:"destination_milvus_embedding_fake"`
	DestinationMilvusEmbeddingFromField       *DestinationMilvusEmbeddingFromField `tfsdk:"destination_milvus_embedding_from_field"`
	DestinationMilvusEmbeddingOpenAI          *DestinationLangchainEmbeddingOpenAI `tfsdk:"destination_milvus_embedding_open_ai"`
	DestinationMilvusUpdateEmbeddingCohere    *DestinationMilvusEmbeddingCohere    `tfsdk:"destination_milvus_update_embedding_cohere"`
	DestinationMilvusUpdateEmbeddingFake      *DestinationLangchainEmbeddingFake   `tfsdk:"destination_milvus_update_embedding_fake"`
	DestinationMilvusUpdateEmbeddingFromField *DestinationMilvusEmbeddingFromField `tfsdk:"destination_milvus_update_embedding_from_field"`
	DestinationMilvusUpdateEmbeddingOpenAI    *DestinationLangchainEmbeddingOpenAI `tfsdk:"destination_milvus_update_embedding_open_ai"`
}
