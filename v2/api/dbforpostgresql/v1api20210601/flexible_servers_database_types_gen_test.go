// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import (
	"encoding/json"
	v1api20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601storage"
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

func Test_FlexibleServersDatabase_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersDatabase to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersDatabase tests if a specific instance of FlexibleServersDatabase round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20210601s.FlexibleServersDatabase
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersDatabase
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

func Test_FlexibleServersDatabase_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersDatabase to FlexibleServersDatabase via AssignProperties_To_FlexibleServersDatabase & AssignProperties_From_FlexibleServersDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersDatabase tests if a specific instance of FlexibleServersDatabase can be assigned to v1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20210601s.FlexibleServersDatabase
	err := copied.AssignProperties_To_FlexibleServersDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersDatabase
	err = actual.AssignProperties_From_FlexibleServersDatabase(&other)
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

func Test_FlexibleServersDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase runs a test to see if a specific instance of FlexibleServersDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase
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

// Generator of FlexibleServersDatabase instances for property testing - lazily instantiated by
// FlexibleServersDatabaseGenerator()
var flexibleServersDatabaseGenerator gopter.Gen

// FlexibleServersDatabaseGenerator returns a generator of FlexibleServersDatabase instances for property testing.
func FlexibleServersDatabaseGenerator() gopter.Gen {
	if flexibleServersDatabaseGenerator != nil {
		return flexibleServersDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase(generators)
	flexibleServersDatabaseGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase{}), generators)

	return flexibleServersDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServers_Database_SpecGenerator()
	gens["Status"] = FlexibleServers_Database_STATUSGenerator()
}

func Test_FlexibleServers_Database_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Database_Spec to FlexibleServers_Database_Spec via AssignProperties_To_FlexibleServers_Database_Spec & AssignProperties_From_FlexibleServers_Database_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Database_Spec, FlexibleServers_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Database_Spec tests if a specific instance of FlexibleServers_Database_Spec can be assigned to v1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Database_Spec(subject FlexibleServers_Database_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20210601s.FlexibleServers_Database_Spec
	err := copied.AssignProperties_To_FlexibleServers_Database_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Database_Spec
	err = actual.AssignProperties_From_FlexibleServers_Database_Spec(&other)
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

func Test_FlexibleServers_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Database_Spec, FlexibleServers_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Database_Spec runs a test to see if a specific instance of FlexibleServers_Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Database_Spec(subject FlexibleServers_Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Database_Spec
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

// Generator of FlexibleServers_Database_Spec instances for property testing - lazily instantiated by
// FlexibleServers_Database_SpecGenerator()
var flexibleServers_Database_SpecGenerator gopter.Gen

// FlexibleServers_Database_SpecGenerator returns a generator of FlexibleServers_Database_Spec instances for property testing.
func FlexibleServers_Database_SpecGenerator() gopter.Gen {
	if flexibleServers_Database_SpecGenerator != nil {
		return flexibleServers_Database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Database_Spec(generators)
	flexibleServers_Database_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Database_Spec{}), generators)

	return flexibleServers_Database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Database_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Database_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServers_Database_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServers_Database_STATUS to FlexibleServers_Database_STATUS via AssignProperties_To_FlexibleServers_Database_STATUS & AssignProperties_From_FlexibleServers_Database_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServers_Database_STATUS, FlexibleServers_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServers_Database_STATUS tests if a specific instance of FlexibleServers_Database_STATUS can be assigned to v1api20210601storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServers_Database_STATUS(subject FlexibleServers_Database_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20210601s.FlexibleServers_Database_STATUS
	err := copied.AssignProperties_To_FlexibleServers_Database_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServers_Database_STATUS
	err = actual.AssignProperties_From_FlexibleServers_Database_STATUS(&other)
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

func Test_FlexibleServers_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServers_Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServers_Database_STATUS, FlexibleServers_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServers_Database_STATUS runs a test to see if a specific instance of FlexibleServers_Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServers_Database_STATUS(subject FlexibleServers_Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServers_Database_STATUS
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

// Generator of FlexibleServers_Database_STATUS instances for property testing - lazily instantiated by
// FlexibleServers_Database_STATUSGenerator()
var flexibleServers_Database_STATUSGenerator gopter.Gen

// FlexibleServers_Database_STATUSGenerator returns a generator of FlexibleServers_Database_STATUS instances for property testing.
// We first initialize flexibleServers_Database_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServers_Database_STATUSGenerator() gopter.Gen {
	if flexibleServers_Database_STATUSGenerator != nil {
		return flexibleServers_Database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS(generators)
	flexibleServers_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Database_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS(generators)
	flexibleServers_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServers_Database_STATUS{}), generators)

	return flexibleServers_Database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServers_Database_STATUS(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServers_Database_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
