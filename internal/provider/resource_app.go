package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = &porterAppResource{}

// porterAppResource is the resource for the Porter application
type porterAppResource struct{}

// NewPorterAppResource creates an application on a porter cluster
func NewPorterAppResource() resource.Resource {
	return &porterAppResource{}
}

func (r *porterAppResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

// Schema defines the schema for the resource.
func (r *porterAppResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create an application on a Porter cluster",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the application",
				Required:    true,
			},
			"image": schema.StringAttribute{
				Description: "The image of the application",
				Required:    true,
			},
			"services": schema.ListNestedAttribute{
				Description: "Services to add to the application, such as web, worker, and jobs",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "The type of the service. Must be either 'web', 'worker', or 'job'",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the service within the application",
							Required:    true,
						},
						"cpu_cores": schema.NumberAttribute{
							Description: "The number of CPU cores to allocate to the service. 0.1 is 10 percent of a single vcpu core",
							Optional:    true,
						},
						"ram_megabytes": schema.NumberAttribute{
							Description: "The number of megabytes of RAM to allocate to the service in Mb",
							Optional:    true,
						},
						"command": schema.StringAttribute{
							Description: "The command to run for the service",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *porterAppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *porterAppResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *porterAppResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *porterAppResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
