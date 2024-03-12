package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = &gkeClusterResource{}

// gkeClusterResource is the resource for the Porter GKE cluster
type gkeClusterResource struct{}

// NewGkeClusterResource creates a new GKE cluster resource
func NewGkeClusterResource() resource.Resource {
	return &gkeClusterResource{}
}

func (r *gkeClusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gke_cluster"
}

// Schema defines the schema for the resource.
func (r *gkeClusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Porter-managed GKE cluster",
		Attributes: map[string]schema.Attribute{
			"is_soc2_compliant": schema.BoolAttribute{
				Description: "Ensure that the cluster is SOC2 compliant",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the cluster",
				Required:    true,
			},
			"version": schema.StringAttribute{
				Description: "The kubernetes version of the cluster",
				Required:    true,
			},
			"region": schema.StringAttribute{
				Description: "The region of the cluster",
				Required:    true,
			},
			"vpc_cidr": schema.StringAttribute{
				Description: "The CIDR of the cluster vpc. Must be a /16",
				Optional:    true,
			},
			"service_cidr": schema.StringAttribute{
				Description: "The CIDR of the kubernetes services within the cluster. Must be a /16 within 172.16.0.0/12",
				Optional:    true,
			},
			"node_groups": schema.ListNestedAttribute{
				Description: "Node groups to add to the cluster, on top of the system and monitoring node groups",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "The type of the node group. Must be either 'application', 'custom', 'monitoring' or 'system'. Monitoring and System will be added by default if unspecified. Use 'custom' for GPU nodes.",
							Required:    true,
						},
						"min_nodes": schema.NumberAttribute{
							Description: "The minimum size of the node group",
							Required:    true,
						},
						"max_nodes": schema.NumberAttribute{
							Description: "The maximum size of the node group",
							Required:    true,
						},
						"instance_type": schema.StringAttribute{
							Description: "The instance type of the node group",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *gkeClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *gkeClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *gkeClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *gkeClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
