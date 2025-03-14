// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

import (
	"encoding/json"
	v1api20181130s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130storage"
	v20181130s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130storage"
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

func Test_UserAssignedIdentity_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity to hub returns original",
		prop.ForAll(RunResourceConversionTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForUserAssignedIdentity tests if a specific instance of UserAssignedIdentity round trips to the hub storage version and back losslessly
func RunResourceConversionTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20181130s.UserAssignedIdentity
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual UserAssignedIdentity
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

func Test_UserAssignedIdentity_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity to UserAssignedIdentity via AssignProperties_To_UserAssignedIdentity & AssignProperties_From_UserAssignedIdentity returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity tests if a specific instance of UserAssignedIdentity can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity
	err := copied.AssignProperties_To_UserAssignedIdentity(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity
	err = actual.AssignProperties_From_UserAssignedIdentity(&other)
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

func Test_UserAssignedIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity runs a test to see if a specific instance of UserAssignedIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity
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

// Generator of UserAssignedIdentity instances for property testing - lazily instantiated by
// UserAssignedIdentityGenerator()
var userAssignedIdentityGenerator gopter.Gen

// UserAssignedIdentityGenerator returns a generator of UserAssignedIdentity instances for property testing.
func UserAssignedIdentityGenerator() gopter.Gen {
	if userAssignedIdentityGenerator != nil {
		return userAssignedIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity(generators)
	userAssignedIdentityGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity{}), generators)

	return userAssignedIdentityGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentity_SpecGenerator()
	gens["Status"] = UserAssignedIdentity_STATUSGenerator()
}

func Test_UserAssignedIdentity_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity_Spec to UserAssignedIdentity_Spec via AssignProperties_To_UserAssignedIdentity_Spec & AssignProperties_From_UserAssignedIdentity_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity_Spec tests if a specific instance of UserAssignedIdentity_Spec can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity_Spec
	err := copied.AssignProperties_To_UserAssignedIdentity_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity_Spec
	err = actual.AssignProperties_From_UserAssignedIdentity_Spec(&other)
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

func Test_UserAssignedIdentity_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_Spec runs a test to see if a specific instance of UserAssignedIdentity_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_Spec
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

// Generator of UserAssignedIdentity_Spec instances for property testing - lazily instantiated by
// UserAssignedIdentity_SpecGenerator()
var userAssignedIdentity_SpecGenerator gopter.Gen

// UserAssignedIdentity_SpecGenerator returns a generator of UserAssignedIdentity_Spec instances for property testing.
// We first initialize userAssignedIdentity_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UserAssignedIdentity_SpecGenerator() gopter.Gen {
	if userAssignedIdentity_SpecGenerator != nil {
		return userAssignedIdentity_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	return userAssignedIdentity_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(UserAssignedIdentityOperatorSpecGenerator())
}

func Test_UserAssignedIdentity_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentity_STATUS to UserAssignedIdentity_STATUS via AssignProperties_To_UserAssignedIdentity_STATUS & AssignProperties_From_UserAssignedIdentity_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentity_STATUS tests if a specific instance of UserAssignedIdentity_STATUS can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentity_STATUS
	err := copied.AssignProperties_To_UserAssignedIdentity_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentity_STATUS
	err = actual.AssignProperties_From_UserAssignedIdentity_STATUS(&other)
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

func Test_UserAssignedIdentity_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_STATUS, UserAssignedIdentity_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_STATUS runs a test to see if a specific instance of UserAssignedIdentity_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_STATUS(subject UserAssignedIdentity_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_STATUS
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

// Generator of UserAssignedIdentity_STATUS instances for property testing - lazily instantiated by
// UserAssignedIdentity_STATUSGenerator()
var userAssignedIdentity_STATUSGenerator gopter.Gen

// UserAssignedIdentity_STATUSGenerator returns a generator of UserAssignedIdentity_STATUS instances for property testing.
func UserAssignedIdentity_STATUSGenerator() gopter.Gen {
	if userAssignedIdentity_STATUSGenerator != nil {
		return userAssignedIdentity_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(generators)
	userAssignedIdentity_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_STATUS{}), generators)

	return userAssignedIdentity_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_STATUS(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentityOperatorSpec to UserAssignedIdentityOperatorSpec via AssignProperties_To_UserAssignedIdentityOperatorSpec & AssignProperties_From_UserAssignedIdentityOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentityOperatorSpec, UserAssignedIdentityOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentityOperatorSpec tests if a specific instance of UserAssignedIdentityOperatorSpec can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentityOperatorSpec(subject UserAssignedIdentityOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentityOperatorSpec
	err := copied.AssignProperties_To_UserAssignedIdentityOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentityOperatorSpec
	err = actual.AssignProperties_From_UserAssignedIdentityOperatorSpec(&other)
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

func Test_UserAssignedIdentityOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityOperatorSpec, UserAssignedIdentityOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityOperatorSpec runs a test to see if a specific instance of UserAssignedIdentityOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityOperatorSpec(subject UserAssignedIdentityOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityOperatorSpec
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

// Generator of UserAssignedIdentityOperatorSpec instances for property testing - lazily instantiated by
// UserAssignedIdentityOperatorSpecGenerator()
var userAssignedIdentityOperatorSpecGenerator gopter.Gen

// UserAssignedIdentityOperatorSpecGenerator returns a generator of UserAssignedIdentityOperatorSpec instances for property testing.
func UserAssignedIdentityOperatorSpecGenerator() gopter.Gen {
	if userAssignedIdentityOperatorSpecGenerator != nil {
		return userAssignedIdentityOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec(generators)
	userAssignedIdentityOperatorSpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityOperatorSpec{}), generators)

	return userAssignedIdentityOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentityOperatorSpec(gens map[string]gopter.Gen) {
	gens["ConfigMaps"] = gen.PtrOf(UserAssignedIdentityOperatorConfigMapsGenerator())
}

func Test_UserAssignedIdentityOperatorConfigMaps_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentityOperatorConfigMaps to UserAssignedIdentityOperatorConfigMaps via AssignProperties_To_UserAssignedIdentityOperatorConfigMaps & AssignProperties_From_UserAssignedIdentityOperatorConfigMaps returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentityOperatorConfigMaps, UserAssignedIdentityOperatorConfigMapsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentityOperatorConfigMaps tests if a specific instance of UserAssignedIdentityOperatorConfigMaps can be assigned to v1beta20181130storage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentityOperatorConfigMaps(subject UserAssignedIdentityOperatorConfigMaps) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20181130s.UserAssignedIdentityOperatorConfigMaps
	err := copied.AssignProperties_To_UserAssignedIdentityOperatorConfigMaps(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentityOperatorConfigMaps
	err = actual.AssignProperties_From_UserAssignedIdentityOperatorConfigMaps(&other)
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

func Test_UserAssignedIdentityOperatorConfigMaps_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityOperatorConfigMaps via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps, UserAssignedIdentityOperatorConfigMapsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps runs a test to see if a specific instance of UserAssignedIdentityOperatorConfigMaps round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityOperatorConfigMaps(subject UserAssignedIdentityOperatorConfigMaps) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityOperatorConfigMaps
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

// Generator of UserAssignedIdentityOperatorConfigMaps instances for property testing - lazily instantiated by
// UserAssignedIdentityOperatorConfigMapsGenerator()
var userAssignedIdentityOperatorConfigMapsGenerator gopter.Gen

// UserAssignedIdentityOperatorConfigMapsGenerator returns a generator of UserAssignedIdentityOperatorConfigMaps instances for property testing.
func UserAssignedIdentityOperatorConfigMapsGenerator() gopter.Gen {
	if userAssignedIdentityOperatorConfigMapsGenerator != nil {
		return userAssignedIdentityOperatorConfigMapsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityOperatorConfigMapsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityOperatorConfigMaps{}), generators)

	return userAssignedIdentityOperatorConfigMapsGenerator
}
