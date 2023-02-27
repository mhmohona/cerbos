// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0
// Source: cerbos/request/v1/request.proto

package requestv1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlanResourcesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_PlanResourcesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceSetRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_CheckResourceSetRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ResourceSet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_ResourceSet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AttributesMap) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_AttributesMap_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceBatchRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_CheckResourceBatchRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourceBatchRequest_BatchEntry) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_CheckResourceBatchRequest_BatchEntry_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_CheckResourcesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CheckResourcesRequest_ResourceEntry) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_CheckResourcesRequest_ResourceEntry_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AuxData) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_AuxData_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AuxData_JWT) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_AuxData_JWT_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *File) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_File_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundValidateRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_PlaygroundValidateRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundTestRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_PlaygroundTestRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundEvaluateRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_PlaygroundEvaluateRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *PlaygroundProxyRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_PlaygroundProxyRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AddOrUpdatePolicyRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_AddOrUpdatePolicyRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListAuditLogEntriesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_ListAuditLogEntriesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListAuditLogEntriesRequest_TimeRange) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_ListAuditLogEntriesRequest_TimeRange_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ListPoliciesRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_ListPoliciesRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetPolicyRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_GetPolicyRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *DisablePolicyRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_DisablePolicyRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *EnablePolicyRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_EnablePolicyRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *AddOrUpdateSchemaRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_AddOrUpdateSchemaRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GetSchemaRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_GetSchemaRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *DeleteSchemaRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_DeleteSchemaRequest_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *ReloadStoreRequest) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_request_v1_ReloadStoreRequest_hashpb_sum(m, hasher, ignore)
	}
}
