// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"encoding/json"
	v1api20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_StorageAccountsQueueService_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsQueueService tests if a specific instance of StorageAccountsQueueService round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20220901s.StorageAccountsQueueService
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsQueueService
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueService_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueService to StorageAccountsQueueService via AssignProperties_To_StorageAccountsQueueService & AssignProperties_From_StorageAccountsQueueService returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueService tests if a specific instance of StorageAccountsQueueService can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220901s.StorageAccountsQueueService
	err := copied.AssignProperties_To_StorageAccountsQueueService(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueService
	err = actual.AssignProperties_From_StorageAccountsQueueService(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueService, StorageAccountsQueueServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueService runs a test to see if a specific instance of StorageAccountsQueueService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueService(subject StorageAccountsQueueService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueService
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccountsQueueService instances for property testing - lazily instantiated by
// StorageAccountsQueueServiceGenerator()
var storageAccountsQueueServiceGenerator gopter.Gen

// StorageAccountsQueueServiceGenerator returns a generator of StorageAccountsQueueService instances for property testing.
func StorageAccountsQueueServiceGenerator() gopter.Gen {
	if storageAccountsQueueServiceGenerator != nil {
		return storageAccountsQueueServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueService(generators)
	storageAccountsQueueServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueService{}), generators)

	return storageAccountsQueueServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccounts_QueueService_SpecGenerator()
	gens["Status"] = StorageAccounts_QueueService_STATUSGenerator()
}

func Test_StorageAccounts_QueueService_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_QueueService_Spec to StorageAccounts_QueueService_Spec via AssignProperties_To_StorageAccounts_QueueService_Spec & AssignProperties_From_StorageAccounts_QueueService_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_QueueService_Spec, StorageAccounts_QueueService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_QueueService_Spec tests if a specific instance of StorageAccounts_QueueService_Spec can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_QueueService_Spec(subject StorageAccounts_QueueService_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220901s.StorageAccounts_QueueService_Spec
	err := copied.AssignProperties_To_StorageAccounts_QueueService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_QueueService_Spec
	err = actual.AssignProperties_From_StorageAccounts_QueueService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccounts_QueueService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_QueueService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_QueueService_Spec, StorageAccounts_QueueService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_QueueService_Spec runs a test to see if a specific instance of StorageAccounts_QueueService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_QueueService_Spec(subject StorageAccounts_QueueService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_QueueService_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccounts_QueueService_Spec instances for property testing - lazily instantiated by
// StorageAccounts_QueueService_SpecGenerator()
var storageAccounts_QueueService_SpecGenerator gopter.Gen

// StorageAccounts_QueueService_SpecGenerator returns a generator of StorageAccounts_QueueService_Spec instances for property testing.
func StorageAccounts_QueueService_SpecGenerator() gopter.Gen {
	if storageAccounts_QueueService_SpecGenerator != nil {
		return storageAccounts_QueueService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec(generators)
	storageAccounts_QueueService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueService_Spec{}), generators)

	return storageAccounts_QueueService_SpecGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_Spec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
}

func Test_StorageAccounts_QueueService_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_QueueService_STATUS to StorageAccounts_QueueService_STATUS via AssignProperties_To_StorageAccounts_QueueService_STATUS & AssignProperties_From_StorageAccounts_QueueService_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_QueueService_STATUS, StorageAccounts_QueueService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_QueueService_STATUS tests if a specific instance of StorageAccounts_QueueService_STATUS can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_QueueService_STATUS(subject StorageAccounts_QueueService_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220901s.StorageAccounts_QueueService_STATUS
	err := copied.AssignProperties_To_StorageAccounts_QueueService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_QueueService_STATUS
	err = actual.AssignProperties_From_StorageAccounts_QueueService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccounts_QueueService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_QueueService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_QueueService_STATUS, StorageAccounts_QueueService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_QueueService_STATUS runs a test to see if a specific instance of StorageAccounts_QueueService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_QueueService_STATUS(subject StorageAccounts_QueueService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_QueueService_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccounts_QueueService_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_QueueService_STATUSGenerator()
var storageAccounts_QueueService_STATUSGenerator gopter.Gen

// StorageAccounts_QueueService_STATUSGenerator returns a generator of StorageAccounts_QueueService_STATUS instances for property testing.
// We first initialize storageAccounts_QueueService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_QueueService_STATUSGenerator() gopter.Gen {
	if storageAccounts_QueueService_STATUSGenerator != nil {
		return storageAccounts_QueueService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_STATUS(generators)
	storageAccounts_QueueService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_STATUS(generators)
	storageAccounts_QueueService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_QueueService_STATUS{}), generators)

	return storageAccounts_QueueService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_QueueService_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_QueueService_STATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
}
