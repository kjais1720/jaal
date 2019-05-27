// Code generated by protoc-gen-graphql. DO NOT EDIT.

package pb

import (
	"context"

	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"go.appointy.com/jaal/gtypes"
	"go.appointy.com/jaal/schemabuilder"
)

func init() {
	RegisterClassType(gtypes.Schema)
	RegisterInputClass_PerInstance(gtypes.Schema)
	RegisterPayloadClass_PerInstance(gtypes.Schema)
	RegisterInputClass_Lumpsum(gtypes.Schema)
	RegisterPayloadClass_Lumpsum(gtypes.Schema)
	RegisterInputClass(gtypes.Schema)
	RegisterPayloadClass(gtypes.Schema)
	RegisterInputCreateClassReq(gtypes.Schema)
	RegisterPayloadCreateClassReq(gtypes.Schema)
	RegisterInputGetClassReq(gtypes.Schema)
	RegisterPayloadGetClassReq(gtypes.Schema)
	RegisterInputServiceProvider(gtypes.Schema)
	RegisterPayloadServiceProvider(gtypes.Schema)
	RegisterInputCreateClassInput(gtypes.Schema)
	RegisterPayloadCreateClassPayload(gtypes.Schema)
}

func RegisterClassType(schema *schemabuilder.Schema) {
	schema.Enum(ClassType(0), map[string]interface{}{
		"INVALID": ClassType(0),
		"REGULAR": ClassType(1),
		"SERIES":  ClassType(2),
	})
}

type UnionClassCharge struct {
	schemabuilder.Union
	*Class_PerInstance
	*Class_Lumpsum
}

func RegisterInputClass_PerInstance(schema *schemabuilder.Schema) {
	input := schema.InputObject("classPerInstance", Class_PerInstance{})
	input.FieldFunc("perInstance", func(target *Class_PerInstance, source *string) {
		target.PerInstance = *source
	})
}

func RegisterPayloadClass_PerInstance(schema *schemabuilder.Schema) {
	payload := schema.Object("Class_PerInstance", Class_PerInstance{})
	payload.FieldFunc("perInstance", func(ctx context.Context, in *Class_PerInstance) string {
		return in.PerInstance
	})
}

func RegisterInputClass_Lumpsum(schema *schemabuilder.Schema) {
	input := schema.InputObject("classLumpsum", Class_Lumpsum{})
	input.FieldFunc("lumpsum", func(target *Class_Lumpsum, source *int32) {
		target.Lumpsum = *source
	})
}

func RegisterPayloadClass_Lumpsum(schema *schemabuilder.Schema) {
	payload := schema.Object("Class_Lumpsum", Class_Lumpsum{})
	payload.FieldFunc("lumpsum", func(ctx context.Context, in *Class_Lumpsum) int32 {
		return in.Lumpsum
	})
}

func RegisterInputClass(schema *schemabuilder.Schema) {
	input := schema.InputObject("Class", Class{})
	input.FieldFunc("id", func(target *Class, source *schemabuilder.ID) {
		target.Id = source.Value
	})
	input.FieldFunc("area", func(target *Class, source *float32) {
		target.Area = *source
	})
	input.FieldFunc("area", func(target *Class, source *float32) {
		target.Area = *source
	})
	input.FieldFunc("strength", func(target *Class, source *int32) {
		target.Strength = *source
	})
	input.FieldFunc("type", func(target *Class, source *ClassType) {
		target.Type = *source
	})
	input.FieldFunc("instructors", func(target *Class, source []*ServiceProvider) {
		target.Instructors = source
	})
	input.FieldFunc("metadata", func(target *Class, source *map[string]string) {
		target.Metadata = *source
	})
	input.FieldFunc("parent", func(target *Class, source *string) {
		target.Parent = *source
	})
	input.FieldFunc("classLumpsum", func(target *Class, source *Class_Lumpsum) {
		target.Charge = source
	})
	input.FieldFunc("classPerInstance", func(target *Class, source *Class_PerInstance) {
		target.Charge = source
	})
	input.FieldFunc("duration", func(target *Class, source *duration.Duration) {
		target.Duration = source
	})
	input.FieldFunc("startDate", func(target *Class, source *timestamp.Timestamp) {
		target.StartDate = source
	})
}

