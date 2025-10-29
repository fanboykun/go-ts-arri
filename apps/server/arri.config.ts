import { defineConfig, servers, generators } from "arri";

export default defineConfig({
  server: servers.goServer(),
  // register client generators here
  generators: [
    generators.typescriptClient({
      clientName: "RpcClient",
      outputFile: "../web/rpc-client.gen.ts",
    }),
  ],
});
