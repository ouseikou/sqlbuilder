package com.holder.sqlbuilder.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.65.1)",
    comments = "Source: proto/api.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SqlBuilderApiGrpc {

  private SqlBuilderApiGrpc() {}

  public static final java.lang.String SERVICE_NAME = "proto.SqlBuilderApi";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest,
      com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> getGenerateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Generate",
      requestType = com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest.class,
      responseType = com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest,
      com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> getGenerateMethod() {
    io.grpc.MethodDescriptor<com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest, com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> getGenerateMethod;
    if ((getGenerateMethod = SqlBuilderApiGrpc.getGenerateMethod) == null) {
      synchronized (SqlBuilderApiGrpc.class) {
        if ((getGenerateMethod = SqlBuilderApiGrpc.getGenerateMethod) == null) {
          SqlBuilderApiGrpc.getGenerateMethod = getGenerateMethod =
              io.grpc.MethodDescriptor.<com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest, com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Generate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response.getDefaultInstance()))
              .setSchemaDescriptor(new SqlBuilderApiMethodDescriptorSupplier("Generate"))
              .build();
        }
      }
    }
    return getGenerateMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SqlBuilderApiStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiStub>() {
        @java.lang.Override
        public SqlBuilderApiStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SqlBuilderApiStub(channel, callOptions);
        }
      };
    return SqlBuilderApiStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SqlBuilderApiBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiBlockingStub>() {
        @java.lang.Override
        public SqlBuilderApiBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SqlBuilderApiBlockingStub(channel, callOptions);
        }
      };
    return SqlBuilderApiBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SqlBuilderApiFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SqlBuilderApiFutureStub>() {
        @java.lang.Override
        public SqlBuilderApiFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SqlBuilderApiFutureStub(channel, callOptions);
        }
      };
    return SqlBuilderApiFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void generate(com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest request,
        io.grpc.stub.StreamObserver<com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGenerateMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service SqlBuilderApi.
   */
  public static abstract class SqlBuilderApiImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return SqlBuilderApiGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service SqlBuilderApi.
   */
  public static final class SqlBuilderApiStub
      extends io.grpc.stub.AbstractAsyncStub<SqlBuilderApiStub> {
    private SqlBuilderApiStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SqlBuilderApiStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SqlBuilderApiStub(channel, callOptions);
    }

    /**
     */
    public void generate(com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest request,
        io.grpc.stub.StreamObserver<com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGenerateMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service SqlBuilderApi.
   */
  public static final class SqlBuilderApiBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<SqlBuilderApiBlockingStub> {
    private SqlBuilderApiBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SqlBuilderApiBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SqlBuilderApiBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response generate(com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGenerateMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service SqlBuilderApi.
   */
  public static final class SqlBuilderApiFutureStub
      extends io.grpc.stub.AbstractFutureStub<SqlBuilderApiFutureStub> {
    private SqlBuilderApiFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SqlBuilderApiFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SqlBuilderApiFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response> generate(
        com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGenerateMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GENERATE = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GENERATE:
          serviceImpl.generate((com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest) request,
              (io.grpc.stub.StreamObserver<com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getGenerateMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.holder.sqlbuilder.proto.SqlBuilderApiProto.BuilderRequest,
              com.holder.sqlbuilder.proto.SqlBuilderApiProto.Response>(
                service, METHODID_GENERATE)))
        .build();
  }

  private static abstract class SqlBuilderApiBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SqlBuilderApiBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.holder.sqlbuilder.proto.SqlBuilderApiProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SqlBuilderApi");
    }
  }

  private static final class SqlBuilderApiFileDescriptorSupplier
      extends SqlBuilderApiBaseDescriptorSupplier {
    SqlBuilderApiFileDescriptorSupplier() {}
  }

  private static final class SqlBuilderApiMethodDescriptorSupplier
      extends SqlBuilderApiBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    SqlBuilderApiMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SqlBuilderApiGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SqlBuilderApiFileDescriptorSupplier())
              .addMethod(getGenerateMethod())
              .build();
        }
      }
    }
    return result;
  }
}
