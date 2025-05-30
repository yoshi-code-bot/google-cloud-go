// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/cloud/securitycenter/v2/backup_disaster_recovery.proto

package securitycenterpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information related to Google Cloud Backup and DR Service findings.
type BackupDisasterRecovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of a Backup and DR template which comprises one or more backup
	// policies. See the [Backup and DR
	// documentation](https://cloud.google.com/backup-disaster-recovery/docs/concepts/backup-plan#temp)
	// for more information. For example, `snap-ov`.
	BackupTemplate string `protobuf:"bytes,1,opt,name=backup_template,json=backupTemplate,proto3" json:"backup_template,omitempty"`
	// The names of Backup and DR policies that are associated with a template
	// and that define when to run a backup, how frequently to run a backup, and
	// how long to retain the backup image. For example, `onvaults`.
	Policies []string `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
	// The name of a Backup and DR host, which is managed by the backup and
	// recovery appliance and known to the management console. The host can be of
	// type Generic (for example, Compute Engine, SQL Server, Oracle DB, SMB file
	// system, etc.), vCenter, or an ESX server. See the [Backup and DR
	// documentation on
	// hosts](https://cloud.google.com/backup-disaster-recovery/docs/configuration/manage-hosts-and-their-applications)
	// for more information. For example, `centos7-01`.
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// The names of Backup and DR applications. An application is a VM, database,
	// or file system on a managed host monitored by a backup and recovery
	// appliance. For example, `centos7-01-vol00`, `centos7-01-vol01`,
	// `centos7-01-vol02`.
	Applications []string `protobuf:"bytes,4,rep,name=applications,proto3" json:"applications,omitempty"`
	// The name of the Backup and DR storage pool that the backup and recovery
	// appliance is storing data in. The storage pool could be of type Cloud,
	// Primary, Snapshot, or OnVault. See the [Backup and DR documentation on
	// storage
	// pools](https://cloud.google.com/backup-disaster-recovery/docs/concepts/storage-pools).
	// For example, `DiskPoolOne`.
	StoragePool string `protobuf:"bytes,5,opt,name=storage_pool,json=storagePool,proto3" json:"storage_pool,omitempty"`
	// The names of Backup and DR advanced policy options of a policy applying to
	// an application. See the [Backup and DR documentation on policy
	// options](https://cloud.google.com/backup-disaster-recovery/docs/create-plan/policy-settings).
	// For example, `skipofflineappsincongrp, nounmap`.
	PolicyOptions []string `protobuf:"bytes,6,rep,name=policy_options,json=policyOptions,proto3" json:"policy_options,omitempty"`
	// The name of the Backup and DR resource profile that specifies the storage
	// media for backups of application and VM data. See the [Backup and DR
	// documentation on
	// profiles](https://cloud.google.com/backup-disaster-recovery/docs/concepts/backup-plan#profile).
	// For example, `GCP`.
	Profile string `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	// The name of the Backup and DR appliance that captures, moves, and manages
	// the lifecycle of backup data. For example, `backup-server-57137`.
	Appliance string `protobuf:"bytes,8,opt,name=appliance,proto3" json:"appliance,omitempty"`
	// The backup type of the Backup and DR image.
	// For example, `Snapshot`, `Remote Snapshot`, `OnVault`.
	BackupType string `protobuf:"bytes,9,opt,name=backup_type,json=backupType,proto3" json:"backup_type,omitempty"`
	// The timestamp at which the Backup and DR backup was created.
	BackupCreateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=backup_create_time,json=backupCreateTime,proto3" json:"backup_create_time,omitempty"`
}

func (x *BackupDisasterRecovery) Reset() {
	*x = BackupDisasterRecovery{}
	mi := &file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupDisasterRecovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupDisasterRecovery) ProtoMessage() {}

func (x *BackupDisasterRecovery) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupDisasterRecovery.ProtoReflect.Descriptor instead.
func (*BackupDisasterRecovery) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescGZIP(), []int{0}
}

func (x *BackupDisasterRecovery) GetBackupTemplate() string {
	if x != nil {
		return x.BackupTemplate
	}
	return ""
}

func (x *BackupDisasterRecovery) GetPolicies() []string {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *BackupDisasterRecovery) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *BackupDisasterRecovery) GetApplications() []string {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *BackupDisasterRecovery) GetStoragePool() string {
	if x != nil {
		return x.StoragePool
	}
	return ""
}

func (x *BackupDisasterRecovery) GetPolicyOptions() []string {
	if x != nil {
		return x.PolicyOptions
	}
	return nil
}

func (x *BackupDisasterRecovery) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *BackupDisasterRecovery) GetAppliance() string {
	if x != nil {
		return x.Appliance
	}
	return ""
}

func (x *BackupDisasterRecovery) GetBackupType() string {
	if x != nil {
		return x.BackupType
	}
	return ""
}

func (x *BackupDisasterRecovery) GetBackupCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BackupCreateTime
	}
	return nil
}

var File_google_cloud_securitycenter_v2_backup_disaster_recovery_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x03, 0x0a, 0x16, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x69, 0x73, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0xf5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x1b, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x44, 0x69, 0x73, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescData = file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_goTypes = []any{
	(*BackupDisasterRecovery)(nil), // 0: google.cloud.securitycenter.v2.BackupDisasterRecovery
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v2.BackupDisasterRecovery.backup_create_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_init() }
func file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_init() {
	if File_google_cloud_securitycenter_v2_backup_disaster_recovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_backup_disaster_recovery_proto = out.File
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_backup_disaster_recovery_proto_depIdxs = nil
}
