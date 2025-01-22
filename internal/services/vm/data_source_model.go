// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/nirvana-terraform/internal/customfield"
)

type VMDataSourceModel struct {
	VMID        types.String                                               `tfsdk:"vm_id" path:"vm_id,required"`
	CreatedAt   types.String                                               `tfsdk:"created_at" json:"created_at,computed"`
	ID          types.String                                               `tfsdk:"id" json:"id,computed"`
	Name        types.String                                               `tfsdk:"name" json:"name,computed"`
	PublicIP    types.String                                               `tfsdk:"public_ip" json:"public_ip,computed"`
	Region      types.String                                               `tfsdk:"region" json:"region,computed"`
	Status      types.String                                               `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                               `tfsdk:"updated_at" json:"updated_at,computed"`
	BootVolume  customfield.NestedObject[VMBootVolumeDataSourceModel]      `tfsdk:"boot_volume" json:"boot_volume,computed"`
	CPUConfig   customfield.NestedObject[VMCPUConfigDataSourceModel]       `tfsdk:"cpu_config" json:"cpu_config,computed"`
	DataVolumes customfield.NestedObjectList[VMDataVolumesDataSourceModel] `tfsdk:"data_volumes" json:"data_volumes,computed"`
	MemConfig   customfield.NestedObject[VMMemConfigDataSourceModel]       `tfsdk:"mem_config" json:"mem_config,computed"`
	VPC         customfield.NestedObject[VMVPCDataSourceModel]             `tfsdk:"vpc" json:"vpc,computed"`
}

type VMBootVolumeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type VMCPUConfigDataSourceModel struct {
	Cores types.Int64 `tfsdk:"cores" json:"cores,computed"`
}

type VMDataVolumesDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Size types.Int64  `tfsdk:"size" json:"size,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type VMMemConfigDataSourceModel struct {
	Size types.Int64 `tfsdk:"size" json:"size,computed"`
}

type VMVPCDataSourceModel struct {
	ID            types.String                                                    `tfsdk:"id" json:"id,computed"`
	CreatedAt     types.String                                                    `tfsdk:"created_at" json:"created_at,computed"`
	FirewallRules customfield.NestedObjectList[VMVPCFirewallRulesDataSourceModel] `tfsdk:"firewall_rules" json:"firewall_rules,computed"`
	Name          types.String                                                    `tfsdk:"name" json:"name,computed"`
	Region        types.String                                                    `tfsdk:"region" json:"region,computed"`
	Status        types.String                                                    `tfsdk:"status" json:"status,computed"`
	Subnet        customfield.NestedObject[VMVPCSubnetDataSourceModel]            `tfsdk:"subnet" json:"subnet,computed"`
	UpdatedAt     types.String                                                    `tfsdk:"updated_at" json:"updated_at,computed"`
}

type VMVPCFirewallRulesDataSourceModel struct {
	ID          types.String                                                           `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String                                                           `tfsdk:"created_at" json:"created_at,computed"`
	Destination customfield.NestedObject[VMVPCFirewallRulesDestinationDataSourceModel] `tfsdk:"destination" json:"destination,computed"`
	Name        types.String                                                           `tfsdk:"name" json:"name,computed"`
	Protocol    types.String                                                           `tfsdk:"protocol" json:"protocol,computed"`
	Source      customfield.NestedObject[VMVPCFirewallRulesSourceDataSourceModel]      `tfsdk:"source" json:"source,computed"`
	Status      types.String                                                           `tfsdk:"status" json:"status,computed"`
	UpdatedAt   types.String                                                           `tfsdk:"updated_at" json:"updated_at,computed"`
	VPCID       types.String                                                           `tfsdk:"vpc_id" json:"vpc_id,computed"`
}

type VMVPCFirewallRulesDestinationDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VMVPCFirewallRulesSourceDataSourceModel struct {
	Address types.String                   `tfsdk:"address" json:"address,computed"`
	Ports   customfield.List[types.String] `tfsdk:"ports" json:"ports,computed"`
}

type VMVPCSubnetDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Cidr      types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
