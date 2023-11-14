// package: mealfriend
// file: protos/mealfriend.proto

var protos_mealfriend_pb = require("../protos/mealfriend_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Mealfriend = (function () {
  function Mealfriend() {}
  Mealfriend.serviceName = "mealfriend.Mealfriend";
  return Mealfriend;
}());

Mealfriend.GetMealPlan = {
  methodName: "GetMealPlan",
  service: Mealfriend,
  requestStream: false,
  responseStream: false,
  requestType: protos_mealfriend_pb.GetMealPlanRequest,
  responseType: protos_mealfriend_pb.GetMealPlanResponse
};

Mealfriend.Scrape = {
  methodName: "Scrape",
  service: Mealfriend,
  requestStream: false,
  responseStream: false,
  requestType: protos_mealfriend_pb.ScrapeRequest,
  responseType: protos_mealfriend_pb.ScrapeResponse
};

exports.Mealfriend = Mealfriend;

function MealfriendClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MealfriendClient.prototype.getMealPlan = function getMealPlan(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mealfriend.GetMealPlan, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealfriendClient.prototype.scrape = function scrape(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mealfriend.Scrape, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.MealfriendClient = MealfriendClient;

