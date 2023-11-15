/**
 * @fileoverview gRPC-Web generated client stub for mealfriend
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v4.25.0
// source: protos/mealfriend.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js')
const proto = {};
proto.mealfriend = require('./mealfriend_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.mealfriend.MealfriendClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.mealfriend.MealfriendPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.mealfriend.GetMealPlanRequest,
 *   !proto.mealfriend.GetMealPlanResponse>}
 */
const methodDescriptor_Mealfriend_GetMealPlan = new grpc.web.MethodDescriptor(
  '/mealfriend.Mealfriend/GetMealPlan',
  grpc.web.MethodType.UNARY,
  proto.mealfriend.GetMealPlanRequest,
  proto.mealfriend.GetMealPlanResponse,
  /**
   * @param {!proto.mealfriend.GetMealPlanRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.mealfriend.GetMealPlanResponse.deserializeBinary
);


/**
 * @param {!proto.mealfriend.GetMealPlanRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.mealfriend.GetMealPlanResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.mealfriend.GetMealPlanResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.mealfriend.MealfriendClient.prototype.getMealPlan =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/mealfriend.Mealfriend/GetMealPlan',
      request,
      metadata || {},
      methodDescriptor_Mealfriend_GetMealPlan,
      callback);
};


/**
 * @param {!proto.mealfriend.GetMealPlanRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.mealfriend.GetMealPlanResponse>}
 *     Promise that resolves to the response
 */
proto.mealfriend.MealfriendPromiseClient.prototype.getMealPlan =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/mealfriend.Mealfriend/GetMealPlan',
      request,
      metadata || {},
      methodDescriptor_Mealfriend_GetMealPlan);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.mealfriend.ScrapeRequest,
 *   !proto.mealfriend.ScrapeResponse>}
 */
const methodDescriptor_Mealfriend_Scrape = new grpc.web.MethodDescriptor(
  '/mealfriend.Mealfriend/Scrape',
  grpc.web.MethodType.UNARY,
  proto.mealfriend.ScrapeRequest,
  proto.mealfriend.ScrapeResponse,
  /**
   * @param {!proto.mealfriend.ScrapeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.mealfriend.ScrapeResponse.deserializeBinary
);


/**
 * @param {!proto.mealfriend.ScrapeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.mealfriend.ScrapeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.mealfriend.ScrapeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.mealfriend.MealfriendClient.prototype.scrape =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/mealfriend.Mealfriend/Scrape',
      request,
      metadata || {},
      methodDescriptor_Mealfriend_Scrape,
      callback);
};


/**
 * @param {!proto.mealfriend.ScrapeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.mealfriend.ScrapeResponse>}
 *     Promise that resolves to the response
 */
proto.mealfriend.MealfriendPromiseClient.prototype.scrape =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/mealfriend.Mealfriend/Scrape',
      request,
      metadata || {},
      methodDescriptor_Mealfriend_Scrape);
};


module.exports = proto.mealfriend;

