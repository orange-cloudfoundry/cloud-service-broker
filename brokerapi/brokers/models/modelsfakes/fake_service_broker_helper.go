// Code generated by counterfeiter. DO NOT EDIT.
package modelsfakes

import (
	"context"
	"sync"

	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceBrokerHelper struct {
	ProvisionStub        func(ctx context.Context, instanceId string, details brokerapi.ProvisionDetails, plan models.ServicePlan) (models.ServiceInstanceDetails, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		ctx        context.Context
		instanceId string
		details    brokerapi.ProvisionDetails
		plan       models.ServicePlan
	}
	provisionReturns struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}
	BindStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (models.ServiceBindingCredentials, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}
	bindReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	BuildInstanceCredentialsStub        func(ctx context.Context, bindRecord models.ServiceBindingCredentials, instanceRecord models.ServiceInstanceDetails) (map[string]string, error)
	buildInstanceCredentialsMutex       sync.RWMutex
	buildInstanceCredentialsArgsForCall []struct {
		ctx            context.Context
		bindRecord     models.ServiceBindingCredentials
		instanceRecord models.ServiceInstanceDetails
	}
	buildInstanceCredentialsReturns struct {
		result1 map[string]string
		result2 error
	}
	buildInstanceCredentialsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	UnbindStub        func(ctx context.Context, details models.ServiceBindingCredentials) error
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		ctx     context.Context
		details models.ServiceBindingCredentials
	}
	unbindReturns struct {
		result1 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 error
	}
	DeprovisionStub        func(ctx context.Context, instance models.ServiceInstanceDetails, details brokerapi.DeprovisionDetails) (operationId *string, err error)
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		ctx      context.Context
		instance models.ServiceInstanceDetails
		details  brokerapi.DeprovisionDetails
	}
	deprovisionReturns struct {
		result1 *string
		result2 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 *string
		result2 error
	}
	PollInstanceStub        func(ctx context.Context, instance models.ServiceInstanceDetails) (bool, error)
	pollInstanceMutex       sync.RWMutex
	pollInstanceArgsForCall []struct {
		ctx      context.Context
		instance models.ServiceInstanceDetails
	}
	pollInstanceReturns struct {
		result1 bool
		result2 error
	}
	pollInstanceReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ProvisionsAsyncStub        func() bool
	provisionsAsyncMutex       sync.RWMutex
	provisionsAsyncArgsForCall []struct{}
	provisionsAsyncReturns     struct {
		result1 bool
	}
	provisionsAsyncReturnsOnCall map[int]struct {
		result1 bool
	}
	DeprovisionsAsyncStub        func() bool
	deprovisionsAsyncMutex       sync.RWMutex
	deprovisionsAsyncArgsForCall []struct{}
	deprovisionsAsyncReturns     struct {
		result1 bool
	}
	deprovisionsAsyncReturnsOnCall map[int]struct {
		result1 bool
	}
	UpdateInstanceDetailsStub        func(ctx context.Context, instance *models.ServiceInstanceDetails) error
	updateInstanceDetailsMutex       sync.RWMutex
	updateInstanceDetailsArgsForCall []struct {
		ctx      context.Context
		instance *models.ServiceInstanceDetails
	}
	updateInstanceDetailsReturns struct {
		result1 error
	}
	updateInstanceDetailsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBrokerHelper) Provision(ctx context.Context, instanceId string, details brokerapi.ProvisionDetails, plan models.ServicePlan) (models.ServiceInstanceDetails, error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		ctx        context.Context
		instanceId string
		details    brokerapi.ProvisionDetails
		plan       models.ServicePlan
	}{ctx, instanceId, details, plan})
	fake.recordInvocation("Provision", []interface{}{ctx, instanceId, details, plan})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(ctx, instanceId, details, plan)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.provisionReturns.result1, fake.provisionReturns.result2
}

func (fake *FakeServiceBrokerHelper) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) ProvisionArgsForCall(i int) (context.Context, string, brokerapi.ProvisionDetails, models.ServicePlan) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].ctx, fake.provisionArgsForCall[i].instanceId, fake.provisionArgsForCall[i].details, fake.provisionArgsForCall[i].plan
}

func (fake *FakeServiceBrokerHelper) ProvisionReturns(result1 models.ServiceInstanceDetails, result2 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) ProvisionReturnsOnCall(i int, result1 models.ServiceInstanceDetails, result2 error) {
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 models.ServiceInstanceDetails
			result2 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails) (models.ServiceBindingCredentials, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}{ctx, instanceID, bindingID, details})
	fake.recordInvocation("Bind", []interface{}{ctx, instanceID, bindingID, details})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(ctx, instanceID, bindingID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}

