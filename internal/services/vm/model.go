// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/apijson"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

type VMModel struct {
	ID            types.String                                     `tfsdk:"id" json:"id,computed"`
	Name          types.String                                     `tfsdk:"name" json:"name,required"`
	NeedPublicIP  types.Bool                                       `tfsdk:"need_public_ip" json:"need_public_ip,required"`
	OsImageID     types.Int64                                      `tfsdk:"os_image_id" json:"os_image_id,required"`
	Region        types.String                                     `tfsdk:"region" json:"region,required"`
	SourceAddress types.String                                     `tfsdk:"source_address" json:"source_address,required"`
	Ports         *[]types.String                                  `tfsdk:"ports" json:"ports,required"`
	SSHKey        *VMSSHKeyModel                                   `tfsdk:"ssh_key" json:"ssh_key,required"`
	SubnetID      types.String                                     `tfsdk:"subnet_id" json:"subnet_id,optional"`
	BootVolume    *VMBootVolumeModel                               `tfsdk:"boot_volume" json:"boot_volume,required"`
	CPU           *VMCPUModel                                      `tfsdk:"cpu" json:"cpu,required"`
	Ram           *VMRamModel                                      `tfsdk:"ram" json:"ram,required"`
	DataVolumes   customfield.NestedObjectList[VMDataVolumesModel] `tfsdk:"data_volumes" json:"data_volumes,computed_optional"`
	CreatedAt     types.String                                     `tfsdk:"created_at" json:"created_at,computed"`
	Kind          types.String                                     `tfsdk:"kind" json:"kind,computed"`
	PublicIP      types.String                                     `tfsdk:"public_ip" json:"public_ip,computed"`
	ResourceID    types.String                                     `tfsdk:"resource_id" json:"resource_id,computed"`
	Status        types.String                                     `tfsdk:"status" json:"status,computed"`
	Type          types.String                                     `tfsdk:"type" json:"type,computed"`
	UpdatedAt     types.String                                     `tfsdk:"updated_at" json:"updated_at,computed"`
	CPUConfig     customfield.NestedObject[VMCPUConfigModel]       `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemConfig     customfield.NestedObject[VMMemConfigModel]       `tfsdk:"mem_config" json:"mem_config,computed"`
	VPC           customfield.NestedObject[VMVPCModel]             `tfsdk:"vpc" json:"vpc,computed"`
}

func (m VMModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m VMModel) MarshalJSONForUpdate(state VMModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type VMSSHKeyModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type VMBootVolumeModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type VMCPUModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,required"`
}

type VMRamModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,required"`
}

type VMDataVolumesModel struct {
	Size types.Int64  `tfsdk:"size" json:"size,required"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}

type VMCPUConfigModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,computed"`
}

type VMMemConfigModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}

type VMVPCModel struct {
	ID            types.String                                          `tfsdk:"id" json:"id,computed"`
	CreatedAt     types.String                                          `tfsdk:"created_at" json:"created_at,computed"`
	FirewallRules customfield.NestedObjectList[VMVPCFirewallRulesModel] `tfsdk:"firewall_rules" json:"firewall_rules,computed"`
	Name          types.String                                          `tfsdk:"name" json:"name,computed"`
	Region        types.String                                          `tfsdk:"region" json:"region,computed"`
	Status        types.String                                          `tfsdk:"status" json:"status,computed"`
	Subnet        customfield.NestedObject[VMVPCSubnetModel]            `tfsdk:"subnet" json:"subnet,computed"`
	UpdatedAt     types.String                                          `tfsdk:"updated_at" json:"updated_at,computed"`
}

type VMVPCFirewallRulesModel struct {
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String                                                 `tfsdk:"created_at" json:"created_at,computed"`
	Destination customfield.NestedObject[VMVPCFirewallRulesDestinationModel] `tfsdk:"destination" json:"destination,computed"`
	Name        types.String                                                 `tfsdk:"name" json:"name,computed"`
	Protocol    types.String                                                 `tfsdk:"protocol" json:"protocol,computed"`
	Source      customfield.NestedObject[VMVPCFirewallRulesSourceModel]      `tfsdk:"source" json:"source,computed"`
	Status      types.String                                                 `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                                 `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID       types.String                                                 `tfsdk:"vpc_id" json:"vpc_id,computed"`
}

type VMVPCFirewallRulesDestinationModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VMVPCFirewallRulesSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VMVPCSubnetModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
