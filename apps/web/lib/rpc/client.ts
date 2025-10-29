import { RpcClient } from "rpc-client-gen";

type RpcClientProps = ConstructorParameters<typeof RpcClient>[0];
const defaultParams: RpcClientProps = {
  baseUrl: "http://localhost:4000",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  options: {
    onRequest: (event) => {
      const cloned = new Request(event.request);
      if (cloned.method === "POST") return;
      const url = new URL(cloned.url);
      Array.from(url.searchParams.entries())
        .filter(([, value]) => value === "null" || value === "undefined")
        .forEach(([key]) => url.searchParams.delete(key));
      const newRequest = new Request(url, {
        method: cloned.method,
        headers: cloned.headers,
      });
      event.request = new Request(newRequest, {
        headers: {
          Authorization: "Bearer 123",
        },
      });
    },
  },
};
export const createRpcClient = (params: Partial<RpcClientProps> = {}) =>
  new RpcClient({ ...defaultParams, ...params });

export const client = createRpcClient(defaultParams);
