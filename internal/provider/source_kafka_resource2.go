package provider

import (
    "airbyte/internal/sdk"
    "context"

    "github.com/hashicorp/terraform-plugin-framework/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

// KafkaResource defines the resource implementation.
type KafkaResource struct {
    // Adicione aqui os campos necessários para interagir com o Kafka, como configurações de conexão.
}

// KafkaResourceModel descreve o modelo de dados do recurso.
type KafkaResourceModel struct {
    Configuration KafkaConfiguration `tfsdk:"configuration"`
    // Adicione outros campos do recurso Kafka aqui.
}

// KafkaConfiguration é um exemplo de como você pode estruturar a configuração do Kafka.
type KafkaConfiguration struct {
    KafkaServer     schema.StringAttribute
    KafkaTopic      schema.StringAttribute
    // Adicione outros atributos de configuração do Kafka aqui.
}

func NewKafkaResource() resource.Resource {
    return &KafkaResource{}
}

func (r *KafkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = "example_kafka"
}

func (r *KafkaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        MarkdownDescription: "Exemplo de recurso Kafka",

        Attributes: map[string]schema.Attribute{
            "configuration": schema.SingleNestedAttribute{
                Required: true,
                Attributes: KafkaConfiguration{
                    KafkaServer: schema.StringAttribute{
                        Required: true,
                        Description: "Endereço do servidor Kafka.",
                    },
                    KafkaTopic: schema.StringAttribute{
                        Required: true,
                        Description: "Tópico do Kafka a ser gerenciado.",
                    },
                    // Adicione outros atributos de configuração do Kafka aqui.
                },
            },
        },
    }
}