func RegisterPayloadClass(schema *schemabuilder.Schema) {
	payload := schema.Object("Class", Class{})
	payload.FieldFunc("id", func(ctx context.Context, in *Class) schemabuilder.ID {
		return schemabuilder.ID{Value: in.Id}
	})
	payload.FieldFunc("area", func(ctx context.Context, in *Class) float32 {
		return in.Area
	})
	payload.FieldFunc("strength", func(ctx context.Context, in *Class) int32 {
		return in.Strength
	})
	payload.FieldFunc("isDeleted", func(ctx context.Context, in *Class) bool {
		return in.IsDeleted
	})
	payload.FieldFunc("type", func(ctx context.Context, in *Class) ClassType {
		return in.Type
	})
	payload.FieldFunc("instructors", func(ctx context.Context, in *Class) []*ServiceProvider {
		return in.Instructors
	})
	payload.FieldFunc("metadata", func(ctx context.Context, in *Class) map[string]string {
		return in.Metadata
	})
	payload.FieldFunc("parent", func(ctx context.Context, in *Class) string {
		return in.Parent
	})
	payload.FieldFunc("charge", func(ctx context.Context, in *Class) *UnionClassCharge {
		switch v := in.Charge.(type) {

		case *Class_PerInstance:
			return &UnionClassCharge{
				Class_PerInstance: v,
			}
		case *Class_Lumpsum:
			return &UnionClassCharge{
				Class_Lumpsum: v,
			}
		}
		return nil
	})
	payload.FieldFunc("startDate", func(ctx context.Context, in *Class) *timestamp.Timestamp {
		return in.StartDate
	})
	payload.FieldFunc("duration", func(ctx context.Context, in *Class) *duration.Duration {
		return in.Duration
	})
	payload.FieldFunc("empty", func(ctx context.Context, in *Class) *empty.Empty {
		return in.Empty
	})
}

func RegisterInputCreateClassReq(schema *schemabuilder.Schema) {
	input := schema.InputObject("CreateClassReq", CreateClassReq{})
	input.FieldFunc("parent", func(target *CreateClassReq, source *string) {
		target.Parent = *source
	})
	input.FieldFunc("class", func(target *CreateClassReq, source *Class) {
		target.Class = source
	})
}

func RegisterPayloadCreateClassReq(schema *schemabuilder.Schema) {
	payload := schema.Object("CreateClassReq", CreateClassReq{})
	payload.FieldFunc("parent", func(ctx context.Context, in *CreateClassReq) string {
		return in.Parent
	})
	payload.FieldFunc("class", func(ctx context.Context, in *CreateClassReq) *Class {
		return in.Class
	})
}

func RegisterInputGetClassReq(schema *schemabuilder.Schema) {
	input := schema.InputObject("GetClassInput", GetClassReq{})
	input.FieldFunc("id", func(target *GetClassReq, source *schemabuilder.ID) {
		target.Id = source.Value
	})
}

func RegisterPayloadGetClassReq(schema *schemabuilder.Schema) {
	payload := schema.Object("GetClassReq", GetClassReq{})
	payload.FieldFunc("id", func(ctx context.Context, in *GetClassReq) schemabuilder.ID {
		return schemabuilder.ID{Value: in.Id}
	})
}

func RegisterInputServiceProvider(schema *schemabuilder.Schema) {
	input := schema.InputObject("ServiceProvider", ServiceProvider{})
	input.FieldFunc("id", func(target *ServiceProvider, source *schemabuilder.ID) {
		target.Id = source.Value
	})
	input.FieldFunc("class", func(target *ServiceProvider, source *string) {
		target.FirstName = *source
	})
}

func RegisterPayloadServiceProvider(schema *schemabuilder.Schema) {
	payload := schema.Object("ServiceProvider", ServiceProvider{})
	payload.FieldFunc("parent", func(ctx context.Context, in *ServiceProvider) schemabuilder.ID {
		return schemabuilder.ID{Value: in.Id}
	})
	payload.FieldFunc("class", func(ctx context.Context, in *ServiceProvider) string {
		return in.FirstName
	})
}

type CreateClassInput struct {
	Parent           string
	Class            *Class
	ClientMutationId string
}

type CreateClassPayload struct {
	Payload          *Class
	ClientMutationId string
}

func RegisterInputCreateClassInput(schema *schemabuilder.Schema) {
	input := schema.InputObject("CreateClassInput", CreateClassInput{})
	input.FieldFunc("input", func(target *CreateClassInput, source *string) {
		target.Parent = *source
	})
	input.FieldFunc("class", func(target *CreateClassInput, source *Class) {
		target.Class = source
	})
	input.FieldFunc("clientMutationId", func(target *CreateClassInput, source *string) {
		target.ClientMutationId = *source
	})
}

func RegisterPayloadCreateClassPayload(schema *schemabuilder.Schema) {
	payload := schema.Object("CreateClassPayload", CreateClassPayload{})
	payload.FieldFunc("payload", func(ctx context.Context, in *CreateClassPayload) *Class {
		return in.Payload
	})
	payload.FieldFunc("clientMutationId", func(ctx context.Context, in *CreateClassPayload) string {
		return in.ClientMutationId
	})
}

func RegisterClassesOperations(schema *schemabuilder.Schema, client ClassesClient) {
	schema.Query().FieldFunc("class", func(ctx context.Context, args struct {
		In *GetClassReq
	}) (*Class, error) {
		return client.GetClass(ctx, args.In)
	})

	schema.Mutation().FieldFunc("createClass", func(ctx context.Context, args struct {
		Input *CreateClassInput
	}) (*CreateClassPayload, error) {
		request := &CreateClassReq{
			Parent: args.Input.Parent,
			Class:  args.Input.Class,
		}

		response, err := client.CreateClass(ctx, request)

		return &CreateClassPayload{
			Payload:          response,
			ClientMutationId: args.Input.ClientMutationId,
		}, err
	})
}
