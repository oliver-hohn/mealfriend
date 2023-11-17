// package: mealfriend
// file: protos/mealfriend.proto

import * as protos_mealfriend_pb from "../protos/mealfriend_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MealfriendGetMealPlan = {
  readonly methodName: string;
  readonly service: typeof Mealfriend;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protos_mealfriend_pb.GetMealPlanRequest;
  readonly responseType: typeof protos_mealfriend_pb.GetMealPlanResponse;
};

type MealfriendScrape = {
  readonly methodName: string;
  readonly service: typeof Mealfriend;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof protos_mealfriend_pb.ScrapeRequest;
  readonly responseType: typeof protos_mealfriend_pb.ScrapeResponse;
};

export class Mealfriend {
  static readonly serviceName: string;
  static readonly GetMealPlan: MealfriendGetMealPlan;
  static readonly Scrape: MealfriendScrape;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class MealfriendClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getMealPlan(
    requestMessage: protos_mealfriend_pb.GetMealPlanRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protos_mealfriend_pb.GetMealPlanResponse|null) => void
  ): UnaryResponse;
  getMealPlan(
    requestMessage: protos_mealfriend_pb.GetMealPlanRequest,
    callback: (error: ServiceError|null, responseMessage: protos_mealfriend_pb.GetMealPlanResponse|null) => void
  ): UnaryResponse;
  scrape(
    requestMessage: protos_mealfriend_pb.ScrapeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: protos_mealfriend_pb.ScrapeResponse|null) => void
  ): UnaryResponse;
  scrape(
    requestMessage: protos_mealfriend_pb.ScrapeRequest,
    callback: (error: ServiceError|null, responseMessage: protos_mealfriend_pb.ScrapeResponse|null) => void
  ): UnaryResponse;
}

