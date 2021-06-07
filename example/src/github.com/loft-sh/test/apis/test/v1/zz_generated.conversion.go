// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	test "github.com/loft-sh/test/apis/test"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterRole)(nil), (*test.ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRole_To_test_ClusterRole(a.(*ClusterRole), b.(*test.ClusterRole), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*test.ClusterRole)(nil), (*ClusterRole)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_test_ClusterRole_To_v1_ClusterRole(a.(*test.ClusterRole), b.(*ClusterRole), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleList)(nil), (*test.ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleList_To_test_ClusterRoleList(a.(*ClusterRoleList), b.(*test.ClusterRoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*test.ClusterRoleList)(nil), (*ClusterRoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_test_ClusterRoleList_To_v1_ClusterRoleList(a.(*test.ClusterRoleList), b.(*ClusterRoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleSpec)(nil), (*test.ClusterRoleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec(a.(*ClusterRoleSpec), b.(*test.ClusterRoleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*test.ClusterRoleSpec)(nil), (*ClusterRoleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec(a.(*test.ClusterRoleSpec), b.(*ClusterRoleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleStatus)(nil), (*test.ClusterRoleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus(a.(*ClusterRoleStatus), b.(*test.ClusterRoleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*test.ClusterRoleStatus)(nil), (*ClusterRoleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus(a.(*test.ClusterRoleStatus), b.(*ClusterRoleStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ClusterRole_To_test_ClusterRole(in *ClusterRole, out *test.ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterRole_To_test_ClusterRole is an autogenerated conversion function.
func Convert_v1_ClusterRole_To_test_ClusterRole(in *ClusterRole, out *test.ClusterRole, s conversion.Scope) error {
	return autoConvert_v1_ClusterRole_To_test_ClusterRole(in, out, s)
}

func autoConvert_test_ClusterRole_To_v1_ClusterRole(in *test.ClusterRole, out *ClusterRole, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_test_ClusterRole_To_v1_ClusterRole is an autogenerated conversion function.
func Convert_test_ClusterRole_To_v1_ClusterRole(in *test.ClusterRole, out *ClusterRole, s conversion.Scope) error {
	return autoConvert_test_ClusterRole_To_v1_ClusterRole(in, out, s)
}

func autoConvert_v1_ClusterRoleList_To_test_ClusterRoleList(in *ClusterRoleList, out *test.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]test.ClusterRole)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ClusterRoleList_To_test_ClusterRoleList is an autogenerated conversion function.
func Convert_v1_ClusterRoleList_To_test_ClusterRoleList(in *ClusterRoleList, out *test.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleList_To_test_ClusterRoleList(in, out, s)
}

func autoConvert_test_ClusterRoleList_To_v1_ClusterRoleList(in *test.ClusterRoleList, out *ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterRole)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_test_ClusterRoleList_To_v1_ClusterRoleList is an autogenerated conversion function.
func Convert_test_ClusterRoleList_To_v1_ClusterRoleList(in *test.ClusterRoleList, out *ClusterRoleList, s conversion.Scope) error {
	return autoConvert_test_ClusterRoleList_To_v1_ClusterRoleList(in, out, s)
}

func autoConvert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec(in *ClusterRoleSpec, out *test.ClusterRoleSpec, s conversion.Scope) error {
	out.Rules = *(*[]string)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec is an autogenerated conversion function.
func Convert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec(in *ClusterRoleSpec, out *test.ClusterRoleSpec, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleSpec_To_test_ClusterRoleSpec(in, out, s)
}

func autoConvert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec(in *test.ClusterRoleSpec, out *ClusterRoleSpec, s conversion.Scope) error {
	out.Rules = *(*[]string)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec is an autogenerated conversion function.
func Convert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec(in *test.ClusterRoleSpec, out *ClusterRoleSpec, s conversion.Scope) error {
	return autoConvert_test_ClusterRoleSpec_To_v1_ClusterRoleSpec(in, out, s)
}

func autoConvert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus(in *ClusterRoleStatus, out *test.ClusterRoleStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus is an autogenerated conversion function.
func Convert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus(in *ClusterRoleStatus, out *test.ClusterRoleStatus, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleStatus_To_test_ClusterRoleStatus(in, out, s)
}

func autoConvert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus(in *test.ClusterRoleStatus, out *ClusterRoleStatus, s conversion.Scope) error {
	return nil
}

// Convert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus is an autogenerated conversion function.
func Convert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus(in *test.ClusterRoleStatus, out *ClusterRoleStatus, s conversion.Scope) error {
	return autoConvert_test_ClusterRoleStatus_To_v1_ClusterRoleStatus(in, out, s)
}
