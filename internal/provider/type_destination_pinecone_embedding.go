// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationPineconeEmbedding struct {
	DestinationPineconeEmbeddingCohere       *DestinationMilvusEmbeddingCohere    `tfsdk:"destination_pinecone_embedding_cohere"`
	DestinationPineconeEmbeddingFake         *DestinationLangchainEmbeddingFake   `tfsdk:"destination_pinecone_embedding_fake"`
	DestinationPineconeEmbeddingOpenAI       *DestinationLangchainEmbeddingOpenAI `tfsdk:"destination_pinecone_embedding_open_ai"`
	DestinationPineconeUpdateEmbeddingCohere *DestinationMilvusEmbeddingCohere    `tfsdk:"destination_pinecone_update_embedding_cohere"`
	DestinationPineconeUpdateEmbeddingFake   *DestinationLangchainEmbeddingFake   `tfsdk:"destination_pinecone_update_embedding_fake"`
	DestinationPineconeUpdateEmbeddingOpenAI *DestinationLangchainEmbeddingOpenAI `tfsdk:"destination_pinecone_update_embedding_open_ai"`
}