func (fake *FakeServiceBrokerHelper) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BindArgsForCall(i int) (context.Context, string, string, brokerapi.BindDetails) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].ctx, fake.bindArgsForCall[i].instanceID, fake.bindArgsForCall[i].bindingID, fake.bindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) BindReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BindReturnsOnCall(i int, result1 models.ServiceBindingCredentials, result2 error) {
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBindingCredentials
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentials(ctx context.Context, bindRecord models.ServiceBindingCredentials, instanceRecord models.ServiceInstanceDetails) (map[string]string, error) {
	fake.buildInstanceCredentialsMutex.Lock()
	ret, specificReturn := fake.buildInstanceCredentialsReturnsOnCall[len(fake.buildInstanceCredentialsArgsForCall)]
	fake.buildInstanceCredentialsArgsForCall = append(fake.buildInstanceCredentialsArgsForCall, struct {
		ctx            context.Context
		bindRecord     models.ServiceBindingCredentials
		instanceRecord models.ServiceInstanceDetails
	}{ctx, bindRecord, instanceRecord})
	fake.recordInvocation("BuildInstanceCredentials", []interface{}{ctx, bindRecord, instanceRecord})
	fake.buildInstanceCredentialsMutex.Unlock()
	if fake.BuildInstanceCredentialsStub != nil {
		return fake.BuildInstanceCredentialsStub(ctx, bindRecord, instanceRecord)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildInstanceCredentialsReturns.result1, fake.buildInstanceCredentialsReturns.result2
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsCallCount() int {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return len(fake.buildInstanceCredentialsArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsArgsForCall(i int) (context.Context, models.ServiceBindingCredentials, models.ServiceInstanceDetails) {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.buildInstanceCredentialsArgsForCall[i].ctx, fake.buildInstanceCredentialsArgsForCall[i].bindRecord, fake.buildInstanceCredentialsArgsForCall[i].instanceRecord
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsReturns(result1 map[string]string, result2 error) {
	fake.BuildInstanceCredentialsStub = nil
	fake.buildInstanceCredentialsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.BuildInstanceCredentialsStub = nil
	if fake.buildInstanceCredentialsReturnsOnCall == nil {
		fake.buildInstanceCredentialsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.buildInstanceCredentialsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Unbind(ctx context.Context, details models.ServiceBindingCredentials) error {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		ctx     context.Context
		details models.ServiceBindingCredentials
	}{ctx, details})
	fake.recordInvocation("Unbind", []interface{}{ctx, details})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(ctx, details)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unbindReturns.result1
}

func (fake *FakeServiceBrokerHelper) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) UnbindArgsForCall(i int) (context.Context, models.ServiceBindingCredentials) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].ctx, fake.unbindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) UnbindReturnsOnCall(i int, result1 error) {
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Deprovision(ctx context.Context, instance models.ServiceInstanceDetails, details brokerapi.DeprovisionDetails) (operationId *string, err error) {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		ctx      context.Context
		instance models.ServiceInstanceDetails
		details  brokerapi.DeprovisionDetails
	}{ctx, instance, details})
	fake.recordInvocation("Deprovision", []interface{}{ctx, instance, details})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(ctx, instance, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deprovisionReturns.result1, fake.deprovisionReturns.result2
}

func (fake *FakeServiceBrokerHelper) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) DeprovisionArgsForCall(i int) (context.Context, models.ServiceInstanceDetails, brokerapi.DeprovisionDetails) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].ctx, fake.deprovisionArgsForCall[i].instance, fake.deprovisionArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) DeprovisionReturns(result1 *string, result2 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) DeprovisionReturnsOnCall(i int, result1 *string, result2 error) {
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 *string
			result2 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 *string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) PollInstance(ctx context.Context, instance models.ServiceInstanceDetails) (bool, error) {
	fake.pollInstanceMutex.Lock()
	ret, specificReturn := fake.pollInstanceReturnsOnCall[len(fake.pollInstanceArgsForCall)]
	fake.pollInstanceArgsForCall = append(fake.pollInstanceArgsForCall, struct {
		ctx      context.Context
		instance models.ServiceInstanceDetails
	}{ctx, instance})
	fake.recordInvocation("PollInstance", []interface{}{ctx, instance})
	fake.pollInstanceMutex.Unlock()
	if fake.PollInstanceStub != nil {
		return fake.PollInstanceStub(ctx, instance)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.pollInstanceReturns.result1, fake.pollInstanceReturns.result2
}

func (fake *FakeServiceBrokerHelper) PollInstanceCallCount() int {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return len(fake.pollInstanceArgsForCall)
}

func (fake *FakeServiceBrokerHelper) PollInstanceArgsForCall(i int) (context.Context, models.ServiceInstanceDetails) {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return fake.pollInstanceArgsForCall[i].ctx, fake.pollInstanceArgsForCall[i].instance
}

func (fake *FakeServiceBrokerHelper) PollInstanceReturns(result1 bool, result2 error) {
	fake.PollInstanceStub = nil
	fake.pollInstanceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) PollInstanceReturnsOnCall(i int, result1 bool, result2 error) {
	fake.PollInstanceStub = nil
	if fake.pollInstanceReturnsOnCall == nil {
		fake.pollInstanceReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.pollInstanceReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsync() bool {
	fake.provisionsAsyncMutex.Lock()
	ret, specificReturn := fake.provisionsAsyncReturnsOnCall[len(fake.provisionsAsyncArgsForCall)]
	fake.provisionsAsyncArgsForCall = append(fake.provisionsAsyncArgsForCall, struct{}{})
	fake.recordInvocation("ProvisionsAsync", []interface{}{})
	fake.provisionsAsyncMutex.Unlock()
	if fake.ProvisionsAsyncStub != nil {
		return fake.ProvisionsAsyncStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.provisionsAsyncReturns.result1
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncCallCount() int {
	fake.provisionsAsyncMutex.RLock()
	defer fake.provisionsAsyncMutex.RUnlock()
	return len(fake.provisionsAsyncArgsForCall)
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncReturns(result1 bool) {
	fake.ProvisionsAsyncStub = nil
	fake.provisionsAsyncReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncReturnsOnCall(i int, result1 bool) {
	fake.ProvisionsAsyncStub = nil
	if fake.provisionsAsyncReturnsOnCall == nil {
		fake.provisionsAsyncReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.provisionsAsyncReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsync() bool {
	fake.deprovisionsAsyncMutex.Lock()
	ret, specificReturn := fake.deprovisionsAsyncReturnsOnCall[len(fake.deprovisionsAsyncArgsForCall)]
	fake.deprovisionsAsyncArgsForCall = append(fake.deprovisionsAsyncArgsForCall, struct{}{})
	fake.recordInvocation("DeprovisionsAsync", []interface{}{})
	fake.deprovisionsAsyncMutex.Unlock()
	if fake.DeprovisionsAsyncStub != nil {
		return fake.DeprovisionsAsyncStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deprovisionsAsyncReturns.result1
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncCallCount() int {
	fake.deprovisionsAsyncMutex.RLock()
	defer fake.deprovisionsAsyncMutex.RUnlock()
	return len(fake.deprovisionsAsyncArgsForCall)
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncReturns(result1 bool) {
	fake.DeprovisionsAsyncStub = nil
	fake.deprovisionsAsyncReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncReturnsOnCall(i int, result1 bool) {
	fake.DeprovisionsAsyncStub = nil
	if fake.deprovisionsAsyncReturnsOnCall == nil {
		fake.deprovisionsAsyncReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.deprovisionsAsyncReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) UpdateInstanceDetails(ctx context.Context, instance *models.ServiceInstanceDetails) error {
	fake.updateInstanceDetailsMutex.Lock()
	ret, specificReturn := fake.updateInstanceDetailsReturnsOnCall[len(fake.updateInstanceDetailsArgsForCall)]
	fake.updateInstanceDetailsArgsForCall = append(fake.updateInstanceDetailsArgsForCall, struct {
		ctx      context.Context
		instance *models.ServiceInstanceDetails
	}{ctx, instance})
	fake.recordInvocation("UpdateInstanceDetails", []interface{}{ctx, instance})
	fake.updateInstanceDetailsMutex.Unlock()
	if fake.UpdateInstanceDetailsStub != nil {
		return fake.UpdateInstanceDetailsStub(ctx, instance)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateInstanceDetailsReturns.result1
}

func (fake *FakeServiceBrokerHelper) UpdateInstanceDetailsCallCount() int {
	fake.updateInstanceDetailsMutex.RLock()
	defer fake.updateInstanceDetailsMutex.RUnlock()
	return len(fake.updateInstanceDetailsArgsForCall)
}

func (fake *FakeServiceBrokerHelper) UpdateInstanceDetailsArgsForCall(i int) (context.Context, *models.ServiceInstanceDetails) {
	fake.updateInstanceDetailsMutex.RLock()
	defer fake.updateInstanceDetailsMutex.RUnlock()
	return fake.updateInstanceDetailsArgsForCall[i].ctx, fake.updateInstanceDetailsArgsForCall[i].instance
}

func (fake *FakeServiceBrokerHelper) UpdateInstanceDetailsReturns(result1 error) {
	fake.UpdateInstanceDetailsStub = nil
	fake.updateInstanceDetailsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) UpdateInstanceDetailsReturnsOnCall(i int, result1 error) {
	fake.UpdateInstanceDetailsStub = nil
	if fake.updateInstanceDetailsReturnsOnCall == nil {
		fake.updateInstanceDetailsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateInstanceDetailsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	fake.provisionsAsyncMutex.RLock()
	defer fake.provisionsAsyncMutex.RUnlock()
	fake.deprovisionsAsyncMutex.RLock()
	defer fake.deprovisionsAsyncMutex.RUnlock()
	fake.updateInstanceDetailsMutex.RLock()
	defer fake.updateInstanceDetailsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceBrokerHelper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ models.ServiceBrokerHelper = new(FakeServiceBrokerHelper)
