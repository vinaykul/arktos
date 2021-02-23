// +build !ignore_autogenerated

/*
Copyright 2020 Authors of Arktos.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.MutatingWebhook)(nil), (*admissionregistration.MutatingWebhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook(a.(*v1.MutatingWebhook), b.(*admissionregistration.MutatingWebhook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.MutatingWebhook)(nil), (*v1.MutatingWebhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook(a.(*admissionregistration.MutatingWebhook), b.(*v1.MutatingWebhook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MutatingWebhookConfiguration)(nil), (*admissionregistration.MutatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(a.(*v1.MutatingWebhookConfiguration), b.(*admissionregistration.MutatingWebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.MutatingWebhookConfiguration)(nil), (*v1.MutatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration(a.(*admissionregistration.MutatingWebhookConfiguration), b.(*v1.MutatingWebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.MutatingWebhookConfigurationList)(nil), (*admissionregistration.MutatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(a.(*v1.MutatingWebhookConfigurationList), b.(*admissionregistration.MutatingWebhookConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.MutatingWebhookConfigurationList)(nil), (*v1.MutatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1_MutatingWebhookConfigurationList(a.(*admissionregistration.MutatingWebhookConfigurationList), b.(*v1.MutatingWebhookConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.Rule)(nil), (*admissionregistration.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Rule_To_admissionregistration_Rule(a.(*v1.Rule), b.(*admissionregistration.Rule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.Rule)(nil), (*v1.Rule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_Rule_To_v1_Rule(a.(*admissionregistration.Rule), b.(*v1.Rule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RuleWithOperations)(nil), (*admissionregistration.RuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(a.(*v1.RuleWithOperations), b.(*admissionregistration.RuleWithOperations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.RuleWithOperations)(nil), (*v1.RuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(a.(*admissionregistration.RuleWithOperations), b.(*v1.RuleWithOperations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServiceReference)(nil), (*admissionregistration.ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceReference_To_admissionregistration_ServiceReference(a.(*v1.ServiceReference), b.(*admissionregistration.ServiceReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ServiceReference)(nil), (*v1.ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ServiceReference_To_v1_ServiceReference(a.(*admissionregistration.ServiceReference), b.(*v1.ServiceReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ValidatingWebhook)(nil), (*admissionregistration.ValidatingWebhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook(a.(*v1.ValidatingWebhook), b.(*admissionregistration.ValidatingWebhook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingWebhook)(nil), (*v1.ValidatingWebhook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook(a.(*admissionregistration.ValidatingWebhook), b.(*v1.ValidatingWebhook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ValidatingWebhookConfiguration)(nil), (*admissionregistration.ValidatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(a.(*v1.ValidatingWebhookConfiguration), b.(*admissionregistration.ValidatingWebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingWebhookConfiguration)(nil), (*v1.ValidatingWebhookConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration(a.(*admissionregistration.ValidatingWebhookConfiguration), b.(*v1.ValidatingWebhookConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ValidatingWebhookConfigurationList)(nil), (*admissionregistration.ValidatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(a.(*v1.ValidatingWebhookConfigurationList), b.(*admissionregistration.ValidatingWebhookConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingWebhookConfigurationList)(nil), (*v1.ValidatingWebhookConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1_ValidatingWebhookConfigurationList(a.(*admissionregistration.ValidatingWebhookConfigurationList), b.(*v1.ValidatingWebhookConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.WebhookClientConfig)(nil), (*admissionregistration.WebhookClientConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(a.(*v1.WebhookClientConfig), b.(*admissionregistration.WebhookClientConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.WebhookClientConfig)(nil), (*v1.WebhookClientConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(a.(*admissionregistration.WebhookClientConfig), b.(*v1.WebhookClientConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook(in *v1.MutatingWebhook, out *admissionregistration.MutatingWebhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]admissionregistration.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.MatchPolicy = (*admissionregistration.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	out.SideEffects = (*admissionregistration.SideEffectClass)(unsafe.Pointer(in.SideEffects))
	out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
	out.AdmissionReviewVersions = *(*[]string)(unsafe.Pointer(&in.AdmissionReviewVersions))
	out.ReinvocationPolicy = (*admissionregistration.ReinvocationPolicyType)(unsafe.Pointer(in.ReinvocationPolicy))
	return nil
}

// Convert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook is an autogenerated conversion function.
func Convert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook(in *v1.MutatingWebhook, out *admissionregistration.MutatingWebhook, s conversion.Scope) error {
	return autoConvert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook(in, out, s)
}

func autoConvert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook(in *admissionregistration.MutatingWebhook, out *v1.MutatingWebhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]v1.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*v1.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.MatchPolicy = (*v1.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	out.SideEffects = (*v1.SideEffectClass)(unsafe.Pointer(in.SideEffects))
	out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
	out.AdmissionReviewVersions = *(*[]string)(unsafe.Pointer(&in.AdmissionReviewVersions))
	out.ReinvocationPolicy = (*v1.ReinvocationPolicyType)(unsafe.Pointer(in.ReinvocationPolicy))
	return nil
}

// Convert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook is an autogenerated conversion function.
func Convert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook(in *admissionregistration.MutatingWebhook, out *v1.MutatingWebhook, s conversion.Scope) error {
	return autoConvert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook(in, out, s)
}

func autoConvert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]admissionregistration.MutatingWebhook, len(*in))
		for i := range *in {
			if err := Convert_v1_MutatingWebhook_To_admissionregistration_MutatingWebhook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Webhooks = nil
	}
	return nil
}

// Convert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration is an autogenerated conversion function.
func Convert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in, out, s)
}

func autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1.MutatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]v1.MutatingWebhook, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_MutatingWebhook_To_v1_MutatingWebhook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Webhooks = nil
	}
	return nil
}

// Convert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1.MutatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration(in, out, s)
}

func autoConvert_v1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]admissionregistration.MutatingWebhookConfiguration, len(*in))
		for i := range *in {
			if err := Convert_v1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_v1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_v1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1.MutatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.MutatingWebhookConfiguration, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_MutatingWebhookConfiguration_To_v1_MutatingWebhookConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1_MutatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1.MutatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1_MutatingWebhookConfigurationList(in, out, s)
}

func autoConvert_v1_Rule_To_admissionregistration_Rule(in *v1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.Scope = (*admissionregistration.ScopeType)(unsafe.Pointer(in.Scope))
	return nil
}

// Convert_v1_Rule_To_admissionregistration_Rule is an autogenerated conversion function.
func Convert_v1_Rule_To_admissionregistration_Rule(in *v1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	return autoConvert_v1_Rule_To_admissionregistration_Rule(in, out, s)
}

func autoConvert_admissionregistration_Rule_To_v1_Rule(in *admissionregistration.Rule, out *v1.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.Scope = (*v1.ScopeType)(unsafe.Pointer(in.Scope))
	return nil
}

// Convert_admissionregistration_Rule_To_v1_Rule is an autogenerated conversion function.
func Convert_admissionregistration_Rule_To_v1_Rule(in *admissionregistration.Rule, out *v1.Rule, s conversion.Scope) error {
	return autoConvert_admissionregistration_Rule_To_v1_Rule(in, out, s)
}

func autoConvert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]admissionregistration.OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_v1_Rule_To_admissionregistration_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations is an autogenerated conversion function.
func Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in, out, s)
}

func autoConvert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1.RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]v1.OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_admissionregistration_Rule_To_v1_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations is an autogenerated conversion function.
func Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(in, out, s)
}

func autoConvert_v1_ServiceReference_To_admissionregistration_ServiceReference(in *v1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Path = (*string)(unsafe.Pointer(in.Path))
	if err := metav1.Convert_Pointer_int32_To_int32(&in.Port, &out.Port, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ServiceReference_To_admissionregistration_ServiceReference is an autogenerated conversion function.
func Convert_v1_ServiceReference_To_admissionregistration_ServiceReference(in *v1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	return autoConvert_v1_ServiceReference_To_admissionregistration_ServiceReference(in, out, s)
}

func autoConvert_admissionregistration_ServiceReference_To_v1_ServiceReference(in *admissionregistration.ServiceReference, out *v1.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Path = (*string)(unsafe.Pointer(in.Path))
	if err := metav1.Convert_int32_To_Pointer_int32(&in.Port, &out.Port, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_ServiceReference_To_v1_ServiceReference is an autogenerated conversion function.
func Convert_admissionregistration_ServiceReference_To_v1_ServiceReference(in *admissionregistration.ServiceReference, out *v1.ServiceReference, s conversion.Scope) error {
	return autoConvert_admissionregistration_ServiceReference_To_v1_ServiceReference(in, out, s)
}

func autoConvert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook(in *v1.ValidatingWebhook, out *admissionregistration.ValidatingWebhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]admissionregistration.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.MatchPolicy = (*admissionregistration.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	out.SideEffects = (*admissionregistration.SideEffectClass)(unsafe.Pointer(in.SideEffects))
	out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
	out.AdmissionReviewVersions = *(*[]string)(unsafe.Pointer(&in.AdmissionReviewVersions))
	return nil
}

// Convert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook is an autogenerated conversion function.
func Convert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook(in *v1.ValidatingWebhook, out *admissionregistration.ValidatingWebhook, s conversion.Scope) error {
	return autoConvert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook(in, out, s)
}

func autoConvert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook(in *admissionregistration.ValidatingWebhook, out *v1.ValidatingWebhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]v1.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*v1.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.MatchPolicy = (*v1.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	out.NamespaceSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*metav1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	out.SideEffects = (*v1.SideEffectClass)(unsafe.Pointer(in.SideEffects))
	out.TimeoutSeconds = (*int32)(unsafe.Pointer(in.TimeoutSeconds))
	out.AdmissionReviewVersions = *(*[]string)(unsafe.Pointer(&in.AdmissionReviewVersions))
	return nil
}

// Convert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook(in *admissionregistration.ValidatingWebhook, out *v1.ValidatingWebhook, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook(in, out, s)
}

func autoConvert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]admissionregistration.ValidatingWebhook, len(*in))
		for i := range *in {
			if err := Convert_v1_ValidatingWebhook_To_admissionregistration_ValidatingWebhook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Webhooks = nil
	}
	return nil
}

// Convert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration is an autogenerated conversion function.
func Convert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in, out, s)
}

func autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1.ValidatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]v1.ValidatingWebhook, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_ValidatingWebhook_To_v1_ValidatingWebhook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Webhooks = nil
	}
	return nil
}

// Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1.ValidatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration(in, out, s)
}

func autoConvert_v1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]admissionregistration.ValidatingWebhookConfiguration, len(*in))
		for i := range *in {
			if err := Convert_v1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_v1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_v1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ValidatingWebhookConfiguration, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1_ValidatingWebhookConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1_ValidatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1_ValidatingWebhookConfigurationList(in, out, s)
}

func autoConvert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
	out.URL = (*string)(unsafe.Pointer(in.URL))
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(admissionregistration.ServiceReference)
		if err := Convert_v1_ServiceReference_To_admissionregistration_ServiceReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Service = nil
	}
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig is an autogenerated conversion function.
func Convert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
	return autoConvert_v1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in, out, s)
}

func autoConvert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1.WebhookClientConfig, s conversion.Scope) error {
	out.URL = (*string)(unsafe.Pointer(in.URL))
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(v1.ServiceReference)
		if err := Convert_admissionregistration_ServiceReference_To_v1_ServiceReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Service = nil
	}
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig is an autogenerated conversion function.
func Convert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1.WebhookClientConfig, s conversion.Scope) error {
	return autoConvert_admissionregistration_WebhookClientConfig_To_v1_WebhookClientConfig(in, out, s)
}