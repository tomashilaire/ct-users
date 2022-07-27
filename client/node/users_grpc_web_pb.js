/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pb = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.AuthenticationClient =
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
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.AuthenticationPromiseClient =
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
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.SignUpRequest,
 *   !proto.pb.SignUpResponse>}
 */
const methodDescriptor_Authentication_SignUp = new grpc.web.MethodDescriptor(
  '/pb.Authentication/SignUp',
  grpc.web.MethodType.UNARY,
  proto.pb.SignUpRequest,
  proto.pb.SignUpResponse,
  /**
   * @param {!proto.pb.SignUpRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SignUpResponse.deserializeBinary
);


/**
 * @param {!proto.pb.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.SignUpResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.SignUpResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.AuthenticationClient.prototype.signUp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.Authentication/SignUp',
      request,
      metadata || {},
      methodDescriptor_Authentication_SignUp,
      callback);
};


/**
 * @param {!proto.pb.SignUpRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.SignUpResponse>}
 *     Promise that resolves to the response
 */
proto.pb.AuthenticationPromiseClient.prototype.signUp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.Authentication/SignUp',
      request,
      metadata || {},
      methodDescriptor_Authentication_SignUp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.SignInRequest,
 *   !proto.pb.SignInResponse>}
 */
const methodDescriptor_Authentication_SignIn = new grpc.web.MethodDescriptor(
  '/pb.Authentication/SignIn',
  grpc.web.MethodType.UNARY,
  proto.pb.SignInRequest,
  proto.pb.SignInResponse,
  /**
   * @param {!proto.pb.SignInRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SignInResponse.deserializeBinary
);


/**
 * @param {!proto.pb.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.SignInResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.SignInResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.AuthenticationClient.prototype.signIn =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.Authentication/SignIn',
      request,
      metadata || {},
      methodDescriptor_Authentication_SignIn,
      callback);
};


/**
 * @param {!proto.pb.SignInRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.SignInResponse>}
 *     Promise that resolves to the response
 */
proto.pb.AuthenticationPromiseClient.prototype.signIn =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.Authentication/SignIn',
      request,
      metadata || {},
      methodDescriptor_Authentication_SignIn);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.AuthenticateRequest,
 *   !proto.pb.AuthenticateResponse>}
 */
const methodDescriptor_Authentication_Authenticate = new grpc.web.MethodDescriptor(
  '/pb.Authentication/Authenticate',
  grpc.web.MethodType.UNARY,
  proto.pb.AuthenticateRequest,
  proto.pb.AuthenticateResponse,
  /**
   * @param {!proto.pb.AuthenticateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.AuthenticateResponse.deserializeBinary
);


/**
 * @param {!proto.pb.AuthenticateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.AuthenticateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.AuthenticateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.AuthenticationClient.prototype.authenticate =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.Authentication/Authenticate',
      request,
      metadata || {},
      methodDescriptor_Authentication_Authenticate,
      callback);
};


/**
 * @param {!proto.pb.AuthenticateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.AuthenticateResponse>}
 *     Promise that resolves to the response
 */
proto.pb.AuthenticationPromiseClient.prototype.authenticate =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.Authentication/Authenticate',
      request,
      metadata || {},
      methodDescriptor_Authentication_Authenticate);
};


module.exports = proto.pb;

