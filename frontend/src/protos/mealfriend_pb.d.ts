// package: mealfriend
// file: protos/mealfriend.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";

export class GetMealPlanRequest extends jspb.Message {
  getRequirementsMap(): jspb.Map<string, number>;
  clearRequirementsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMealPlanRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMealPlanRequest): GetMealPlanRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMealPlanRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMealPlanRequest;
  static deserializeBinaryFromReader(message: GetMealPlanRequest, reader: jspb.BinaryReader): GetMealPlanRequest;
}

export namespace GetMealPlanRequest {
  export type AsObject = {
    requirementsMap: Array<[string, number]>,
  }
}

export class GetMealPlanResponse extends jspb.Message {
  clearRecipesList(): void;
  getRecipesList(): Array<Recipe>;
  setRecipesList(value: Array<Recipe>): void;
  addRecipes(value?: Recipe, index?: number): Recipe;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMealPlanResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMealPlanResponse): GetMealPlanResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMealPlanResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMealPlanResponse;
  static deserializeBinaryFromReader(message: GetMealPlanResponse, reader: jspb.BinaryReader): GetMealPlanResponse;
}

export namespace GetMealPlanResponse {
  export type AsObject = {
    recipesList: Array<Recipe.AsObject>,
  }
}

export class ScrapeRequest extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ScrapeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ScrapeRequest): ScrapeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ScrapeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ScrapeRequest;
  static deserializeBinaryFromReader(message: ScrapeRequest, reader: jspb.BinaryReader): ScrapeRequest;
}

export namespace ScrapeRequest {
  export type AsObject = {
    url: string,
  }
}

export class ScrapeResponse extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): Recipe | undefined;
  setRecipe(value?: Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ScrapeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ScrapeResponse): ScrapeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ScrapeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ScrapeResponse;
  static deserializeBinaryFromReader(message: ScrapeResponse, reader: jspb.BinaryReader): ScrapeResponse;
}

export namespace ScrapeResponse {
  export type AsObject = {
    recipe?: Recipe.AsObject,
  }
}

export class Recipe extends jspb.Message {
  getCode(): string;
  setCode(value: string): void;

  getName(): string;
  setName(value: string): void;

  getSource(): string;
  setSource(value: string): void;

  clearIngredientsList(): void;
  getIngredientsList(): Array<string>;
  setIngredientsList(value: Array<string>): void;
  addIngredients(value: string, index?: number): string;

  clearTagsList(): void;
  getTagsList(): Array<TagMap[keyof TagMap]>;
  setTagsList(value: Array<TagMap[keyof TagMap]>): void;
  addTags(value: TagMap[keyof TagMap], index?: number): TagMap[keyof TagMap];

  hasCookTime(): boolean;
  clearCookTime(): void;
  getCookTime(): google_protobuf_duration_pb.Duration | undefined;
  setCookTime(value?: google_protobuf_duration_pb.Duration): void;

  getImageUrl(): string;
  setImageUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Recipe.AsObject;
  static toObject(includeInstance: boolean, msg: Recipe): Recipe.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Recipe, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Recipe;
  static deserializeBinaryFromReader(message: Recipe, reader: jspb.BinaryReader): Recipe;
}

export namespace Recipe {
  export type AsObject = {
    code: string,
    name: string,
    source: string,
    ingredientsList: Array<string>,
    tagsList: Array<TagMap[keyof TagMap]>,
    cookTime?: google_protobuf_duration_pb.Duration.AsObject,
    imageUrl: string,
  }
}

export interface TagMap {
  TAG_UNSPECIFIED: 0;
  TAG_BEEF: 1;
  TAG_DAIRY: 2;
  TAG_EGG: 3;
  TAG_FISH: 4;
  TAG_FRUIT: 5;
  TAG_GRAIN: 6;
  TAG_LEGUMES: 7;
  TAG_PASTA: 8;
  TAG_PORK: 9;
  TAG_POULTRY: 10;
  TAG_SHELLFISH: 11;
  TAG_VEGETABLE: 12;
}

export const Tag: TagMap;

