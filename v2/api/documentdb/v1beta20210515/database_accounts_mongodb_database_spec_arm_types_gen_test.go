// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"encoding/json"
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

func Test_DatabaseAccounts_MongodbDatabase_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabase_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec_ARM, DatabaseAccounts_MongodbDatabase_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec_ARM runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabase_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec_ARM(subject DatabaseAccounts_MongodbDatabase_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabase_Spec_ARM
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

// Generator of DatabaseAccounts_MongodbDatabase_Spec_ARM instances for property testing - lazily instantiated by
// DatabaseAccounts_MongodbDatabase_Spec_ARMGenerator()
var databaseAccounts_MongodbDatabase_Spec_ARMGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabase_Spec_ARMGenerator returns a generator of DatabaseAccounts_MongodbDatabase_Spec_ARM instances for property testing.
// We first initialize databaseAccounts_MongodbDatabase_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabase_Spec_ARMGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabase_Spec_ARMGenerator != nil {
		return databaseAccounts_MongodbDatabase_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM(generators)
	databaseAccounts_MongodbDatabase_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM(generators)
	databaseAccounts_MongodbDatabase_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_Spec_ARM{}), generators)

	return databaseAccounts_MongodbDatabase_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseCreateUpdateProperties_ARMGenerator())
}

func Test_MongoDBDatabaseCreateUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseCreateUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties_ARM, MongoDBDatabaseCreateUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties_ARM runs a test to see if a specific instance of MongoDBDatabaseCreateUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseCreateUpdateProperties_ARM(subject MongoDBDatabaseCreateUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseCreateUpdateProperties_ARM
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

// Generator of MongoDBDatabaseCreateUpdateProperties_ARM instances for property testing - lazily instantiated by
// MongoDBDatabaseCreateUpdateProperties_ARMGenerator()
var mongoDBDatabaseCreateUpdateProperties_ARMGenerator gopter.Gen

// MongoDBDatabaseCreateUpdateProperties_ARMGenerator returns a generator of MongoDBDatabaseCreateUpdateProperties_ARM instances for property testing.
func MongoDBDatabaseCreateUpdateProperties_ARMGenerator() gopter.Gen {
	if mongoDBDatabaseCreateUpdateProperties_ARMGenerator != nil {
		return mongoDBDatabaseCreateUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties_ARM(generators)
	mongoDBDatabaseCreateUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseCreateUpdateProperties_ARM{}), generators)

	return mongoDBDatabaseCreateUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseCreateUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_ARMGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResource_ARMGenerator())
}

func Test_CreateUpdateOptions_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions_ARM, CreateUpdateOptions_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions_ARM runs a test to see if a specific instance of CreateUpdateOptions_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions_ARM(subject CreateUpdateOptions_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions_ARM
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

// Generator of CreateUpdateOptions_ARM instances for property testing - lazily instantiated by
// CreateUpdateOptions_ARMGenerator()
var createUpdateOptions_ARMGenerator gopter.Gen

// CreateUpdateOptions_ARMGenerator returns a generator of CreateUpdateOptions_ARM instances for property testing.
// We first initialize createUpdateOptions_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptions_ARMGenerator() gopter.Gen {
	if createUpdateOptions_ARMGenerator != nil {
		return createUpdateOptions_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_ARM(generators)
	createUpdateOptions_ARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_ARM(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions_ARM(generators)
	createUpdateOptions_ARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_ARM{}), generators)

	return createUpdateOptions_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions_ARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions_ARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_ARMGenerator())
}

func Test_MongoDBDatabaseResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResource_ARM, MongoDBDatabaseResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResource_ARM runs a test to see if a specific instance of MongoDBDatabaseResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResource_ARM(subject MongoDBDatabaseResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResource_ARM
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

// Generator of MongoDBDatabaseResource_ARM instances for property testing - lazily instantiated by
// MongoDBDatabaseResource_ARMGenerator()
var mongoDBDatabaseResource_ARMGenerator gopter.Gen

// MongoDBDatabaseResource_ARMGenerator returns a generator of MongoDBDatabaseResource_ARM instances for property testing.
func MongoDBDatabaseResource_ARMGenerator() gopter.Gen {
	if mongoDBDatabaseResource_ARMGenerator != nil {
		return mongoDBDatabaseResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_ARM(generators)
	mongoDBDatabaseResource_ARMGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource_ARM{}), generators)

	return mongoDBDatabaseResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_AutoscaleSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_ARM, AutoscaleSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_ARM runs a test to see if a specific instance of AutoscaleSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_ARM(subject AutoscaleSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_ARM
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

// Generator of AutoscaleSettings_ARM instances for property testing - lazily instantiated by
// AutoscaleSettings_ARMGenerator()
var autoscaleSettings_ARMGenerator gopter.Gen

// AutoscaleSettings_ARMGenerator returns a generator of AutoscaleSettings_ARM instances for property testing.
func AutoscaleSettings_ARMGenerator() gopter.Gen {
	if autoscaleSettings_ARMGenerator != nil {
		return autoscaleSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_ARM(generators)
	autoscaleSettings_ARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_ARM{}), generators)

	return autoscaleSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_ARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}
