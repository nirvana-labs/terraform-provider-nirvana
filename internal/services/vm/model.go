// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/apijson"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/customfield"
)

type VMModel struct {
	ID            types.String                                     `tfsdk:"id" json:"id,computed"`
	Name          types.String                                     `tfsdk:"name" json:"name,required"`
	NeedPublicIP  types.Bool                                       `tfsdk:"need_public_ip" json:"need_public_ip,required"`
	OSImageName   types.String                                     `tfsdk:"os_image_name" json:"os_image_name,required"`
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
	VPCID         types.String                                     `tfsdk:"vpc_id" json:"vpc_id,computed"`
	CPUConfig     customfield.NestedObject[VMCPUConfigModel]       `tfsdk:"cpu_config" json:"cpu_config,computed"`
	MemConfig     customfield.NestedObject[VMMemConfigModel]       `tfsdk:"mem_config" json:"mem_config,computed"`
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
