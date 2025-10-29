import { RpcClient } from "rpc-client-gen";

type RpcClientProps = ConstructorParameters<typeof RpcClient>[0];
const defaultParams: RpcClientProps = {
  baseUrl: "http://localhost:4000",
  headers: {
    "Content-Type": "application/json",
  },
};
export const createRpcClient = (params: Partial<RpcClientProps> = {}) =>
  new RpcClient({ ...defaultParams, ...params });

export const client = createRpcClient(defaultParams);
